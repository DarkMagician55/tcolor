package tcolor

import (
	"fmt"
)

const (
	FullColorTpl   = "\x1b[%sm%s\x1b[0m"
	FullColorTplLn = "\x1b[%sm%s\x1b[0m\n"
)

//todo windows
//var env  = "linux"

//// There are basic and light foreground color aliases
var (
	Red     = TColor{FgRed, 0, 0}
	Cyan    = TColor{FgCyan, 0, 0}
	Gray    = TColor{FgDarkGray, 0, 0}
	Blue    = TColor{FgBlue, 0, 0}
	Black   = TColor{FgBlack, 0, 0}
	Green   = TColor{FgGreen, 0, 0}
	White   = TColor{FgWhite, 0, 0}
	Yellow  = TColor{FgYellow, 0, 0}
	Magenta = TColor{FgMagenta, 0, 0}
	// special
	Bold   = TColor{OpBold, 0, 0}
	Normal = TColor{FgDefault, 0, 0}
	// extra light
	LightRed     = TColor{FgLightRed, 0, 0}
	LightCyan    = TColor{FgLightCyan, 0, 0}
	LightBlue    = TColor{FgLightBlue, 0, 0}
	LightGreen   = TColor{FgLightGreen, 0, 0}
	LightWhite   = TColor{FgLightWhite, 0, 0}
	LightYellow  = TColor{FgLightYellow, 0, 0}
	LightMagenta = TColor{FgLightMagenta, 0, 0}
)

type TColor struct {
	FgColor Color
	BgColor Color
	Op      Color
}

func NewTColor(color ...Color) TColor {
	switch len(color) {
	case 1:
		return TColor{FgColor: color[0]}
	case 2:
		return TColor{FgColor: color[0], BgColor: color[1]}
	case 3:
		return TColor{FgColor: color[0], BgColor: color[1], Op: color[2]}
	default:
		return TColor{}
	}
}

// Render messages by color setting
// Usage:
// 		fmt.Println(color.FgGreen.Render("message"))
func (c TColor) Render(msg string) string {
	return fmt.Sprintf(FullColorTpl, c.Code(), msg)
}

func (c TColor) Println(msg string) {
	fmt.Printf(FullColorTplLn, c.Code(), msg)
}

func (c TColor) Code() string {
	if c.Op != 0 {
		return fmt.Sprintf("%d;%d;%d", c.FgColor, c.BgColor, c.Op)
	} else if c.BgColor != 0 {
		return fmt.Sprintf("%d;%d", c.FgColor, c.BgColor)
	} else {
		return fmt.Sprintf("%d", c.FgColor)
	}
}
