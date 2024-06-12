// gstools_bootstrap.go
package gstools

import (
	"fmt"
	"runtime"
)

// Bootstrap is responsible for bootstrapping the gstools project
func Bootstrap() {
	printLogo()

	system := runtime.GOOS
	architecture := runtime.GOARCH

	fmt.Println("System Info")
	fmt.Printf("System: %s\n", system)
	fmt.Printf("Architecture: %s\n", architecture)
	fmt.Println("=================================")
}

func printLogo() {
	fmt.Println("=================================")
	fmt.Println("   _____    _____   _______    ____     ____    _         _____ ")
	fmt.Println("  / ____|  / ____| |__   __|  / __ \\   / __ \\  | |       / ____|")
	fmt.Println(" | |  __  | (___      | |    | |  | | | |  | | | |      | (___  ")
	fmt.Println(" | | |_ |  \\___ \\     | |    | |  | | | |  | | | |       \\___ \\ ")
	fmt.Println(" | |__| |  ____) |    | |    | |__| | | |__| | | |____   ____) |")
	fmt.Println("  \\_____| |_____/     |_|     \\____/   \\____/  |______| |_____/ ")
	fmt.Println("                                                                ")
	fmt.Println("=================================")
}
