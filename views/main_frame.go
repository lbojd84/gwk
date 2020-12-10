// Copyright 2013 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"gwk/vango"
	"gwk/views/R"
)

type MainFrame struct {
	View
	canvas_bkg *vango.Canvas

	left_panel Viewer
	main_panel Viewer

	left_pos int
}

func NewMainFrame() *MainFrame {
	var v = new(MainFrame)
	v.SetID("main_frame")
	v.SetLayouter(v)
	v.SetXYWH(0, 0, 50, 50)
	v.canvas_bkg = R.LoadCanvas("data/texture.png")
	v.Canvas().DrawTexture(v.Canvas().Bounds(), v.canvas_bkg, v.canvas_bkg.Bounds())
	v.left_pos = 150
	return v
}

func (v *MainFrame) MockUp(ui UIMap) {
	if left_panel, ok := ui.UIMap("left_panel"); ok {
		v.left_panel = MockUp(left_panel)
		v.AddChild(v.left_panel)
	}

	if main_panel, ok := ui.UIMap("main_panel"); ok {
		v.main_panel = MockUp(main_panel)
		v.AddChild(v.main_panel)
	}
}

func (v *MainFrame) Layout(parent Viewer) {
	v.Canvas().DrawTexture(v.Canvas().Bounds(), v.canvas_bkg, v.canvas_bkg.Bounds())
	x, y, w, h := v.X(), v.Y(), v.W(), v.H()
	if v.left_panel != nil {
		v.left_panel.SetXYWH(x+5, y+5, v.left_pos-10, h-10)
	}
	if v.main_panel != nil {
		v.main_panel.SetXYWH(v.left_pos, y+5, w-v.left_pos-5, h-10)
	}
}
