# Chalk

Chalk is a Go Package which can be used for making terminal output more vibrant with text colors, text styles and background colors.

### Documentation
Check out godoc for some example usage: http://godoc.org/github.com/golang-demos/chalk


### Usage

![Instructions](/images/Instructions.png)


## Sample Code
```go
package main

import (
	"fmt"

	"github.com/golang-demos/chalk"
)

func main() {
	fmt.Println(chalk.BlackLight("BlackLight-Text"))

	fmt.Println(chalk.Blue("Blue-Text"))

	fmt.Println(chalk.Underline("Underline-Text"))

	fmt.Println(chalk.Red().Italic("Red Italic Text"))
	fmt.Println(chalk.Green().Strikethrough("Green Strikethrough text"))
	fmt.Println(chalk.Cyan().Underline().BgBlackLight("Cyan Underline text on BlackLight Background"))
	fmt.Println(chalk.Yellow().BgRed().Inverse("Yellow text on Red background with inverted colors"))

	// For Existing Code
	fmt.Print(chalk.Green())
	fmt.Println("Data Sent Successfully")
	fmt.Print(chalk.Reset())
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

