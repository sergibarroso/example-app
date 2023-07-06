# Example App

## Locally

### Build

```shell
go mod init example-app
go mod tidy
```

```shell
go build
```

### Test

```shell
go test -v -bench=.
```

### Run

```shell
go run main.go
```

## Docker

### Build

```shell
docker build -t example-app .
```

### Run

```shell
docker run -it --rm -p 8080:8080 --name example-app example-app
```

## Helm

### Render template

```shell
helm template example-app charts
```

### Install

```shell
helm install example-app charts --namespace default
```

### Test

```shell
helm test example-app
```
