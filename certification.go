package trakt

type Certification struct {
	Name        string `json:"name"`
	Slug        Slug   `json:"slug"`
	Description string `json:"description"`
}

type CertificationIterator struct{ Iterator }

func (c *CertificationIterator) Certification() *Certification { return c.Current().(*Certification) }
