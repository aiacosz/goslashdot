package utils

import (
	"fmt"
)

// SetColorRed set red to a text
func SetColorRed(str string) string {
	red := "\033[0;31m"
	reset := "\033[0m"
	return red + str + reset
}

// SetColorGreen set green to a text
func SetColorGreen(str string) string {
	green := "\033[0;32m"
	reset := "\033[0m"
	return green + str + reset
}

// SetColorBlue set blue to a text
func SetColorBlue(str string) string {
	blue := "\033[0;34m"
	reset := "\033[0m"
	return blue + str + reset
}

// SetColorYela set Yellow to a text
func SetColorYela(str string) string {
	yela := "\033[0;33m"
	reset := "\033[0m"
	return yela + str + reset
}

// Banner this function return a banner of tool
func Banner() {

	fmt.Println("	▄██████▄   ▄██████▄     ▄████████  ▄█          ▄████████    ▄█    █▄    ████████▄   ▄██████▄      ███     ")
	fmt.Println("	███    ███ ███    ███   ███    ███ ███         ███    ███   ███    ███   ███   ▀███ ███    ███ ▀█████████▄ ")
	fmt.Println("	███    █▀  ███    ███   ███    █▀  ███         ███    ███   ███    ███   ███    ███ ███    ███    ▀███▀▀██ ")
	fmt.Println("	▄███       ███    ███   ███        ███         ███    ███  ▄███▄▄▄▄███▄▄ ███    ███ ███    ███     ███   ▀ ")
	fmt.Println("	▀▀███ ████ ███    ███ ▀███████████ ███       ▀███████████ ▀▀███▀▀▀▀███▀  ███    ███ ███    ███     ███")
	fmt.Println("	███    ███ ███    ███          ███ ███         ███    ███   ███    ███   ███    ███ ███    ███     ███     ")
	fmt.Println("	███    ███ ███    ███    ▄█    ███ ███▌    ▄   ███    ███   ███    ███   ███   ▄███ ███    ███     ███     ")
	fmt.Println("	████████▀   ▀██████▀   ▄████████▀  █████▄▄██   ███    █▀    ███    █▀    ████████▀   ▀██████▀     ▄████▀   ")
	fmt.Println("								   ▀                                                                       ")

	fmt.Println("by: AiacosZ")

}
