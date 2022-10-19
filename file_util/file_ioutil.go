package file_util

import (
	"io/ioutil"
	"strings"
)

//返回文件内容所有行
func (frw *FileRW) IoUtilReader() ([]string, error) {
	var lines = make([]string, 0, 0)
	data, err := ioutil.ReadFile(frw.FileName)
	if err != nil {
		return nil, err
	} else {
		contents := string(data)
		lines = strings.Split(contents, "\n")
		return lines, nil
	}
}
