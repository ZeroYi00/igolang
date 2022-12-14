package practice

import (
	"fmt"
	"sort"
)

//golang 标准库sort学习
//用来实现排序相关的操作，实现了四种基本的排序算法：插入排序（insertionSort）、归并排序（symMerge）、堆排序（heapSort）和快速排序（quickSort）
//sort 包会依据实际数据自动选择最优的排序算法会对原值进行改变

//Note:
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
func SortFloat64s() {
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

//sort.StringSlice
//基于”字典序“ 排序
func SortStrings() {
	//arrNum := sort.StringSlice{
	//	"100",
	//	"42",
	//	"41",
	//	"3",
	//	"2",
	//}
	//fmt.Println("原始数据：", arrNum)
	//sort.Strings(arrNum)
	//fmt.Println("排序(升序)数据：", arrNum)

	arrStr := []string{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Println("原始数据：", arrStr)
	sort.Strings(arrStr)
	fmt.Println("排序(升序)数据：", arrStr)
	//是否有序
	//bo := sort.StringsAreSorted(arrStr)
	//fmt.Println("是否有序:", bo)
	//search
	x := "ac"
	idx := sort.SearchStrings(arrStr, x)
	fmt.Printf("值%s 下标为：%v\n", x, idx)
	//降序
	//sort.Sort(sort.Reverse(sort.StringSlice(arrStr)))
	//fmt.Println("排序(降序)数据：", arrStr)

	//arrHan := sort.StringSlice{
	//	"啊",
	//	"博",
	//	"次",
	//	"得",
	//	"饿",
	//	"周",
	//}
	//fmt.Println("原始数据：", arrHan)
	//sort.Strings(arrHan)
	//fmt.Println("排序(升序)数据：", arrHan)
	//
	//for _, v := range arrHan {
	//	fmt.Println(v, []byte(v))
	//}

}

type Int2dSlice [][]int

//自定义实现sort.Interface
func (x Int2dSlice) Len() int { return len(x) }

//基于每个切片的第2个元素比较并排序
func (x Int2dSlice) Less(i, j int) bool { return x[i][1] < x[j][1] }
func (x Int2dSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//复杂结构 [][]int
func Sort2dInts() {
	arrIntSlice := Int2dSlice{
		{1, 4},
		{9, 3},
		{7, 5},
	}

	fmt.Println(arrIntSlice)
	//sort.Sort(arrIntSlice)
	//fmt.Println(arrIntSlice)
	bo := sort.IsSorted(arrIntSlice)
	fmt.Println("是否有序：", bo)

}

type MapSlice []map[string]float64

func (x MapSlice) Len() int           { return len(x) }
func (x MapSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x MapSlice) Less(i, j int) bool { return x[i]["a"] < x[j]["a"] } // 按照"a"对应的值排序

//复杂结构 []map[string]int
func SortMaps() {
	arrMap := MapSlice{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}

	fmt.Println(arrMap)
	//sort.Sort(arrMap)
	//fmt.Println(arrMap)
	bo := sort.IsSorted(arrMap)
	fmt.Println("是否有序：", bo)
}

type People struct {
	Name string
	Age  int
}
type PeopleSlice []People

func (x PeopleSlice) Len() int           { return len(x) }
func (x PeopleSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x PeopleSlice) Less(i, j int) bool { return x[i].Age < x[j].Age } // 按照"a"对应的值排序

//复杂结构体：[]struct
func SortStructs() {
	arrPeople := PeopleSlice{
		{Name: "n1", Age: 12},
		{Name: "n2", Age: 11},
		{Name: "n3", Age: 10},
	}
	fmt.Println(arrPeople)
	//sort.Sort(arPeople)
	//fmt.Println(arrPeople)
	sort.Sort(sort.Reverse(arrPeople))
	fmt.Println(arrPeople)

}
