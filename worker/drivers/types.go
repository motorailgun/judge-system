package drivers

import (
	"worker/pb"
)

const PROBLEM_FILES_PATH string = "/judge/problems/"

type LanguageDriver interface {
	Run(problem_id uint, code string) pb.JudgeResult
}

var LangServiceTable = map[string]LanguageDriver{
	"nodejs": NewNodejsDriver(),
	"cpp":    NewCppDriver(),
	"python": NewPythonDriver(),
	"ruby":   NewRubyDriver(),
}
