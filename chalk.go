package chalk

import (
	"fmt"
	"strings"
)

type Color struct {
	color   int
	style   int
	bgcolor int
}

func (c Color) String() string {
	var ext string
	var parts []string

	if c.bgcolor > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.bgcolor))
	}
	if c.color > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.color))
	}
	if c.style > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.style))
	}

	ext = strings.Join(parts[:], ";")

	return "\u001b[" + ext + "m"
}

// Colors
func (c *Color) Black() *Color {
	c.color = 30
	return c
}
func (c *Color) Red() *Color {
	c.color = 31
	return c
}
func (c *Color) Green() *Color {
	c.color = 32
	return c
}
func (c *Color) Yellow() *Color {
	c.color = 33
	return c
}
func (c *Color) Blue() *Color {
	c.color = 34
	return c
}
func (c *Color) Magenta() *Color {
	c.color = 35
	return c
}
func (c *Color) Cyan() *Color {
	c.color = 36
	return c
}
func (c *Color) White() *Color {
	c.color = 37
	return c
}

func (c *Color) BlackLight() *Color {
	c.color = 90
	return c
}
func (c *Color) RedLight() *Color {
	c.color = 91
	return c
}
func (c *Color) GreenLight() *Color {
	c.color = 92
	return c
}
func (c *Color) YellowLight() *Color {
	c.color = 93
	return c
}
func (c *Color) BlueLight() *Color {
	c.color = 94
	return c
}
func (c *Color) MagentaLight() *Color {
	c.color = 95
	return c
}
func (c *Color) CyanLight() *Color {
	c.color = 96
	return c
}
func (c *Color) WhiteLight() *Color {
	c.color = 97
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
func (c *Color) BgBlack() *Color {
	c.bgcolor = 40
	return c
}
func (c *Color) BgRed() *Color {
	c.bgcolor = 41
	return c
}
func (c *Color) BgGreen() *Color {
	c.bgcolor = 42
	return c
}
func (c *Color) BgYellow() *Color {
	c.bgcolor = 43
	return c
}
func (c *Color) BgBlue() *Color {
	c.bgcolor = 44
	return c
}
func (c *Color) BgMagenta() *Color {
	c.bgcolor = 45
	return c
}
func (c *Color) BgCyan() *Color {
	c.bgcolor = 46
	return c
}
func (c *Color) BgWhite() *Color {
	c.bgcolor = 47
	return c
}

func (c *Color) BgBlackLight() *Color {
	c.bgcolor = 100
	return c
}
func (c *Color) BgRedLight() *Color {
	c.bgcolor = 101
	return c
}
func (c *Color) BgGreenLight() *Color {
	c.bgcolor = 102
	return c
}
func (c *Color) BgYellowLight() *Color {
	c.bgcolor = 103
	return c
}
func (c *Color) BgBlueLight() *Color {
	c.bgcolor = 104
	return c
}
func (c *Color) BgMagentaLight() *Color {
	c.bgcolor = 105
	return c
}
func (c *Color) BgCyanLight() *Color {
	c.bgcolor = 106
	return c
}
func (c *Color) BgWhiteLight() *Color {
	c.bgcolor = 107
	return c
}

// For all reset. BackgroundColor and TextColor
func Reset() string {
	return fmt.Sprintf("\033[%dm", 0)
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
	Black   = func() *Color { return newColor().Black() }
	Red     = func() *Color { return newColor().Red() }
	Green   = func() *Color { return newColor().Green() }
	Yellow  = func() *Color { return newColor().Yellow() }
	Blue    = func() *Color { return newColor().Blue() }
	Magenta = func() *Color { return newColor().Magenta() }
	Cyan    = func() *Color { return newColor().Cyan() }
	White   = func() *Color { return newColor().White() }

	// Text Light Colors
	BlackLight   = func() *Color { return newColor().BlackLight() }
	RedLight     = func() *Color { return newColor().RedLight() }
	GreenLight   = func() *Color { return newColor().GreenLight() }
	YellowLight  = func() *Color { return newColor().YellowLight() }
	BlueLight    = func() *Color { return newColor().BlueLight() }
	MagentaLight = func() *Color { return newColor().MagentaLight() }
	CyanLight    = func() *Color { return newColor().CyanLight() }
	WhiteLight   = func() *Color { return newColor().WhiteLight() }

	// Text Styles
	Bold          = func() *Color { return newColor().Bold() }
	Dim           = func() *Color { return newColor().Dim() }
	Italic        = func() *Color { return newColor().Italic() }
	Underline     = func() *Color { return newColor().Underline() }
	Inverse       = func() *Color { return newColor().Inverse() }
	Hidden        = func() *Color { return newColor().Hidden() }
	Strikethrough = func() *Color { return newColor().Strikethrough() }

	// Background Colors
	BgBlack   = func() *Color { return newColor().BgBlack() }
	BgRed     = func() *Color { return newColor().BgRed() }
	BgGreen   = func() *Color { return newColor().BgGreen() }
	BgYellow  = func() *Color { return newColor().BgYellow() }
	BgBlue    = func() *Color { return newColor().BgBlue() }
	BgMagenta = func() *Color { return newColor().BgMagenta() }
	BgCyan    = func() *Color { return newColor().BgCyan() }
	BgWhite   = func() *Color { return newColor().BgWhite() }

	BgBlackLight   = func() *Color { return newColor().BgBlackLight() }
	BgRedLight     = func() *Color { return newColor().BgRedLight() }
	BgGreenLight   = func() *Color { return newColor().BgGreenLight() }
	BgYellowLight  = func() *Color { return newColor().BgYellowLight() }
	BgBlueLight    = func() *Color { return newColor().BgBlueLight() }
	BgMagentaLight = func() *Color { return newColor().BgMagentaLight() }
	BgCyanLight    = func() *Color { return newColor().BgCyanLight() }
	BgWhiteLight   = func() *Color { return newColor().BgWhiteLight() }
)
