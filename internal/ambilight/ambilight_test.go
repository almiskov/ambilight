package ambilight

import (
	"fmt"
	"image/color"
	"testing"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func TestAreas(t *testing.T) {
	cfg := Config{
		X:       17,
		Y:       9,
		Depth:   50,
		Display: 0,
	}
	ss, err := screenshot.CaptureDisplay(cfg.Display)
	if err != nil {
		t.Fatal(err)
	}

	sides := getSidesAreas(ss.Rect, cfg.X, cfg.Y, cfg.Depth)

	gc := draw2dimg.NewGraphicContext(ss)
	gc.SetLineWidth(1.0)

	for i, s := range sides {
		switch i {
		case TOP, BOTTOM:
			gc.SetStrokeColor(color.RGBA{255, 0, 0, 255})
		case LEFT, RIGHT:
			gc.SetStrokeColor(color.RGBA{0, 255, 234, 255})
		}

		for _, area := range s {
			draw2dkit.Rectangle(gc, float64(area.Min.X), float64(area.Min.Y), float64(area.Max.X), float64(area.Max.Y))
		}

		gc.Stroke()
	}

	draw2dimg.SaveToPngFile(fmt.Sprintf("test_%d.png", time.Now().UnixMilli()), ss)
}

func BenchmarkAreas17x9(b *testing.B) {
	cfg := Config{
		X:       17,
		Y:       9,
		Depth:   50,
		Display: 0,
	}
	ss, err := screenshot.CaptureDisplay(cfg.Display)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getSidesAreas(ss.Rect, cfg.X, cfg.Y, cfg.Depth)
	}
}

func BenchmarkAreas10x10(b *testing.B) {
	cfg := Config{
		X:       10,
		Y:       10,
		Depth:   50,
		Display: 0,
	}
	ss, err := screenshot.CaptureDisplay(cfg.Display)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getSidesAreas(ss.Rect, cfg.X, cfg.Y, cfg.Depth)
	}
}

func BenchmarkAreas20x20(b *testing.B) {
	cfg := Config{
		X:       20,
		Y:       20,
		Depth:   50,
		Display: 0,
	}
	ss, err := screenshot.CaptureDisplay(cfg.Display)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getSidesAreas(ss.Rect, cfg.X, cfg.Y, cfg.Depth)
	}
}

func BenchmarkAreas30x30(b *testing.B) {
	cfg := Config{
		X:       30,
		Y:       30,
		Depth:   50,
		Display: 0,
	}
	ss, err := screenshot.CaptureDisplay(cfg.Display)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getSidesAreas(ss.Rect, cfg.X, cfg.Y, cfg.Depth)
	}
}
