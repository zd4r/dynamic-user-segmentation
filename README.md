# dynamic-user-segmentation
Сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент).

## Запуск сервиса
Выполните команду для клонирования репозитория:
```bash
$ git clone git@github.com:zd4r/dynamic-user-segmentation.git
```
Перейдите в папку проекта:
```bash
$ cd dynamic-user-segmentation
```
Выполните команду для запуска сервиса:
```bash
$ make compose-build-up
```
Примечание: схема в БД создается при помощи инициализирующего скрипта `init.sql`, расположенного в корневой папке проекта.

## Примеры запросов / ответов
Ознакомиться с примерами запросов и ответов можно также в `swagger` документации.

Она доступна по ссылке http://localhost:8080/docs/ после запуска сервиса. Также файлы спецификации в формате `.yaml` и `.json` находятся в папке `docs` в корневой папке проекта.

### Создание сегмента
**Запрос:**
```bash
$ curl -X 'POST' \
'http://localhost:8080/v1/segment' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '{
      "slug": "avito_discount_50",
      "usersPercent": 50
    }'
```
**Ответ:**

`status code`: `201`

`response body`: `null`

### Удаление сегмента
**Запрос:**
```bash
$ curl -X 'DELETE' \
  'http://localhost:8080/v1/segment/AVITO_DISCOUNT_110' \
  -H 'accept: application/json'
```
**Ответ:**

`status code`: `200`

`response body`: `null`

### Создание пользователя
**Запрос:**
```bash
$ curl -X 'POST' \
  'http://localhost:8080/v1/user' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
        "id": 1
      }'
```
**Ответ:**

`status code`: `201`

`response body`: `null`

### Удаление пользователя
**Запрос:**
```bash
$ curl -X 'DELETE' \
  'http://localhost:8080/v1/user/1' \
  -H 'accept: application/json'
```
**Ответ:**

`status code`: `200`

`response body`: `null`

### Изменение набора сегментов пользователя
**Запрос:**

`expireAt` - опциональное поле, формат: `YYYY-MM-DDTHH:MM:SSZ`.

```bash
$ curl -X 'POST' \
  'http://localhost:8080/v1/user/1/segments' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "segmentsToAdd": [
    {
      "slug": "avito_discount_50",
      "expireAt": "2023-09-15T00:00:00Z"
    },
    {
      "slug": "avito_voice_messages"
    }
  ],
  "segmentsToRemove": [
    {
      "slug": "avito_voice_messages"
    }
  ]
}'
```
**Ответ:**

`status code`: `200`

`response body`: `null`

### Получение сегментов, к которым принадлежит пользователь
**Запрос:**
```bash
$ curl -X 'GET' \
  'http://localhost:8080/v1/user/1/segments' \
  -H 'accept: application/json'
```
**Ответ:**

`status code`: `200`

`response body`: 
```
{
    "segments": [
        "avito_discount_50"
    ]
}
```

## Доп. задание 1. 
Получение отчета по пользователю за определенный период времени.

### Решение
Для получения отчета в формате `CSV` по пользователю за определенный период времени, необходимо выполнить следующий запрос (формат опциональных `query` параметров: `YYYY-MM`, при отсутствии параметров, предоставляется весь доступный отчет): 
```bash
curl -X 'GET' \
  'http://localhost:8080/v1/user/1/report?from=2023-08&to=2023-09' \
  -H 'accept: application/json'
```
В результате данного запроса начнется загрузка файла с отчетом по пользователю в формате `CSV`. Подробнее с методом можно ознакомиться после запуска сервиса в [swagger документации](http://localhost:8080/docs/index.html#/user/get-user-report).

## Доп. задание 2.
Реализовать возможность задавать TTL (время автоматического удаления пользователя из сегмента)
### Решение
Опциональное добавление времени жизни сегмента для пользователя. Далее для соответствующих методов сервис проверяет, является ли сегмент актуальным для пользователя или нет, на основании этого поля. Также реализована регулярная очистка устаревших сегментов пользователя в БД.

**Запрос:**

`expireAt` - опциональное поле, формат (UTC): `YYYY-MM-DDTHH:MM:SSZ`.

```bash
$ curl -X 'POST' \
  'http://localhost:8080/v1/user/1/segments' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "segmentsToAdd": [
    {
      "slug": "avito_discount_50",
      "expireAt": "2023-09-15T00:00:00Z"
    },
    {
      "slug": "avito_voice_messages"
    }
  ],
  "segmentsToRemove": [
    {
      "slug": "avito_voice_messages"
    }
  ]
}'
```
**Ответ:**

`status code`: `200`

`response body`: `null`

Подробнее с методом можно ознакомиться после запуска сервиса в [swagger документации](http://localhost:8080/docs/index.html#/user/update-user-segmentsreport).

## Доп. задание 3.
В методе создания сегмента, добавить опцию указания процента пользователей, которые будут попадать в сегмент автоматически.
### Решение
В метод создания сегмента используется опциональный параметр `usersPercent`. При его указании заданному проценту текущих пользователей будет присвоен создаваемый сегмент (пользователи выбираются случайным образом без повторения).

**Запрос:**
```bash
$ curl -X 'POST' \
'http://localhost:8080/v1/segment' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '{
      "slug": "avito_discount_50",
      "usersPercent": 50
    }'
```
**Ответ:**

`status code`: `201`

`response body`: `null`

## TODO:
* Unit и интеграционные тесты