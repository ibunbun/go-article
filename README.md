###

A simple article web service with cache written in Golang.

### Usage
```
docker-compose build
docker-compose up
```

### Post Data
```
curl --location --request POST 'localhost:8080/articles' \
--header 'Content-Type: application/json' \
--data-raw '{
"author": "Ibnu",
"title": "Berita Satu",
"body": "Isi berita 1"
}'
```


### Get Data

```
curl --location --request GET 'http://localhost:8080/articles?author=ibnu&query=berita'
```

### Test

```
go test ./handler
```
