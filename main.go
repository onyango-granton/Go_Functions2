package main

import "fmt"

// import "fmt"

func main(){
	mainStr := "A God morning is a Good Morning"


	s1 := [] string {}
	s2 := "Good"

	resString := ""



	intArr := []int{}


	for i:=0; i<len(mainStr);i++{
		s1 = append(s1, string(mainStr[i]))
	}

	// res := ""

	
	for i:=0; i < len(s1);{
		for j:=0; j<len(s2);{
			if string(s2[j]) == s1[i]{
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

	for i:=0; i < len(intArr);i++{
		// fmt.Println("Here")
		s1[intArr[i]] = "\033[34m"+s1[intArr[i]]+"\033[0m"
	}

	for i:=0; i<len(s1);i++{
		resString += s1[i]
	}

	fmt.Println(resString)

	
}




// if indexDiff == 1{
// 	// 	// fmt.Println(currIndex, prevIndex, string(s2[j]))
// 	// }
// 	// if indexDiff > 1{
// 	// 	intArr = []int{}
// 	// 	j = 0
// 	// }