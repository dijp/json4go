// Copyright 2014 The "json4go" Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json4go

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Json struct {
	m map[string]interface{}
}

func (this *Json) GetJson(key string) *Json {
	p := this.m[key].(*Json)
	return p
}

func (this *Json) GetMap() map[string]interface{} {
	return this.m
}

func (this *Json) GetJsonArray(key string) *JsonArray {
	p := this.m[key].(*JsonArray)
	return p
}

func (this *Json) GetString(key string) string {
	v := this.m[key]
	var r string
	switch v.(type) {
	case string:
		r = v.(string)
	case float64:
		r = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case bool:
		r = strconv.FormatBool(v.(bool))
	case *Json:
		r = v.(*Json).ToString()
	case *JsonArray:
		r = v.(*JsonArray).ToString()
	}
	return r
}

func (this *Json) GetFloat(key string) float64 {
	v := this.m[key]
	var r float64
	switch v.(type) {
	case string:
		r, _ = strconv.ParseFloat(v.(string), 64)
	case float64:
		r = v.(float64)
	}
	return r
}

func (this *Json) GetInt(key string) int64 {
	r := this.GetFloat(key)
	return int64(r)
}

func (this *Json) GetBool(key string) bool {
	v := this.m[key]
	var r bool
	switch v.(type) {
	case bool:
		r = v.(bool)
	case float64:
		if v.(float64) == 1 {
			r = true
		} else {
			r = false
		}
	}
	return r
}

func (this *Json) GetInterface(key string) interface{} {
	p := this.m[key]
	return p
}

func (this *Json) ToString() string {
	sb := make([]byte, 0)
	sb = append(sb, '{')
	notfirst := false
	for k, v := range this.m {
		if notfirst {
			sb = append(sb, ',')
		} else {
			notfirst = true
		}
		sb = append(sb, '"')
		sb = append(sb, k...)
		sb = append(sb, '"')
		sb = append(sb, ':')
		switch v.(type) {
		case *Json:
			sb = append(sb, (v.(*Json)).ToString()...)
		case string:
			sb = append(sb, '"')
			sb = append(sb, v.(string)...)
			sb = append(sb, '"')
		case *JsonArray:
			sb = append(sb, (v.(*JsonArray)).ToString()...)
		case float64:
			sb = append(sb, strconv.FormatFloat(v.(float64), 'f', -1, 64)...)
		case bool:
			sb = append(sb, strconv.FormatBool(v.(bool))...)
		}
	}

	sb = append(sb, '}')
	return string(sb)
}

func (this *Json) Put(key string, value interface{}) {
	this.m[key] = value
}

func (this *Json) Del(key string) {
	delete(this.m, key)
}

func (this *Json) GetLength() int {
	return len(this.m)
}

func NewJson() *Json {
	j := new(Json)
	j.m = make(map[string]interface{})
	return j
}

type JsonArray struct {
	arr []interface{}
}

func (this *JsonArray) GetJson(index int) *Json {
	p := this.arr[index].(*Json)
	return p
}

func (this *JsonArray) GetArray() []interface{} {
	return this.arr
}

func (this *JsonArray) GetJsonArray(index int) *JsonArray {
	p := this.arr[index].(*JsonArray)
	return p
}

func (this *JsonArray) GetString(index int) string {
	v := this.arr[index]
	var r string
	switch v.(type) {
	case string:
		r = v.(string)
	case float64:
		r = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case bool:
		r = strconv.FormatBool(v.(bool))
	case *Json:
		r = v.(*Json).ToString()
	case *JsonArray:
		r = v.(*JsonArray).ToString()
	}
	return r
}

func (this *JsonArray) GetFloat(index int) float64 {
	v := this.arr[index]
	var r float64
	switch v.(type) {
	case string:
		r, _ = strconv.ParseFloat(v.(string), 64)
	case float64:
		r = v.(float64)
	}
	return r
}

func (this *JsonArray) GetInt(index int) int64 {
	r := this.GetFloat(index)
	return int64(r)
}

func (this *JsonArray) GetBool(index int) bool {
	v := this.arr[index]
	var r bool
	switch v.(type) {
	case bool:
		r = v.(bool)
	case float64:
		if v.(float64) == 1 {
			r = true
		} else {
			r = false
		}
	}
	return r
}

func (this *JsonArray) GetInterface(index int) interface{} {
	p := this.arr[index]
	return p
}

func (this *JsonArray) ToString() string {
	sb := make([]byte, 0)
	sb = append(sb, '[')
	notfirst := false
	for _, v := range this.arr {
		if notfirst {
			sb = append(sb, ',')
		} else {
			notfirst = true
		}
		switch v.(type) {
		case *Json:
			sb = append(sb, (v.(*Json)).ToString()...)
		case string:
			sb = append(sb, '"')
			sb = append(sb, v.(string)...)
			sb = append(sb, '"')
		case *JsonArray:
			sb = append(sb, (v.(*JsonArray)).ToString()...)
		case float64:
			sb = append(sb, strconv.FormatFloat(v.(float64), 'f', -1, 64)...)
		case bool:
			sb = append(sb, strconv.FormatBool(v.(bool))...)
		}
	}
	sb = append(sb, ']')
	return string(sb)
}

func (this *JsonArray) Add(value interface{}) {
	this.arr = append(this.arr, value)
}

func (this *JsonArray) Del(index int) {
	this.arr = append(this.arr[:index], this.arr[index+1:]...)
}

func (this *JsonArray) GetLength() int {
	return len(this.arr)
}

func NewJsonArray() *JsonArray {
	j := new(JsonArray)
	j.arr = make([]interface{}, 0)
	return j
}

//developer:DaiJingPei
//email:dijp@qq.com
//url:https://github.com/dijp/json4go
func Parse2Json(str string) *Json {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	scan := newScanner(str)

	s := scan.nextToken()

	if s[0] == '{' {
		return createJson(scan)
	}
	return nil
}

func Parse2JsonArray(str string) *JsonArray {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	scan := newScanner(str)

	s := scan.nextToken()

	if s[0] == '[' {
		return createJsonArray(scan)
	}
	return nil
}

func ToJson(o interface{}) *Json {
	body, err := json.Marshal(o)
	if err != nil {
		panic(err.Error())
	}
	return Parse2Json(string(body))
}

func createJson(scan *Scanner) *Json {
	j := NewJson()
	var key string
	var value interface{}
for1:
	for {
		s := scan.nextToken()
		if s[0] == '"' {
			key = s[1 : len(s)-1]
		} else if s[0] == '}' {
			return j
		} else {
			key = s
		}
		s = scan.nextToken()
		if s[0] == ':' {
			s = scan.nextToken()
			switch s[0] {
			case '"':
				value = s[1 : len(s)-1]
				j.Put(key, value)
			case '{':
				value = createJson(scan)
				j.Put(key, value)
			case '[':
				value = createJsonArray(scan)
				j.Put(key, value)
			case 't':
				if s == "true" {
					value = true
					j.Put(key, value)
				}
			case 'f':
				if s == "false" {
					value = false
					j.Put(key, value)
				}
			case 'n':
				if s == "null" {
				}
			default:
				value, _ = strconv.ParseFloat(s, 64)
				j.Put(key, value)
			}
			s = scan.nextToken()
			switch s[0] {
			case '}':
				return j
			case ',':
				continue for1
			default:
				return nil
			}
		} else {
			return nil
		}
	}
	return nil
}

func createJsonArray(scan *Scanner) *JsonArray {
	a := NewJsonArray()
	var value interface{}
for1:
	for {
		s := scan.nextToken()
		switch s[0] {
		case '"':
			value = s[1 : len(s)-1]
			a.Add(value)
		case '{':
			value = createJson(scan)
			a.Add(value)
		case '[':
			value = createJsonArray(scan)
			a.Add(value)
		case 't':
			if s == "true" {
				value = true
				a.Add(value)
			}
		case 'f':
			if s == "false" {
				value = false
				a.Add(value)
			}
		case 'n':
			if s == "null" {
			}
		case ']':
			return a
		default:
			value, _ = strconv.ParseFloat(s, 64)
			a.Add(value)
		}
		s = scan.nextToken()
		switch s[0] {
		case ',':
			continue for1
		case ']':
			return a
		default:
			return nil
		}
	}
	return nil
}

type Scanner struct {
	r string
	i int
}

func newScanner(str string) *Scanner {

	return &Scanner{str, 0}
}

func (this *Scanner) nextToken() string {
	switch this.r[this.i] {
	case ' ', '\f', '\n', '\r', '\t', '\v':
		this.i++
	for2:
		for {
			switch this.r[this.i] {
			case ' ', '\f', '\n', '\r', '\t', '\v':
				this.i++
			default:
				break for2
			}
		}
		return this.nextToken()
	case '{', '}', '[', ']', ':', ',':
		start := this.i
		this.i += 1
		return this.r[start:this.i]
	case '"':
		start := this.i
		this.i += 1
		for {
			if this.r[this.i] == '"' {
				break
			} else {
				this.i += 1
			}
		}
		this.i += 1
		return this.r[start:this.i]
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		start := this.i
		this.i += 1
		for {
			e := this.r[this.i]
			if e != '.' && (e < '0' || e > '9') {
				break
			} else {
				this.i += 1
			}
		}
		return this.r[start:this.i]

	default:
		start := this.i
		this.i++
	for3:
		for {
			switch this.r[this.i] {
			case ' ', '\f', '\n', '\r', '\t', '\v', '{', '}', '[', ']', ':', ',':
				break for3
			default:
				this.i++
			}
		}
		return this.r[start:this.i]
	}

}
