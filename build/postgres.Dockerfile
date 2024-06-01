# Используем официальный образ PostgreSQL
FROM postgres:latest

# Устанавливаем необходимые пакеты, включая postgresql-contrib и postgis
RUN apt-get update && apt-get install -y postgresql-contrib postgis && rm -rf /var/lib/apt/lists/*

# Копируем файл конфигурации PostgreSQL в контейнер
COPY HW2/postgresql.conf /etc/postgresql/postgresql.conf
COPY HW2/pg_hba.conf /etc/postgresql/pg_hba.conf

# Копируем скрипты инициализации базы данных в директорию /docker-entrypoint-initdb.d/
COPY HW2/initdb.sql /docker-entrypoint-initdb.d/
COPY db/migrations/*.sql /docker-entrypoint-initdb.d/