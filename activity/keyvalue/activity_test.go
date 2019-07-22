package arrayfiler

import (
	"encoding/json"
	"fmt"
	"github.com/project-flogo/core/activity"
	_ "github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	v := `
   {
      "appt_id":1,
      "practice_id":1,
      "patient_id":1,
      "pateint_refrence_id":"last inserted patient id"
   }`
	var in map[string]interface{}

	json.Unmarshal([]byte(v), &in)
	input := &Input{Object: in}

	tc.SetInputObject(input)

	act.Eval(tc)
	s, _ := json.Marshal(tc.GetOutput("output"))
	fmt.Println(string(s))
}
