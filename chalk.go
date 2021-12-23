package chalk

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
