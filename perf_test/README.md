## Основная сущность
Основной сущностью у нас являются объявления.

## Инструмент тестирования
Мы использовали для тестирования https://github.com/wg/wrk

## Генерация данных
Для генерации данных мы написали скрипт на python, который подключается к контейнеру docker с 
postgresql и отправляет туда запросы на заполнение.

Для тестирования мы заполнили таблицы advert, user_data, favourite_advert.

## Что тестируем
Тестируем насколько будет существенна разница между рассчётом количества лайков у объявления через смежную таблицу, и
если это значение запоминать в отдельном поле likes в advert. И изменять его при помощи триггеров.

## Выполнение тестирования

Было создано 100 000 сущностей advert user_data и favourite_advet(лайков).

После скриптом на python (функция main_bench) было проведено на машине несколько
нагрузочных тестирований. Результаты тестирования:
```
Test Result count likes:
Running 30s test @ http://tean.homes/api/test/count/63671
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.84ms    7.50ms 292.36ms   98.86%
    Req/Sec     3.95k   362.91     7.73k    91.58%
  471336 requests in 30.03s, 146.52MB read
  Non-2xx or 3xx responses: 471336
Requests/sec:  15697.20
Transfer/sec:      4.88MB

Test Result fast get:
Running 30s test @ http://tean.homes/api/test/fast/63671
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.75ms    7.03ms 306.08ms   98.93%
    Req/Sec     3.94k   315.80     7.83k    89.33%
  470087 requests in 30.02s, 146.13MB read
  Non-2xx or 3xx responses: 470087
Requests/sec:  15657.21
Transfer/sec:      4.87MB




Test Result count likes:
Running 30s test @ http://tean.homes/api/test/count/97774
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.76ms    7.59ms 323.87ms   99.09%
    Req/Sec     3.94k   316.76     6.50k    82.67%
  470002 requests in 30.03s, 146.10MB read
  Non-2xx or 3xx responses: 470002
Requests/sec:  15653.68
Transfer/sec:      4.87MB

Test Result fast get:
Running 30s test @ http://tean.homes/api/test/fast/97774
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.69ms    6.18ms 277.93ms   98.83%
    Req/Sec     3.92k   328.46     7.07k    86.42%
  468659 requests in 30.02s, 145.68MB read
  Non-2xx or 3xx responses: 468659
Requests/sec:  15610.25
Transfer/sec:      4.85MB




Test Result count likes:
Running 30s test @ http://tean.homes/api/test/count/80434
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.73ms    6.76ms 295.81ms   98.89%
    Req/Sec     3.94k   323.75     7.30k    88.75%
  470922 requests in 30.02s, 146.39MB read
  Non-2xx or 3xx responses: 470922
Requests/sec:  15685.08
Transfer/sec:      4.88MB

Test Result fast get:
Running 30s test @ http://tean.homes/api/test/fast/80434
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.85ms    7.19ms 312.27ms   99.05%
    Req/Sec     3.87k   284.24     6.45k    86.91%
  461453 requests in 30.02s, 143.44MB read
  Non-2xx or 3xx responses: 461453
Requests/sec:  15369.38
Transfer/sec:      4.78MB
```
Здесь в Test Result count likes написана статистика для запросов через подсчёт лайков через смежную таблицу,
а в Test Result fast get, получение из аттрибута likes в advert.

Как можно заметить, есть разница примерно на 40-50, но на последнем в 300 запросов в секунду.
Из этого можно сделать вывод, что разница между способами сильно зависит от того, как много лайков у объявления,
следовательно на больших объёмах данных данная оптимизация значительно улучшит работу сервиса.