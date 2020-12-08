package components

import "html/template"

type Component interface {
	GetContent() template.HTML
}
