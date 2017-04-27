package algo

import (

)

func LinearSearch(array []int,key int) (result bool){
	for i:= range array {
		if array[i]==key{
			return true
		}
	}
	return false
}