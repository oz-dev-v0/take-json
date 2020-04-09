# Take JSON

pretty-print json of linux pipes or file.

## install

```
$ go install
$ ln -s $GOPATH/bin/take-json /usr/local/bin/
$ take-json -h
```

## example

```
$ take-json -h
pretty-print json of linux pipes or file

Usage:
  take-json [flags]

Flags:
  -u, --decodeURI       print, decodeURI.
  -p, --fallbackPrint   fallback, just input text print.
  -f, --file string     path to the file.
  -h, --help            help for take-json
```

### dummy.log

```sh
$ cat ./dummy/dummy.log 
start
2020-04-10 02:12:00 +0000 system.access: { "key-string": "s", "key-int": 1, "key-bool": true, "key-float": 1.1 }
2020-04-10 02:13:00 +0000 system.access: { "key-string": "ss", "key-int": 2, "key-bool": true, "key-float": 2.2 }
2020-04-10 02:14:00 +0000 system.access: { "key-string": "ss", "key-int": 2, "key-bool": true, "key-float": 2.2 } hogehoge
end
```

```sh
$ tail -f dummy/dummy.log | take-json
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
{
        "key-string": "ss",
        "key-int": 2,
        "key-bool": true,
        "key-float": 2.2
}
{
        "key-string": "ss",
        "key-int": 2,
        "key-bool": true,
        "key-float": 2.2
}
```

```sh
$ tail -f dummy/dummy.log | take-json -p
start
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
{
        "key-string": "ss",
        "key-int": 2,
        "key-bool": true,
        "key-float": 2.2
}
{
        "key-string": "ss",
        "key-int": 2,
        "key-bool": true,
        "key-float": 2.2
}
end
```

### dummy-encodeuri.log

```sh
$ cat dummy/dummy-encodeuri.log 
start
2020-04-10 02:12:00 +0000 system.access: /log?r=%7B%20%22key-string%22:%20%22s%22,%20%22key-int%22:%201,%20%22key-bool%22:%20true,%20%22key-float%22:%201.1%20%7D
2020-04-10 02:13:00 +0000 system.access: /log?r=%7B%20%22key-string%22:%20%22s%22,%20%22key-int%22:%201,%20%22key-bool%22:%20true,%20%22key-float%22:%201.1%20%7D
2020-04-10 02:14:00 +0000 system.access: /log?r=%7B%20%22key-string%22:%20%22s%22,%20%22key-int%22:%201,%20%22key-bool%22:%20true,%20%22key-float%22:%201.1%20%7D hogehoge
end
```

```sh
$ tail -f dummy/dummy-encodeuri.log | take-json -p
start
2020-04-10 02:12:00 +0000 system.access: /log?r=%7B%20%22key-string%22:%20%22s%22,%20%22key-int%22:%201,%20%22key-bool%22:%20true,%20%22key-float%22:%201.1%20%7D
2020-04-10 02:13:00 +0000 system.access: /log?r=%7B%20%22key-string%22:%20%22s%22,%20%22key-int%22:%201,%20%22key-bool%22:%20true,%20%22key-float%22:%201.1%20%7D
2020-04-10 02:14:00 +0000 system.access: /log?r=%7B%20%22key-string%22:%20%22s%22,%20%22key-int%22:%201,%20%22key-bool%22:%20true,%20%22key-float%22:%201.1%20%7D hogehoge
end
```

```sh
$ tail -f dummy/dummy-encodeuri.log | take-json

```

```sh
$ tail -f dummy/dummy-encodeuri.log | take-json -u
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
```

```sh
$ tail -f dummy/dummy-encodeuri.log | take-json -p -u
start
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
{
        "key-string": "s",
        "key-int": 1,
        "key-bool": true,
        "key-float": 1.1
}
end
```

## docker

```sh
$ docker-compose up -d
$ docker-compose exec golang bash
# or docker-compose exec golang /bin/bash
```

## Todo

- debug mode
- add test code
