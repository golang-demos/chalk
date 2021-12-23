package main

import "fmt"

type Color struct {
	color   int
	style   int
	bgcolor int
}

func (c Color) String() string {
	styleText := fmt.Sprintf("\u001b[%dm", c.color)
	if c.style > 0 {
		styleText = fmt.Sprintf("\u001b[%d;%dm", c.style, c.color)
	}
	return fmt.Sprintf("\u001b[%dm", c.bgcolor) + styleText
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

var (
	// Text Colors
	Black  = (&Color{}).Black
	Red    = (&Color{}).Red
	Green  = (&Color{}).Green
	Orange = (&Color{}).Orange
	Yellow = (&Color{}).Yellow
	Blue   = (&Color{}).Blue
	Purple = (&Color{}).Purple
	Violet = (&Color{}).Violet
	Cyan   = (&Color{}).Cyan
	White  = (&Color{}).White

	// Text Styles
	Bold   = (&Color{}).Bold
	Dim    = (&Color{}).Dim
	Italic = (&Color{}).Italic
)

// For all reset. BackgroundColor and TextColor
func Reset() string {
	return fmt.Sprintf("\u001b[0;%dm", 0)
}

func main() {

	fmt.Println("--", Red(), "Red + Regular", Reset(), "--")
	fmt.Println("--", Red().Bold(), "Red + Bold", Reset(), "--")
	fmt.Println("--", Red().Italic(), "Red + Italic", Reset(), "--")
}
