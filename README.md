### json-pretty-print (or jpp)

"jpp" is an abbreviation from "json-pretty-print".\
The purpose is to make a high speed/throughput solution for pretty printing json files in golang.


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

