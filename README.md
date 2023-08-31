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
      "slug": "AVITO_DISCOUNT_110"
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
```bash
$ curl -X 'POST' \
  'http://localhost:8080/v1/user/1/segments' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
      "segmentsToAdd": [
        "avito_discount_50",
        "avito_voice_messages"
      ],
      "segmentsToRemove": [
        "avito_performance_vas"
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
        "avito_discount_50",
        "avito_voice_messages"
    ]
}
```