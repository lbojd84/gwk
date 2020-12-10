// Copyright 2013 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"gwk/vango"
	"image"
	"log"
)

type RootView struct {
	View
	host_window *HostWindow

	mouse_move_handler Viewer
}

func NewRootView(bounds image.Rectangle) *RootView {
	var v = new(RootView)
	v.SetID("cn.ustc.edu/gwk/root_view")
	v.SetBounds(bounds)
	return v
}

func (r *RootView) AddChild(child Viewer) {
	log.Printf("RootView.AddChild %v + %v", r.ID(), child.ID())
	r.children = append(r.children, child)
	child.SetParent(r)
}

func (r *RootView) SetHostWindow(h *HostWindow) {
	r.host_window = h
	if h.RootView() != r {
		h.SetRootView(r)
	}
}

func (r *RootView) ToAbsPt(pt image.Point) image.Point {
	return pt
}

func (r *RootView) ToAbsRect(rc image.Rectangle) image.Rectangle {
	return rc
}

func (r *RootView) ToDevicePt(pt image.Point) image.Point {
	log.Printf("NOTIMPLEMENT")
	return pt
}

func (r *RootView) ToDeviceRect(rc image.Rectangle) image.Rectangle {
	log.Printf("NOTIMPLEMENT")
	return rc
}

func (r *RootView) HostWindow() *HostWindow {
	return r.host_window
}

func DispatchDraw(v Viewer, canvas *vango.Canvas, dirty_rect image.Rectangle) {
	view_rect := v.Bounds()
	src_rect := view_rect.Intersect(dirty_rect)

	if src_rect.Empty() {
		return
	}

	v.OnDraw(v.Canvas())

	src_rect = src_rect.Sub(view_rect.Min)
	if v.Canvas().Opaque() {
		canvas.AlphaBlendCanvas(v.X(), v.Y(), v.Canvas(), &src_rect)
	} else {
		canvas.DrawCanvas(v.X(), v.Y(), v.Canvas(), &src_rect)
	}

	sub_canvas := canvas.SubCanvas(v.Bounds())
	for _, child := range v.Children() {
		child_rect := src_rect.Intersect(child.Bounds())
		if !child_rect.Empty() {
			DispatchDraw(child, sub_canvas, child_rect)
		}
	}
}

func (r *RootView) DispatchDraw(dirty_rect image.Rectangle) {
	if r.Children() != nil && len(r.Children()) >= 1 {
		DispatchDraw(r.Children()[0], r.Canvas(), dirty_rect)
	}
}

func DispatchLayout(v Viewer) {
	if v.Layouter() != nil {
		v.Layouter().Layout(v)
	}

	for _, child := range v.Children() {
		DispatchLayout(child)
	}
}

func (r *RootView) DispatchLayout() {
	new_rect := r.Bounds()

	if r.Children() == nil {
		return
	}

	r.Children()[0].SetXYWH(0, 0, new_rect.Dx(), new_rect.Dy())
	DispatchLayout(r.Children()[0])

	r.DispatchDraw(r.Bounds())
}

func get_event_handler_for_point(v Viewer, pt image.Point) Viewer {
	pt.X, pt.Y = pt.X-v.X(), pt.Y-v.Y()

	for _, child := range v.Children() {
		rect := child.Bounds()
		if rect.Min.X < pt.X && rect.Min.Y < pt.Y && rect.Max.X > pt.X &&
			rect.Max.Y > pt.Y {
			return get_event_handler_for_point(child, pt)
		}
	}

	return v
}

func (r *RootView) DispatchMouseMove(pt image.Point) {
	v := get_event_handler_for_point(r, pt)

	// for v != r.mouse_move_handler {
	// 	v = v.Parent()
	// }

	if v != nil && v != r && v != r.mouse_move_handler {
		old_handler := r.mouse_move_handler
		r.mouse_move_handler = v
		if old_handler != nil {
			old_handler.OnMouseLeave()
		}

		if r.mouse_move_handler != nil {
			r.mouse_move_handler.OnMouseEnter()
		}
	}

	if r.mouse_move_handler != nil {
		// r.mouse_move_handler.OnMouseMove()
	}
}

func (r *RootView) UpdateRect(rect image.Rectangle) {
	rect = rect.Intersect(r.Bounds())
	r.DispatchDraw(rect)
	r.host_window.InvalidateRect(rect)
}
