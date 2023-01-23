package drivers

import (
	"worker/pb"
)

type LanguageDriver interface {
	Run(problem_id uint, code string) pb.JudgeResult
}

var LangServiceTable = map[string]LanguageDriver{
	"nodejs": NewNodejsDriver(),
	"cpp":    NewCppDriver(),
	"python": NewPythonDriver(),
	"ruby":   NewRubyDriver(),
}
