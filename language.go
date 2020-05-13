package trakt

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type LanguageIterator struct{ Iterator }

func (c *LanguageIterator) Language() *Language { return c.Current().(*Language) }
