# Используем официальный образ PostgreSQL
FROM postgres:latest

# Устанавливаем необходимые пакеты, включая postgresql-contrib и postgis
RUN apt-get update && apt-get install -y postgresql-contrib postgis

# Копируем файл конфигурации PostgreSQL в контейнер
COPY HW2-3/postgresql.conf /etc/postgresql/postgresql.conf

# Копируем скрипты инициализации базы данных в директорию /docker-entrypoint-initdb.d/
COPY db/migrations/*.sql /docker-entrypoint-initdb.d/

# Добавляем команду для создания расширения PostGIS при инициализации базы данных
RUN echo "CREATE EXTENSION IF NOT EXISTS \"postgis\";" >> /docker-entrypoint-initdb.d/init.sql
RUN echo "CREATE EXTENSION IF NOT EXISTS \"pg_stat_statements\";" >> /docker-entrypoint-initdb.d/init.sql
RUN echo "CREATE EXTENSION IF NOT EXISTS \"auto_explain\";" >> /docker-entrypoint-initdb.d/init.sql
RUN echo "SET auto_explain.log_analyze = ON;" >> /docker-entrypoint-initdb.d/init.sql
RUN echo "SET auto_explain.log_min_duration \= \'100ms\';" >> /docker-entrypoint-initdb.d/init.sql
RUN echo "SET auto_explain.log_timing = ON;" >> /docker-entrypoint-initdb.d/init.sql
