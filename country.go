package trakt

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountryIterator struct{ Iterator }

func (c *CountryIterator) Country() *Country { return c.Current().(*Country) }
