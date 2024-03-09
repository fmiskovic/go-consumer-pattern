// Description: This is a whisky consumer. Doesn't care about other drinks, just wants a whisky.
package whisky

type Bar interface {
	IWantWhisky() string
}

type Whisky struct {
	bar Bar
}

func New(b Bar) Whisky {
	return Whisky{b}
}

func (w Whisky) Order() string {
	return w.bar.IWantWhisky()
}
