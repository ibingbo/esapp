/*
 * Date: 2022/01/11
 * File: query_test.go
 */

// Package model TODO package function desc
package model

import (
	"testing"
)

func TestNewQuery(t *testing.T) {
	res := NewQuery().
		Match("source", "ugc").
		Term("user_id", "4132699562").
		Terms("type", []interface{}{"word"}).
		Must().
		Range("create_time", GTE, 1637942400).
		Filter().
		Range("create_time", LTE, 1637943404).
		Filter().
		Sort("create_time", "desc").
		From(0).
		Size(2).Build()
	t.Log(res)
}
