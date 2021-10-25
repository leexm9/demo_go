package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	// 解析已知结构的 json 串
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN", "serverIp":"127.0.0.1"},
{"serverName":"Beijing_VPN", "serverIp":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println()

	// 解析未知结构的 json 串
	b := []byte(`{"Name":"Wednesday", "Age":6, "Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", v)
		case int:
			fmt.Println(k, "is int", v)
		case float64:
			fmt.Println(k, "is float64", v)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	fmt.Println()

	// simpleJson
	js, err := simplejson.NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"int": 10,
			"float":5.150,
			"bignum":94537509283045027452,
			"string":"simplejson",
			"bool": true
		}
}`))
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms, _ := js.Get("test").Get("string").String()
	fmt.Println(arr)
	fmt.Println(i)
	fmt.Println(ms)
	fmt.Println()

	// 输出 json 串
	var ss Serverslice
	ss.Servers = append(ss.Servers, Server{ServerName: "Shanghai", ServerIp: "127.0.0.1"})
	ss.Servers = append(ss.Servers, Server{ServerName: "Beijing", ServerIp: "127.0.0.2"})
	b2, err := json.Marshal(ss)
	if err != nil {
		fmt.Println("json error:", err)
	}
	fmt.Println(string(b2))
}
