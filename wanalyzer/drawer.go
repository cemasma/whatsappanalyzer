package wanalyzer

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

// Drawer struct for drawing
type Drawer struct {
	ImageName string
}

// NewGraph creates a Drawer
func NewGraph(imageName string) Drawer {
	return Drawer{ImageName: imageName}
}

// DrawFrequency draws a graph of frequency by time periods
func (dr Drawer) DrawFrequency(frequency map[string]int) {
	sizeVariables := getTimeFrequencySizeVariables(frequency)
	width, height := sizeVariables[0]*100+70, (sizeVariables[1]/10)+100

	outputImage := image.NewRGBA(image.Rect(0, 0, width, height))
	defer dr.saveImage(outputImage)

	dr.drawRectangle(0, 0, width, height, outputImage, color.RGBA{
		R: 230,
		G: 250,
		B: 230,
		A: 255,
	})

	i := 0
	columnAndTextColor := color.RGBA{
		R: 53,
		G: 55,
		B: 52,
		A: 255,
	}
	for key, value := range frequency {
		x1 := 100 * (i + 1)
		y1 := height - 50 - (value / 10)
		dr.drawRectangle(x1, y1, x1+35, y1+(value/10), outputImage, columnAndTextColor)
		dr.drawString(key, 100*(i+1)-5, height-25, outputImage, columnAndTextColor)
		dr.drawString(strconv.Itoa(value), 100*(i+1), height-50-(value/10)-10, outputImage, columnAndTextColor)
		i++
	}
}

func getTimeFrequencySizeVariables(frequency map[string]int) []int {
	wVar, maxVal := 0, 0
	for _, value := range frequency {
		if value > 0 {
			wVar++
		}
		if value > maxVal {
			maxVal = value
		}
	}
	return []int{wVar, maxVal}
}

func (dr Drawer) getHeight(frequency []MessageFrequence) (max int) {
	for _, elem := range frequency {
		if elem.Count > max {
			max = elem.Count
		}
	}
	return
}

func (dr Drawer) saveImage(rgba *image.RGBA) {
	f, err := os.Create(dr.ImageName)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := f.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	err = png.Encode(f, rgba)
	if err != nil {
		log.Fatal(err)
	}
}

func (dr Drawer) drawRectangle(x1, y1, x2, y2 int, rgba *image.RGBA, color color.Color) {
	for i := x1; i <= x2; i++ {
		for j := y1; j < y2; j++ {
			rgba.Set(i, j, color)
		}
	}
}

func (dr Drawer) drawString(text string, x, y int, rgba *image.RGBA, color color.Color) {
	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(color),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)},
	}
	d.DrawString(text)
}
