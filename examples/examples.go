package main

import (
	"fmt"

	"github.com/golang-demos/chalk"
)

func main() {
	// Colors
	fmt.Println("\n> Colors")
	fmt.Println(chalk.Black(), "chalk.Black()", chalk.Reset())
	fmt.Println(chalk.BlackLight(), "chalk.BlackLight()", chalk.Reset())
	fmt.Println(chalk.Red(), "chalk.Red()", chalk.Reset())
	fmt.Println(chalk.RedLight(), "chalk.RedLight()", chalk.Reset())
	fmt.Println(chalk.Green(), "chalk.Green()", chalk.Reset())
	fmt.Println(chalk.GreenLight(), "chalk.GreenLight()", chalk.Reset())
	fmt.Println(chalk.Yellow(), "chalk.Yellow()", chalk.Reset())
	fmt.Println(chalk.YellowLight(), "chalk.YellowLight()", chalk.Reset())
	fmt.Println(chalk.Blue(), "chalk.Blue()", chalk.Reset())
	fmt.Println(chalk.BlueLight(), "chalk.BlueLight()", chalk.Reset())
	fmt.Println(chalk.Magenta(), "chalk.Magenta()", chalk.Reset())
	fmt.Println(chalk.MagentaLight(), "chalk.MagentaLight()", chalk.Reset())
	fmt.Println(chalk.Cyan(), "chalk.Cyan()", chalk.Reset())
	fmt.Println(chalk.CyanLight(), "chalk.CyanLight()", chalk.Reset())
	fmt.Println(chalk.White(), "chalk.White()", chalk.Reset())
	fmt.Println(chalk.WhiteLight(), "chalk.WhiteLight()", chalk.Reset())

	// Style
	fmt.Println("\n> Text Style")
	fmt.Println(chalk.Bold(), "chalk.Bold()", chalk.Reset())
	fmt.Println(chalk.Dim(), "chalk.Dim()", chalk.Reset())
	fmt.Println(chalk.Italic(), "chalk.Italic()", chalk.Reset())
	fmt.Println(chalk.Underline(), "chalk.Underline()", chalk.Reset())
	fmt.Println(chalk.Inverse(), "chalk.Inverse()", chalk.Reset())
	fmt.Println(chalk.Hidden(), "chalk.Hidden()", chalk.Reset())
	fmt.Println(chalk.Strikethrough(), "chalk.Strikethrough()", chalk.Reset())

	// Background Colors
	fmt.Println("\n> Background Colors")
	fmt.Println(chalk.BgBlack(), "chalk.BgBlack()", chalk.Reset())
	fmt.Println(chalk.BgBlackLight(), "chalk.BgBlackLight()", chalk.Reset())
	fmt.Println(chalk.Black().BgRed(), "chalk.BgRed()", chalk.Reset())
	fmt.Println(chalk.Black().BgRedLight(), "chalk.BgRedLight()", chalk.Reset())
	fmt.Println(chalk.Black().BgGreen(), "chalk.BgGreen()", chalk.Reset())
	fmt.Println(chalk.Black().BgGreenLight(), "chalk.BgGreenLight()", chalk.Reset())
	fmt.Println(chalk.Black().BgYellow(), "chalk.BgYellow()", chalk.Reset())
	fmt.Println(chalk.Black().BgYellowLight(), "chalk.BgYellowLight()", chalk.Reset())
	fmt.Println(chalk.Black().BgBlue(), "chalk.BgBlue()", chalk.Reset())
	fmt.Println(chalk.Black().BgBlueLight(), "chalk.BgBlueLight()", chalk.Reset())
	fmt.Println(chalk.Black().BgMagenta(), "chalk.BgMagenta()", chalk.Reset())
	fmt.Println(chalk.Black().BgMagentaLight(), "chalk.BgMagentaLight()", chalk.Reset())
	fmt.Println(chalk.Black().BgCyan(), "chalk.BgCyan()", chalk.Reset())
	fmt.Println(chalk.Black().BgCyanLight(), "chalk.BgCyanLight()", chalk.Reset())
	fmt.Println(chalk.Black().BgWhite(), "chalk.BgWhite()", chalk.Reset())
	fmt.Println(chalk.Black().BgWhiteLight(), "chalk.BgWhiteLight()", chalk.Reset())

	// Mixed
	fmt.Println("\n> Mixed")
	fmt.Println(chalk.Red().Italic(), "chalk.Red().Italic()", chalk.Reset())
	fmt.Println(chalk.Green().Strikethrough(), "chalk.Green().Strikethrough()", chalk.Reset())
	fmt.Println(chalk.Cyan().Underline().BgBlackLight(), "chalk.Cyan().Underline().BgBlackLight()", chalk.Reset())
	fmt.Println(chalk.Yellow().BgRed().Inverse(), "chalk.Yellow().BgRed().Inverse()", chalk.Reset())

}
