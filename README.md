### json-pretty-print (or jpp)

"jpp" is an abbreviation from "json-pretty-print".\
The purpose is to make a high speed/throughput solution with good memory usage
written in golang for pretty printing .json files.\
Developed as a replacement for `python -mjson.tool`


### How to build

```
go build -o bin/jpp jpp.go
```

### How to use

```
echo "<your json here>" | ./bin/jpp

ex.:
echo "{}" | ./bin/jpp
```

#### Quick shortcut for dev
```
go build -ldflags="-s -w" -o bin/jpp jpp.go && echo "{}" | ./bin/jpp
go build -ldflags="-s -w" -o bin/jpp jpp.go && echo "{abc:asdf,"a":{a:\"b{{\\\"{\",a:[\"[:[[\"],a:b}}" | ./bin/jpp
```



