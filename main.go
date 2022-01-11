/*
 * Date: 2021/11/28
 * File: main.go
 */

// Package awesomeProject TODO package function desc
package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/Andrew-M-C/go.emoji"
)

func hasChinese(word string) bool {
	for _, char := range word {
		if unicode.Is(unicode.Han, char) {
			return true
		}
	}
	return false
}
func length(word string) int64 {
	charLength := int64(0)
	for _, char := range word {
		if unicode.Is(unicode.Han, char) {
			charLength += 2
		} else {
			charLength += 1
		}
	}
	return charLength
}

// str:="HelloWord"
// l1:=len([]rune(str))
// l2:=bytes.Count([]byte(str),nil)-1)
// l3:=strings.Count(str,"")-1
// l4:=utf8.RuneCountInString(str)
func main() {
	str2 := "hello,bill"
	str1 := "男子深夜找工。作饭店老板端上热饭"
	fmt.Println(hasChinese("hello"), hasChinese("hell,好"), hasChinese("你好"), hasChinese("111"), hasChinese("。1"))
	fmt.Println("****", length(str1), length(str2), len([]rune(str1)), len([]rune(str2)))
	fmt.Println(utf8.RuneCountInString(str1), utf8.RuneCountInString(str2), bytes.Count([]byte(str1), nil), bytes.Count([]byte(str2), nil))
	str := "✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊✊"
	str11 := "1一丨一`"
	fmt.Println("-----", hasChinese(str11))
	str12 := "禿             禿"
	fmt.Println(strings.ReplaceAll(str12, " ", ""))
	set := map[int32]struct{}{}
	for _, s := range str {
		set[s] = struct{}{}
		fmt.Println(s, reflect.TypeOf(s))
	}
	fmt.Println(set, len(set), repeated(str), emoji.ReplaceAllEmojiFunc(str, func(emoji string) string {
		return ""
	}))
	timeStr := time.Now().Format("2006-01-02")
	timeDayBegin, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	fmt.Println(timeDayBegin, timeDayBegin.Unix())

}

func repeated(content string) bool {
	set := map[int32]struct{}{}
	for _, item := range content {
		set[item] = struct{}{}
	}
	return len(set) == 1
}
