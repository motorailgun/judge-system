package judge

import "distributor/pb"

type LangDriver interface {
	RunJudge(*pb.Submission) (*pb.JudgeResult, error)
}

type Judge struct {
	lang map[string]LangDriver
}
