package main

import (
	"fmt"

	"github.com/golang-demos/chalk"
)

func main() {
	// chalk.LIBRARY_DEVELOPER_MODE = true

	fmt.Println("--", chalk.Red(), "Red + Regular", chalk.Reset(), "--")
	fmt.Println("--", chalk.Red().Bold(), "Red + Bold", chalk.Reset(), "--")
	fmt.Println("--", chalk.Red().Italic(), "Red + Italic", chalk.Reset(), "--")

	fmt.Println("--", chalk.White().BgRed(), "Red + Regular", chalk.Reset(), "--")

	fmt.Println("--", chalk.White().Bold().BgRed(), "White + Bold + BgRed", chalk.Reset(), "--")
	fmt.Println("--", chalk.White().Bold().BgRed(true), "White + Bold + BgRed(I)", chalk.Reset(), "--")

	fmt.Println("--", chalk.Red().BgBlue(), "Red + Regular + BgBlue(I)", chalk.Reset(), "--")
	fmt.Println("--", chalk.Red().BgBlue(), "Red(I) + Regular + BgBlue(I)", chalk.Reset(), "--")

	fmt.Println("--", chalk.Black().Italic().BgGreen(), "Black + Italic + BgGreen", chalk.Reset(), "--")
	fmt.Println("--", chalk.Black().Italic().BgGreen(true), "Black + Italic + BgGreen(I)", chalk.Reset(), "--")

	fmt.Println("--", chalk.Cyan().BgBlack(true), "Cyan + Regular + BgBlack(I)", chalk.Reset(), "--")
	fmt.Println("--", chalk.Cyan(true).BgBlack(true), "Cyan(I) + Regular + BgBlack(I)", chalk.Reset(), "--")

	fmt.Println(chalk.Red().Bold(), "My Name is ", chalk.Green(), "Vinay", chalk.Reset())
}
