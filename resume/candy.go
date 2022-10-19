package main

import "fmt"

//常数空间遍历
//糖果总是尽量少给，且从 11 开始累计，每次要么比相邻的同学多给一个，要么重新置为1

//func candy(ratings []int) int {
//	n := len(ratings)
//	ans, inc, dec, pre := 1, 1, 0, 1
//	for i := 1; i < n; i++ {
//		if ratings[i] >= ratings[i-1] {
//			dec = 0
//			if ratings[i] == ratings[i-1] {
//				pre = 1
//			} else {
//				pre++
//			}
//			ans += pre
//			inc = pre
//		} else {
//			dec++
//			if dec == inc {
//				dec++
//			}
//			ans += dec
//			pre = 1
//		}
//	}
//	return ans
//}

func candy(ratings []int) int {
	var (
		pre int //左边同学糖果
		inc int //递增序列的长度
		dec int //递减序列的长度
	)

	//n := len(ratings)
	var ans = 0
	for i, _ := range ratings {
		//如果当前同学得分>上一个
		if i == 0 { //第1个同学，发1个糖
			pre, ans = 1, 1
		} else if ratings[i] > ratings[i-1] { //当前处于递增序列
			inc, dec = 1, 0
			ans = pre + 1 + dec
		} else { //当前同学得分<=前一个时：直接赋1

		}
	}
	return ans
}

func main() {
	//ratings := []int{1, 2, 2}
	//ratings := []int{1, 3, 5, 2, 3, 3}
	ratings := []int{1, 0, 2}
	ans := candy(ratings)

	fmt.Println(ans)
}
