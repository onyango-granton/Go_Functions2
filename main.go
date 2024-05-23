package main

import "fmt"

// import "fmt"

func main(){
	s1 := "A God morning is a Good Morning"
	s2 := "Good"

	intArr := []int{}

	// res := ""

	
	for i:=0; i < len(s1);{
		for j:=0; j<len(s2);{
			if s2[j] == s1[i]{
				intArr = append(intArr, i)

				//indexes
				currIndex := intArr[len(intArr) - 1]
				prevIndex := 0
				if len(intArr) < 2{
					prevIndex = i-1
				} else {
					prevIndex = intArr[len(intArr) - 2]
				}

				indexDiff := currIndex - prevIndex

				fmt.Println(currIndex, indexDiff, string(s2[j]))
			
				j++
				
				if indexDiff > 1{
					j = 0
					intArr = []int{}
				}

				
			}

			i++


			//break statement
			if i == len(s1){
				break
			}
		}
	}

	fmt.Println(intArr)

	
}




// if indexDiff == 1{
// 	// 	// fmt.Println(currIndex, prevIndex, string(s2[j]))
// 	// }
// 	// if indexDiff > 1{
// 	// 	intArr = []int{}
// 	// 	j = 0
// 	// }