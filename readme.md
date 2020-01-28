# Create Crud and login with Golang (Rest API)

## Installation

- Clone this repo, you can put in or outside your `$GOPATH` because we use [go modules][gomod]
- Create empty database
- Configure the `.env`
- Run this command : `go run main.go`

[gomod]: https://blog.golang.org/using-go-modules

## How this app working ?
When first time this app running it will create dummy data to your database base on you config on `.env` file.
after that you can check the data to your database or you can check it with postman, Insomnia, HTTPie etc.
Check the list of route on `api/controllers/routes.go` to test all endpoint on this app.
just want to remind to test the CRUD, according to the action `(POST, GET, PUT, DELETE)` you want to do.
**Don't forget to login first and use `Bearer Token` if you want to update and delete data.**

## Thank You