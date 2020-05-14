package trakt

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountryIterator struct{ BasicIterator }

func (c *CountryIterator) Country() *Country { return c.Current().(*Country) }
