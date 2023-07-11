# Example App

This is the example application. Developed by the example team.

Use the following development workflow to code features or fixes for this app.

- Create a feature/fix branch
- Code the new feature. You can test it locally by using the commands defined in the [local development](#local-development) section.
- Create a Merge Request to the target environment:
  - Production: main branch
  - Staging: staging branch
  - Dev: dev branch
- On every new commit, the code will be lint and tested and you could get your feedback in the Actions section of GitHub.
- When the Merge Request is reviewed and merged, the new version of the application will be deployed into the target environment cluster.

## Local Development

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

### Docker

#### Build

```shell
docker build -t example-app .
```

#### Run

```shell
docker run -it --rm -p 8080:8080 --name example-app example-app
```

### Helm

#### Render template

```shell
helm template example-app charts
```

#### Install

```shell
helm install example-app charts --namespace default
```

#### Test

```shell
helm test example-app
```