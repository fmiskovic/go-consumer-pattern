package wine

type Bar interface {
	IWantWine() string
}

type Wine struct {
	bar Bar
}

func New(b Bar) Wine {
	return Wine{b}
}

func (w Wine) Order() string {
	return w.bar.IWantWine()
}
