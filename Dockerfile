# Use latest golang version
FROM golang

MAINTAINER Vaibhav Magon

# create app directory in container
RUN mkdir -p $GOPATH/src/go-todo

# set /go-todo directory as default working directory
WORKDIR $GOPATH/src/go-todo

# --pure-lockfile
RUN go get -u github.com/gorilla/mux && go get -u github.com/go-martini/martini && go get -u go.mongodb.org/mongo-driver/mongo && go get -u github.com/spf13/viper

# copy all file from current dir
COPY . .

# expose port 8080
EXPOSE 8080

# cmd to start service
CMD [ "go", "run", "main.go" ]