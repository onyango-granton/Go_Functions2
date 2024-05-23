package main

import "fmt"

// import "fmt"

func FindSubString(mainString, subString string) {
	// mainString := "A God morning is A Good God glory Morning Good"

	s1 := []string{}
	// subString := "Good"

	resString := ""

	intArr := []int{}

	intArrArr := [][]int{}

	for i := 0; i < len(mainString); i++ {
		s1 = append(s1, string(mainString[i]))
	}

	// res := ""

	for i := 0; i < len(s1); {
		for j := 0; j < len(subString); {
			if string(subString[j]) == s1[i] {
				intArr = append(intArr, i)

				// indexes
				currIndex := intArr[len(intArr)-1]
				prevIndex := 0
				if len(intArr) < 2 {
					prevIndex = i - 1
				} else {
					prevIndex = intArr[len(intArr)-2]
				}

				indexDiff := currIndex - prevIndex

				fmt.Println(currIndex, indexDiff, string(subString[j]))

				j++

				if indexDiff > 1 && len(subString) > 1 {
					j = 0
					intArr = []int{}
				}

			}
			// fmt.Println(j)

			if j == len(subString) {
				fmt.Println("Here")
				intArrArr = append(intArrArr, intArr)
			}

			i++

			// break statement
			if i == len(s1) {
				break
			}
		}
	}

	fmt.Println(intArrArr)

	for _, ch := range intArrArr {
		for i := 0; i < len(ch); i++ {
			// fmt.Println("Here")
			s1[ch[i]] = "\033[34m" + s1[ch[i]] + "\033[0m"
		}
	}

	// for i:=0; i < len(intArr);i++{
	// 	// fmt.Println("Here")
	// 	s1[intArr[i]] = "\033[34m"+s1[intArr[i]]+"\033[0m"
	// }

	for i := 0; i < len(s1); i++ {
		resString += s1[i]
	}

	fmt.Println(resString)
}

func main() {
	FindSubString()
}

// if indexDiff == 1{
// 	// 	// fmt.Println(currIndex, prevIndex, string(subString[j]))
// 	// }
// 	// if indexDiff > 1{
// 	// 	intArr = []int{}
// 	// 	j = 0
// 	// }
