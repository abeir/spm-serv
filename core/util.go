package core

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

//CloseQuietly 安静的调用Close()
func CloseQuietly(closer io.Closer){
	_ = closer.Close()
}

func IsExists(p string) bool{
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//左补位
// s为原字符串，pad为补位字符，len为补位后的长度
// 请注意，pad只能是长度为1的字符串
func PadLeft(s string, pad string, len int) (string, error){
	if utf8.RuneCountInString(pad) != 1 {
		return "", errors.New("pad length must be 1")
	}
	if s=="" {
		return strings.Repeat(pad, len), nil
	}
	sLen := utf8.RuneCountInString(s)
	if sLen >= len {
		return s, nil
	}
	pads := strings.Repeat(pad, len - sLen)
	return pads + s, nil
}

// 生成uuid，注意：并没有去除"-"分隔符
func UUID() string{
	return uuid.NewV4().String()
}
