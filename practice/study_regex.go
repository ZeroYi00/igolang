package practice

import (
	"bytes"
	"fmt"
	"regexp"
)

func MatchByRegex() {
	//MatchString返回的第一个参数是bool类型即匹配结果，第二个参数是error类型
	sourceStr := `my email is gerrylon@163.com`
	//\w 的释义指包含大小写字母数字和下划线 相当于([0-9a-zA-Z_])
	//(?:X)在正则中表示所匹配的子组X不作为结果输出
	matched, _ := regexp.MatchString(`[\w-]+@[\w]+(?:\.[\w]+)+`, sourceStr)
	fmt.Printf("%v\n", matched) // true
}

//Compile函数或MustCompile函数是将正则表达式进行编译，返回优化的 Regexp 结构体
//区别:
//1、Compile返回两个参数*Regexp,error类型
//2、MustCompile只返回*Regexp类型

//regexp.Compile
func RegCompile() {
	r, _ := regexp.Compile("p([a-z]+)ch")
	b := r.MatchString("peach")
	fmt.Println(b) //结果为true

	//查找匹配的字符串
	str1 := r.FindString("peach punch")
	fmt.Printf("FindString Return:%s\n", str1)

	//查找匹配字符串开始和结束位置的索引，而不是匹配内容[0 5]
	loc := r.FindStringIndex("peach punch")
	fmt.Printf("FindStringIndex Return:%v\n", loc)

	//返回完全匹配和局部匹配的字符串，例如，这里会返回  p([a-z]+)ch 和 `([a-z]+) 的信息
	ret := r.FindStringSubmatch("peach punch") //打印结果：[peach ea]
	fmt.Printf("FindStringSubmatch Return:%v\n", ret)
}

func RegMustCompile() {
	r := regexp.MustCompile("p([a-z]+)ch")
	b := r.MatchString("peach")
	fmt.Println(b)                        //结果为true
	fmt.Println(r.Match([]byte("peach"))) //打印结果：true

	//将匹配的结果，替换成新输入的结果
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //打印结果： a <fruit>

	//Func 变量允许传递匹配内容到一个给定的函数中，
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Printf(string(out)) //打印结果：   a PEACH

	//返回完全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //结果： [0 5 1 3]

	//返回所有的匹配项，而不仅仅是首次匹配项。正整数用来限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", -1)) //打印结果：[peach punch pinch]
	fmt.Println(r.FindAllString("peach punch pinch", 2))  //匹配两次   打印结果：[peach punch]

	//返回所有的完全匹配和局部匹配的索引位置
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	//打印结果： [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	//	var sourceStr string = `
	//test text     lljflsdfjdskal
	//gerrylon@163.com
	//abc@gmail.com
	//someone@sina.com.cn`
	//	re := regexp.MustCompile(`[\w-]+@([\w]+(?:\.[\w]+)+)`)
	//	matched := re.FindAllStringSubmatch(sourceStr, -1)
	//	for _, match := range matched {
	//		fmt.Printf("email is: %s, domain is: %s\n", match[0], match[1])
	//	}
	//email is: gerrylon@163.com, domain is: 163.com
	//email is: abc@gmail.com, domain is: gmail.com
	//email is: someone@sina.com.cn, domain is: sina.com.cn
}
