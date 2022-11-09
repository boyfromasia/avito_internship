# Микросервис для работы с балансом пользователей

## Стек

* Язык разработки: Golang.
* БД: PostgreSQL
* REST API фреймворк: Gin

## Структура Проекта

![alt text](img/1.png)

## Архитектура БД

* Orders - История только покупок юзеров, без пополнения счета (Могут иметь три статуса: _waited, cancelled, approved_)
* Purchases - Наименование покупок
* HistoryUser - История только подтвержденных транзакций для всех юзеров
* User - Баланс юзера

![alt text](img/db.png)

## Реализовано 

* Метод начисления средств на баланс.

* Метод получения баланса пользователя.

* Метод резервирования средств с основного баланса на отдельном счете.

* Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии.

## Примеры

* **POST /user**

Метод начисления средств на баланс. Если пользователя не было в базе, создастся новый с указанным балансом,
если пользоватеь есть, средства начилсятся на баланс. Поле _balance_ валидируется, он может быть только > 0

Пример:


```json
{
    "user_id": 111,
    "balance": 321
}
```

Возвращает: 

```json
{
    "status": "OK"
}
```


* **GET /user** 

Метод получения баланса пользователя.

Пример:


```json
{
  "user_id": 111
}
```

Возвращает:

```json
{
  "balance": 121
}
```


* **POST /order/create**

Метод резервирования средств с основного баланса на отдельном счете. Поле _price_ валидируется. Он может быть только > 0. Пример:


```json
{
  "order_id": 4,
  "user_id": 111,
  "purchase_id": 1,
  "price": 1
}
```

Возвращает:

```json
{
    "status": "OK"
}
```


* **POST /order/decision**

Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии.
Проверяет по всем полям есть ли такой заказ. Поля _decision_ может быть _"cancelled"_
или _"approved"_. Поле _decision_ валидируется, то есть уже признанные покупки не могут поменять статус. Пример:


```json
{
  "order_id": 4,
  "user_id": 111,
  "purchase_id": 1,
  "decision": "cancelled",
  "price": 1
}
```

Возвращает:

```json
{
    "status": "OK"
}
```
