package mimg

import (
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"strings"
)

func getMaxLen(text string) int {
	lines := strings.Split(text, "\n")
	max := 0
	for _, line := range lines {
		lineLen := len(line)
		if lineLen > max {
			max = lineLen
		}
	}
	return max
}

var (
	DefaultLinuxFontFile = "/usr/share/fonts/truetype/freefont/FreeMono.ttf"
)

func TextToImage(text string, fontSize float64, fontFile string, fontColor color.Color) (image.Image, error) {
	lines := len(strings.Split(text, "\n"))
	maxLen := getMaxLen(text)
	margin := 10
	width := maxLen*int(fontSize) + margin*2
	height := int(fontSize)*lines + margin*2

	var fontDPI float64 = 72

	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	bg := image.Transparent

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(newImg, newImg.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(newImg.Bounds())
	c.SetDst(newImg)
	c.SetSrc(image.NewUniform(fontColor))
	// Draw the text.
	for i, s := range strings.Split(text, "\n") {
		Y := int(fontSize) + i*int(fontSize)
		pt := freetype.Pt(margin, Y)
		_, err := c.DrawString(s, pt)
		if err != nil {
			return nil, err
		}
	}
	return newImg, nil
}

func TextToImageFile(text string, fontSize float64, fontFile string, fontColor color.Color, outFile string) error {
	newImg, err := TextToImage(text, fontSize, fontFile, fontColor)
	if err != nil {
		return err
	}
	return WriteImageFile(outFile, newImg)
}
