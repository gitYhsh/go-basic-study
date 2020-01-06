package main

import (
	"fmt"
)

func main() {

	var mapdemo = make(map[string]int, 0)

	mapdemo["sex"] = 0

	mapdemo["name"] = 2

	var map3 = make(map[string]interface{}, 0)

	map3["map1"] = mapdemo

	map3["map3"] = "呵呵呵"
	map3["map2"] = 123

	for _, v := range map3 {

		if v, ok := v.(int); ok {
			fmt.Println("数字int64", v, ok)
		}
		if v, ok := v.(map[string]int); ok {
			fmt.Println("map集合", v, ok)
		}
		if v, ok := v.(string); ok1 {
			fmt.Println("string:", v, ok)
		}
	}

}
