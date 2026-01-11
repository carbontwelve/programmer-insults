# Programmer Insults
Source code for programmerinsults.com

To build and run locally:

```bash
docker build -t ghcr.io/carbontwelve/programmer-insults:latest .
docker run --rm -p 8080:8080 ghcr.io/carbontwelve/programmer-insults:latest
```

The application is built with symbols and debugging information stripped and then passed through [`upx`](https://upx.github.io/) in order to produce the smallest binary possible. The resulting image size for this *basic* application is ~21MB with the binary compressed to ~3MB. 