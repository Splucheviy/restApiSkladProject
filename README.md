# Rest Api для склада товаров из Китaя и Казахстана

## Реализация

1. Сервер - hhtp + mux  
2. База данных - postgre  
3. Миграции - goose  
4. Swagger - swaggo  

## Ручки

http://localhost:8080/ - Домашняя страница  
http://localhost:8080/swagger/index.html - Страничка swaggera  

Get /goods - вывод всех товаров на складе  
Get /goods/{id} - вывод товара со склада по ID-шнику  
POST /goods - создание товара  
PUT /goods/{id} - изменение товара по заданному id-шнику  
DELETE /goods/{id} - удаление товара из списка товаров  

Основная реализованная модель:  
ID - ид товара  
Name - его название  
Provider - поставщик (1 - Kz / 2 - Китай)  
Price - его цена  
Quantity - наличиие на складе  

## Build
Должен быть свободен :5431 порт для БД  

Запуск docker compose:  
```
docker compose up -d
```
Остановка контейнера:  
```
docker compose down
```

БД заполняется в main из мока, но можно накатить миграцию через:  
```
goose -dir pkg/db/migrations postgres "postgresql://john:ramzesik@127.0.0.1:5431/goodsDB?sslmode=disable" up
```

Почистить миграцию через:  
```
goose -dir pkg/db/migrations postgres "postgresql://john:ramzesik@127.0.0.1:5431/goodsDB?sslmode=disable" down
```

Билд и запуск сервера через make:  
```
make; ./rest_api_sklad_project  
```

## TODO

1. Реализовать логирование - slog / logrus.  
2. Конфиг - viper.  
3. Сервер - gin / echo / fiber.  
4. Тесты - tistify    
