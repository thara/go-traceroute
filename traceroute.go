package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func run(hostname string) error {

	port := 33434

	udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Printf("traceroute to %s (%s)\n", hostname, udpAddr.IP.String())

	curTTL := 1
	maxTTL := 64

	retry := 3

	for {
		fmt.Printf("%3d ", curTTL)

		for i := 0; i < retry; i++ {
			sendConn, err := dialSendConn(udpAddr, port, curTTL)
			if err != nil {
				return errors.WithStack(err)
			}
			recvConn, err := listenICMP()
			if err != nil {
				return errors.WithStack(err)
			}

			_, err = sendConn.Write([]byte{})
			if err != nil {
				return errors.WithStack(err)
			}

			rb := make([]byte, 64)
			_, from, err := recvConn.ReadFrom(rb)
			if err != nil {
				if os.IsTimeout(err) {
					// retry
					fmt.Printf("* ")
					continue
				}
				return errors.WithStack(err)
			}

			msg, err := icmp.ParseMessage(ipv4.ICMPTypeTimeExceeded.Protocol(), rb)
			if err != nil {
				return errors.WithStack(err)
			}

			names, _ := net.LookupAddr(from.String())
			if 0 < len(names) {
				fmt.Printf("%s (%s)", names[0], from.String())
			} else {
				fmt.Printf(from.String())
			}

			if from.String() == udpAddr.IP.String() {
				return nil
			}
			if msg.Type == ipv4.ICMPTypeDestinationUnreachable {
				return nil
			}

			break
		}
		fmt.Printf("\n")

		curTTL += 1
		port += 1

		if maxTTL < curTTL {
			break
		}
	}

	return nil
}

func dialSendConn(udpAddr *net.UDPAddr, port int, ttl int) (*net.UDPConn, error) {
	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	raw, err := conn.SyscallConn()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var sysErr error
	err = raw.Control(func(fd uintptr) {
		sysErr = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_TTL, ttl)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if sysErr != nil {
		return nil, fmt.Errorf("syscall.SetsockoptInt failed: %+v", sysErr)
	}

	return conn, nil
}

func listenICMP() (*net.IPConn, error) {
	conn, err := net.ListenIP("ip4:icmp", nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return conn, nil
}
