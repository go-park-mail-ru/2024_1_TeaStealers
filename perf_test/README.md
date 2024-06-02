## Основная сущность
Основной сущностью у нас являются объявления.

## Инструмент тестирования
Мы использовали для тестирования https://github.com/wg/wrk

## Как запустить скрипт
Командой в директории с кодом: python3 perf_test.py </br >
При этом можно закомментировать вызов main() и оставить только
main_bench(), чтобы не генерировать заново базу!

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
Running 30s test @ https://tean.homes/api/test/count/49904
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   858.05ms  345.62ms   1.95s    69.02%
    Req/Sec    30.12     18.04   171.00     61.70%
  3336 requests in 30.08s, 2.62MB read
  Socket errors: connect 0, read 0, write 0, timeout 36
  Non-2xx or 3xx responses: 589
Requests/sec:    110.89
Transfer/sec:     89.17KB

Test Result fast get:
Running 30s test @ https://tean.homes/api/test/fast/49904
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   426.88ms  488.40ms   1.99s    78.86%
    Req/Sec   113.30     96.96   570.00     79.16%
  12585 requests in 30.09s, 10.85MB read
  Socket errors: connect 0, read 0, write 0, timeout 18
Requests/sec:    418.25
Transfer/sec:    369.22KB




Test Result count likes:
Running 30s test @ https://tean.homes/api/test/count/26971
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   853.42ms  362.77ms   1.98s    68.19%
    Req/Sec    30.05     16.83   121.00     62.88%
  3373 requests in 30.09s, 2.62MB read
  Socket errors: connect 0, read 0, write 0, timeout 31
  Non-2xx or 3xx responses: 633
Requests/sec:    112.08
Transfer/sec:     89.23KB

Test Result fast get:
Running 30s test @ https://tean.homes/api/test/fast/26971
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   421.86ms  486.15ms   2.00s    79.09%
    Req/Sec   116.45     97.45   580.00     76.38%
  12591 requests in 30.09s, 10.81MB read
  Socket errors: connect 0, read 0, write 0, timeout 20
Requests/sec:    418.48
Transfer/sec:    367.79KB




Test Result count likes:
Running 30s test @ https://tean.homes/api/test/count/58075
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   857.43ms  363.62ms   1.97s    67.62%
    Req/Sec    29.93     16.27   101.00     63.05%
  3347 requests in 30.06s, 2.45MB read
  Socket errors: connect 0, read 0, write 0, timeout 34
  Non-2xx or 3xx responses: 626
Requests/sec:    111.36
Transfer/sec:     83.38KB

aTest Result fast get:
Running 30s test @ https://tean.homes/api/test/fast/58075
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   421.35ms  486.06ms   1.97s    78.89%
    Req/Sec   117.75    110.11   575.00     80.77%
  12421 requests in 30.04s, 9.95MB read
Requests/sec:    413.51
Transfer/sec:    339.20KB
```
Здесь в Test Result count likes написана статистика для запросов через подсчёт лайков через смежную таблицу,
а в Test Result fast get, получение из аттрибута likes в advert.

Как можно заметить, есть разница в 4 раза по количеству запросов в секунду!
Из этого можно сделать вывод, что разница между способами сильно зависит от того, как много лайков у объявления,
следовательно, на больших объёмах данных данная оптимизация значительно улучшит работу сервиса.