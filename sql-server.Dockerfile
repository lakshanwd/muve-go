FROM mysql:latest

COPY backup.sql.gz /docker-entrypoint-initdb.d/backup.sql.gz

ENV MYSQL_ROOT_PASSWORD=123
ENV MYSQL_USER=developer
ENV MYSQL_PASSWORD=123
ENV MYSQL_DATABASE=XYZ_lakshan_dissanayake

EXPOSE 3306