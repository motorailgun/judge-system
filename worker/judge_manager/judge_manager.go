package judge_manager

import (
	"context"
	"worker/dist_worker"
	"worker/drivers"
	"worker/pb"
)

type JudgeManager struct {
	driver drivers.LanguageDriver
	dist_worker.UnimplementedLangServiceServer
}

func NewJudgeManager(lang string) *JudgeManager {
	return &JudgeManager{
		driver: drivers.LangServiceTable[lang],
	}
}

func (d *JudgeManager) RunJudge(ctx context.Context, submission *pb.Submission) (*pb.JudgeResult, error) {
	jtype := 
}
