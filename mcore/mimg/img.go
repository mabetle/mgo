package mimg

import (
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

type FileType int

const (
	UNKNOWN_TYPE FileType = iota
	JPEG
	PNG
)

// ImageRGBA 按照分辨率创建一张空白图片对象
func ImageRGBA(width, height int) *image.RGBA {
	// 建立图像,image.Rect(最小X,最小Y,最大X,最小Y)
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// getFileExt
func getFileExt(path string) string {
	extA := strings.Split(path, ".")
	lenExt := len(extA)
	// lenExt should greater than 2
	if lenExt < 2 {
		return ""
	}
	ext := extA[lenExt-1]
	ext = strings.ToLower(ext)
	return ext
}

// OpenImage open Image with file type
func OpenImageFile(path string) (image.Image, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	ft := getFileType(path)
	return OpenImage(f, ft)
}

func OpenImage(in io.Reader, ft FileType) (image.Image, error) {
	switch ft {
	case JPEG:
		return jpeg.Decode(in)
	case PNG:
		return png.Decode(in)
	default:
	}
	return nil, errors.New("unknown file type")
}

// 将图片绘制到图片
func ImageDrawRGBA(img *image.RGBA, imgcode image.Image, x, y int) {
	// 绘制图像
	// image.Point A点的X,Y坐标,轴向右和向下增加{0,0}
	// image.ZP ZP is the zero Point
	// image.Pt Pt is shorthand for Point{X, Y}
	draw.Draw(img, img.Bounds(), imgcode, image.Pt(x, y), draw.Over)
}

func getFileType(path string) FileType {
	ext := getFileExt(path)
	ft := UNKNOWN_TYPE
	switch ext {
	case "jpg", "jpeg":
		ft = JPEG
	case "png":
		ft = PNG
	default:
	}
	return ft
}

// WriteImage encode by file ext
func WriteImageFile(path string, img image.Image) error {
	f, errNew := os.Create(path)
	defer f.Close()
	if errNew != nil {
		return errNew
	}
	ft := getFileType(path)
	return WriteImage(img, ft, f)
}

// WriteImage
func WriteImage(img image.Image, ft FileType, out io.Writer) error {
	switch ft {
	case JPEG:
		return jpeg.Encode(out, img, &jpeg.Options{100})
	case PNG:
		return png.Encode(out, img)
	default:
	}
	return errors.New("unknown output format.")
}
