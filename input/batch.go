package input

import "strings"

type BatchResult struct {
	Outputs []string
	Errors  []string
	Stopped bool
}

func ExecuteScript(executor Executor, script string) BatchResult {
	result := BatchResult{Outputs: make([]string, 0), Errors: make([]string, 0)}
	for _, line := range strings.Split(script, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		command, err := Parse(line)
		if err != nil {
			result.Errors = append(result.Errors, err.Error())
			continue
		}
		output, stopped, err := executor.Execute(command)
		if err != nil {
			result.Errors = append(result.Errors, err.Error())
		}
		if output != "" {
			result.Outputs = append(result.Outputs, output)
		}
		if stopped {
			result.Stopped = true
			break
		}
	}
	return result
}

func JoinOutputs(result BatchResult) string {
	return strings.Join(result.Outputs, "\n")
}

func HasFailures(result BatchResult) bool {
	return len(result.Errors) > 0
}

func CommandCount(script string) int {
	count := 0
	for _, line := range strings.Split(script, "\n") {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}
	return count
}

func NormalizeScript(script string) string {
	lines := make([]string, 0)
	for _, line := range strings.Split(script, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			lines = append(lines, trimmed)
		}
	}
	return strings.Join(lines, "\n")
}
