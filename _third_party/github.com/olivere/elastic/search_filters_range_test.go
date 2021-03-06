package elastic

import (
	"encoding/json"
	"testing"
)

func TestRangeFilter(t *testing.T) {
	f := NewRangeFilter("postDate").From("2010-03-01").To("2010-04-01")
	f = f.Cache(true)
	f = f.CacheKey("MyAndFilter")
	f = f.FilterName("MyFilterName")
	f = f.Execution("index")
	data, err := json.Marshal(f.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"_cache":true,"_cache_key":"MyAndFilter","_name":"MyFilterName","execution":"index","postDate":{"from":"2010-03-01","include_lower":true,"include_upper":true,"to":"2010-04-01"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

/*
func TestRangeFilterGte(t *testing.T) {
	f := NewRangeFilter("postDate").Gte("2010-03-01")
	data, err := json.Marshal(f.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"postDate":{"gte":"2010-03-01"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
*/
