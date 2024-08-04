package main

import "fmt"

func initializeCard () string {
	return "functionCalled"
}

func main(){
	
	cards:=[]string{"1"}

	
	cards=append(cards,"2")

	fmt.Println(cards)
}