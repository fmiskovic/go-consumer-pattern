// Description: This is a beer consumer. Doesn't care about other drinks, just wants a beer.
package beer

type Bar interface {
	IWantBeer() string
}

type Beer struct {
	bar Bar
}

func New(b Bar) Beer {
	return Beer{b}
}

func (b Beer) Order() string {
	return b.bar.IWantBeer()
}
