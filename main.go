package esclient

import (
	"encoding/json"
	"fmt"
)

// Range
type Range struct {
	Range map[string]interface{} `json:"range"`
}

// multi_match
type MultiMatch struct {
	MultiMatch []map[string]interface{} `json:"multi_match"`
}

// Query
type Query map[string]interface{}

func (q Query) Marshal() []byte {
	by, _ := json.Marshal(q)
	return by
}

func (q Query) MatchAll() []byte {
	q["query"] = map[string]interface{}{
		"match_all": map[string]interface{}{},
	}
	return q.Marshal()
}

func (q Query) MatchOne(match Match) []byte {
	q["query"] = map[string]interface{}{
		"match": match,
	}
	return q.Marshal()
}

func (q Query) Must(must Must) []byte {
	q["query"] = map[string]interface{}{
		"bool": must,
	}
	return q.Marshal()
}

type Must struct {
	Must []map[string]interface{} `json:"must"`
}

func (must *Must) Match(match Match) {
	mm := map[string]interface{}{
		"match": map[string]interface{}{
			match.Field: match.Value}}
	must.Must = append(must.Must, mm)
}

func (must *Must) Wildcard(wildcard Wildcard) {
	must.Must = append(must.Must, map[string]interface{}{
		"wildcard": map[string]interface{}{
			wildcard.Field: fmt.Sprintf("*%s*", wildcard.Value),
		},
	})
}

type Should map[string]interface{}
type MustNot map[string]interface{}
type Filter map[string]interface{}
type MinimumShouldMatch int

// Match
type Match struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// Wildcard
type Wildcard struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// SortASC
func SortASC(keys ...string) interface{} {
	var sort []map[string]interface{}
	/*
	 "sort": [{"field": "asc"}],
	*/
	for _, key := range keys {
		sort = append(sort, map[string]interface{}{
			key: "asc",
		})
	}
	return sort
}

// SortDESC
func SortDESC(keys ...string) interface{} {
	var sort []map[string]interface{}
	for _, key := range keys {
		sort = append(sort, map[string]interface{}{
			key: "desc",
		})
	}
	return sort
}
