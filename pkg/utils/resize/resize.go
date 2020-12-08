package resize

import (
	"image"
	"image/png"
	"math"
	"os"

	"github.com/nfnt/resize"
)

type MakeThumb struct {
	MaxWidth  float64 // 最大宽度
	MaxHeight float64 // 最大高度
	IsHeight  bool    // 是否以高度为准
	ImagePath string  // 原图片路径
	SavePath  string  // 保存图片路径
}

func New(imagePath, savePath string) *MakeThumb {
	return &MakeThumb{
		MaxHeight: 180,
		MaxWidth:  180,
		IsHeight:  true,
		ImagePath: imagePath,
		SavePath:  savePath,
	}
}

// 计算图片缩放后的尺寸
func (m *MakeThumb) calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	var ratio float64
	if m.IsHeight {
		ratio = m.MaxHeight / float64(srcHeight)
	} else {
		ratio = math.Min(m.MaxWidth/float64(srcWidth), m.MaxHeight/float64(srcHeight))
	}
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

func (m *MakeThumb) MakeThumbnail() error {
	file, err := os.Open(m.ImagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := m.calculateRatioFit(width, height)

	// 调用resize库进行图片缩放
	mm := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	imgFile, err := os.Create(m.SavePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	if err = png.Encode(imgFile, mm); err != nil {
		return err
	}

	return nil
}
