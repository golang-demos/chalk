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
func (c *Color) Red(makeIntense ...bool) *Color {
	c.color = getColorIntensity(makeIntense) + 1
	return c
}

var (
	// Text Colors
	Red = (&Color{}).Red
)

// For all reset. BackgroundColor and TextColor
func Reset() string {
	return fmt.Sprintf("\u001b[0;%dm", 0)
}

func main() {

	fmt.Println("--", Red(), "Red + Regular", Reset(), "--")
}
