package component

import "html/template"

type Component interface {
	GetContent() template.HTML
}
