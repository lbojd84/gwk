// Copyright 2014 By Jshi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package main

import (
	. "gwk/views"
	"image"
	"math/rand"
	"time"
)

func makeMainUIMap() UIMap {
	var main_uimap = UIMap{
		"type": "main_frame",
		"left_panel": UIMap{
			"type":   "image_view",
			"color":  0xe6e6e6,
			"width":  100,
			"height": "fill_parent",
		},
		"main_panel": UIMap{
			"type":  "image_view",
			"color": 0x272822,
			"children": []UIMap{
				{
					"type":   "image_view",
					"left":   50,
					"top":    120,
					"width":  570,
					"height": 300,
				},
			},
		},
	}

	return main_uimap
}

func main() {
	rand.Seed(time.Now().Unix())

	var hv = NewHostView(image.Rect(0, 0, 840, 610))

	v := MockUp(makeMainUIMap())
	hv.RootView.AddChild(v)

	hv.Show()
	hv.Run()
}
