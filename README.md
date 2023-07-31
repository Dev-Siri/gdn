# Pretext

Yes I did use ChatGPT to write this, but please I don't want to write READMEs I hate it.

# GDN - Go Delivery Network

GDN is an incredibly fast CDN (Content Delivery Network) written in the Go programming language. It is designed to efficiently deliver content to end-users, providing a high-performance and reliable content delivery solution.

## Features

GDN comes with a wide range of powerful features, making it a versatile CDN suitable for various use cases:

1. **Extreme Speed**: GDN is optimized for speed, ensuring lightning-fast content delivery to users, resulting in reduced latency and improved user experiences.

2. **Efficient Caching**: The CDN employs intelligent caching mechanisms to store frequently accessed content locally. This reduces the need to repeatedly fetch content from the origin server, leading to faster responses.

3. **Customizable Configuration**: The configuration file, `gdn.json`, allows users to tailor various aspects of the CDN, including cache directory location, logging preferences, and the origin server URL.

4. **Structured Logging**: GDN supports structured logging, providing detailed insights into its performance and request handling. Users can enable or disable logging based on their requirements.

5. **Dynamic Request Handling**: GDN can dynamically adjust its request handling based on traffic conditions, ensuring optimal performance during periods of high demand.

6. **High Scalability**: GDN can scale horizontally by deploying multiple edge servers across various regions, ensuring efficient content delivery to a global audience.

7. **Efficient File Serving**: The CDN efficiently serves static assets such as images, videos, stylesheets, and scripts, reducing server load and improving content delivery times.

8. **Content Origin Configuration**: The origin server URL is a required configuration option in `gdn.json`. This allows GDN to fetch files from the specified server when they are not available in the cache.

10. **User-Friendly Index Page**: GDN provides a user-friendly index page that can also serve as a health-check route, allowing users to verify the status of their CDN.

13. **Optimized Concurrency**: GDN utilizes Go's inherent concurrency features to handle multiple requests concurrently, ensuring efficient resource utilization.

## Getting Started

To use GDN, you need to provide a `gdn.json` configuration file with the following options:

```json
{
  "origin_server": "https://your-origin-server.com",
  "log": true,
  "cache_dir": ".cache"
}
```

- `origin_server`: The URL of the origin server from which GDN will fetch files when they are not available in the cache. This option is required.
- `log`: A boolean value to enable or disable structured logging. Set to `true` to enable logging or `false` to disable it.
- `cache_dir`: The directory location where GDN will store cached files.

## Conclusion

GDN is an exceptional CDN solution written in Go, designed to deliver content at extreme speed and with high scalability. Its efficient caching mechanisms, rate limiting capabilities, and user-friendly configuration make it an excellent choice for improving content delivery across the web.

Please note that GDN is an open-source project, and contributions from the developer community are welcome. Feel free to explore and contribute to the GDN repository to enhance its features and functionalities.

For more information and detailed instructions, visit the [GDN GitHub repository](https://github.com/your-gdn-repo-link).
