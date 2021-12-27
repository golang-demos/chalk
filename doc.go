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
//   fmt.Println(chalk.Black(), "Black", chalk.Reset())
//   fmt.Println(chalk.BlackLight(), "BlackLight", chalk.Reset())
//   fmt.Println(chalk.Red(), "Red", chalk.Reset())
//   fmt.Println(chalk.RedLight(), "RedLight", chalk.Reset())
//   fmt.Println(chalk.Green(), "Green", chalk.Reset())
//   fmt.Println(chalk.GreenLight(), "GreenLight", chalk.Reset())
//   fmt.Println(chalk.Orange(), "Orange", chalk.Reset())
//   fmt.Println(chalk.OrangeLight(), "OrangeLight", chalk.Reset())
//   fmt.Println(chalk.Yellow(), "Yellow", chalk.Reset())
//   fmt.Println(chalk.YellowLight(), "YellowLight", chalk.Reset())
//   fmt.Println(chalk.Blue(), "Blue", chalk.Reset())
//   fmt.Println(chalk.BlueLight(), "BlueLight", chalk.Reset())
//   fmt.Println(chalk.Purple(), "Purple", chalk.Reset())
//   fmt.Println(chalk.PurpleLight(), "PurpleLight", chalk.Reset())
//   fmt.Println(chalk.Violet(), "Violet", chalk.Reset())
//   fmt.Println(chalk.VioletLight(), "VioletLight", chalk.Reset())
//   fmt.Println(chalk.Cyan(), "Cyan", chalk.Reset())
//   fmt.Println(chalk.CyanLight(), "CyanLight", chalk.Reset())
//   fmt.Println(chalk.White(), "White", chalk.Reset())
//   fmt.Println(chalk.WhiteLight(), "WhiteLight", chalk.Reset())
//
//
// Text Style
//
// There are seven text styles. You can use them as follow:
//
//   fmt.Println(chalk.Bold(), "Bold", chalk.Reset())
//   fmt.Println(chalk.Dim(), "Dim", chalk.Reset())
//   fmt.Println(chalk.Italic(), "Italic", chalk.Reset())
//   fmt.Println(chalk.Underline(), "Underline", chalk.Reset())
//   fmt.Println(chalk.Inverse(), "Inverse", chalk.Reset())
//   fmt.Println(chalk.Hidden(), "Hidden", chalk.Reset())
//   fmt.Println(chalk.Strikethrough(), "Strikethrough", chalk.Reset())
//
//
// Background Color
//
// There are sixteen background colors. You can use them as follow:
//
//   fmt.Println(chalk.BgBlack(), "BgBlack", chalk.Reset())
//   fmt.Println(chalk.BgBlackLight(), "BgBlackLight", chalk.Reset())
//   fmt.Println(chalk.BgRed(), "BgRed", chalk.Reset())
//   fmt.Println(chalk.BgRedLight(), "BgRedLight", chalk.Reset())
//   fmt.Println(chalk.BgGreen(), "BgGreen", chalk.Reset())
//   fmt.Println(chalk.BgGreenLight(), "BgGreenLight", chalk.Reset())
//   fmt.Println(chalk.BgYellow(), "BgYellow", chalk.Reset())
//   fmt.Println(chalk.BgYellowLight(), "BgYellowLight", chalk.Reset())
//   fmt.Println(chalk.BgBlue(), "BgBlue", chalk.Reset())
//   fmt.Println(chalk.BgBlueLight(), "BgBlueLight", chalk.Reset())
//   fmt.Println(chalk.BgPurple(), "BgPurple", chalk.Reset())
//   fmt.Println(chalk.BgPurpleLight(), "BgPurpleLight", chalk.Reset())
//   fmt.Println(chalk.BgCyan(), "BgCyan", chalk.Reset())
//   fmt.Println(chalk.BgCyanLight(), "BgCyanLight", chalk.Reset())
//   fmt.Println(chalk.BgWhite(), "BgWhite", chalk.Reset())
//   fmt.Println(chalk.BgWhiteLight(), "BgWhiteLight", chalk.Reset())
//
// You can use above functions in combinations. Following are some
// examples to refer:
//
//   fmt.Println(chalk.Red().Italic(), "Red + Italic", chalk.Reset())
//   fmt.Println(chalk.Green().Strikethrough(), "Green + Strikethrough", chalk.Reset())
//   fmt.Println(chalk.Cyan().Underline().BgBlackLight(), "Cyan + Underline + BgBlackLight", chalk.Reset())
//   fmt.Println(chalk.Yellow().BgRed().Inverse(), "Yellow + BgRed + Inverse", chalk.Reset())
//
package chalk
