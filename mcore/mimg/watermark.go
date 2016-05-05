package mimg

import (
	"image"
	"image/draw"
	"io"
)

func WaterMarkImage(img image.Image, watermark image.Image) image.Image {
	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	newImg := image.NewNRGBA(b)

	draw.Draw(newImg, b, img, image.ZP, draw.Src)
	draw.Draw(newImg, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)
	return newImg
}

// WaterMark
func OutWaterMark(oldPic io.Reader, waterMark io.Reader, out io.Writer, ft FileType) error {
	img, err := OpenImage(oldPic, ft)
	if err != nil {
		return err
	}
	watermark, errMark := OpenImage(waterMark, ft)
	if errMark != nil {
		return err
	}
	newImg := WaterMarkImage(img, watermark)
	return WriteImage(newImg, ft, out)
}

// WaterMarkJpgFile
func WaterMarkFile(inputFile, watermarkFile, newFile string) error {
	img, errInput := OpenImageFile(inputFile)
	if errInput != nil {
		return errInput
	}

	wm, errMark := OpenImageFile(watermarkFile)
	if errMark != nil {
		return errMark
	}
	newImg := WaterMarkImage(img, wm)
	return WriteImageFile(newFile, newImg)
}
