package gstools

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	environmentCheck()
}

// environmentCheck ensures that the Go version is 1.16+
func environmentCheck() {
	const requiredGoVersion = 1.16
	currentGoVersion := float64(runtime.Version()[2]-'0') + 0.1*float64(runtime.Version()[4]-'0')

	if currentGoVersion < requiredGoVersion {
		fmt.Printf("Go 1.16+ is needed to run gstools\n")
		fmt.Printf("Current Go version %s is too old\n", runtime.Version())
		os.Exit(1)
	}
}

// Entrypoint is the main entry point for gstools
func Entrypoint() {
	for {
		Bootstrap()
		fmt.Println("Do you want to restart the program? (yes/no)")
		var choice string
		fmt.Scanln(&choice)
		if choice == "no" {
			break
		}
	}

	// CLI dispatch logic goes here
	cliDispatch()
}

func cliDispatch() {
	// Placeholder for CLI dispatch logic
	fmt.Println("CLI dispatch completed")
}
