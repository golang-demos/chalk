package chalk

import (
	"fmt"
)

type Color struct {
	color   int
	style   int
	bgcolor int
}

var LIBRARY_DEVELOPER_MODE = false

func (c Color) String() string {
	fmt.Println("\n", c.color, c.style, c.bgcolor)
	if LIBRARY_DEVELOPER_MODE {
		return fmt.Sprintf("\\u001b[%dm", c.bgcolor) + fmt.Sprintf("\\u001b[%d;%dm", c.style, c.color)
	} else {
		styleText := fmt.Sprintf("\u001b[0;%dm", c.color)
		if c.style > 0 {
			styleText = fmt.Sprintf("\u001b[%d;%dm", c.style, c.color)
		}
		return fmt.Sprintf("\u001b[0;%dm", c.bgcolor) + styleText
	}
}

// Colors
func getColorIntensity(intensity []bool) (intensityVal int) {
	intensityVal = 30
	if len(intensity) > 0 && intensity[0] {
		intensityVal = 90
	}
	return
}
func (c *Color) Black(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 0
	return c
}
func (c *Color) Red(makeIntense ...bool) *Color {
	// log.Println(">", c.color, c.style, c.bgcolor)
	c.color = getColorIntensity(makeIntense) + 1
	return c
}
func (c *Color) Green(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 2
	return c
}
func (c *Color) Orange(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 3
	return c
}
func (c *Color) Yellow(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 3
	return c
}
func (c *Color) Blue(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 4
	return c
}
func (c *Color) Purple(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 5
	return c
}
func (c *Color) Violet(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 5
	return c
}
func (c *Color) Cyan(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 6
	return c
}
func (c *Color) White(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 7
	return c
}

// Text Styles
func (c *Color) Bold() *Color {
	c.style = 1
	return c
}
func (c *Color) Dim() *Color {
	c.style = 2
	return c
}
func (c *Color) Italic() *Color {
	c.style = 3
	return c
}
func (c *Color) Underline() *Color {
	c.style = 4
	return c
}
func (c *Color) Inverse() *Color {
	c.style = 7
	return c
}
func (c *Color) Hidden() *Color {
	c.style = 8
	return c
}
func (c *Color) Strikethrough() *Color {
	c.style = 9
	return c
}

// Background Colors
func getBackgroundIntensity(intensity []bool) (intensityVal int) {
	intensityVal = 40
	if len(intensity) > 0 && intensity[0] {
		intensityVal = 100
	}
	return
}

func (c *Color) BgBlack(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 0
	return c
}
func (c *Color) BgRed(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 1
	return c
}
func (c *Color) BgGreen(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 2
	return c
}
func (c *Color) BgYellow(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 3
	return c
}
func (c *Color) BgBlue(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 4
	return c
}
func (c *Color) BgPurple(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 5
	return c
}
func (c *Color) BgCyan(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 6
	return c
}
func (c *Color) BgWhite(makeIntense ...bool) *Color {
	c.bgcolor = getBackgroundIntensity(makeIntense) + 7
	return c
}

// For all reset. BackgroundColor and TextColor
func Reset() string {
	if LIBRARY_DEVELOPER_MODE {
		return fmt.Sprintf("\\033[0;%dm\\033[%dm", 0, 49)
	} else {
		return fmt.Sprintf("\033[0;%dm\033[%dm", 0, 49)
	}
}

func newColor() *Color {
	var c Color
	c.color = 0
	c.style = 0
	c.bgcolor = 0
	return &c
}

var (
	// Text Colors
	Black  = func() *Color { return newColor().Black() }
	Red    = func() *Color { return newColor().Red() }
	Green  = func() *Color { return newColor().Green() }
	Orange = func() *Color { return newColor().Orange() }
	Yellow = func() *Color { return newColor().Yellow() }
	Blue   = func() *Color { return newColor().Blue() }
	Purple = func() *Color { return newColor().Purple() }
	Violet = func() *Color { return newColor().Violet() }
	Cyan   = func() *Color { return newColor().Cyan() }
	White  = func() *Color { return newColor().White() }

	// Text Styles
	Bold          = func() *Color { return newColor().Bold() }
	Dim           = func() *Color { return newColor().Dim() }
	Italic        = func() *Color { return newColor().Italic() }
	Underline     = func() *Color { return newColor().Underline() }
	Inverse       = func() *Color { return newColor().Inverse() }
	Hidden        = func() *Color { return newColor().Hidden() }
	Strikethrough = func() *Color { return newColor().Strikethrough() }

	// Background Colors
	BgBlack  = func() *Color { return newColor().BgBlack() }
	BgRed    = func() *Color { return newColor().BgRed() }
	BgGreen  = func() *Color { return newColor().BgGreen() }
	BgYellow = func() *Color { return newColor().BgYellow() }
	BgBlue   = func() *Color { return newColor().BgBlue() }
	BgPurple = func() *Color { return newColor().BgPurple() }
	BgCyan   = func() *Color { return newColor().BgCyan() }
	BgWhite  = func() *Color { return newColor().BgWhite() }
)
