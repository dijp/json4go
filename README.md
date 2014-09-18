json4go
=======

json parser go golang json4go

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file. 

~~~ go
package main

import (
  "github.com/dijp/json4go"
  "fmt"
  )

func main() {
  str := `{"function":"200","servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	j := json4go.Parse2Json(str)
	fmt.Println(j.GetString("function"))
	fmt.Println(j.GetJsonArray("servers").GetJson(0).GetString("serverName"))
	fmt.Println(j.GetJsonArray("servers").GetJson(1).ToString())
}
~~~

Then install the json4go package (**go 1.1** and greater is required):
~~~
go get github.com/dijp/json4go
~~~
