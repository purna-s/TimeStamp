package timestamp

import (
	"fmt"
    "time"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-timestamp")

// MyActivity is a stub for your Activity implementation
type timestamp struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &timestamp{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *timestamp) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *timestamp) Eval(ctx activity.Context) (done bool, err error) {
	
	dt := time.Now()

	activityLog.Debugf("Activity has listed out the latest file Successfully")
	fmt.Println("Activity has listed out the latest file Successfully")
	
	ctx.SetOutput("TimeStamp", dt)
	return true, nil
}

