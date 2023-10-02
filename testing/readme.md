generate coverage doc

```
go test  -coverprofile=coverage.out
```

obtain info in console

```
go tool cover -func=coverage.out
```

obtain info in html
```
go  tool cover -html=coverage.out -o coverage.html
``` 