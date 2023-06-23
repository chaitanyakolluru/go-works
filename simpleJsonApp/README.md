# Simple Json App:

The idea behind this app is simple, it creates a new data.json file in current working directory, and starts exposing APIs to perform CRUD operations on the created json file.

## usage

Run project

```bash
$ go run main.go
```

Swagger page [here](http://localhost:8081/swagger/index.html)

Swagger home:
![Swagger page](./images/swagger1.png)

Get your JWT token:
![Swagger page](./images/swagger4.png)

and authenticate:
![Swagger page](./images/swagger5.png)

Examples: (auth token not shown in curl request)
![Swagger page](./images/swagger2.png)

![Swagger page](./images/swagger3.png)
