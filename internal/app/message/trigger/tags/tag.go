package tags

type Tagger interface {
	Handle(text string)
}

func MapTagger(text string) Tagger {
	m := map[string]Tagger{
		"issue": NewIssue(),
		"todo":  NewTodo(),
	}
	if t, ok := m[text]; ok {
		return t
	}
	return nil
}
