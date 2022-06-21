# Chalk

Chalk is a Go Package which can be used for making terminal output more vibrant with text colors, text styles and background colors.

### Documentation
Check out godoc for some example usage: http://godoc.org/github.com/golang-demos/chalk


### Sample Output
![Instructions](/images/Instructions.png)


## Colors
| Colors       | TextColor       | Background-Color       |
|--------------|--------------|--------------|
| Black        | `chalk.Black("Black-Text")` | `chalk.BgBlack("Black-Background")` |
| BlackLight   | `chalk.BlackLight("BlackLight-Text")` | `chalk.BgBlackLight("BlackLight-Background")` |
| Red          | `chalk.Red("Red-Text")` | `chalk.BgRed("Red-Background")` |
| RedLight     | `chalk.RedLight("RedLight-Text")` | `chalk.BgRedLight("RedLight-Background")` |
| Green        | `chalk.Green("Green-Text")` | `chalk.BgGreen("Green-Background")` |
| GreenLight   | `chalk.GreenLight("GreenLight-Text")` | `chalk.BgGreenLight("GreenLight-Background")` |
| Yellow       | `chalk.Yellow("Yellow-Text")` | `chalk.BgYellow("Yellow-Background")` |
| YellowLight  | `chalk.YellowLight("YellowLight-Text")` | `chalk.BgYellowLight("YellowLight-Background")` |
| Blue         | `chalk.Blue("Blue-Text")` | `chalk.BgBlue("Blue-Background")` |
| BlueLight    | `chalk.BlueLight("BlueLight-Text")` | `chalk.BgBlueLight("BlueLight-Background")` |
| Magenta      | `chalk.Magenta("Magenta-Text")` | `chalk.BgMagenta("Magenta-Background")` |
| MagentaLight | `chalk.MagentaLight("MagentaLight-Text")` | `chalk.BgMagentaLight("MagentaLight-Background")` |
| Cyan         | `chalk.Cyan("Cyan-Text")` | `chalk.BgCyan("Cyan-Background")` |
| CyanLight    | `chalk.CyanLight("CyanLight-Text")` | `chalk.BgCyanLight("CyanLight-Background")` |
| White        | `chalk.White("White-Text")` | `chalk.BgWhite("White-Background")` |
| WhiteLight   | `chalk.WhiteLight("WhiteLight-Text")` | `chalk.BgWhiteLight("WhiteLight-Background")` |


## Basic Formatting
| Formatting    | Use          |
|---------------|--------------|
| Bold  | `chalk.Bold("Bold-Text")` |
| Dim  | `chalk.Dim("Dim-Text")` |
| Italic  | `chalk.Italic("Italic-Text")` |
| Underline  | `chalk.Underline("Underline-Text")` |
| Inverse  | `chalk.Inverse("Inverse-Text")` |
| Hidden  | `chalk.Hidden("Hidden-Text")` |
| Strikethrough  | `chalk.Strikethrough("Strikethrough-Text")` |


## Sample Code
```go
package main

import (
	"fmt"

	"github.com/golang-demos/chalk"
)

func main() {
	// Colors
	fmt.Println(chalk.BlackLight("BlackLight-Text"))
	fmt.Println(chalk.Blue("Blue-Text"))

	// Basic Formatting
	fmt.Println(chalk.Underline("Underline-Text"))
	fmt.Println(chalk.Bold("Bold-Text"))

	// Background Coloring
	fmt.Println(chalk.BgYellow("Yellow-Background-Text"))

	// Mixed
	fmt.Println(chalk.Red().Italic("Red Italic Text"))
	fmt.Println(chalk.Green().Strikethrough("Green Strikethrough text"))
	fmt.Println(chalk.Cyan().Underline().BgBlackLight("Cyan Underline text on BlackLight Background"))
	fmt.Println(chalk.Yellow().BgRed().Inverse("Yellow text on Red background with inverted colors"))

	// For Existing Code
	fmt.Print(chalk.Green())
	fmt.Println("Data Sent Successfully")
	fmt.Print(chalk.Reset())

	// Reusable Configurations
	SuccessMessage := chalk.Green().Bold("SUCCESS : ")
	fmt.Println(SuccessMessage.Apply("Completed successfully"))
	fmt.Println(SuccessMessage.Apply("Process Complete"))
}

```


## Source
- https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html


## Author
### Vinay Jeurkar

<p>
  <a href="https://www.linkedin.com/in/vinay-jeurkar/" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white" /></a> 
	&nbsp; 
  <a href="https://github.com/vinay03" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white" /></a> 
	&nbsp; 
  <a href="https://twitter.com/Vinay_Jeurkar" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/Twitter-1DA1F2?style=flat&logo=twitter&logoColor=white" /></a>
</p>

