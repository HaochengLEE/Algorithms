package main

import "fmt"

func Merge(arr []int) {
	//判断退出条件
	if len(arr)<=1 {
		return
	}
	var arr1,arr2 []int
	value:=0
	lenth:=len(arr)
	for i:=0;i<lenth/2;i++ {
		arr1[i]=arr[i]
	}
	for i:=lenth/2;i<lenth;i++ {
		arr2[value]=arr[i]
		value++
	}

	Merge(arr1)
	Merge(arr2)
}
func MergeSort(arr []int) []int {
	Merge(arr)
	return arr

}

func main() {
	arr:=[]int{4,5,6,3,2,1}
	fmt.Print(MergeSort(arr))
}
