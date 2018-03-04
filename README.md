### json-pretty-print (or jpp)

"jpp" is an abbreviation from "json-pretty-print".\
The purpose is to make a high speed/throughput solution for pretty printing json files in golang.\
As a replacement for `python -mjson.tool`


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

#### quick shortcut for dev
```
go build -ldflags="-s -w" -o bin/jpp jpp.go && echo "{}" | ./bin/jpp
go build -ldflags="-s -w" -o bin/jpp jpp.go && echo "{abc:asdf,"a":{a:\"b{{\\\"{\",a:[\"[:[[\"],a:b}}" | ./bin/jpp
```



