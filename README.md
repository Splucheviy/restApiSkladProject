# Rest Api для склада товаров из Китaя и Казахстана

## Реализация

Стандартный пакет http и mux, для swager использовался swaggo  

## Ручки

http://localhost:8080/ - Домашняя страница  
http://localhost:8080/swagger/index.html - Страничка swaggera  

Get /goods - вывод всех товаров на складе  
Get /goods/{id} - вывод товара со склада по ID-шнику  
POST /goods - создание товара  
PUT /goods/{id} - изменение товара по заданному id-шнику  
DELETE /goods/{id} - удаление товара из списка товаров  

Приложуха работает на 80:80 порту  

Основная реализованная модель:  
ID - ид товара  
Name - его название  
Provider - поставщик (1 - Kz / 2 - Китай)  
Price - его цена  
Quantity - наличиие на складе  

## Build

make; ./rest_api_sklad_project  

## TODO

Вариант с postgre смотреть в psql branch
