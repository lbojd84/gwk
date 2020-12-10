// Copyright 2013 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"gwk/vango"
	"image"
)

type Viewer interface {
	ID() string
	SetID(id string)

	AddChild(child Viewer)
	Children() []Viewer

	Parent() Viewer
	SetParent(parent Viewer)

	Canvas() *vango.Canvas
	SetCanvas(canvas *vango.Canvas)

	ToAbsPt(pt image.Point) image.Point
	ToAbsRect(rc image.Rectangle) image.Rectangle
	ToDevicePt(pt image.Point) image.Point
	ToDeviceRect(rc image.Rectangle) image.Rectangle

	OnDraw(canvas *vango.Canvas)

	OnMouseIn()
	OnMouseOut()

	UpdateView()
	UpdateRect(dirty image.Rectangle)
	NeedsUpdate() bool
	SetNeedsUpdate(needs_update bool)

	X() int
	Y() int
	W() int
	H() int
	XYWH() (x, y, w, h int)
	SetXYWH(x, y, w, h int)

	Left() int
	Top() int
	Width() int
	Height() int
	SetLeft(left int)
	SetTop(top int)
	SetWidth(width int)
	SetHeight(height int)

	Bounds() image.Rectangle
	SetBounds(bounds image.Rectangle)

	Attrs() UIMap
	SetAttrs(attrs UIMap)
	MockUp(ui UIMap)

	Layouter() Layouter
	SetLayouter(l Layouter)
}
