# go-traceroute

simple traceroute ports in Go.

## Installation

```
$ go install github.com/thara/go-traceroute@latest
```

## Usage

```
$ go-traceroute -h
Usage: go-traceroute [OPTIONS] HOSTNAME

  -f int
        first TTL (default 1)
  -m int
        max TTL (default 64)
  -n int
        retry count (default 3)
  -port int
        port (default 33434)
```

NOTE: macOS requires root access to run it correctly.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details

## Author

Tomochika Hara ([thara](https://thara.dev))
