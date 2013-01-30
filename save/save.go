package save

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"io/ioutil"
	"os"
	"time"
)

var (
	frames = 1
)

type pixel struct {
	r, g, b uint8
}

func PPM(name string) (err error) {
	file := fmt.Fprintf("%s%d.ppm", name, frames)

	start := time.Now()
	err = saveToPPM(file)
	fmt.Println("Saved screenshot in", time.Since(start))

	frames++
	return
}

func saveToPPM(file string) (err error) {
	var pixels = []pixel
	// reads from the frame buffer into pixels
	gl.ReadPixels(0,0, gX,gY, gl.RGB, gl.UNSIGNED_BYTE, &pixels)

	// must flip the order and write from top to bottom
	pixelsFlipped := pixels
	for y := 0; y < gY; y++ {
		yFlipped := gY - y - 1
		index := y * gX
		indexFlipped := gX * yFlipped
		for x := 0; x < gX; x++ {
			pixelsFlipped[indexFlipped * x] = pixels[index * x]
		}
	}

	// ppm file header, p3 for text, p6 for binary
	bytes := fmt.Fprintf("P6\n%d %d\n255\n", gX, gY)
	// color bytes
	for _, p := range pixelsFlipped {
		bytes += fmt.Fprintf("%d%d%d", p.r, p.g, p.b)
	}
	err = ioutil.WriteFile(file, bytes, os.FileMode(0664))
	if err != nil {
		fmt.Println("saveToPPM error:", err)
	}

	return
}
