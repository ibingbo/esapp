/*
 * Date: 2021/11/28
 * File: main.go
 */

// Package awesomeProject TODO package function desc
package main

import (
	"bytes"
	"encoding/json"
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
	str11:="1一丨一`"
	fmt.Println("-----",hasChinese(str11))
	str12:= "禿             禿"
	fmt.Println(strings.ReplaceAll(str12," ",""))
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
	res := NewEsQuery().
		Match("source", "ugc_task").
		Term("user_id", "4132699562").
		Terms("type", []interface{}{"word"}).
		Must().
		Range("create_time", "gte", 1637942400).
		Filter().
		Range("create_time", "lte", 1637943404).
		Filter().
		Sort("create_time", "desc").
		From(0).
		Size(2).build()
	fmt.Println(res)
}

func repeated(content string) bool {
	set := map[int32]struct{}{}
	for _, item := range content {
		set[item] = struct{}{}
	}
	return len(set) == 1
}

type EsQuery struct {
	must    []map[string]interface{}
	filter  []map[string]interface{}
	term    []Term
	terms   []Terms
	between []Range
	match   []Match
	sort    []Sort
	from    int64
	size    int64
}

func NewEsQuery() *EsQuery {
	return &EsQuery{}
}
func (query *EsQuery) Match(key string, value interface{}) *EsQuery {
	query.match = append(query.match, Match{key, value})
	return query
}
func (query *EsQuery) Term(key string, value interface{}) *EsQuery {
	query.term = append(query.term, Term{key, value})
	return query
}

func (query *EsQuery) Terms(key string, value []interface{}) *EsQuery {
	query.terms = append(query.terms, Terms{key, value})
	return query
}

func (query *EsQuery) Range(key string, operator string, value interface{}) *EsQuery {
	query.between = append(query.between, Range{key, value, operator})
	return query
}
func (query *EsQuery) Sort(key string, order string) *EsQuery {
	query.sort = append(query.sort, Sort{key, order})
	return query
}
func (query *EsQuery) From(from int64) *EsQuery {
	query.from = from
	return query
}
func (query *EsQuery) Size(size int64) *EsQuery {
	query.size = size
	return query
}
func (query *EsQuery) Filter() *EsQuery {
	if len(query.terms) > 0 {
		for _, item := range query.terms {
			query.filter = append(query.filter, map[string]interface{}{
				"terms": map[string]interface{}{item.Key: item.Value},
			})
		}
		query.terms = []Terms{}
	}
	if len(query.term) > 0 {
		for _, item := range query.term {
			query.filter = append(query.filter, map[string]interface{}{
				"term": map[string]interface{}{item.Key: item.Value},
			})
		}
		query.term = []Term{}
	}
	if len(query.match) > 0 {
		for _, item := range query.match {
			query.filter = append(query.filter, map[string]interface{}{
				"match": map[string]interface{}{item.Key: item.Value},
			})
		}
		query.match = []Match{}
	}
	if len(query.between) > 0 {
		items := map[string]map[string]interface{}{}
		for _, item := range query.between {
			if _, ok := items[item.Key]; ok {
				items[item.Key][item.Operator] = item.Value
			} else {
				items[item.Key] = map[string]interface{}{item.Operator: item.Value}
			}
		}
		if len(items) > 0 {
			query.filter = append(query.filter, map[string]interface{}{"range": items})
		}
		query.between = []Range{}
	}
	return query
}
func (query *EsQuery) Must() *EsQuery {
	if len(query.terms) > 0 {
		for _, item := range query.terms {
			query.must = append(query.must, map[string]interface{}{
				"terms": map[string]interface{}{item.Key: item.Value},
			})
		}
		query.terms = []Terms{}
	}
	if len(query.term) > 0 {
		for _, item := range query.term {
			query.must = append(query.must, map[string]interface{}{
				"term": map[string]interface{}{item.Key: item.Value},
			})
		}
		query.term = []Term{}
	}
	if len(query.match) > 0 {
		for _, item := range query.match {
			query.must = append(query.must, map[string]interface{}{
				"match": map[string]interface{}{item.Key: item.Value},
			})
		}
		query.match = []Match{}
	}
	if len(query.between) > 0 {
		items := map[string]map[string]interface{}{}
		for _, item := range query.between {
			if _, ok := items[item.Key]; ok {
				items[item.Key][item.Operator] = item.Value
			} else {
				items[item.Key] = map[string]interface{}{item.Operator: item.Value}
			}
		}
		if len(items) > 0 {
			for _, item := range items {
				query.must = append(query.must, map[string]interface{}{"range": item})
			}
		}
		query.between = []Range{}
	}
	return query
}

func (query *EsQuery) build() string {
	condition := map[string]interface{}{
		"query": map[string]interface{}{},
	}
	boolMap := map[string]interface{}{}
	if len(query.filter) > 0 {
		boolMap["filter"] = query.filter
	}
	if len(query.must) > 0 {
		boolMap["must"] = query.must
	}
	if len(boolMap) > 0 {
		condition["query"] = map[string]interface{}{"bool": boolMap}
	}
	if len(query.sort) > 0 {
		sorts := []map[string]map[string]interface{}{}
		for _, item := range query.sort {
			sorts = append(sorts, map[string]map[string]interface{}{
				item.Key: {"order": item.Order},
			})
		}
		if len(sorts) > 0 {
			condition["sort"] = sorts
		}
	}
	if query.from >= 0 {
		condition["from"] = query.from
	}
	if query.size > 0 {
		condition["size"] = query.size
	}
	res, _ := json.Marshal(condition)
	return string(res)
}

type Term struct {
	Key   string
	Value interface{}
}
type Terms struct {
	Key   string
	Value []interface{}
}

type Match struct {
	Key   string
	Value interface{}
}

type Range struct {
	Key      string
	Value    interface{}
	Operator string
}

type Sort struct {
	Key   string
	Order string
}
