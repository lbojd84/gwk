package views

import (
	"log"
)

func MockUp(ui UIMap) Viewer {
	typ, ok := ui.String("type")
	if !ok {
		return nil
	}

	var v Viewer
	switch typ {
	case "main_frame":
		v = NewMainFrame()
	case "button":
		v = NewButton()
	case "image_view":
		v = NewImageView()
	default:
		log.Printf("Can't find view type %s", typ)
		return nil
	}

	if intval, ok := ui.Int("width"); ok {
		v.SetWidth(intval)
	} else if strval, ok := ui.String("width"); ok {
		// empty
		log.Printf("%v", strval)
	}

	if intval, ok := ui.Int("height"); ok {
		v.SetHeight(intval)
	} else if strval, ok := ui.String("height"); ok {
		// empty
		log.Printf("%v", strval)
	}

	if intval, ok := ui.Int("left"); ok {
		v.SetLeft(intval)
	}

	if intval, ok := ui.Int("top"); ok {
		v.SetTop(intval)
	}

	v.SetUIMap(ui)

	// if view has specified attrs.
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
