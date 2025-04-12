package main
import "fmt"

func split(a int)(x,y int){
	x = a/10
	y = a%10

	return
}

func main(){
	fmt.Println(split(13))
}