package webserver

import (
	"time"

	svg "github.com/ajstarks/svgo"
)

var (
	width     = 800
	height    = 400
	startTime = time.Now().UnixNano()
)

func drawPoint(osvg *svg.SVG, pnt int, process int) {
	sec := time.Now().UnixNano()

	diff := (int64(sec) - int64(startTime)) / 100000

	pointLocation := diff
	pointLocationV := 0

	color := "#000000"

	switch {
	case process == 1:
		pointLocationV = 60
		color = "#cc6666"
	default:
		pointLocationV = 180
		color = "#66cc66"
	}

	osvg.Rect(int(pointLocation), pointLocationV, 3, 5, "fill:"+color+";stroke:none;")
	time.Sleep(time.Millisecond * 150)
}
