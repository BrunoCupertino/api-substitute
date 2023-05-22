# api-substitute
CLI for easily provide web API locally, it is usefull when you need mock or test an API

## build
```shell
make build
```

## usage
```shell
api-substitute --port 2020 --routes /hello --responses {result:\"world\"}
```