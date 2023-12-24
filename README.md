// На Ubuntu 22.04.2 LTS отработало успешно.

Используется безголовый хром с помощью `go-rod/rod`. 

Если в `~/.cache/rod/browser/chromium-12345` хрома нет то загрузится сам при первом запуске.
У меня загружаться не торопился, рекомендую отойти колотить чай иль еще чего; после из кеша удалить.

Запускается так:

`$ ./dromtest -file=filename -concurrency=1`

concurrency опциональный параметр

Предусмотрено тестов в `./run_tests.sh`. Советую запускать это.
Там билд, и подготовка тестовых данных, и нужные проверки. Если скрипт в консоль ничего не написал - значит тесты пройдены успешно.

Что в коде делается:<br>
одна горутина читает файл построчно и пишет номера в небуф.канал<br>
из этого канала читает n = concurrency горутин-обработчиков<br>
&nbsp;каждый обработчик<br>
&nbsp;&nbsp;range по каналу<br>
&nbsp;&nbsp;добывает нужный json с помощью go-rod<br>
&nbsp;&nbsp;парсит, опционально загружает изображение<br>
&nbsp;&nbsp;записывает в папку<br>


Минусы решения:<br>
Обработка ошибок дурная (при ненахождении элемента в браузере будет паника к примеру), на обработку каждого номера запускается браузер (без переиспользования): стабильность, удобство, скорость - не про этот код, от этого кода не просили большего.
