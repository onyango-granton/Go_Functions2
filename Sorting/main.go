package main

import "fmt"

func main(){
		maxInt := 0
		minInt := 900

		numArr := []int{4,9,5,3,2,8}
		length := len(numArr)
		copyArr := make([]int,length)
		copy(copyArr, numArr)

		for _,ch := range copyArr{
				if ch > maxInt{
						maxInt = ch
				}
				if ch < minInt{
						minInt = ch
				}
		}

		countArr := make([]int,maxInt-minInt+1)

		for _, ch := range copyArr{
				countArr[ch - minInt]++
		
		}

		total := 0
		for i := range countArr{
				subtotal := countArr[i]
				countArr[i] = total
				total += subtotal
		}

		for _,ch := range numArr{
				copyArr[countArr[ch - minInt]] = ch
				countArr[ch-minInt]++
		}

		fmt.Println(copyArr)

}

