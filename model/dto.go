/*
 * Date: 2022/01/11
 * File: dto.go
 */

// Package model TODO package function desc
package model

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