/*
 * Date: 2022/01/11
 * File: query.go
 */

// Package model TODO package function desc
package model

import "encoding/json"

type Query struct {
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

func NewQuery() *Query {
	return &Query{}
}
func (query *Query) Match(key string, value interface{}) *Query {
	query.match = append(query.match, Match{key, value})
	return query
}
func (query *Query) Term(key string, value interface{}) *Query {
	query.term = append(query.term, Term{key, value})
	return query
}

func (query *Query) Terms(key string, value []interface{}) *Query {
	query.terms = append(query.terms, Terms{key, value})
	return query
}

func (query *Query) Range(key string, operator string, value interface{}) *Query {
	query.between = append(query.between, Range{key, value, operator})
	return query
}
func (query *Query) Sort(key string, order string) *Query {
	query.sort = append(query.sort, Sort{key, order})
	return query
}
func (query *Query) From(from int64) *Query {
	query.from = from
	return query
}
func (query *Query) Size(size int64) *Query {
	query.size = size
	return query
}
func (query *Query) Filter() *Query {
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
func (query *Query) Must() *Query {
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

func (query *Query) Build() string {
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
		var sorts []map[string]map[string]interface{}
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
