package main 

import (
	"linearsearch/algo"
	"fmt"
)

func main() {
	array:=[]int{2, 3, 5, 7, 11, 13}
	
	fmt.Println(algo.LinearSearch(array,7))
	
	fmt.Println(algo.LinearSearch(array,1))
	
	fmt.Println(algo.LinearSearch(array,13))
	
	
	fmt.Println(algo.LinearSearch(array,4))
}
