package trakt

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type LanguageIterator struct{ BasicIterator }

func (c *LanguageIterator) Language() *Language { return c.Current().(*Language) }
