package utils

import (
	"image/color"
)

type SLedFrame struct {
	Section int
	Row     int
	Col     int
	R       uint8
	G       uint8
	B       uint8
}

type SLedSection struct {
	Id         int
	H          int
	W          int
	Serpentine bool
}

type SLedProject struct {
	FrameDelay int
	Looping    bool
	FrameDrop  bool
	Frames     [][]SLedFrame
	Sections   []SLedSection
}

type SLedConfig struct {
	Version string
	Author  string
	Project SLedProject
}

func SLed(config SLedConfig, frameIndex int) (ledColors []color.RGBA, totalFrames int, frameDelay int, err error) {

	totalLeds := 0
	var sectionTotals []int
	var sectionOffsets []int

	for index, section := range config.Project.Sections {
		sectionTotals[index] = section.H * section.W
		sectionOffsets[index] = totalLeds
		totalLeds += sectionTotals[index]
	}

	var leds []color.RGBA

	for i := 0; i < totalLeds; i++ {
		leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}

	frameCount := len(config.Project.Frames)

	if frameIndex > frameCount {
		return leds, frameCount, config.Project.FrameDelay, nil
	}

	frame := config.Project.Frames[frameIndex]

	for _, pixel := range frame {
		// find section offset
		sectionOffset := sectionOffsets[pixel.Section]

		thisSection := config.Project.Sections[pixel.Section]

		// find row offset
		rowOffset := thisSection.W * pixel.Row

		// find section col index
		colOffset := pixel.Col
		if thisSection.Serpentine && (pixel.Row%2 == 0) {
			colOffset = thisSection.W - pixel.Col
		}

		// combine offset with row # and col index
		pixelIndex := sectionOffset + rowOffset + colOffset

		// assign in the array
		leds[pixelIndex] = color.RGBA{R: pixel.R, G: pixel.G, B: pixel.B}
	}

	return leds, frameCount, config.Project.FrameDelay, nil
}
