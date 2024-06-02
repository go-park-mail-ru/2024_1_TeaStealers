# Используем официальный образ PostgreSQL
FROM postgres:latest

RUN apt-get update && \
    apt-get install -y postgresql-contrib postgis

RUN apt-cache show postgresql-auto-explain | grep -E "Version:\s" | awk '{print $2}' | head -n 1 | xargs apt-get install -y

# Копируем файлы во временную директорию
COPY HW2/inituser.sql /tmp/inituser.sql
COPY HW2/initdb.sql /tmp/initdb.sql
# COPY db/migrations/*.sql /tmp/

# Создаем объединенный файл SQL
RUN cat /tmp/*.sql >> /docker-entrypoint-initdb.d/combined.sql && \
    rm -f /tmp/*.sql

# Копируем файл конфигурации PostgreSQL в контейнер
COPY HW2/postgresql.conf /etc/postgresql/postgresql.conf
COPY HW2/pg_hba.conf /etc/postgresql/pg_hba.conf