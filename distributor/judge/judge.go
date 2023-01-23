package judge

import (
	"distributor/pb"
)

var languages []string = []string{
	"nodejs", "python", "cpp", "ruby",
}

type Judger struct {
	drivers map[string]LangDriver
}

type LangDriver interface {
	Judge(*pb.Job) (*pb.JudgeResult, error)
}

func NewJudger() Judger {
	judger := Judger{make(map[string]LangDriver)}

	return judger
}

func (j *Judger) Judge(job *pb.Job) (*pb.JudgeResult, error) {
	lang_driver := j.drivers[job.Submission.Language.String()]
	if lang_driver != nil {
		return {
			
		}
	}
}
