package config

import (
	"bufio"
	"bytes"
	"container/list"
	"logger"
	"os"
	"strconv"
	"unicode"
)

type element struct {
	key   string // 如果是属性行这里表示属性的key
	value string // 行的内容,如果是注释注释引导符也包含在内
}

// Information The properties document in memory.
type Information struct {
	elems *list.List
	props map[string]*list.Element
}

func VLGParsing(filename string) {
	vlgFile, err := os.Open(filename)
	if err != nil {
		logger.Sugar.Error("无法读取 virtual liver gachi 配置文件", logger.FormatTitle("WRONG"), err)
	}

	lineCount := -1
	lastColon := -1
	scanner := bufio.NewScanner(vlgFile)

	for scanner.Scan() {
		//  逐行读取
		line := scanner.Bytes()
		lineCount++

		//  遇到空行
		if 0 == len(line) {
			continue
		}

		//  找到第一个非空白字符
		pos := bytes.IndexFunc(line, func(r rune) bool {
			return !unicode.IsSpace(r)
		})

		//  找到第一个冒号的位置
		end := bytes.IndexFunc(line[pos+1:], func(r rune) bool {
			return ':' == r
		})

		if -1 != end {
			lastColon = lineCount
		}

		//  遇到空白行
		if -1 == pos {
			continue
		}

		//  遇到注释行
		if '#' == line[pos] {
			continue
		}

		element := element{}

		if -1 != end {
			element.key = string(bytes.TrimRightFunc(line[pos:pos+1+end], func(r rune) bool {
				return unicode.IsSpace(r)
			}))

			element.value = string(bytes.TrimSpace(line[pos+1+end+1:]))
		}

	}

	//  完全没有冒号，说明该配置文件错误
	if -1 == lastColon {
		logger.Sugar.Error("无法解析 virtual liver gachi 配置文件", logger.FormatTitle("WRONG"), err)
	}
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//err = yaml.Unmarshal(yamlFile, &items)
//if err != nil {
//	logger.Error("无法解析 virtual liver gachi 配置文件", false, err)
//}

//}
