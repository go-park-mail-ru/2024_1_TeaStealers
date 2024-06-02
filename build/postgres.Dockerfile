# Используем официальный образ PostgreSQL
FROM postgres:latest

RUN apt-get update && \
    apt-get install -y postgresql-contrib postgis

RUN apt-cache show postgresql-auto-explain | grep -E "Version:\s" | awk '{print $2}' | head -n 1 | xargs apt-get install -y

COPY HW2/postgresql.conf /etc/postgresql/postgresql.conf
COPY HW2/pg_hba.conf /etc/postgresql/pg_hba.conf

COPY HW2/initdb.sql /docker-entrypoint-initdb.d/
COPY HW2/inituser.sql /docker-entrypoint-initdb.d/
COPY db/migrations/*.sql /docker-entrypoint-initdb.d/
