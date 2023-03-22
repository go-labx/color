package color

import (
	"fmt"
	"os"
)

// Attribute represents a color or style attribute.
type Attribute int

const escape = "\x1b"

// Constants for color and style attributes.
const (
	Reset        Attribute = iota // Reset all attributes
	Bold                          // Bold font
	Faint                         // Faint font (decreased intensity)
	Italic                        // Italic font
	Underline                     // Underline font
	BlinkSlow                     // Blinking font (slow)
	BlinkRapid                    // Blinking font (rapid)
	ReverseVideo                  // Inverse video (background becomes foreground and vice versa)
	Concealed                     // Concealed font (hidden)
	CrossedOut                    // Crossed-out font
)

// Constants for foreground colors.
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// wrapString wraps the input string with the given color.
func wrapString(color Attribute, format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)

	if !IsTTY() {
		return s
	}
	return fmt.Sprintf("%s[%dm%s%s[%dm", escape, color, s, escape, Reset)
}

// BlackString returns a black string.
func BlackString(format string, a ...interface{}) string {
	return wrapString(FgBlack, format, a...)
}

// RedString returns a red string.
func RedString(format string, a ...interface{}) string {
	return wrapString(FgRed, format, a...)
}

// GreenString returns a green string.
func GreenString(format string, a ...interface{}) string {
	return wrapString(FgGreen, format, a...)
}

// YellowString returns a yellow string.
func YellowString(format string, a ...interface{}) string {
	return wrapString(FgYellow, format, a...)
}

// BlueString returns a blue string.
func BlueString(format string, a ...interface{}) string {
	return wrapString(FgBlue, format, a...)
}

// MagentaString returns a magenta string.
func MagentaString(format string, a ...interface{}) string {
	return wrapString(FgMagenta, format, a...)
}

// CyanString returns a cyan string.
func CyanString(format string, a ...interface{}) string {
	return wrapString(FgCyan, format, a...)
}

// WhiteString returns a white string.
func WhiteString(format string, a ...interface{}) string {
	return wrapString(FgWhite, format, a...)
}

// IsTTY checks if the output is a TTY device.
func IsTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}
