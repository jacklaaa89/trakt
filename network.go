package trakt

type Network struct {
	Name string `json:"name"`
}

type NetworkIterator struct{ BasicIterator }

func (n *NetworkIterator) Network() *Network { return n.Current().(*Network) }
