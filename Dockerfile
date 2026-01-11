FROM ghcr.io/eriksoderblom/alpine-apache-php:25.12

WORKDIR /htdocs
COPY index.php insults.txt styles.css ./

EXPOSE 80