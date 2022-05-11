# Golang API Server Boilerplate (Work In Progress)

A boilerplate/starter project for quickly building APIs using Golang with Clean Architechture 

## 🎨 Features

- **Web Framework**: Switchable between [Mux](https://github.com/gorilla/mux), [Gin](https://github.com/gin-gonic/gin).
- **Logging**: Switchable between [Zap](https://github.com/uber-go/zap) or plain logging.
- **SQL database**: Using [GORM](https://gorm.io/index.html) for ORM library. (will be added more later...)
- **Testing**: Use [Testify](https://pkg.go.dev/github.com/stretchr/testify) for testing framework.
- **[Docker](https://docker.com/)** support: 
  - Using multi-stage build to reduce **production** image size
  - Using [nodemon](https://nodemon.io/) to do live-reload when **development**

## 🐳 How to run using Docker
- Start production enviroment
```console
$ docker-compose up -d production
// View logs
$ docker-compose logs --tail 100 -f production
```

- Start development enviroment 
```console
$ docker-compose up development
```

- Re-building docker
```console
$ docker-compose build --no-cache
```

- Attach to bash
```console
$ docker-compose exec <production|development> sh
```


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
+-- docker
```

## ❓ How to
- Build docker image with cutom tag: `docker build -t go-boilerplate .`
- Build docker image with custom directory: `docker build -f ./docker/Dockerfile .`
- Remove **none:none** image: `docker rmi $(docker images -f "dangling=true" -q)`
- Run docker: `docker run -p 8000:8000 go-boilerplate`
- Sync local branch with remote: `git fetch -p`
- Ignoring files that are already tracked: `git update-index —assume-unchanged <file>`
- To get undo/show dir's/files that are set to assume-unchanged run this: `git update-index --no-assume-unchanged <file>`
- To get a list of dir's/files that are assume-unchanged run this: `git ls-files -v|grep '^h'`
- Auto reload using nodemon: `nodemon --exec go run main.go --signal SIGTERM`


## 💡 Reference
- [Clean architechture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

![Clean Architechture Image](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)


## 👮 License

[MIT](LICENSE)
