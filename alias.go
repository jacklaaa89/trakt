package trakt

type Alias struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type AliasIterator struct{ BasicIterator }

func (a *AliasIterator) Alias() *Alias { return a.Current().(*Alias) }
