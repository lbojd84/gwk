// Copyright 2013 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"gwk/vango"
	"image/color"
	"log"
	"math/rand"
)

type ImageView struct {
	View
	clr color.RGBA
}

func NewImageView() *ImageView {
	var v = new(ImageView)
	v.SetID("image_view")
	v.SetLayouter(v)
	v.SetXYWH(0, 0, 50, 50)
	return v
}

func (v *ImageView) MockUp(ui UIMap) {
	if clr, ok := ui.Int("color"); ok {
		val := uint(clr)
		v.clr.R = byte((val & 0xff0000) >> 16)
		v.clr.G = byte((val & 0x00ff00) >> 8)
		v.clr.B = byte(val & 0x0000ff)
		v.clr.A = 0x00
	}
}

func (v *ImageView) Layout(parent Viewer) {

}

func (v *ImageView) OnMouseIn() {
	v.clr.R = uint8(rand.Intn(255) + 1)
	v.clr.G = uint8(rand.Intn(255) + 1)
	v.clr.B = uint8(rand.Intn(255) + 1)
	v.clr.A = uint8(rand.Intn(255) + 1)
	v.UpdateView()
}

func (v *ImageView) OnDraw(canvas *vango.Canvas) {
	log.Printf("image view draw")
	canvas.DrawColor(v.clr.R, v.clr.G, v.clr.B)
}
