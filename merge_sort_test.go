package mdbrasilianalphabetsort

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

	var testArray2, realResult2 []string

	testArray2 = append(testArray2, "bcd", "dfe", "def", "abc")
	realResult2 = append(realResult2, "abc", "bcd", "def", "dfe")
	var testResult2 []string = MergeSort(testArray2)

	if !reflect.DeepEqual(testResult2, realResult2) {
		t.Errorf("Failure. \nResult expected: \t%s\nActual result: \t\t%s", realResult2, testResult2)
	}

}
