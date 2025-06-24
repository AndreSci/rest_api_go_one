# rest_api_go_one
ju test


- Get Books\n
GET 127.0.0.1:8081/books\n

- Get Book By Id\n
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