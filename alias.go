package trakt

type Alias struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type AliasIterator struct{ BasicIterator }

func (a *AliasIterator) Alias() (*Alias, error) {
	rcv := &Alias{}
	return rcv, a.Scan(rcv)
}
