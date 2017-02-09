package splashes

import "fmt"

// Splash basically it grabs a splash from splashes and prints it out
func Splash(s string) {
	switch s {
	case "forge":
		fmt.Println(FORGE)
	}
}
