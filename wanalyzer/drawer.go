package wanalyzer

import (
	"strconv"

	"github.com/fogleman/gg"
)

// Drawer struct for drawing
type Drawer struct {
	ImageName string
}

// NewGraph creates a Drawer
func NewGraph(imageName string) Drawer {
	return Drawer{ImageName: imageName}
}

// DrawFrequence draws a graph of the messaging frequence
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

	context.SavePNG(dr.ImageName)
}

// DrawTimeFrequence draws a graph of messaging frequence by time periods
func (dr Drawer) DrawTimeFrequence(frequence map[string]int) {
	sizeVariables := getTimeFrequenceSizeVariables(frequence)
	width, height := sizeVariables[0]*100+70, (sizeVariables[1]/10)+100

	context := gg.NewContext(width, height)

	context.SetHexColor("#e6fae6")
	context.DrawRectangle(0, 0, float64(width), float64(height))
	context.Fill()

	i := 0
	context.SetHexColor("#353734")
	for key, value := range frequence {
		context.DrawRectangle(100*float64(i+1), float64((height)-50-(value/10)), 35, float64(value/10))
		context.DrawString(key, (100*float64(i+1))-5, float64(height-25))
		context.DrawString(strconv.Itoa(value), (100 * float64(i+1)), float64(height-50-(value/10)-10))

		context.Fill()
		i++
	}

	context.SavePNG(dr.ImageName)
}

func getTimeFrequenceSizeVariables(frequence map[string]int) []int {
	wVar, maxVal := 0, 0
	for _, value := range frequence {
		if value > 0 {
			wVar++
		}
		if value > maxVal {
			maxVal = value
		}
	}
	return []int{wVar, maxVal}
}

func (dr Drawer) getHeight(frequence []MessageFrequence) (max int) {
	for _, elem := range frequence {
		if elem.Count > max {
			max = elem.Count
		}
	}
	return
}
