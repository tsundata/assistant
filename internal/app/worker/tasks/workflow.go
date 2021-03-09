package tasks

import "time"

func RunWorkflow(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	time.Sleep(time.Hour)
	return sum, nil
}
