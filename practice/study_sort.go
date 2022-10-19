package practice

import (
	"fmt"
	"sort"
)

//golang 标准库sort学习
//用来实现排序相关的操作，实现了四种基本的排序算法：插入排序（insertionSort）、归并排序（symMerge）、堆排序（heapSort）和快速排序（quickSort）
//sort 包会依据实际数据自动选择最优的排序算法会对原值进行改变

//Note:
//一、排序
//对数据集合（包括自定义数据类型的集合）排序，需要实现sort.Interface接口的三个方法

//sort.IntSlice 实现了 sort.Interface
//sort.Reverse 基于Less() 重新实现了逻辑
//sort.SearchInts 封装 sort.Search  二分查找
//sort.IntsAreSorted 封装 sort.IsSorted
func SortInts() {
	arr := []int{3, 5, 1, 2, 4}
	fmt.Printf("排序前arr: %v\n", arr)

	//默认升序.
	//sort.Ints 本质是包装了 sort.Sort 方法
	sort.Ints(arr)
	fmt.Printf("排序(升序)后arr: %v\n", arr)

	bo := sort.IntsAreSorted(arr)
	fmt.Printf("是否排序：%v\n", bo)

	//降序
	//sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	//fmt.Printf("排序(降序)后arr: %v\n", arr)

	val := 5
	//查找位置
	//SearchInts使用条件：切片已升序排序
	r := sort.SearchInts(arr, val)
	fmt.Printf("值%d 索引位置为: %v\n", val, r)

}

//sort.Float64Slice
func SortFloat64() {
	arr := []float64{1.1, 4.4, 5.5, 3.3, 2.2}
	fmt.Printf("排序前的arr: %v\n", arr)

	//sort.Float64s(arr)
	//fmt.Printf("排序(升序)后的arr: %v\n", arr)

	sort.Sort(sort.Reverse(sort.Float64Slice(arr)))
	fmt.Printf("排序(降序)后的arr: %v\n", arr)
	
	//返回FALSE；原因：依赖于实现的LESS()和LEN()方法
	bo := sort.Float64sAreSorted(arr)
	fmt.Printf("是否排序：%v\n", bo)

	var x float64 = 5.5
	r := sort.SearchFloat64s(arr, x)
	fmt.Printf("值%f 索引位置为: %v\n", x, r)
}
