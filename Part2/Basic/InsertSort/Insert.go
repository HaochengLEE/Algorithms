package main

import (
	"fmt"
)

func InsertSort(arr []int) []int {
	//i 每增加 1，就有一个未排序元素变成了已排序元素
	for i:=1;i<len(arr);i++ {
		value:=arr[i]
		j:=i-1
		//找到 arr[i] 的插入位置
		for ;j>=0;j-- {
			if arr[j]>value{
				//后移数据
				arr[j+1]=arr[j]
			}else {
				break
			}
		}
		//插入值
		arr[j+1]=value
	}
	return arr

}
func main() {
	arr:=[]int{4,5,6,3,2,1}
	fmt.Print(InsertSort(arr))

}
