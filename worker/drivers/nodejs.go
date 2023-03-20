package drivers

import (
	"fmt"
	"os"
	"worker/pb"
)

type NodeJSDriver struct{}

func NewNodejsDriver() NodeJSDriver {
	return NodeJSDriver{}
}

func (d *NodeJSDriver) Run(problem_id uint, code string) pb.JudgeResult {
	file, err := os.OpenFile(
		fmt.Sprintf("%s%d", PROBLEM_FILES_PATH, problem_id),
		os.O_RDONLY, os.ModeCharDevice)

	if err != nil {
		return pb.JudgeResult{
			ResultCode: pb.JudgeInfo_err,
			Submission: nil,
		}
	}

}
