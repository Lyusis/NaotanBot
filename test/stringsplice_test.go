package test

import (
	"fmt"
	"github.com/Lyusis/NaotanBot/utils"
	"testing"
)

func TestExtractContent(t *testing.T) {
	for index, item := range utils.ExtractContent(" 订阅   眠大佐 22816111") {
		fmt.Printf("%d: %s\n", index, item)
	}
}
