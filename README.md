# mh-z19b-exporter

二酸化炭素濃度計「MH-Z19B」のPrometheus Exporter

## インストール

```shell
$ go get github.com/macaron/mh-z19b-exporter
```

## 使い方

```shell
Usage of ./mh-z19b-exporter:
  -dev string
    	MH-Z19B device (default "/dev/serial0")
  -interval int
    	The frequency in seconds in which to gather data (default 60)
  -path string
    	Path for metrics (default "/metrics")
  -port string
    	Address for this exporter run (default "8080")
```
