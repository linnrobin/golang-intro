package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 1; i < 10; i++ {

		switch {
		case i == 1:
			fmt.Println("Satu")
		case i == 2:
			fmt.Println("Dua")
		case i == 3:
			fmt.Println("Tiga")
		default:
			fmt.Println("Nope")
		}

		switch i {
		case 4:
			fmt.Println("Empat")
		case 5:
			fmt.Println("Lima")
		case 6:
			fmt.Println("Enam")
		default:
			fmt.Println("NopeNope")
		}
	}

	OS := runtime.GOOS

	switch OS {
	case "darwin":
		fmt.Println("MAC")
	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Println("WIN????")
	}
}
