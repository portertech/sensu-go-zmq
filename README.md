sensu-go
========

Playing with #golang

```
$ go get -tags zmq_2_1 github.com/alecthomas/gozmq
$ go run client.go
{
  "servers": [
    {
      "host": "127.0.0.1",
      "port": 5000
    }
  ]
}

Config: {[{127.0.0.1 5000}]}
Connecting to server: tcp://127.0.0.1:5000
Command output: foo
```
