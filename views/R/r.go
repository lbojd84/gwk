// Copyright 2014 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package R

import (
	"gwk/vango"
	"image"
	"image/png"
	"os"
)

func LoadCanvas(filename string) *vango.Canvas {
	var fd, err = os.Open(filename)
	if err != nil {
		return nil
	}

	defer fd.Close()

	return LoadCanvasFile(fd)
}

func LoadCanvasFile(fd *os.File) *vango.Canvas {
	var png, err = png.Decode(fd)
	if err != nil {
		return nil
	}

	if nrgba, ok := png.(*image.NRGBA); ok {
		var canvas = vango.NewCanvas(nrgba.Rect.Dx(), nrgba.Rect.Dy())

		canvas.DrawImageNRGBA(0, 0, nrgba, nil)

		canvas.SetOpaque(true)

		return canvas
	}

	if rgba, ok := png.(*image.RGBA); ok {
		var canvas = vango.NewCanvas(rgba.Rect.Dx(), rgba.Rect.Dy())

		canvas.DrawImageRGBA(0, 0, rgba, nil)

		return canvas
	}

	return nil
}
