package run

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gonrs/leetcode-go/common/models"
)

// func Run(Code string, Tests []models.Test) (int, error) {
// 	testIndex := 0
// 	for testIndex < len(Tests) {

//			testIndex += 1
//		}
//		return testIndex, nil
//	}
func Run(LanguageCode string, StartCode string, Code string, Tests []models.Test) (int, string, error) {
	testIndex := 0
	file := "./internal/runCode/lib/tmp/temp.go"
	if LanguageCode == "py" {
		file = "./internal/runCode/lib/tmp/temp.py"
	}
	newCode := StartCode + "\n" + Code
	if LanguageCode == "py" {
		newCode = Code + "\n" + StartCode
	}
	err := os.WriteFile(file, []byte(newCode), 0644)
	if err != nil {
		return testIndex, "", fmt.Errorf("error open file: %v", err)
	}
	for testIndex < len(Tests) {
		var cmd *exec.Cmd
		if LanguageCode == "go" {
			cmd = exec.Command("go", "run", file)
		} else if LanguageCode == "py" {
			cmd = exec.Command("python", file)
		}
		cmd.Stdin = strings.NewReader(Tests[testIndex].Input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			// fmt.Println(output)
			return testIndex, "", fmt.Errorf("error running code: %v", err)
		}

		actualOutput := strings.TrimSpace(string(output))
		expectedOutput := strings.TrimSpace(Tests[testIndex].Output)
		if actualOutput != expectedOutput {
			// return fmt.Errorf("%v/%v test failed: expected %q, got %q", testIndex+1, len(Tests), expectedOutput, actualOutput), testIndex, actualOutput
			return testIndex, actualOutput, fmt.Errorf("expected %q, got %q", expectedOutput, actualOutput)
		}
		testIndex += 1
	}
	return testIndex, "", nil
}
