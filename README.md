# Golang API Server Boilerplate (Work In Progress)

A boilerplate/starter project for quickly building APIs using Golang with Clean Architechture 

## 🎨 Features

- **Web Framework**: Switchable between [Mux](https://github.com/gorilla/mux). (will be added more later...)
- **Logging**: Switchable between [Zap](https://github.com/uber-go/zap) or plain logging.
- **SQL database**: Using [GORM](https://gorm.io/index.html) for ORM library. (will be added more later...)

## 🌲 Project Folder Structure

```
+-- app
|  +-- controller
|  +-- entity
|  +-- error
|  +-- infrastructure
|  |   +-- gorm
|  |   |   +-- database
|  |   |   +-- repository
|  |   +-- http
|  |   +-- logging
|  +-- usecase
+-- cmd
```

## 👮 License

[MIT](LICENSE)