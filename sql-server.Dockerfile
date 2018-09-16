FROM mysql:latest

COPY backup.sql.gz /docker-entrypoint-initdb.d/backup.sql.gz

EXPOSE 3306