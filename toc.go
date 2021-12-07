package main

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
	"net/url"
)

// HistoryIDs custom implementation of goldmark/parser.IDs
type HistoryIDs struct {
	values map[string]bool
}

func newIDs() *HistoryIDs {
	return &HistoryIDs{
		values: map[string]bool{},
	}
}

// Generate generates a new element id by encode the title
func (s *HistoryIDs) Generate(value []byte, _ ast.NodeKind) []byte {
	value = util.TrimLeftSpace(value)
	value = util.TrimRightSpace(value)
	result := []byte(url.QueryEscape(string(value)))
	if _, ok := s.values[util.BytesToReadOnlyString(result)]; !ok {
		s.values[util.BytesToReadOnlyString(result)] = true
		return result
	}
	for i := 1; ; i++ {
		newResult := fmt.Sprintf("%s-%d", result, i)
		if _, ok := s.values[newResult]; !ok {
			s.values[newResult] = true
			return []byte(newResult)
		}
	}
}

// Put puts a given element id to the used ids table.
func (s *HistoryIDs) Put(value []byte) {
	s.values[util.BytesToReadOnlyString(value)] = true
}
