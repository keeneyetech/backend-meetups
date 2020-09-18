# Say Hello

## How to use

1. Start a local redis. You can use docker for this purpose.

```
$ docker run --name my-redis-container -p 6379:6379 -d redis
```

2. Build the project

```
$ go install ./...
```

3. Run the project

```
$ ~/go/sayhello
```

4. Target the server from anywhere with an HTTP request. Locally you can use the
following command

```
$ curl --header "X-Forwarded-For: 81.24.35.12" localhost:8080
```

Then connect to your redis to see what was saved.

```
docker exec -it $(docker ps -qf name=my-redis-container) redis-cli

127.0.0.1> keys *

```