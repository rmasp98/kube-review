package search_test

import (
	"kube-review/mocks"
	"kube-review/search"
	"reflect"
	"testing"
)

//tests
// Should not start with anything other than Any, Key, Value
//can return possible keywords
//correctly proccess input

func TestBasicCommandExecutesCorrectFunctions(t *testing.T) {
	intelligent, _ := search.NewIntelligent("Any==\"Test\"")
	mock := mocks.NodeListMock{}
	intelligent.Execute(&mock)
	expected := []string{"GetNodesMatching", "ApplyFilter"}
	actual := mock.Calls
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestComplexCommandExecutesCorrectFunctions(t *testing.T) {
	intelligent, _ := search.NewIntelligent("Any==\"Test\" + (ParentHasChildKey==\"Test2\")")
	mock := mocks.NodeListMock{}
	intelligent.Execute(&mock)
	expected := []string{"GetNodesMatching", "GetParentChildrenMatching", "GetParentChildrenMatching", "GetParentChildrenMatching", "ApplyFilter"}
	actual := mock.Calls
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestComplexCommandReturnsCorrectOutput(t *testing.T) {
	intelligent, _ := search.NewIntelligent("Any==\"Test\" + (ParentHasChildKey==\"Test2\")")
	mock := mocks.NodeListMock{}
	intelligent.Execute(&mock)
	expected := []int{1, 2, 5, 2, 7, 2, 7, 2, 7}
	actual := mock.Args[len(mock.Args)-1][0]
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestSubtractCommandReturnsCorrectOutput(t *testing.T) {
	intelligent, _ := search.NewIntelligent("Any==\"Test\" - Any==\"Test2\"")
	mock := mocks.NodeListMock{}
	intelligent.Execute(&mock)
	actual := mock.Args[2][0].([]int)
	if len(actual) != 0 {
		t.Errorf("Expected empty array but got '%v'", actual)
	}
}

func TestBracketsCommandReturnsCorrectOutput(t *testing.T) {
	intelligent, _ := search.NewIntelligent("Any==\"Test\" - (ParentHasChildKey==\"Test2\" + ChildHasValue==\"Test3\") + Key!=\"Test4\"")
	mock := mocks.NodeListMock{}
	intelligent.Execute(&mock)
	expected := []int{5, 1, 2, 5}
	actual := mock.Args[len(mock.Args)-1][0]
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}