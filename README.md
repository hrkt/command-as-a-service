# cmd-exec-server

## It does...

command-as-a-service executes the command with options specified in the url, with STDIN from HTTP request body, and returns STDOUT in HTTP response body.

## It does not...

- authenticate request
- limit rate

## future work

see GitHub project : https://github.com/hrkt/command-as-a-service/projects

# How to run

0. edit app-settings.json

see: "app-settings.json" paragraph in this README. 

1. execute server.

```
$ make run
```

2. make HTTP POST request

```
$ curl "http://localhost:8080/bin/date"
2019年 7月27日 土曜日 15時27分43秒 JST
```


Please point the command you want to execcute with url-path same as filesystem-path.

i.e.
```
URL                                Executed
http://localhost:8080/bin/date ==> /bin/date
```

Options for the command are given by query parameter. 

i.e.
```
URL                                Executed
http://localhost:8080/usr/bin/sort?-rn ==> /bin/sort -rn
```

Standard input for the command can be passed throw Request body.

i.e.
```
URL
$ curl "http://localhost:8080/usr/bin/sort?-r&-n" --data-binary @testdata/test.txt 
```



# prerequisites

- dep as dependency manager
- linux, *nix like platforms - ("endless" shows error message  on windows platform, at this point of moment)

# app-settings.json

specify "whitelist", "port".
enable "dangerousMode"(that does not use whitelist. In other worrds, you can execute 'rm -rf ' via http request), if it is needed.

```
{
    "dangerousMode": false,
    "port": 8080,
    "whitelist": [
        "/bin/date",
        "/bin/echo",
        "/bin/hostname",
        "/bin/sleep",
        :
        :
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

[![CircleCI](https://circleci.com/gh/hrkt/command-as-a-service.svg?style=svg)](https://circleci.com/gh/hrkt/command-as-a-service)