// Copyright 2013 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vango

import (
	"image"
	"log"
)

type Canvas struct {
	pix    []byte
	bounds image.Rectangle
	stride int
	opaque bool
}

func NewCanvas(width int, height int) *Canvas {
	log.Printf("NewCanvas")
	var c Canvas
	c.bounds = image.Rect(0, 0, width, height)
	c.pix = make([]byte, c.W()*c.H()*4)
	c.stride = c.W() * 4
	return &c
}

func (c *Canvas) X() int {
	return c.bounds.Min.X
}

func (c *Canvas) Y() int {
	return c.bounds.Min.Y
}

func (c *Canvas) W() int {
	return c.bounds.Dx()
}

func (c *Canvas) H() int {
	return c.bounds.Dy()
}

func (c *Canvas) Stride() int {
	return c.stride
}

func (c *Canvas) SetStride(stride int) {
	c.stride = stride
}

func (c *Canvas) Pix() []byte {
	return c.pix
}

func (c *Canvas) SetPix(pix []byte) {
	c.pix = pix
}

func (c *Canvas) Bounds() image.Rectangle {
	return c.bounds
}

func (c *Canvas) SetBounds(bounds image.Rectangle) {
	c.bounds = bounds
}

func (c *Canvas) Opaque() bool {
	return c.opaque
}

func (c *Canvas) SetOpaque(opaque bool) {
	c.opaque = opaque
}

func (c *Canvas) SubCanvas(r image.Rectangle) *Canvas {
	r = r.Intersect(c.Bounds())
	if r.Empty() {
		return &Canvas{}
	}
	i := c.PixOffset(r.Min.X, r.Min.Y)
	return &Canvas{
		pix:    c.pix[i:],
		stride: c.stride,
		bounds: r,
	}
}

func (c *Canvas) PixOffset(x int, y int) int {
	return x*4 + y*c.Stride()
}

func (dst *Canvas) DrawColor(r, g, b byte) {
	var i = dst.PixOffset(dst.X(), dst.Y())
	var W = i + dst.W()*4
	var p = dst.Pix()
	var j = 0

	for j < dst.H() {
		for i < W {
			p[i+0] = b
			p[i+1] = g
			p[i+2] = r
			p[i+3] = 255
			i += 4
		}
		i = i + dst.Stride() - dst.W()*4
		W = W + dst.Stride()
		j++
	}
}

func (dst *Canvas) DrawLine(from image.Point, to image.Point) {
	return
}

func (dst *Canvas) DrawCanvas(x int, y int, src *Canvas, srcRc *image.Rectangle) {
	if srcRc == nil {
		var tmpRc = src.Bounds()
		srcRc = &tmpRc
	}
	log.Printf("dstRc srcRc %v %v", dst.Bounds(), srcRc)
	log.Printf("src_rc %v src_bds %v dst_bds %v", srcRc, src.Bounds(), dst.Bounds())

	var srcX, srcY = srcRc.Min.X, srcRc.Min.Y
	var dstX, dstY = x, y

	var bltW, bltH = srcRc.Dx(), srcRc.Dy()
	if bltW > dst.Bounds().Dx()-x {
		bltW = dst.Bounds().Dx() - x
	}
	if bltH > dst.Bounds().Dy()-y {
		bltH = dst.Bounds().Dy() - y
	}

	var srcI = src.PixOffset(srcX, srcY)
	var dstI = dst.PixOffset(dstX, dstY)

	var srcPix = src.Pix()
	var dstPix = dst.Pix()

	var srcStride = src.Stride()
	var dstStride = dst.Stride()

	var i, j = 0, 0
	for j < bltH {
		i = 0
		for i < bltW*4 {
			// DIB
			dstPix[dstI+i+0] = srcPix[srcI+i+0]
			dstPix[dstI+i+1] = srcPix[srcI+i+1]
			dstPix[dstI+i+2] = srcPix[srcI+i+2]
			dstPix[dstI+i+3] = srcPix[srcI+i+3]
			i += 4
		}
		srcI = srcI + srcStride
		dstI = dstI + dstStride
		j++
	}

	if src.Opaque() && !dst.Opaque() {
		dst.SetOpaque(true)
	}
}

func (dst *Canvas) AlphaBlendCanvas(x int, y int, src *Canvas, srcRc *image.Rectangle) {
	if srcRc == nil {
		var tmpRc = src.Bounds()
		srcRc = &tmpRc
	}

	var srcX, srcY = srcRc.Min.X, srcRc.Min.Y
	var dstX, dstY = x, y

	var bltW, bltH = srcRc.Dx(), srcRc.Dy()

	var srcI = src.PixOffset(srcX, srcY)
	var dstI = dst.PixOffset(dstX, dstY)

	var srcPix = src.Pix()
	var dstPix = dst.Pix()

	var srcStride = src.Stride()
	var dstStride = dst.Stride()

	var i, j = 0, 0

	for j < bltH {
		i = 0
		for i < bltW*4 {
			var sa = srcPix[srcI+i+3]
			if sa == 255 {
				dstPix[dstI+i+0] = srcPix[srcI+i+0]
				dstPix[dstI+i+1] = srcPix[srcI+i+1]
				dstPix[dstI+i+2] = srcPix[srcI+i+2]
				dstPix[dstI+i+3] = srcPix[srcI+i+3]
			} else if sa != 0 {
				// http://archive.gamedev.net/archive/reference/articles/article817.html
				var sr, sg, sb = &srcPix[srcI+i+0], &srcPix[srcI+i+1], &srcPix[srcI+i+2]
				var dr, dg, db = &dstPix[dstI+i+0], &dstPix[dstI+i+1], &dstPix[dstI+i+2]

				var alpha = int32(srcPix[srcI+i+3])

				*dr = byte((alpha*(int32(*sr)-int32(*dr)))/256) + *dr
				*dg = byte((alpha*(int32(*sg)-int32(*dg)))/256) + *dg
				*db = byte((alpha*(int32(*sb)-int32(*db)))/256) + *db

				dstPix[dstI+i+3] = 255
			}
			i += 4
		}
		srcI = srcI + srcStride
		dstI = dstI + dstStride
		j++
	}
}

func (dst *Canvas) DrawImageNRGBA(x int, y int, src *image.NRGBA, srcRc *image.Rectangle) {
	if srcRc == nil {
		srcRc = &(src.Rect)
	}

	var srcX, srcY = srcRc.Min.X, srcRc.Min.Y
	var dstX, dstY = x, y

	var bltW, bltH = srcRc.Dx(), srcRc.Dy()

	var srcI = src.PixOffset(srcX, srcY)
	var dstI = dst.PixOffset(dstX, dstY)

	var srcPix = src.Pix
	var dstPix = dst.Pix()

	var srcStride = src.Stride
	var dstStride = dst.Stride()

	var i, j = 0, 0

	for j < bltH {
		i = 0
		for i < bltW*4 {
			dstPix[dstI+i+0] = srcPix[srcI+i+0]
			dstPix[dstI+i+1] = srcPix[srcI+i+1]
			dstPix[dstI+i+2] = srcPix[srcI+i+2]
			dstPix[dstI+i+3] = srcPix[srcI+i+3]
			i += 4
		}
		srcI = srcI + srcStride
		dstI = dstI + dstStride
		j++
	}
}

func (dst *Canvas) DrawImageRGBA(x int, y int, src *image.RGBA, srcRc *image.Rectangle) {
	if srcRc == nil {
		srcRc = &(src.Rect)
	}

	var srcX, srcY = srcRc.Min.X, srcRc.Min.Y
	var dstX, dstY = x, y

	var bltW, bltH = srcRc.Dx(), srcRc.Dy()

	var srcI = src.PixOffset(srcX, srcY)
	var dstI = dst.PixOffset(dstX, dstY)

	var srcPix = src.Pix
	var dstPix = dst.Pix()

	var srcStride = src.Stride
	var dstStride = dst.Stride()

	var i, j = 0, 0

	for j < bltH {
		i = 0
		for i < bltW*4 {
			dstPix[dstI+i+0] = srcPix[srcI+i+0]
			dstPix[dstI+i+1] = srcPix[srcI+i+1]
			dstPix[dstI+i+2] = srcPix[srcI+i+2]
			dstPix[dstI+i+3] = srcPix[srcI+i+3]
			i += 4
		}
		srcI = srcI + srcStride
		dstI = dstI + dstStride
		j++
	}
}

func (dst *Canvas) DrawTexture(dstRc image.Rectangle, tex *Canvas, texRc image.Rectangle) {
	var texX, texY = texRc.Min.X, texRc.Min.Y
	var dstX, dstY = dstRc.Min.X, texRc.Min.Y

	var texW, texH = texRc.Dx(), texRc.Dy()
	var dstW, dstH = dstRc.Dx(), dstRc.Dy()

	var texI = tex.PixOffset(texX, texY)
	var dstI = dst.PixOffset(dstX, dstY)

	var texStride = tex.Stride()
	var dstStride = dst.Stride()

	var texPix = tex.Pix()
	var dstPix = dst.Pix()

	var i0, i1 = 0, 0
	var j0, j1 = 0, 0

	for {
		dstPix[dstI+i1+0] = texPix[texI+i0+0]
		dstPix[dstI+i1+1] = texPix[texI+i0+1]
		dstPix[dstI+i1+2] = texPix[texI+i0+2]
		dstPix[dstI+i1+3] = texPix[texI+i0+3]

		i0, i1 = i0+4, i1+4

		if i0 >= texW*4 {
			i0 = 0
		}

		if i1 >= dstW*4 {
			i0, i1 = 0, 0

			texI = texI + texStride
			dstI = dstI + dstStride

			j0, j1 = j0+1, j1+1

			if j0 >= texH {
				j0 = 0
				texI = tex.PixOffset(texX, texY)
			}

			if j1 >= dstH {
				break
			}
		}
	}
}
