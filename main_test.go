package brazil_merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var testArray, realResult []string

	testArray = append(testArray, "e", "d", "a", "b", "c", "f")
	realResult = append(realResult, "a", "b", "c", "d", "e", "f")
	var testResult []string = MergeSort(testArray)

	if !reflect.DeepEqual(testResult, realResult) {
		t.Errorf("Failure. \nResult expected: \t%s\nActual result: \t\t%s", realResult, testResult)
	}

}
