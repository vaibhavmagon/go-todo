package main

import (
    "fmt"
    "log"
    "net/http"
    "go-todo/server/router"
    "go-todo/server/utils"
)

func main() {
    r := router.Router()
    Env := utils.ViperEnvVariable("ENVIRONMENT")
    fmt.Println("Starting " + Env + " server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

