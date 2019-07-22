package arrayfiler

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{})
}

type Input struct {
	Object map[string]interface{} `md:"object"` // The message to log
}

type Output struct {
	Output []interface{} `md:"output"` // The result of the counter operation
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"object": i.Object,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Object, err = coerce.ToObject(values["object"])
	if err != nil {
		return err
	}
	return nil
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)

	var keyValueMapArray []map[string]interface{}

	for k, v := range input.Object {
		m := make(map[string]interface{})
		m["Key"] = k
		m["Value"] = v
		keyValueMapArray = append(keyValueMapArray, m)
	}

	err = ctx.SetOutput("output", keyValueMapArray)
	if err != nil {
		panic(err)
	}
	return true, nil
}
