# Programmer Insults
Source code for programmerinsults.com

To build and run locally:

```bash
docker build -t ghcr.io/carbontwelve/programmer-insults:latest .
docker run --rm -p 8080:80 ghcr.io/carbontwelve/programmer-insults:latest
```

The application image makes use of [`eriksoderblom/alpine-apache-php`](https://github.com/eriksoderblom/alpine-apache-php) as its base.