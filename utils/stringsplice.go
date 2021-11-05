package utils

import (
	"container/list"
	"strconv"
	"strings"
)

func SingleFrontInt(num int, msg0 string) string {
	return strconv.Itoa(num) + msg0
}

func SingleBackInt(msg0 string, num int) string {
	return msg0 + strconv.Itoa(num)
}

func MiddleInt(msg0 string, num int, msg1 string) string {
	return msg0 + strconv.Itoa(num) + msg1
}

// CheckCurlyBraces 查看文本中是否有花括号
func CheckCurlyBraces(str string) bool {
	var (
		left  = '{'
		right = '}'
	)

	for _, ch := range str {
		if ch == left || ch == right {
			return true
		}
	}

	return false
}

// ExtractContent 提取内容, 根据空格分割, 有引号保护除外
func ExtractContent(str string) []string {
	var (
		chStack list.List
		result  []string
		maskFlg bool
		//prevCh  rune

		space = ' '
	)

	str = strings.TrimSpace(str)
	for _, ch := range str {
		if space == ch {
			if maskFlg {
				// 遇到空格且在符号中
				chStack.PushBack(string(ch))
			} else {
				// 遇到空格但是不在符号中
				item := chStackToString(&chStack)
				if item != "" && item != " " {
					result = append(result, item)
				}
				chStack.Init()
			}
			continue
		}
		if rangeCheck(ch) {
			maskFlg = !maskFlg
			continue
		}
		chStack.PushBack(string(ch))
	}
	if maskFlg {
		return nil
	}
	result = append(result, chStackToString(&chStack))
	return result
}

func rangeCheck(ch rune) bool {
	var (
		left   = '“'
		middle = '"'
		right  = '”'
	)

	if ch == left || ch == middle || ch == right {
		return true
	} else {
		return false
	}
}

func chStackToString(chStack *list.List) string {
	var result string
	length := chStack.Len()
	for i := 0; i < length; i++ {
		result += chStack.Front().Value.(string)
		chStack.Remove(chStack.Front())
	}
	return result
}
