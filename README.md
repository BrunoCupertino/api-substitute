# api-substitute
CLI for easily provide web API locally, it is usefull when you need mock or test an API and check the request made.

## build
```shell
make build
```

## usage
```shell
api-substitute --port 2020 --routes /hello --responses {result:\"world\"}
```

after running api-substitute will be able to call the routes provided with their respectives responses, using the example above:
```shell
curl http://localhost:2020/hello
```
it will return:
```json
{message: "world"}
```

besides the api-subsitute will print all informations about the request, including the raw body it is usefull to check the request.