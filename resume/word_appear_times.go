package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
)

//读取文本，统计每个单词出现次数，并逆序打印
//文本存在tab、逗号、空格等

var Str string

func ReadTxt() {
	text := "Hello World,\tHello Golang\tI love you\tGolang is the best language"
	reader := strings.NewReader(text)
	bufReader := bufio.NewReader(reader)
	for {
		str, _, err := bufReader.ReadLine()
		//fmt.Println(string(str))
		if err != nil {
			break
		}
		Str = Str + string(str) + " "
	}

}

func SliceBaseOnWord() {
	strings.Replace(Str, "\t", " ", -1)

	fmt.Println(Str)

}

func main() {
	//ReadTxt()
	//SliceBaseOnWord()

	str := "Hello World"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		fmt.Println(reflect.TypeOf(str[i]))
		fmt.Printf("%c\n", str[i])

	}

}
