# Scanner

A simple port scanner written in Go

## Getting Started


### Installing

`go get github.com/xzased/scanner`


## Running Scans

Scan a host/port 

```
scanner scan <host> <port>
```

The command above will output the host/port state as well as any information available
in the initial packet, for example, scanning a local mysql instance returns the following:

```
go run main.go scan 0.0.0.0 3306
{"level":"debug","hostport":"0.0.0.0:3306","time":"2020-10-21T16:30:08-05:00","message":"scanning"}
{"level":"info","state":"open","host":"0.0.0.0","port":3306,"data":"[\u0000\u0000\u0000\n5.7.31-0ubuntu0.18.04.1\u0000\u001b\u0000\u0000\u0000\u0001`2k\"c2\u0005\u0000\ufffd\ufffd\b\u0002\u0000\ufffd\ufffd\u0015\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u00000a.P\u0012I<a\u0016\u0016k\u001c\u0000mysql_native_password\u0000","time":"2020-10-21T16:30:08-05:00","message":"found service"}
```

Scanning other services
```
go run main.go scan censys.io 80
{"level":"debug","hostport":"censys.io:80","time":"2020-10-21T16:35:09-05:00","message":"scanning"}
{"level":"info","state":"open","host":"censys.io","port":80,"data":"","time":"2020-10-21T16:35:11-05:00","message":"found service"}
```