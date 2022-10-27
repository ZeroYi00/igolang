package algorithm

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := permute(nums)
	fmt.Println(res)
	//n!=n*(n-1)*...*1
	//5*4*3*2*1=120
	fmt.Println(len(res))
}
