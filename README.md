# rest_api_go_one
ju test


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


Friendly remind -------------
// go get -u github.com/lib/pq
// требуется для установки драйвера коннекта к БД postgres
// _ "github.com/lib/pq" // Имптор для сторонних эффектов

docker-compose up -d