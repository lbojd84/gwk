package views

import (
	"log"
)

type NewViewFunc func() Viewer

var g_mock_up_map map[string]NewViewFunc = make(map[string]NewViewFunc)

func RegisterNewFuncToMockUp(typ string, new_func func() Viewer) {
	g_mock_up_map[typ] = new_func
}

func init_mockup() {
	g_mock_up_map["view"] = func() Viewer { return NewView() }
	g_mock_up_map["image_view"] = func() Viewer { return NewImageView() }
	g_mock_up_map["button"] = func() Viewer { return NewButton() }
	g_mock_up_map["panel"] = func() Viewer { return NewPanel() }
	g_mock_up_map["main_frame"] = func() Viewer { return NewMainFrame() }
	g_mock_up_map["toolbar"] = func() Viewer { return NewToolbar() }
}

func MockUp(ui UIMap) Viewer {
	typ, ok := ui.String("type")
	if !ok {
		return nil
	}

	var v Viewer

	new_view_func := g_mock_up_map[typ]
	if new_view_func != nil {
		v = new_view_func()
	} else {
		log.Printf("WARNING: Can't find view func for view type %s", typ)
		return nil
	}

	if id, ok := ui.String("id"); ok {
		v.SetID(id)
	}

	if intval, ok := ui.Int("width"); ok {
		v.SetWidth(intval)
	} else if strval, ok := ui.String("width"); ok {
		log.Printf("%v", strval)
	}

	if intval, ok := ui.Int("height"); ok {
		v.SetHeight(intval)
	} else if strval, ok := ui.String("height"); ok {
		log.Printf("%v", strval)
	}

	if intval, ok := ui.Int("left"); ok {
		v.SetLeft(intval)
	}

	if intval, ok := ui.Int("top"); ok {
		v.SetTop(intval)
	}

	// process the layout attribute.
	if val, ok := ui["layout"]; ok {
		if str_val, ok := val.(string); ok {
			if str_val == "vertical" {
				v.SetLayouter(g_vertical_layouter)
			} else if str_val == "horizontal" {
				v.SetLayouter(g_horizontal_layouter)
			} else {
				log.Panicf("WARNING: Unknown layout %v", str_val)
			}
		} else if func_val, ok := val.(LayoutFunc); ok {
			v.SetLayouter(NewFuncLayouter(func_val))
		}
	}

	v.SetUIMap(ui)

	if delegate, ok := ui.UIMap("delegate"); ok {
		v.SetDelegate(delegate)
	}

	// If the view has some view specifc attributes.
	v.MockUp(ui)

	children, ok := ui.UIMaps("children")
	for _, child := range children {
		typ, ok := child.String("type")
		if !ok {
			continue
		}
		if typ == "custom_view" {
			child_view, ok := child.Viewer("custom_view")
			if ok {
				v.AddChild(child_view)
			}
		} else {
			child_view := MockUp(child)
			if child_view != nil {
				v.AddChild(child_view)
			}
		}
	}

	return v
}

func hierarchy_mockup() {

}
