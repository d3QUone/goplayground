### Description

GoPlayground is a project to learn GO when implementing simple API for `???`

### Options

* `-port` specify port to run server, default `8080`
* ...

### Usage

Create database with parameters from `congif.json`. In my case: 
``` bash
$ createdb -h 127.0.0.1 -p 5433 -U vladimir gobase
```

```bash
$ go build goplayground
$ ./goplayground
```
