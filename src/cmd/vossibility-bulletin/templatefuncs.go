package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	"github.com/mattbaird/elastigo/core"
)

var templateFuncs = template.FuncMap{
	"count":      fnCount,
	"formatDate": fnFormatDate,
	"map":        fnMap,
	"one":        fnOne,
	"search":     fnSearch,
}

func fnCount(index, queryFile string) (interface{}, error) {
	res, err := executeQuery(index, queryFile)
	if err != nil {
		return nil, err
	}
	return res.Hits.Total, nil
}

func fnFormatDate(date string) (interface{}, error) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}
	return t.Format("02 Jan 06"), nil
}

func fnMap(v interface{}) (interface{}, error) {
	var m map[string]interface{}
	if j, ok := v.(json.RawMessage); ok {
		err := json.Unmarshal(j, &m)
		return m, err
	} else if j, ok := v.(*json.RawMessage); ok {
		err := json.Unmarshal(*j, &m)
		return m, err
	}
	return nil, fmt.Errorf("invalid argument for fnMap: %#v", v)
}

func fnOne(v interface{}) (interface{}, error) {
	s, ok := v.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected an array parameter, got %#v", v)
	}
	if len(s) != 1 {
		return nil, fmt.Errorf("expected a single element array, got %#v", v)
	}
	return s[0], nil
}

func fnSearch(index, queryFile string) (interface{}, error) {
	res, err := executeQuery(index, queryFile)
	if err != nil {
		return nil, err
	}

	var aggs map[string]interface{}
	if res.Aggregations != nil {
		if err := json.Unmarshal(res.Aggregations, &aggs); err != nil {
			return nil, err
		}
	}

	return struct {
		Hits         core.Hits
		Aggregations map[string]interface{}
	}{
		Hits:         res.Hits,
		Aggregations: aggs,
	}, nil
}
