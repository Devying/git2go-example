package main

import "fmt"


func main() {
	a:= make(map[interface{}]interface{})
	a["name"]="huangby"
	a["age"]=28
	a["addr"]="Beijing Chaina"
	fmt.Println(a)
}