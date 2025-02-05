package main

import (
	"bytes"
	"fmt"
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
	code := `package main

	import (
		"fmt"
		"os"
	)
	
	func main() {
		var input string
		fmt.Scanln(&input) // Чтение входных данных
		fmt.Println(input) // Вывод входных данных
		fmt.Println("Special characters: \" \\ \\n \\t")
	}`
	tests := []Test{
		{Input: "10", Output: "10\nSpecial characters: \" \\ \\n \\t\n"},
		{Input: "11", Output: "11\nSpecial characters: \" \\ \\n \\t\n"},
	}
	resultIndex, err := Run(code, tests)
	if err != nil {
		fmt.Printf("Test %d failed: %v\n", resultIndex, err)
	} else {
		fmt.Println("All tests passed!")
	}
}

func Run(Code string, Tests []Test) (int, error) {
	testIndex := 0
	for testIndex < len(Tests) {
		input := Tests[testIndex].Input
		expectedOutput := Tests[testIndex].Output

		cmd := exec.Command("go", "run", "./test/run/code.go")
		cmd.Stdin = strings.NewReader(input)

		var out bytes.Buffer
		cmd.Stdout = &out
		var errOut bytes.Buffer
		cmd.Stderr = &errOut

		err := cmd.Run()
		if err != nil {
			fmt.Println(out.String())
			return testIndex + 1, fmt.Errorf("error running code: %v", err)
		}

		returnedOutput := out.String()
		if returnedOutput != expectedOutput {
			return testIndex + 1, fmt.Errorf("expected: %s, returned: %s", expectedOutput, returnedOutput)
		}
		testIndex += 1
	}
	return testIndex, nil
}
