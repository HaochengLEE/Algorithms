# 2.1.2 选择排序
我们将数据分为已排序和未排序两个部分。

先找到最小的数，和第一个元素交换位置。此时第一个位置就是已排序的部分，而后面的位置就是未排序的部分。

接着从未排序的数据中找出最小值，将其与未排序的第一个元素交换位置。

如此反复，排序就完成了。

```
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
```
**总结**
1. go 语言可以直接使用
```
arr[key],arr[i]=arr[i],arr[key]
```
进行值的交换，不需要再用另外的变量

# 2.1.3 插入排序
