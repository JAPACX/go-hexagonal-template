
### Quick start gqlgen

1 - go mod init [example]
2 - go get github.com/99designs/gqlgen 
3 - printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
4 - go mod tidy
5 - go run github.com/99designs/gqlgen init
6 - go mod tidy

### for regenerate models after new changes
```go run github.com/99designs/gqlgen generate```


### run server 

```go run server.go```


