# gin-rest-api-example
Rest API with Go and Gin Gonic. This example needs a postgres server running with a database created. The url to connect with the postgres server must be configured through a config file.


## Configuration

#### 1. Make a copy of .env file provided and fill it with values for each env var

```
cp pkg/common/envs/.env.example pkg/common/envs/.env
```

#### 2. Install dependencies with the following command

```
make install
```

<br>

## Running the server

```
make server
```

### Available endpoints
<br>

* Create a book

POST /books/

Request body:

```
{
    "title": "My first book",
    "author": "Me",
    "description": "This is my first book"
}
```

<br>

* List all books

GET /books/

<br>

* Get a book by ID

GET /books/:id

<br>

* Update a book

PUT /books/:id

Request body:

```
{
    "title": "Title changed",
    "author": "Great author",
    "description": "Great description here!"
}
```

<br>

* Delete a book by ID

DELETE /books/:id