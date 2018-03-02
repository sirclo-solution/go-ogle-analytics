package ga

import (
	"net/url"
)

//Event Hit Type
type Event struct {
	category        string
	action          string
	label           string
	labelSet        bool
	value           int64
	valueSet        bool
	fieldsObject    map[string]string
	fieldsObjectSet bool
}

// NewEvent creates a new Event Hit Type.
// Specifies the event category.
// Specifies the event action.

func NewEvent(category string, action string) *Event {
	h := &Event{
		category: category,
		action:   action,
	}
	return h
}

func (h *Event) addFields(v url.Values) error {
	v.Add("ec", h.category)
	v.Add("ea", h.action)
	if h.labelSet {
		v.Add("el", h.label)
	}
	if h.valueSet {
		v.Add("ev", int2str(h.value))
	}
	if h.fieldsObjectSet {
		for name, val := range h.fieldsObject {
			v.Add(name, val)
		}
	}

	return nil
}

// Specifies the event label.
func (h *Event) Label(label string) *Event {
	h.label = label
	h.labelSet = true
	return h
}

// Specifies the event value. Values must be non-negative.
func (h *Event) Value(value int64) *Event {
	h.value = value
	h.valueSet = true
	return h
}

//Specify additional fields to the event
//Currently supporting dimension only, Protocol name Format :"cd<index No.>"
func (h *Event) FieldsObject(obj map[string]string) *Event {
	h.fieldsObject = obj
	h.fieldsObjectSet = true
	return h
}

func (h *Event) Copy() *Event {
	c := *h
	return &c
}
