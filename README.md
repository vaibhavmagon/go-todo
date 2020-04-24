# Golang To-Do Application

This repo is mainly for the purpose of starting to code in Go for production. It shows how folder should be structured for a production level app in a proper MVC format. It has 2 models:
1. User
2. TodoList
(Built in such a manner that it can be extended across others as well)

To run the application (on local without docker):
1. Go works with a specific folder structure that can be studied while installing go! Ideally install Go and pull this repo in the src folder.
2. Then run - go run main.go (to run)

It has an .env file to take in the values based on environment. One can create multiple .env files like .env.staging etc. for different envinroments.

The Dockerfile is also provided to make the app dockerised and run using docker-compose file (as it helps setup a mongo instance as well).
Commands:
1. docker build -t go-todo . (creates image)
2. docker-compose up (this run's the application without detach mode)
