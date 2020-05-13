package trakt

type Network struct {
	Name string `json:"name"`
}

type NetworkIterator struct{ Iterator }

func (n *NetworkIterator) Network() *Network { return n.Current().(*Network) }
