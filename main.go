package main


func initializeCard () string {
	return "functionCalled"
}

func main(){
	
	cards:=deck{"Ace of clubs"}

	
	cards=append(cards,"Jack of diamonds")
	cards=append(cards,"Jack of spades")

	cards.print(5)
}