//go:generate fyne bundle -o bundled.go -package gui ./resources/NotoSansJP-Medium.ttf

package examples

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var errorColor = color.NRGBA{0xf4, 0x43, 0x36, 0xff}
var palette = map[fyne.ThemeColorName]color.Color{
	theme.ColorNameBackground:      color.NRGBA{0xff, 0xff, 0xff, 0xff},
	theme.ColorNameButton:          color.Transparent,
	theme.ColorNameDisabled:        color.NRGBA{0x0, 0x0, 0x0, 0x42},
	theme.ColorNameDisabledButton:  color.NRGBA{0xe5, 0xe5, 0xe5, 0xff},
	theme.ColorNameError:           errorColor,
	theme.ColorNameForeground:      color.NRGBA{0x21, 0x21, 0x21, 0xff},
	theme.ColorNameHover:           color.NRGBA{0x0, 0x0, 0x0, 0x0f},
	theme.ColorNameInputBackground: color.NRGBA{0x0, 0x0, 0x0, 0x19},
	theme.ColorNamePlaceHolder:     color.NRGBA{0x88, 0x88, 0x88, 0xff},
	theme.ColorNamePressed:         color.NRGBA{0x0, 0x0, 0x0, 0x19},
	theme.ColorNameScrollBar:       color.NRGBA{0x0, 0x0, 0x0, 0x99},
	theme.ColorNameShadow:          color.NRGBA{0x0, 0x0, 0x0, 0x33},
	theme.ColorNamePrimary:         color.NRGBA{0x42, 0x85, 0xF4, 0xff},
}

type DefaultTheme struct{}

func (m *DefaultTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return palette[n]
}

func (m *DefaultTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
	// return resourceNotoSansJPMediumTtf
}

func (m *DefaultTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *DefaultTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	default:
		log.Printf("ThemeColorName: %s Not Found", n)
		return theme.DefaultTheme().Size(n)
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameInlineIcon:
		return 9
	case theme.SizeNamePadding:
		return 2
	case theme.SizeNameScrollBar:
		return 10
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameText:
		return 10
	case theme.SizeNameCaptionText:
		return 8
	case theme.SizeNameInputBorder:
		return 1
	}

}
