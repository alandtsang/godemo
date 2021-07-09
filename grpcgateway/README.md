# Prepare

You need to install googleapis, if it's already installed, skip this step.
```
go get github.com/googleapis/googleapis
```

# Usage

1. Generate pb files and compile:
```
$ make build
```

Now you have the gateway and grpc server program.

Run in a terminal:
```
$ ./rpcserver
listen on :50001
```

Run in another terminal:
```
$ ./gateway 
Listen on :8081
```

Now test:
```
$ curl -XPOST http://127.0.0.1:8081/v1/example/echo -d '{"name":"Alan", "age":18}'
{"greeting":"Hello Alan, your age is 18."}
```