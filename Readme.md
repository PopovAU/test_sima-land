# Тестовое задание HTTP сервис sima-land

[само задание](https://github.com/sima-land/intern-test/blob/master/golang/rest-service/README.md)


 - src/api/main.go - main файл:
    - Инициирует чтение настроек из environment;
    - Инициирует пакет db;
    - Запускает myhttp
 - src/api/db/ - пакет работы с БД:
    - Подключается к MySQL;
    - Создает базу, если ее еще нет;
    - Создает таблицу, если ее еще нет;
    - Описывает структуру User с возможностями CRUD;
 - src/api/setting/ - пакет настроек;
 - src/api/myhttp/ - пакет запускает API Listener. API есть handler:
    - [GET] / статический фай для проверке работы API;
    - [GET]    /user?id=:id - запрашивает информацию о пользователе. Если пользователь найден то вернется JSON с данными пользователя. Если не найдется то 404.
    - [POST]   /user - создание нового пользователя;
    - [PUT]    /user - изменение имени пользователя;
    - [DELETE] /user - удаление пользователя;

Самый простой способ запустить приложение - docker-compose up -d

Для теста запущен [тут](http://185.185.68.195:9080/) до 11.04.2018
