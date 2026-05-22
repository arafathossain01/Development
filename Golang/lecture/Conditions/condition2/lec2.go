package main
import ("fmt")

func main(){
	var temp float32

	fmt.Print("Enter the temperature: ")
	fmt.Scan(&temp)

	if temp > 30{
		fmt.Println("It is hot day")
	} else{
		fmt.Println("It is cool day")
	}

}