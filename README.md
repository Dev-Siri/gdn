# GDN

GDN is a fast CDN (Content Delivery Network) written in Go.

## Getting Started

Make sure you have Go installed on your system.

- Clone the repository

```sh
$ git clone https://github.com/Dev-Siri/gdn
```

- Run it with Go.

```sh
$ go run main.go
```

- Or, run the `build.sh` script to compile Go for all architectures

```sh
$ ./build.sh
```

To use GDN, you need to provide a `gdn.json` configuration file which accepts the following options:

```json
{
  "origin_server": "https://your-origin-server.com",
  "log": true,
  "cache_dir": ".cache"
}
```

- `origin_server`: The URL of the origin server from which GDN will fetch files when they are not available in the cache. (Required)
- `log`: A boolean value to enable or disable structured logging. Set to `true` to enable logging or `false` to disable it.
- `cache_dir`: The directory where GDN will store cached files.

## License

This project is [MIT](LICENSE.md) Licensed.
