# Hostgator challenge

## Stack of tech

Here is a list of stack of tech used in this project

* Golang version 1.14.1 linux/amd64
* Docker CE
* Sqlite3
* Jq
* httpie
* golangci-lint

## Dependencies of project

You need install some libs for running this project.

### Httpie for curl sintact sugar

Linux based on **Debian**, you need first install **httpie**, so you can use `sudo apt install httpie`, and you need a **plugin** for **jwt** request, so you can install with this command `pip3 install -U httpie-jwt-auth`, in my case I'm using python3.

### Jq (JQuery for command line)

On Debin linux, you can install with `sudo apt install jq`

### Swagger API doc

`go get -u github.com/swaggo/swag/cmd/swag`

for update the documentation just run command `cd api/cmd` and after run the swag for find and write a swagger documentation. `swag init -g main.go`

Endpoint for documentation: `http://localhost:8080/swagger/index.html`

## Test project

In command line inside folder of project, run `go test ./... -v` or makefile `make test`

## Check lint of code

You need the `golangci-lint`, for the installation I recommend you see the documentation about that. Link for doc [Here](https://github.com/golangci/golangci-lint)

## Start project

Run `go run api/cmd/main.go` in command line.
Or you can make the binary files using `makefile build`, this command you going generate all binaries(***raw and zip***) files for **linux**, **macOs** and **raspBery Pie** in folder `out`.

## Generate docker image

For create a image with **Docker-CE** just running in folder of projet the command `docker build -t thiagozs/challenge .`

Runing the image after build `sudo docker run --rm --name=challenge --publish=8080:8080 thiagozs/challenge:latest`

## Docker Healthcheck

On the construction of image we have a change to put a little **healtcheck** on API.

```sh
Step 17/17 : HEALTHCHECK --interval=5s --timeout=2s --start-period=2s --retries=5 CMD [ "curl", "--silent", "--fail", "http://localhost:8080/ping" ]
```

## Testing API in easy mode

After started the api, you can acess the endpoint for create/show a **account**, this method you going generate a login with password. At this time you don't need login with a **JWT** token, to create/show. To create your credential follow the staps below.

### Create registry (@usernamea and password)

In your terminal: `http post http://localhost:8080/account username=thiagozs password=123`

```sh
HTTP/1.1 201 Created
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Length
Access-Control-Max-Age: 86400
Content-Length: 267
Content-Type: application/json; charset=utf-8
Date: Wed, 13 May 2020 06:41:51 GMT

{
    "data": {
        "CreatedAt": "2020-05-13T03:41:51.230253909-03:00",
        "DeletedAt": null,
        "ID": 1,
        "Password": "$argon2id$v=19$m=65536,t=3,p=2$JWpPZfr1lNXQdEXndEt0yQ$/vwTThbrnNulO/61v1lLSeQCuH2qsLecaLS+aW8C5e0",
        "UpdatedAt": "2020-05-13T03:41:51.230253909-03:00",
        "UserName": "thiagozs"
    }
}
```

### Show a registry inside DB

In your terminal: `http get http://localhost:8080/account/1`

```sh
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Length
Access-Control-Max-Age: 86400
Content-Length: 258
Content-Type: application/json; charset=utf-8
Date: Wed, 13 May 2020 06:42:32 GMT

{
    "CreatedAt": "2020-05-13T03:41:51.230253909-03:00",
    "DeletedAt": null,
    "ID": 1,
    "Password": "$argon2id$v=19$m=65536,t=3,p=2$JWpPZfr1lNXQdEXndEt0yQ$/vwTThbrnNulO/61v1lLSeQCuH2qsLecaLS+aW8C5e0",
    "UpdatedAt": "2020-05-13T03:41:51.230253909-03:00",
    "UserName": "thiagozs"
}

```

### Login with your @username and get JWT token

In you terminal: `http post http://localhost:8080/login username=thiagozs password=123 | jq -r .token > mytoken.txt`

```sh
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Length
Access-Control-Max-Age: 86400
Content-Length: 219
Content-Type: application/json; charset=utf-8
Date: Wed, 13 May 2020 06:46:13 GMT

{
    "code": 200,
    "expire": "2020-05-13T04:46:13-03:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODkzNTU5NzMsIm9yaWdfaWF0IjoxNTg5MzUyMzczLCJ1dWlkIjoidGhpYWdvenMifQ.Dq2yKkvouUK-smsq9acwihZs5vHYSQ1Nk39U_HY1hug"
}
```

### Cosume the endpoint with JWT token

In your terminal: `http --auth-type=jwt --auth="$(cat mytoken.txt)" get http://localhost:8080/breeds/ango`

```sh
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Length
Access-Control-Max-Age: 86400
Content-Length: 1177
Content-Type: application/json; charset=utf-8
Date: Wed, 13 May 2020 06:47:33 GMT

{
    "data": [
        {
            "CreatedAt": "2020-05-13T03:47:33.81681774-03:00",
            "DeletedAt": null,
            "UpdatedAt": "2020-05-13T03:47:33.81681774-03:00",
            "adaptability": 5,
            "affection_level": 5,
            "alt_names": "Ankara",
            "child_friendly": 4,
            "country_code": "TR",
            "country_codes": "TR",
            "description": "This is a smart and intelligent cat which bonds well with humans. With its affectionate and playful personality the Angora is a top choice for families. The Angora gets along great with other pets in the home, but it will make clear who is in charge, and who the house belongs to",
            "dog_friendly": 5,
            "energy_level": 5,
            "experimental": 0,
            "grooming": 2,
            "hairless": 0,
            "health_issues": 2,
            "hypoallergenic": 0,
            "id": "tang",
            "indoor": 0,
            "intelligence": 5,
            "life_span": "15 - 18",
            "name": "Turkish Angora",
            "natural": 1,
            "origin": "Turkey",
            "rare": 0,
            "rex": 0,
            "shedding_level": 2,
            "short_legs": 0,
            "social_needs": 5,
            "stranger_friendly": 5,
            "suppressed_tail": 0,
            "temperament": "Affectionate, Agile, Clever, Gentle, Intelligent, Playful, Social",
            "vcahospitals_url": "https://vcahospitals.com/know-your-pet/cat-breeds/turkish-angora",
            "vocalisation": 3,
            "weight": {
                "imperial": "5 - 10",
                "metric": "2 - 5"
            },
            "wikipedia_url": "https://en.wikipedia.org/wiki/Turkish_Angora"
        }
    ]
}
```

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the [LICENSE](https://github.com/thiagozs/hostgator-challenge/blob/master/LICENCE) file

**2020, thiagozs**
