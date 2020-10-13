package main

import "fmt"

func SelectionSort(arr []int) []int {


	for i:=0;i<len(arr);i++ {
		min:=arr[i]
		key:=i
		//找出最小的值
		for j:=i;j<len(arr);j++{
			if arr[j]<min {
				min=arr[j]
				key=j
			}

		}
		//最小的值交换
		arr[key],arr[i]=arr[i],arr[key]
	}
	return arr
}

func main() {
	arr:=[]int{4,5,6,3,2,1}
	fmt.Print(SelectionSort(arr))

	
}
