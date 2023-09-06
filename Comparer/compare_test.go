package Comparer

import (
	"reflect"
	"sort"
	"testing"
)
var parentDir = "../"
func TestComparer(t *testing.T) {
	expected := map[string][]string{
		"Inserts": {
					"42|Shikhar Dhawan|29",
					"24|Ajinkya Rahane|26",
					},
		"Updates": {
					"18|Virat Kohli|26",
					"7|MS Dhoni|33",
					},
		"Deletes": {
					"44|Virender Sehwag|32",
					"23|Gautam Gambhir|29",
					},
	}
	inserts, updates, deletes := Compare(parentDir + "players_v1.csv", parentDir + "players_v2.csv", "|")
	// Sorting the data before comparing
	for key := range expected {
		sort.Slice(expected[key], func(i, j int) bool {
			return expected[key][i] < expected[key][j]
		})
	}
	actual := map[string][]string{
		"Inserts": inserts,
		"Updates": updates,
		"Deletes": deletes,
	} 
	// Sorting the data before comparing
	for key := range actual {
		sort.Slice(actual[key], func(i, j int) bool {
			return actual[key][i] < actual[key][j]
		})
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Test failed, expected: '%v', actual:  '%v'", expected, actual)
	}
}