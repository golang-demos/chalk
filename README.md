# Chalk

Chalk is a Go Package which can be used for making terminal output more vibrant with text colors, text styles and background colors.

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
}

```


## Author
### Vinay Jeurkar

<p>
  <a href="https://www.linkedin.com/in/vinay-jeurkar/" rel="nofollow noreferrer" style="text-decoration:none;">
    <img src="https://content.linkedin.com/content/dam/me/business/en-us/amp/brand-site/v2/bg/LI-Bug.svg.original.svg" height="12px" alt="linkedin"> LinkedIn
  </a> &nbsp; 
  <a href="https://github.com/vinay03" rel="nofollow noreferrer" style="text-decoration:none;">
    <img src="https://github.githubassets.com/images/modules/logos_page/Octocat.png" height="12px" alt="github"> Github
  </a>
</p>
