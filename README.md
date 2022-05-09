# user_authentication

# Installation
## Database & Migrations
There's a docker compose to create a db and migrate all the migrations, just do the following and you'll be good to go.

```docker-compose up -d```

## Running The Project
Run the following command in the root folder of the application.

```go get ```

And then:

```go run main.go```

And to run the tests all together:

`go test ./...`


# Docs
## List Users
### List Users with filters
The pagination starts at 0

`GET` `localhost:8080/user/?page=0&cancelled=false`

### List All Users

`GET` `localhost:8080/user/`

## Create Users
The signup method and the create method do pretty much the same thing because there is no validation on who gives what data so... :D

### Create 
`POST` `localhost:8080/user/`
Request body should be like:
```
{
    "first_name": "name",
    "last_name": "last name",
    "national_id": "some string and numbers"
}
```

### Signup
`POST` `localhost:8080/user/signup`
Request body should be like:
```
{
    "first_name": "name",
    "last_name": "last name",
    "national_id": "some string and numbers"
}
```

## Cancel Signup

`POST` `localhost:8080/user/signup/cancel/{{{some-user-uuid}}}`

## Delete User

`DELETE` `localhost:8080/user/{{{some-user-uuid}}}`

