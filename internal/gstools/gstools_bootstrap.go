// gstools_bootstrap.go
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

	printLogo()

	time.Sleep(1000 * time.Millisecond)

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

func printLogo() {
	fmt.Println("   _____    _____   _______    ____     ____    _         _____ ")
	fmt.Println("  / ____|  / ____| |__   __|  / __ \\   / __ \\  | |       / ____|")
	fmt.Println(" | |  __  | (___      | |    | |  | | | |  | | | |      | (___  ")
	fmt.Println(" | | |_ |  \\___ \\     | |    | |  | | | |  | | | |       \\___ \\ ")
	fmt.Println(" | |__| |  ____) |    | |    | |__| | | |__| | | |____   ____) |")
	fmt.Println("  \\_____| |_____/     |_|     \\____/   \\____/  |______| |_____/ ")
	fmt.Println("                                                                ")
}
