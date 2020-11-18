# Golang To Do Application

<img src="https://img.shields.io/static/v1?label=Go&message=1.3&color=<COLOR>"> <img src="https://img.shields.io/static/v1?label=Mongo&message=3.4&color=<COLOR>"> <img src="https://img.shields.io/static/v1?label=Build&message=Passing&color=<COLOR>">

This repo is mainly for the purpose of starting to code in Go for production. It shows how folder should be structured for a production level app in a proper MVC format. It has 2 models: <br/>
- User
- TodoList
(Built in such a manner that it can be extended across others as well)


## To run the application (on local without docker):
- Go works with a specific folder structure that can be studied while installing go! Ideally install Go and pull this repo in the src folder.
- To run:
```console
go run main.go
```

It has an .env file to take in the values based on environment. One can create multiple .env files like .env.staging etc. for different envinroments.
The Dockerfile is also provided to make the app dockerised and run using docker-compose file (as it helps setup a mongo instance as well).


## Commands:
```console
docker build -t go-todo . (creates image)
docker-compose up (this run's the application without detach mode)
```


## Maintainer
- Vaibhav Magon
