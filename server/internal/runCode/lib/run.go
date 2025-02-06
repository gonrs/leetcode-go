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
func Run(StartCode string, Code string, Tests []models.Test) (int, error) {
	testIndex := 0
	goFile := "./internal/runCode/lib/tmp/temp.go"
	err := os.WriteFile(goFile, []byte(StartCode+"\n"+Code), 0644)
	if err != nil {
		return testIndex + 1, fmt.Errorf("error open file: %v", err)
	}
	for testIndex < len(Tests) {
		cmd := exec.Command("go", "run", goFile)
		cmd.Stdin = strings.NewReader(Tests[testIndex].Input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(output)
			return testIndex + 1, fmt.Errorf("error running code: %v", err)
		}

		actualOutput := strings.TrimSpace(string(output))
		expectedOutput := strings.TrimSpace(Tests[testIndex].Output)
		if actualOutput != expectedOutput {
			return testIndex + 1, fmt.Errorf("test failed: expected %q, got %q", expectedOutput, actualOutput)
		}
		testIndex += 1
	}
	return testIndex, nil
}
