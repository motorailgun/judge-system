package judge

import (
	"os"
	"strings"
)

var languages []string = []string{
	"nodejs", "python", "cpp", "ruby",
}

type Judger struct {
	port map[string]string
}

func NewJudger() Judger {
	judger := Judger{make(map[string]string)}
	for _, lang := range languages {
		env_key := strings.ToUpper(lang) + "_port"

		judger.port[lang] = os.Getenv(env_key)
	}

	return judger
}
