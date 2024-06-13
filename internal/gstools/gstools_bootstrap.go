package gstools

import (
	"fmt"
	"gstools/internal/gstools/config"
	"runtime"
	"time"
)

// Bootstrap is responsible for bootstrapping the gstools project
func Bootstrap() {
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("GSTools %s is starting up\n", config.Version)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("GSTools is open source, you can find it here: https://github.com/Simp-Community/gstools\n")

	fmt.Println()

	time.Sleep(1000 * time.Millisecond)

	// Print System Info
	system := runtime.GOOS
	architecture := runtime.GOARCH

	fmt.Println("=================================")
	fmt.Println("System Info")
	fmt.Printf("System: %s\n", system)
	fmt.Printf("Architecture: %s\n", architecture)
	fmt.Println("=================================")

	fmt.Println("")

	time.Sleep(500 * time.Millisecond)
}