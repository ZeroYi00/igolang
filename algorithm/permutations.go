package algorithm

//力扣(全排列46题)：https://leetcode.cn/problems/permutations/
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//示例:
//  输入：nums = [1,2,3]
//  输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

func permute(nums []int) [][]int {
	//结果存储
	var (
		res   = make([][]int, 0, 0)
		track = make([]int, 0, 0)
	)

	res = backtrack(res, track, nums)

	return res
}

func backtrack(res [][]int, track []int, nums []int) [][]int {
	//递归结束条件
	if len(track) == len(nums) {
		//copy 记录到结果集
		copySlice := make([]int, len(nums), len(nums))
		copy(copySlice, track)
		res = append(res, copySlice)
		//返回
		return res
	}

	for i := 0; i < len(nums); i++ {
		//过滤：确认nums[i]是否在track中
		if contain(track, nums[i]) {
			continue
		}
		//做选择：添加元素
		track := append(track, nums[i])
		//回溯
		res = backtrack(res, track, nums)
		//撤销选择：删除最后一个元素
		track = track[:len(track)-1]
	}
	return res

}

func contain(s []int, x int) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

//func main() {
//	nums := []int{1, 2, 3, 4, 5}
//	res := permute(nums)
//	fmt.Println(res)
//}
