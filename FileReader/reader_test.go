package FileReader

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func getLength()(int){
	if runtime.GOOS == "windows" {
		return 2
	} else {
		return 1
	}
}

var parentDirectory = "../"

func TestGetNewLineCharacterLength(t *testing.T) {
	filePath := parentDirectory + "Employees.csv"
	var expected int = getLength()
	actual := GetNewLineCharacterLength(filePath, ",")
	if actual != expected {
		t.Error("Expected: ", expected, "Actual: ", actual)
	}
}

func TestGetOffsets(t *testing.T) {
	filePath := parentDirectory + "Employees.csv"
	expected := []string{
		"198,Donald,OConnell,DOCONNEL,650.507.9833,21-JUN-07,SH_CLERK,2600, - ,124,50",
		"199,Douglas,Grant,DGRANT,650.507.9844,13-JAN-08,SH_CLERK,2600, - ,124,50",
		"200,Jennifer,Whalen,JWHALEN,515.123.4444,17-SEP-03,AD_ASST,4400, - ,101,10",
		"201,Michael,Hartstein,MHARTSTE,515.123.5555,17-FEB-04,MK_MAN,13000, - ,100,20",
		"202,Pat,Fay,PFAY,603.123.6666,17-AUG-05,MK_REP,6000, - ,201,20",
	}
	offsets := GetOffsets(filePath, ",")
	actual := GetDataFromOffsets(filePath, offsets[1:6])
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected: ", expected, "Actual: ", actual)
	}
	fmt.Println(actual)
}