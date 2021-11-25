package utils

import (
	"container/list"
)

func PopUp(list *list.List) string {
	if list.Len() > 0 {
		result := list.Front().Value.(string)
		list.Remove(list.Front())
		return result
	}
	return ""
}

func ArrayToStack(array []string) *list.List {
	var (
		stack *list.List
	)

	for _, str := range array {
		stack.PushBack(str)
	}
	return stack
}
