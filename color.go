package tcolor

import (
	"fmt"
)

const (
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
)

type Env struct {
	env string
}

type TColor struct {
	FgColor Color
	BgColor Color
	Op      Color
}

func NewTColor(color ...Color) TColor {
	switch len(color) {
	case 1:
		return TColor{FgColor: color[0]}
	default:
		return TColor{}
	}
}

func (c *TColor) Render(msg string) string {
	return fmt.Sprintf(FullColorTpl, c.Code(), msg)
}

//func Red(msg string) string {
//	c := NewTColor(FgRed)
//	return c.Render(msg)
//}

func Render(fg Color, msg string) string {
	c := NewTColor(fg)
	return c.Render(msg)
}

func (c TColor) Code() string {
	if c.Op != 0 {
		return fmt.Sprintf("%d;%d,%d", c.FgColor, c.BgColor, c.Op)
	} else if c.BgColor != 0 {
		return fmt.Sprintf("%d;%d", c.FgColor, c.BgColor)
	} else {
		return fmt.Sprintf("%d", c.FgColor)
	}
}
