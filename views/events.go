// Copyright 2013 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"gwk/vango"
	"image"
)

type PaintEvent struct {
	Canvas    *vango.Canvas
	DirtyRect image.Rectangle
}
