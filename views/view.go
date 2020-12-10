package views

import (
	"gwk/vango"
	"image"
	"log"
)

type View struct {
	id string

	canvas *vango.Canvas

	x, y int // Relative to Parent.
	w, h int

	children []Viewer
	parent   Viewer

	attrs    UIMap
	layouter Layouter

	needs_update bool
}

func NewView() *View {
	var v = new(View)
	return v
}

func (v *View) ID() string {
	return v.id
}

func (v *View) SetID(id string) {
	v.id = id
}

func (v *View) AddChild(child Viewer) {
	v.children = append(v.children, child)
	child.SetParent(v)
	child.SetNeedsUpdate(true)
}

func (v *View) Children() []Viewer {
	return v.children
}

func (v *View) Parent() Viewer {
	return v.parent
}

func (v *View) SetParent(parent Viewer) {
	v.parent = parent
}

func (v *View) Canvas() *vango.Canvas {
	if v.canvas == nil {
		v.canvas = vango.NewCanvas(v.W(), v.H())
	}

	canvas_bounds := v.canvas.Bounds()

	if canvas_bounds.Dx() < v.W() || canvas_bounds.Dy() < v.H() {
		v.canvas = vango.NewCanvas(v.W(), v.H())
	} else {
		return v.canvas.SubCanvas(image.Rect(0, 0, v.W(), v.H()))
	}

	return v.canvas
}

func (v *View) SetCanvas(canvas *vango.Canvas) {
	v.canvas = canvas
}

func (v *View) ToAbsPt(pt image.Point) image.Point {
	if v.Parent() == nil {
		return pt
	}
	pt.X = pt.X + v.X()
	pt.Y = pt.Y + v.Y()
	return v.Parent().ToAbsPt(pt)
}

func (v *View) ToAbsRect(rc image.Rectangle) image.Rectangle {
	if v.Parent() == nil {
		return rc
	}
	rc.Min.X = rc.Min.X + v.X()
	rc.Min.Y = rc.Min.Y + v.Y()
	rc.Max.X = rc.Max.X + v.X()
	rc.Max.Y = rc.Max.Y + v.Y()
	return v.Parent().ToAbsRect(rc)
}

func (v *View) ToDevicePt(pt image.Point) image.Point {
	pt = pt.Add(image.Pt(v.X(), v.Y()))
	return v.Parent().ToDevicePt(pt)
}

func (v *View) ToDeviceRect(r image.Rectangle) image.Rectangle {
	r = r.Add(image.Pt(v.X(), v.Y()))
	return v.Parent().ToDeviceRect(r)
}

func update_rect(v Viewer, rect image.Rectangle) {
	if v.Parent() == nil {
		v.UpdateRect(rect)
		return
	}

	rect = rect.Add(image.Pt(v.X(), v.Y()))
	update_rect(v.Parent(), rect)
}

func (v *View) UpdateView() {
	v.SetNeedsUpdate(true)
	if v.Parent() != nil {
		update_rect(v.Parent(), v.Bounds())
	}
}

func (v *View) UpdateRect(rect image.Rectangle) {
	v.SetNeedsUpdate(true)
	if v.Parent() != nil {
		update_rect(v.Parent(), rect)
	}
}

func (v *View) NeedsUpdate() bool {
	return v.needs_update
}

func (v *View) SetNeedsUpdate(needs_update bool) {
	v.needs_update = needs_update
}

func (v *View) X() int {
	return v.x
}

func (v *View) Y() int {
	return v.y
}

func (v *View) W() int {
	return v.w
}

func (v *View) H() int {
	return v.h
}

func (v *View) XYWH() (x, y, w, h int) {
	x, y, w, h = v.x, v.y, v.w, v.h
	return
}

func (v *View) SetXYWH(x, y, w, h int) {
	v.x, v.y, v.w, v.h = x, y, w, h
}

func (v *View) Width() int {
	return v.W()
}

func (v *View) SetWidth(width int) {
	v.w = width
}

func (v *View) Height() int {
	return v.H()
}

func (v *View) SetHeight(height int) {
	v.h = height
}

func (v *View) Left() int {
	return v.X()
}

func (v *View) SetLeft(left int) {
	v.x = left
}

func (v *View) Top() int {
	return v.Y()
}

func (v *View) SetTop(top int) {
	v.y = top
}

func (v *View) Bounds() image.Rectangle {
	return image.Rect(v.x, v.y, v.x+v.w, v.y+v.h)
}

func (v *View) SetBounds(bounds image.Rectangle) {
	v.x, v.y = bounds.Min.X, bounds.Min.Y
	v.w, v.h = bounds.Dx(), bounds.Dy()
}

func (v *View) Layouter() Layouter {
	return v.layouter
}

func (v *View) SetLayouter(layouter Layouter) {
	v.layouter = layouter
}

func (v *View) OnDraw(canvas *vango.Canvas) {
	log.Printf("View.OnDraw()")
}

func (v *View) OnMouseIn() {
	log.Printf("View.OnMouseIn()")
}

func (v *View) OnMouseOut() {
	log.Printf("View.OnMouseOut()")
}

func (v *View) SetAttrs(attrs UIMap) {
	v.attrs = attrs
}

func (v *View) Attrs() UIMap {
	return v.attrs
}

func (v *View) MockUp(ui UIMap) {
	return
}
