package color

import (
	"testing"
)

func TestWrapString(t *testing.T) {
	expected := "\x1b[31mred\x1b[0m"
	result := wrapString(FgRed, "red")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestBlackString(t *testing.T) {
	expected := "\x1b[30mblack\x1b[0m"
	result := BlackString("black")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestRedString(t *testing.T) {
	expected := "\x1b[31mred\x1b[0m"
	result := RedString("red")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestGreenString(t *testing.T) {
	expected := "\x1b[32mgreen\x1b[0m"
	result := GreenString("green")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestYellowString(t *testing.T) {
	expected := "\x1b[33myellow\x1b[0m"
	result := YellowString("yellow")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestBlueString(t *testing.T) {
	expected := "\x1b[34mblue\x1b[0m"
	result := BlueString("blue")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestMagentaString(t *testing.T) {
	expected := "\x1b[35mmagenta\x1b[0m"
	result := MagentaString("magenta")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestCyanString(t *testing.T) {
	expected := "\x1b[36mcyan\x1b[0m"
	result := CyanString("cyan")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestWhiteString(t *testing.T) {
	expected := "\x1b[37mwhite\x1b[0m"
	result := WhiteString("white")
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

func TestIsTTY(t *testing.T) {
	// we cannot test this directly, so just ensure it doesn't panic
	_ = IsTTY()
}
