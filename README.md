# gin-gonic.com-docs-examples-upload-file-single-file

- Single file

- Reference: https://gin-gonic.com/docs/examples/upload-file/single-file/

## gvm

```sh
gvm use go1.23.5
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

- POST /file

```sh
curl --location 'localhost:8080/file' \
--form 'file=@"example.png"' \
--header "Content-Type: multipart/form-data"
```
