package file_util

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type FileRW struct {
	FileName string
	Flag     int
	Perm     fs.FileMode
}

//按行读取
func (frw *FileRW) BufioReadByline() ([]string, error) {
	lines := make([]string, 0, 0)
	f, err := os.Open(frw.FileName)

	defer f.Close()
	if err != nil {
		return nil, err
	}
	bufReader := bufio.NewReader(f)
	for {
		line, _, eof := bufReader.ReadLine()
		if eof == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

//写字符串
func (frw *FileRW) BufioWriteString(content string) error {
	//os.O_RDONLY|os.O_APPEND, 0666
	f, err := os.OpenFile(frw.FileName, frw.Flag, frw.Perm)
	defer f.Close()
	if err != nil {
		return err
	}
	bufWriter := bufio.NewWriter(f)
	_, err = bufWriter.WriteString(content)
	if err != nil {
		return err
	}
	bufWriter.Flush()
	f.Sync() //持久化
	return nil
}

//文件内容读至p中
//返回bytes 大小
func (frw *FileRW) BufioNewReader(p []byte) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recover")
		}
	}()
	f, err := os.Open(frw.FileName)
	defer f.Close()
	if err != nil {
		return 0, err
	}
	bufReader := bufio.NewReader(f)

	n, erR := bufReader.Read(p)
	return n, erR
}
