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
func Run(StartCode string, Code string, Tests []models.Test) (error, int, string) {
	testIndex := 0
	goFile := "./internal/runCode/lib/tmp/temp.go"
	err := os.WriteFile(goFile, []byte(StartCode+"\n"+Code), 0644)
	if err != nil {
		return fmt.Errorf("error open file: %v", err), testIndex, ""
	}
	for testIndex < len(Tests) {
		cmd := exec.Command("go", "run", goFile)
		cmd.Stdin = strings.NewReader(Tests[testIndex].Input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(output)
			return fmt.Errorf("error running code: %v", err), testIndex, ""
		}

		actualOutput := strings.TrimSpace(string(output))
		expectedOutput := strings.TrimSpace(Tests[testIndex].Output)
		if actualOutput != expectedOutput {
			// return fmt.Errorf("%v/%v test failed: expected %q, got %q", testIndex+1, len(Tests), expectedOutput, actualOutput), testIndex, actualOutput
			return fmt.Errorf("expected %q, got %q", expectedOutput, actualOutput), testIndex, actualOutput
		}
		testIndex += 1
	}
	return nil, testIndex, ""
}
