package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Test struct {
	Input     string `json:"input"`
	Output    string `json:"output"`
	ProblemID uint   `json:"problem_id"`
	Type      int    `json:"type"`
}

func main() {
	startCode := "package main\n\nimport (\n\t\"fmt\" \n)\n\nfunc main() {\n\tvar (\n\t\ta int\n\t\tb int\n\t)\n\tfmt.Scanln(&a, &b)\n\tfmt.Println(Code(a, b))\n}\n"
	code := "func Code(a int, b int) int{\n return a + b\n}\n"
	tests := []Test{
		{Input: "10 1", Output: "11"},
		{Input: "11 11", Output: "22"},
		{Input: "10 5", Output: "1"},
	}
	resultIndex, err := Run(startCode, code, tests)
	if err != nil {
		fmt.Printf("Test %d failed: %v\n", resultIndex, err)
	} else {
		fmt.Println("All tests passed!")
	}
}

func Run(StartCode string, Code string, Tests []Test) (int, error) {
	testIndex := 0
	goFile := "./test/go/temp.go"
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
