# Concurrency Management using Golang

A simple web server that fetches data from an API and saves it in a database so that the next time it's requested it can be served without querying the external API.

- Built in Go version 1.19
- Uses [gorilla/mux](https://github.com/gorilla/mux) for routing.
- Uses [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver).

Tested the API using [hey](https://github.com/rakyll/hey). Here's a screenshot:
![hey load test](https://user-images.githubusercontent.com/89561537/226923899-d5c9a91f-a72e-43d2-8505-84c2c507ca7a.png)

The output for the stated example test:
![output](https://user-images.githubusercontent.com/89561537/226924193-b2a89806-c7f6-4506-9ee8-bacf2d36ef1b.png)
