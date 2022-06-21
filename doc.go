// Chalk is a Go Package which can be used for making terminal output more vibrant with text colors, text styles and background colors.
// The styling factors are:
// - Text Color
// - Text Style
// - Background Color
//
//
// Text Color
//
// There are sixteen Colors. You can use them as follow :
//
//   fmt.Println(chalk.Black("Black-Text"))
//   fmt.Println(chalk.BlackLight("BlackLight-Text"))
//   fmt.Println(chalk.Red("Red-Text"))
//   fmt.Println(chalk.RedLight("RedLight-Text"))
//   fmt.Println(chalk.Green("Green-Text"))
//   fmt.Println(chalk.GreenLight("GreenLight-Text"))
//   fmt.Println(chalk.Yellow("Yellow-Text"))
//   fmt.Println(chalk.YellowLight("YellowLight-Text"))
//   fmt.Println(chalk.Blue("Blue-Text"))
//   fmt.Println(chalk.BlueLight("BlueLight-Text"))
//   fmt.Println(chalk.Magenta("Magenta-Text"))
//   fmt.Println(chalk.MagentaLight("MagentaLight-Text"))
//   fmt.Println(chalk.Cyan("Cyan-Text"))
//   fmt.Println(chalk.CyanLight("CyanLight-Text"))
//   fmt.Println(chalk.White("White-Text"))
//   fmt.Println(chalk.WhiteLight("WhiteLight-Text"))
//
//
// Text Style
//
// There are seven text styles. You can use them as follow:
//
//   fmt.Println(chalk.Bold("Bold-Text"))
//   fmt.Println(chalk.Dim("Dim-Text"))
//   fmt.Println(chalk.Italic("Italic-Text"))
//   fmt.Println(chalk.Underline("Underline-Text"))
//   fmt.Println(chalk.Inverse("Inverse-Text"))
//   fmt.Println(chalk.Hidden("Hidden-Text"))
//   fmt.Println(chalk.Strikethrough("Strikethrough-Text"))
//
//
// Background Color
//
// There are sixteen background colors. You can use them as follow:
//
//   fmt.Println(chalk.WhiteLight().BgBlack("WhiteLight-Text-on-BgBlack"))
//   fmt.Println(chalk.WhiteLight().BgBlackLight("WhiteLight-Text-on-BgBlackLight"))
//   fmt.Println(chalk.Black().BgRed("Black-Text-on-BgRed"))
//   fmt.Println(chalk.Black().BgRedLight("Black-Text-on-BgRedLight"))
//   fmt.Println(chalk.Black().BgGreen("Black-Text-on-BgGreen"))
//   fmt.Println(chalk.Black().BgGreenLight("Black-Text-on-BgGreenLight"))
//   fmt.Println(chalk.Black().BgYellow("Black-Text-on-BgYellow"))
//   fmt.Println(chalk.Black().BgYellowLight("Black-Text-on-BgYellowLight"))
//   fmt.Println(chalk.Black().BgBlue("Black-Text-on-BgBlue"))
//   fmt.Println(chalk.Black().BgBlueLight("Black-Text-on-BgBlueLight"))
//   fmt.Println(chalk.Black().BgMagenta("Black-Text-on-BgMagenta"))
//   fmt.Println(chalk.Black().BgMagentaLight("Black-Text-on-BgMagentaLight"))
//   fmt.Println(chalk.Black().BgCyan("Black-Text-on-BgCyan"))
//   fmt.Println(chalk.Black().BgCyanLight("Black-Text-on-BgCyanLight"))
//   fmt.Println(chalk.Black().BgWhite("Black-Text-on-BgWhite"))
//   fmt.Println(chalk.Black().BgWhiteLight("Black-Text-on-BgWhiteLight"))
//
// You can use above functions in combinations. Following are some
// examples to refer:
//
//   fmt.Println(chalk.Red().Italic("Red-Italic-Text"))
//   fmt.Println(chalk.Green().Strikethrough("Green-Strikethrough-Text"))
//   fmt.Println(chalk.Cyan().Underline().BgBlackLight("Cyan-underline-text-on-BgBlackLight"))
//   fmt.Println(chalk.Yellow().BgRed().Inverse("Yellow-text-on-red-inverted"))
//
// You can save formatting configuration in a variable and reuse them multiple times.
// examples to refer:
//
//   SuccessMessage := chalk.Green().Bold("SUCCESS : ")
//   WarningMessage := chalk.YellowLight().Bold("WARNING : ")
//   ErrorMessage := chalk.RedLight().Bold("ERROR   : ")
//   fmt.Println(SuccessMessage.Apply("Completed successfully"))
//   fmt.Println(SuccessMessage.Apply("Process Complete"))
//   fmt.Println(WarningMessage.Apply("Call Deprecated"))
//   fmt.Println(ErrorMessage.Apply("Fatal error ocurred"))
//
package chalk
