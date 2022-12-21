package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	createFile()

	//TODO: exporting metrics to prometheus
}

func createFile() {
	usr, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err)
	}
	filePath := filepath.Join(usr, "Desktop")

	// createfile, err := os.OpenFile(filepath.Join(filePath, "output.txt"), os.O_RDWR|os.O_CREATE, 0644)
	absPath := filepath.Join(filePath, "output.txt")
	createfile, err := os.Create(absPath)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(createfile)

	populater(usr, filePath, absPath)
}

// populates output.txt file with metrics from top utility
// this function extracts pid,cpu metrics but it can track other metrics as well, to see options available try `top --help`
func populater(usr, filePath, absPath string) {
	getrawstats := exec.Command("top", "-l1", "-n10", "-stats", "pid,cpu")
	output, err := getrawstats.CombinedOutput()
	if err != nil {
		fmt.Print(err)
	}
	// fmt.Print(string(output))

	rawpopulate := os.WriteFile(absPath, output, 0755)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(rawpopulate)

	skimmedstats := exec.Command("sed", "1,10d", absPath)
	finaloutput, err := skimmedstats.CombinedOutput()
	if err != nil {
		fmt.Print(err)
	}

	finalpopulate := os.WriteFile(absPath, finaloutput, 0755)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(finalpopulate)

}
