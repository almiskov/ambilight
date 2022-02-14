package ambilight

import (
	"image"

	"github.com/kbinani/screenshot"
)

func Run(cfg Config) error {
	b := screenshot.GetDisplayBounds(cfg.Display)
	areas := getSidesAreas(b, cfg.X, cfg.Y, cfg.Depth)

	_ = areas

	return nil
}

func getSidesAreas(rect image.Rectangle, x, y, depth int) [4][]image.Rectangle {
	sides := getSides(rect, depth)
	return splitSides(sides, x, y)
}

func getSides(rect image.Rectangle, depth int) [4]image.Rectangle {
	var rects [4]image.Rectangle

	rects[TOP] = image.Rectangle{rect.Min, image.Pt(rect.Max.X, depth)}
	rects[RIGHT] = image.Rectangle{image.Pt(rect.Max.X-depth, 0), rect.Max}
	rects[BOTTOM] = image.Rectangle{image.Pt(0, rect.Max.Y-depth), rect.Max}
	rects[LEFT] = image.Rectangle{rect.Min, image.Pt(depth, rect.Max.Y)}

	return rects
}

func splitSides(sides [4]image.Rectangle, x, y int) [4][]image.Rectangle {
	var split [4][]image.Rectangle
	area := make([]image.Rectangle, 2*x+2*y)

	split[TOP] = area[:x]
	split[RIGHT] = area[x : x+y]
	split[BOTTOM] = area[x+y : 2*x+y]
	split[LEFT] = area[2*x+y:]

	for i, s := range sides {
		switch i {
		case TOP, BOTTOM:
			var (
				w   = s.Dx() / x
				mod = s.Dx() % x
				ls  = mod / 2
				rs  = mod - ls
			)

			for j := 0; j < x; j++ {
				min := image.Pt(s.Min.X+w*j+ls, s.Min.Y)
				max := image.Pt(s.Min.X+w*(j+1)+ls, s.Max.Y)

				if j == 0 {
					min.X -= ls
				}

				if j == x-1 {
					max.X += rs
				}

				split[i][j] = image.Rectangle{min, max}
			}
		case LEFT, RIGHT:
			var (
				h   = s.Dy() / y
				mod = s.Dy() % y
				ts  = mod / 2
				bs  = mod - ts
			)

			for j := 0; j < y; j++ {
				min := image.Pt(s.Min.X, s.Min.Y+h*j+ts)
				max := image.Pt(s.Max.X, s.Min.Y+h*(j+1)+ts)

				if j == 0 {
					min.Y -= ts
				}

				if j == y-1 {
					max.Y += bs
				}

				split[i][j] = image.Rectangle{min, max}
			}
		}
	}

	return split
}
