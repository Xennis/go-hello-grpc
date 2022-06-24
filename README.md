# go-hello-grpc

## Local development

```shell
make setup generate
make run-helloworld
make run-restgateway
curl localhost:6001/echo -d '{"message": "Hello World"}'
```
