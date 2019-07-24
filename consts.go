package tcolor

import "fmt"

/*************************************************************
 * Basic 16 color definition
 *************************************************************/

type Color uint8

// Foreground colors. basic foreground colors 30 - 37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta // 品红
	FgCyan    // 青色
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

// Extra foreground color 90 - 97(非标准)
const (
	FgDarkGray Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is alias of FgDarkGray
	FgGray Color = 90 // 亮黑（灰）
)

// Background colors. basic background colors 40 - 47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault revert default BG
	BgDefault Color = 49
)

// Extra background color 100 - 107(非标准)
const (
	BgDarkGray Color = iota + 99
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is alias of BgDarkGray
	BgGray Color = 100
)

// Option settings
const (
	OpReset         Color = iota // 0 重置所有设置
	OpBold                       // 1 加粗
	OpFuzzy                      // 2 模糊(不是所有的终端仿真器都支持)
	OpItalic                     // 3 斜体(不是所有的终端仿真器都支持)
	OpUnderscore                 // 4 下划线
	OpBlink                      // 5 闪烁
	OpFastBlink                  // 5 快速闪烁(未广泛支持)
	OpReverse                    // 7 颠倒的 交换背景色与前景色
	OpConcealed                  // 8 隐匿的
	OpStrikethrough              // 9 删除的，删除线(未广泛支持)
)

//// There are basic and light foreground color aliases
//const (
//	Red     = FgRed
//	Cyan    = FgCyan
//	Gray    = FgDarkGray // is light Black
//	Blue    = FgBlue
//	Black   = FgBlack
//	Green   = FgGreen
//	White   = FgWhite
//	Yellow  = FgYellow
//	Magenta = FgMagenta
//	// special
//	Bold   = OpBold
//	Normal = FgDefault
//	// extra light
//	LightRed     = FgLightRed
//	LightCyan    = FgLightCyan
//	LightBlue    = FgLightBlue
//	LightGreen   = FgLightGreen
//	LightWhite   = FgLightWhite
//	LightYellow  = FgLightYellow
//	LightMagenta = FgLightMagenta
//)

// String to code string. eg "35"
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}
