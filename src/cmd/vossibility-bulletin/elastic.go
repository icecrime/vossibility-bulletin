package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattbaird/elastigo/core"
)

const QueriesStore = "queries"

func executeQuery(index, queryFile string) (core.SearchResult, error) {
	f, err := os.Open(filepath.Join(QueriesStore, queryFile))
	if err != nil {
		return core.SearchResult{}, fmt.Errorf("fail to open query file %q", queryFile)
	}
	defer f.Close()
	return core.SearchRequest(index, "", map[string]interface{}{}, f)
}
