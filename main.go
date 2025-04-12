package main

import "fmt"

func swap (a,b int) (int,int){
	return b,a
}

func main (){
	a:=5
	b:=9

	a,b=swap(a,b)
	fmt.Println(a)
	fmt.Println(b)
}