package main

import "fmt"

// import "fmt"

// func main() {
	// mainString := "Hello I like mellow"
	// subString := "el"


	// // Avoiding autoincrement //
	// // for _,ch := range mainString{
	// // 	for _,sch := range subString{
	// // 		if 
	// // 	}
	// // }

	// indexArr := []int{}

	// for i:=0; i<len(mainString);{
	// 	for j:=0; j<len(subString);{
	// 		if mainString[i] == subString[j]{
	// 			indexArr = append(indexArr, i)
	// 			j++
	// 		}
	// 		i++

	// 		if i == len(mainString){
	// 			break
	// 		}
			
	// 	}
	// 	// i++
	// }

	// for i:=0; i < len(indexArr);i++{
	// 	fmt.Print(string(mainString[indexArr[i]]))
	// }

	// fmt.Println(indexArr)
// 	fmt.Println(isDuplicate("ele", "ele"))
// }


func isDuplicate(s1, s2 string) bool {
	count := 0
	for i:=0; i < len(s1);{
		for j:=0; j < len(s2);{
			if s1[i] == s2[j]{
				count ++
				j++
				
			}
			i ++
			if i == len(s1){
				break
			}
		}
	}
	if count == len(s1) && count == len(s2){
		return true
	}
	return false
}


func findArr(mainString, subString string) []int {
	res := []int{}
	for i:=0; i<len(mainString);i++{
		if  i+len(subString) < len(mainString) && isDuplicate(subString, mainString[i:i+len(subString)]){
			for j:=i ; j < i+len(subString);j++{
				res = append(res, j)
			}
		}
	}
	return res
}


func main() {
	mainString := "Hello I like mellow"
	subString := "el"

	intArr := []int{}

	for i:=0; i<len(mainString);i++{
		if  i+len(subString) < len(mainString) && isDuplicate(subString, mainString[i:i+len(subString)]){
			for j:=i ; j < i+len(subString);j++{
				fmt.Println(j)
				intArr = append(intArr, j)
			}
			fmt.Println(mainString[i:i+len(subString)])
		}
	}
	fmt.Println(intArr)
}