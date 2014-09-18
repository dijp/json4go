// Copyright 2014 The "json4go" Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/dijp/json4go"
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
	str := `{"function":200,"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	j := json4go.Parse2Json(str)
	fmt.Println(j.GetInt("function"))
	fmt.Println(j.GetJsonArray("servers").GetJson(0).GetString("serverName"))
	fmt.Println(j.GetJsonArray("servers").GetJson(1).ToString())
}

//developer:DaiJingPei
//email:dijp@qq.com
//url:https://github.com/dijp/json4go
func main() {
	test1()
	test2()
}
