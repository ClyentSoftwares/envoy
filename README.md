# Envoy

Envoy is a lightning-fast, environment variable-based URL redirector written in Go. It allows you to configure multiple redirects entirely from environment variables, making it perfect for simple redirects where you don't want to manage a config file or database.

## Features

- Configure redirects using environment variables
- Set redirect status code via environment variable
- Optimized for speed and low resource usage
- Docker support with multi-architecture builds (AMD64 and ARM64)

## Usage

1. Set your redirects as environment variables:

   ```
   REDIRECT_TWITTER=https://twitter.com/clyentsoftwares
   REDIRECT_GITHUB=https://github.com/clyentsoftwares
   ```

2. (Optional) Set the redirect status code:

   ```
   REDIRECT_STATUS=302
   ```

   If not set, redirects will use status 301 (Moved Permanently) by default. You can set any valid HTTP redirect status code (300-399).

3. Run the application:

   ```
   docker run -p 8080:8080 \
     -e REDIRECT_TWITTER=https://twitter.com/clyentsoftwares \
     -e REDIRECT_GITHUB=https://github.com/clyentsoftwares \
     -e REDIRECT_STATUS=302 \
     ghcr.io/clyentsoftwares/envoy:latest
   ```

4. Access your redirects:

   ```
   http://localhost:8080/twitter -> redirects to https://twitter.com/clyentsoftwares
   http://localhost:8080/github -> redirects to https://github.com/clyentsoftwares
   ```

## Common Redirect Status Codes

- 301: Moved Permanently (default)
- 302: Found (temporary redirect)
- 307: Temporary Redirect
- 308: Permanent Redirect

## Building

To build the application locally:

```
go build -o envoy
```

## Docker

To build the Docker image locally:

```
docker build -t envoy .
```

Note: The GitHub Actions workflow builds multi-architecture images (AMD64 and ARM64) automatically.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
