package pipe

import (
	_ "embed"
	g "github.com/AllenDang/giu"
)

const FontSize = 18

//go:embed segoeuisl.ttf
var fontBytes []byte

func SetFont() {
	g.Context.FontAtlas.SetDefaultFontFromBytes(fontBytes, FontSize)
}
