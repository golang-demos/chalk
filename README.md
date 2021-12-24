# Chalk

Chalk is a Go Package which can be used for making terminal output more vibrant

## API Syntax
```
chalk.[<TextColor>()][.<TextStyle>()][.<BackgroundColor>()]
```

For e.g.
```
chalk.Red().Bold().BgWhite()
```

Use following functions as per your need.

### Text Colors
- `Black()`
- `Red()`
- `Green()`
- `Orange()`
- `Yellow()`
- `Blue()`
- `Purple()`
- `Violet()`
- `Cyan()`
- `White()`


### Text Styles
- `Bold()`
- `Dim()`
- `Italic()`
- `Underline()`
- `Inverse()`
- `Hidden()`
- `Strikethrough()`


### Background Colors
- `BgBlack()`
- `BgRed()`
- `BgGreen()`
- `BgYellow()`
- `BgBlue()`
- `BgPurple()`
- `BgCyan()`
- `BgWhite()`

### Reset
- `chalk.Reset()` will reset all styles.


## Examples
```go
package main

import (
	"github.com/golang-demos/chalk"
)

func main() {
	fmt.Println(chalk.Red().Bold(), "Text in bold red color", chalk.Reset())
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
