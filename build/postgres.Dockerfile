# Используем официальный образ PostgreSQL
FROM postgres:latest

# Устанавливаем необходимые пакеты, включая postgresql-contrib и postgis
RUN apt-get update && apt-get install -y postgresql-contrib postgis

# Копируем файл конфигурации PostgreSQL в контейнер
COPY postgresql.conf /etc/postgresql/postgresql.conf

# Копируем скрипты инициализации базы данных в директорию /docker-entrypoint-initdb.d/
COPY db/migrations/*.sql /docker-entrypoint-initdb.d/

# Добавляем команду для создания расширения PostGIS при инициализации базы данных
RUN echo "CREATE EXTENSION IF NOT EXISTS \"postgis\";" >> /docker-entrypoint-initdb.d/init.sql
