package basic

import "fmt"

func Data(){
	sum := 0

	for i := 0; i < 10; i++{
		sum += i
	}
	
	fmt.Println(sum)
}