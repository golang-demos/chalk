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
	fmt.Println(chalk.BlackLight(), "chalk.BlackLight()", chalk.Reset())

	fmt.Println(chalk.Blue(), "chalk.Blue()", chalk.Reset())

	fmt.Println(chalk.Underline(), "chalk.Underline()", chalk.Reset())

	fmt.Println(chalk.Red().Italic(), "chalk.Red().Italic()", chalk.Reset())
	fmt.Println(chalk.Green().Strikethrough(), "chalk.Green().Strikethrough()", chalk.Reset())
	fmt.Println(chalk.Cyan().Underline().BgBlackLight(), "chalk.Cyan().Underline().BgBlackLight()", chalk.Reset())
	fmt.Println(chalk.Yellow().BgRed().Inverse(), "chalk.Yellow().BgRed().Inverse()", chalk.Reset())

	// For Existing Code
	fmt.Print(chalk.Green())
	fmt.Prinln("Data Sent Successfully")
	fmt.Print(chalk.Reset())
}

```


## Author
### Vinay Jeurkar

<p>
  <a href="https://www.linkedin.com/in/vinay-jeurkar/" rel="nofollow noreferrer" style="text-decoration:none;">
    LinkedIn
  </a> &nbsp; 
  <a href="https://github.com/vinay03" rel="nofollow noreferrer" style="text-decoration:none;">
    Github
  </a>
</p>


## Source
- https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html