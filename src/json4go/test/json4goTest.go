package main

import (
	"fmt"
	"json4go"
)

type Person struct {
	Name string
	Age  int
	Sex  bool
}

type User struct {
	UserId int
	P      Person
	Map    map[string]int
	Arr    []string
}

func test1() {
	u := User{1, Person{"dog", 30, true}, map[string]int{"a": 1, "b": 2, "c": 3}, []string{"aaa", "bbb", "ccc"}}
	j := json4go.ToJson(u)
	fmt.Println(j.GetJson("P").GetString("Name"))
}

func test2() {
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	j := json4go.Parse2Json(str)
	fmt.Println(j.GetJsonArray("servers").GetJson(0).GetString("serverName"))
	fmt.Println(j.GetJsonArray("servers").GetJson(1).ToString())
}

func main() {
	test1()
	test2()
}
