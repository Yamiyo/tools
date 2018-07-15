package timetool

import (
	"fmt"
	"testing"
	"time"
)

func Test_GetIntervalBefore(t *testing.T) {

	srcTime := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)

	target, err := GetIntervalBefore(srcTime, 1, "d")
	fmt.Printf("%+v,  %+v\n", target, err)
}
