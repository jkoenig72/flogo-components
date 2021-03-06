package splitjson

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/flow/test"
	"io/ioutil"
	"testing"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestCreate(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("splitting json string into separate name value pairs: '{\"distance\":150, \"status\":\"optimal\"}'")

	//setup attrs
	tc.SetInput("input", "{\"distance\":150, \"status\":\"optimal\"}")

	act.Eval(tc)

	result := tc.GetOutput("result")
	name1 := tc.GetOutput("name1")
	value1 := tc.GetOutput("value1")
	name2 := tc.GetOutput("name2")
	value2 := tc.GetOutput("value2")

	fmt.Println("result: ", result)
	fmt.Println("name1: ", name1)
	fmt.Println("value1: ", value1)
	fmt.Println("name2: ", name2)
	fmt.Println("value2: ", value2)

	if result == nil {
		t.Fail()
	}
}
