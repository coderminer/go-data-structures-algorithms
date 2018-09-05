package main

import(
	"fmt"
	"go-data-structures-algorithms/dictionary/dictionary"
)

func main(){
	d := dictionary.Dictionary{}
	d.Set("1","first")
	d.Set("2","Second")
	d.Set("3","third")
	d.Set("4","forth")
	d.Set("5","fifth")

	fmt.Println(d.Get("2"))
	fmt.Println(d.Has("4"))
	fmt.Println(d.Has("6"))
	fmt.Println(d.Delete("5"))
	fmt.Println(d.Size())
	fmt.Println(d.Keys())
	fmt.Println(d.Values())
}