package judge

import (
	"distributor/pb"
	"distributor/worker_client"
	"fmt"
	"os"
	"strings"
)

var languages []string = []string{
	"nodejs", "python", "cpp", "ruby",
}

type Judger struct {
	ports map[string]string
}

type LangDriver interface {
	Judge(*pb.Job) (*pb.JudgeResult, error)
}

func NewJudger() Judger {
	judger := Judger{make(map[string]string)}
	for _, lang := range languages {
		judger.ports[lang] = os.Getenv(strings.ToUpper(lang) + "_PORT")
	}

	return judger
}

func (j *Judger) Judge(submission *pb.Submission) (*pb.JudgeResult, error) {
	port := j.ports[submission.Language.String()]
	if port == "" {
		fmt.Fprintf(os.Stderr, "got unknown language: %s: possibly not implemented yet?\n", submission.Language.String())

		return &pb.JudgeResult{
			ResultCode: pb.JudgeInfo_err,
			Submission: submission,
		}, nil
	}

	client := worker_client.NewWorkerClient(port)
	return client.RunJudge(submission)
}
