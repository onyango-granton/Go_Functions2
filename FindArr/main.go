package main

import "fmt"

func main() {
	mainString := "Hello I like mellow"
	subString := "el"


	// Avoiding autoincrement //
	// for _,ch := range mainString{
	// 	for _,sch := range subString{
	// 		if 
	// 	}
	// }

	indexArr := []int{}

	for i:=0; i<len(mainString);{
		for j:=0; j<len(subString);{
			if mainString[i] == subString[j]{
				indexArr = append(indexArr, i)
				j++
			}
			i++

			if i == len(mainString){
				break
			}
			
		}
		// i++
	}

	for i:=0; i < len(indexArr);i++{
		fmt.Print(string(mainString[indexArr[i]]))
	}

	fmt.Println(indexArr)
}