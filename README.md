# rest_api_go_one
Тестовое задание. Реализация CRUD API.

Есть не нужен Автотест нужно удалить в main.go 	

// AUTO-TEST
go unittest_test.RunTests()

Пример запросов для Postman:

- Get Books

    GET 127.0.0.1:8081/books

- Get Book By Id

    GET 127.0.0.1:8081/book?id=3

- Add Book

    POST 127.0.0.1:8081/book

    Json 

    {
    "name": "NewBookFromRequest",
    "author": "NewAuthor"
}

- Update Book by ID

    PUT 127.0.0.1:8081/book

    Json

    {
    "id": 2,
    "name": "UpdateookFromRequest",
    "author": "UpdateAuthor"
}

- Delete Book by ID
    DELETE 127.0.0.1:8081/book?id=3


Friendly remind <-------------

go get -u github.com/lib/pq

требуется для установки драйвера коннекта к БД postgres

_ "github.com/lib/pq" // Имптор для сторонних эффектов

Развернуть базу PostgresSQL в Docker контейнере
docker-compose up -d

