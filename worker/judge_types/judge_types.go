package judge_types

import "worker/pb"

type 

type JudgeTypes interface {
	Judge(*pb.Submission) (*pb.JudgeResult, error)
}

var JudgeTypesTable = map[string]JudgeTypes{
	"normal": NormalJudge{},
}

type NormalJudge struct{}

func (j *NormalJudge) Judge(submission *pb.Submission) (*pb.JudgeResult, error) {

}
