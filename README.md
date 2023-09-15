## BookStore Management Apis (Beginner Friendly)
___
* Database - Mysql
* GORM
* Json marshall, unmarshall
* Project structure
* Gorilla Mux

```go
# 初始化项目
go mod init github.com/lipaysamart/withGoBookStore

# 导入需要的依赖包
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"
```


## Directer Striction
___
```
├── cmd
│   └── main
│       ├── main.exe
│       └── main.go
├── go.mod
├── go.sum
├── pkg
│   ├── config
│   │   └── app.go
│   ├── controllers
│   │   └── book-controller.go
│   ├── models
│   │   └── book.go
│   ├── routes
│   │   └── bookstore-routes.go
│   └── utils
│       └── utils.go
└── README.md
```