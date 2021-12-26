package main

import (
	"fmt"

	"github.com/golang-demos/chalk"
)

func main() {
	// chalk.LIBRARY_DEVELOPER_MODE = true

	// Colors
	fmt.Println("--", chalk.Black(), "Black", chalk.Reset(), "--")
	fmt.Println("--", chalk.Red(), "Red", chalk.Reset(), "--")
	fmt.Println("--", chalk.Green(), "Green", chalk.Reset(), "--")
	fmt.Println("--", chalk.Orange(), "Orange", chalk.Reset(), "--")
	fmt.Println("--", chalk.Yellow(), "Yellow", chalk.Reset(), "--")
	fmt.Println("--", chalk.Blue(), "Blue", chalk.Reset(), "--")
	fmt.Println("--", chalk.Purple(), "Purple", chalk.Reset(), "--")
	fmt.Println("--", chalk.Violet(), "Violet", chalk.Reset(), "--")
	fmt.Println("--", chalk.Cyan(), "Cyan", chalk.Reset(), "--")
	fmt.Println("--", chalk.White(), "White", chalk.Reset(), "--")

	// Style
	fmt.Println("--", chalk.Bold(), "Bold", chalk.Reset(), "--")
	fmt.Println("--", chalk.Dim(), "Dim", chalk.Reset(), "--")
	fmt.Println("--", chalk.Italic(), "Italic", chalk.Reset(), "--")
	fmt.Println("--", chalk.Underline(), "Underline", chalk.Reset(), "--")
	fmt.Println("--", chalk.Inverse(), "Inverse", chalk.Reset(), "--")
	fmt.Println("--", chalk.Hidden(), "Hidden", chalk.Reset(), "--")
	fmt.Println("--", chalk.Strikethrough(), "Strikethrough", chalk.Reset(), "--")

	// Background Colors
	fmt.Println("--", chalk.BgBlack(), "BgBlack", chalk.Reset(), "--")
	fmt.Println("--", chalk.BgRed(), "BgRed", chalk.Reset(), "--")
	fmt.Println("--", chalk.BgGreen(), "BgGreen", chalk.Reset(), "--")
	fmt.Println("--", chalk.BgYellow(), "BgYellow", chalk.Reset(), "--")
	fmt.Println("--", chalk.BgBlue(), "BgBlue", chalk.Reset(), "--")
	fmt.Println("--", chalk.BgPurple(), "BgPurple", chalk.Reset(), "--")
	fmt.Println("--", chalk.BgCyan(), "BgCyan", chalk.Reset(), "--")
	fmt.Println("--", chalk.BgWhite(), "BgWhite", chalk.Reset(), "--")

	// Misc
	/* fmt.Println("--", chalk.Red(), "Red + Regular", chalk.Reset(), "--")
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

	fmt.Println("--", chalk.Red().Bold(), "My Name is ", chalk.Green(), "Vinay", chalk.Reset()) */
}
