package wanalyzer

import (
	"strconv"

	"github.com/fogleman/gg"
)

type Drawer struct {
	ImageName string
	Frequence []MessageFrequence
}

func NewGraph(imageName string) Drawer {
	return Drawer{ImageName: imageName}
}

func (dr Drawer) DrawFrequence(frequence []MessageFrequence) {
	width, height := len(frequence)*100+70, (dr.getHeight(frequence)/10)+100
	context := gg.NewContext(width, height)

	context.SetHexColor("#e6fae6")
	context.DrawRectangle(0, 0, float64(width), float64(height))
	context.Fill()

	context.SetHexColor("#353734")
	for i, elem := range frequence {
		context.DrawRectangle(100*float64(i+1), float64((height)-50-(elem.Count/10)), 35, float64(elem.Count/10))
		context.DrawString(elem.Date, (100*float64(i+1))-15, float64(height-25))
		context.DrawString(strconv.Itoa(elem.Count), (100*float64(i+1))+5, float64(height-50-(elem.Count/10)-10))

		context.Fill()
	}

	context.Rotate(90)
	context.SavePNG(dr.ImageName)

}

func (dr Drawer) getHeight(frequence []MessageFrequence) (max int) {
	for _, elem := range frequence {
		if elem.Count > max {
			max = elem.Count
		}
	}
	return
}
