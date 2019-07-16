# cmd-exec-server

## It does...

cmd-exec-server executes the command specified by app-settings.json, with STDIN from HTTP request body, and returns STDOUT in HTTP response body.

## It does not...

- authenticate request
- limit rate

## future work

see GitHub project : https://github.com/hrkt/cmd-exec-server/projects

# How to run

0. edit app-settings.json

see: "app-settings.json" paragraph in this README. 

1. execute server.

```
$ make run
```

2. make HTTP POST request

```
$ curl -X POST localhost:8080/api/exec -d "some input"
{"result":"SOME INPUT"}
```


# prerequisites

- dep as dependency manager
- linux, *nix like platforms - ("endless" shows error message  on windows platform, at this point of moment)

# app-settings.json

specify "command" and "arguments"

```
{
    "command": "tr",
    "arguments": [
        "a-z",
        "A-Z"
    ]
}
```


# usage

## run in dev mode

```
$ make run
```

## run in release mode

```
$ make release-run
```

## build

```
$ make run
```

## test

```
$ make test
```

## format

```
$ make fmt
```

## build container

```
$ make build-container
```

## run container

```
$ make run-container
```

# License
MIT

# CI

[![CircleCI](https://circleci.com/gh/hrkt/cmd-exec-server.svg?style=svg)](https://circleci.com/gh/hrkt/cmd-exec-server)