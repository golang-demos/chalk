package main

import (
	"fmt"

	"github.com/golang-demos/chalk"
)

func main() {
	fmt.Println("--", chalk.Red(), "Red + Regular", chalk.Reset(), "--")
	fmt.Println("--", chalk.Red().Bold(), "Red + Bold", chalk.Reset(), "--")
	fmt.Println("--", chalk.Red().Italic(), "Red + Italic", chalk.Reset(), "--")

	fmt.Println("--", chalk.White().BgRed(), "Red + Regular", chalk.Reset(), "--")

	fmt.Println("--", chalk.White().Bold().BgRed(), "White + Bold + BgRed", chalk.Reset(), "--")
	fmt.Println("--", chalk.White().Bold().BgRed(true), "White + Bold + BgRed(I)", chalk.Reset(), "--")

	fmt.Println("--", chalk.Red().BgWhite(true), "Red + Regular + BgWhite(I)", chalk.Reset(), "--")
	fmt.Println("--", chalk.Red(true).BgWhite(true), "Red(I) + Regular + BgWhite(I)", chalk.Reset(), "--")

	fmt.Println("--", chalk.Black().Italic().BgGreen(), "Black + Italic + BgGreen", chalk.Reset(), "--")
	fmt.Println("--", chalk.Black().Italic().BgGreen(true), "Black + Italic + BgGreen(I)", chalk.Reset(), "--")

	fmt.Println("--", chalk.Cyan().BgBlack(true), "Cyan + Regular + BgBlack(I)", chalk.Reset(), "--")
	fmt.Println("--", chalk.Cyan(true).BgBlack(true), "Cyan(I) + Regular + BgBlack(I)", chalk.Reset(), "--")
}
