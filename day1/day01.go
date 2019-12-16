package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/bitly/go-simplejson"
)

type persion struct {
	Name string `json:"name"`
	Age  int16  `json:"-"`
}

func main() {

	p := persion{Name: "张三", Age: 15}
	data, _ := json.Marshal(&p)

	fmt.Println(string(data))

	res, _ := simplejson.NewJson([]byte(data))

	fmt.Println(res)
	num := 123
	var num2str = strconv.Itoa(num)
	fmt.Println(reflect.TypeOf(num2str))
	atr, _ := strconv.ParseInt(num2str, 10, 64)
	fmt.Println(reflect.TypeOf(atr))
}
