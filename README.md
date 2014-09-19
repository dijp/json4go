json4go
=======

A JSON parser for Go. keywords json,parser,go,golang,json4go.

## Getting Started

~~~ go
package main

import (
  "github.com/dijp/json4go"
  "fmt"
  )

func main() {
	str := `{"function":200,"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	j := json4go.Parse2Json(str)
	fmt.Println(j.GetInt("function"))//200
	fmt.Println(j.GetJsonArray("servers").GetJson(0).GetString("serverName"))//Shanghai_VPN
	fmt.Println(j.GetJsonArray("servers").GetJson(1).ToString())//{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}
}
~~~

## Installation

~~~
go get github.com/dijp/json4go
~~~

### Documentation

Visit the docs on [gopkgdoc](http://godoc.org/github.com/dijp/json4go)
