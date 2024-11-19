package etree

import (
	"strings"
)

func (e *Element) TrimmedText() string {
	var text string
	for _, child := range e.Child {
		if cd, ok := child.(*CharData); ok && !cd.whitespace {
			text = text + cd.Data
		}
	}
	return strings.TrimSpace(text)
}
