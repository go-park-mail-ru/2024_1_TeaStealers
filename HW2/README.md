## Защита от SQL Injections
Для защиты от sql инъекций библиотека  jackc/pgx/v4 поддерживает использование placeholders для запросов.

Пример запроса для таблицы advert:
```
insert := `INSERT INTO 
advert (user_id, type_placement, title, description, phone, is_agent, priority)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
```

## Пулл соединений и параметры соединений

Для пула соединений используем jackc/pgx/v4/pgxpool, при старте сервиса создаём пул соединений. И используем его в
функциях в repo сервисов.

Конфиг пула создаётся в "2024_1_TeaStealers/internal/pkg/config", и вызывается в main каждого микросервиса.

max_connections поставим немного больше, чем потенциально все сервисы могут установить одновременно. <br />
max_connections = 150. <br />
Сервис(repo) - количество соединений: <br />
adverts - 50 <br />
users - 30 <br />
auth - 30 <br />
complexes - 10 <br />
images - 20 <br />
questionnaire - 5

в сумме получается: 145

Значение max_connections в postgresql.conf должно быть чуть больше, чем максимальное количество соединений в connection pool.
Так всегда будет в запасе несколько свободных соединений к базе.

listen_addresses = 'localhost' - так как все микросервисы на одной машине.


## Таймауты
d
Мы поставили таймаут в 15 секунд, потому что исходя из бизнес-логики не имеет смысла так долго ждать выполнения запроса, он уже уйдёт к этому времени...
Ставить ограничение больше смысла не имеет.
 
```
statement_timeout = 15s 
lock_timeout = 15s
```

В случае DOS-атаки, злоумышленники будут получать timeout и сервер не будет нагружен слишком большим запросом долго. <br />
В случае DDOS-атаки, пул соединений у микросервиса будет израсходован злоумышленниками и микросервис не будет работать. (ещё зависит какой тип запроса будет к базе во время атаки)


## Логгирование и протоколирование медленных запросов

Эти параметры в `postgresql.conf` настраивают логирование.

```
# %t: временная метка, %p: идентификатор процесса
log_line_prefix = '%t [%p]: '

# Включает сбор логов в файлы, а не вывод в консоль.
logging_collector = on

# Определяет каталог, в котором будут храниться файлы журналов.
log_directory = 'log'

# Определяет шаблон имени файлов журналов, включая дату и время.
log_filename = 'postgresql-%Y-%m-%d_%H%M%S.log'

# Журналирует длительность каждого выполненного запроса, занимающего больше указанного времени (в миллисекундах).
log_min_duration_statement = 7

# Устанавливает уровень подробности сообщений об ошибках в журналах.
log_error_verbosity = verbose

# Журналирует все SQL-запросы.
log_statement = 'all'

# Журналирует длительность каждого SQL-запроса.
log_duration = on

# Журналирует ожидания блокировок, длительность которых превышает deadlock_timeout.
log_lock_waits = on

# Журналирует подключения клиентов к базе данных.
log_connections = on

# Журналирует отключения клиентов от базы данных.
log_disconnections = on

# Журналирует создание временных файлов запросами.
log_temp_files = 0
```

## Скрипт для создания пользователя
Находится в initdb.sql
Приведу переменные, которые в env находятся:
```
JWT_SECRET=some_secret
DB_HOST=db
DB_NAME=test
DB_PORT=5432
DB_USER=test
DB_PASS=test
POSTGRES_USER=serv_acc
POSTGRES_PASSWORD=randomPasswordSec
POSTGRES_DB=test
BASE_DIR=/Users/maksimshagaev/GolandProjects/Papka

GRPC_AUTH_PORT=8081
GRPC_AUTH_CONTAINER_IP=tean-auth

GRPC_USER_PORT=8082
GRPC_USER_CONTAINER_IP=tean-user

GRPC_ADVERT_PORT=8083
GRPC_ADVERT_CONTAINER_IP=tean-advert

GRPC_QUESTION_PORT=8084
GRPC_QUESTION_CONTAINER_IP=tean-question

GRPC_COMPLEX_PORT=8085
GRPC_COMPLEX_CONTAINER_IP=tean-complex

GRAFANA_DIR=/Users/maksimshagaev/GolandProjects/Papka
```