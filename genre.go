package trakt

type Genre struct {
	Name string `json:"name"`
	Slug Slug   `json:"slug"`
}

type GenreIterator struct{ Iterator }

func (c *GenreIterator) Genre() *Genre { return c.Current().(*Genre) }
