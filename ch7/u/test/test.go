package test

type FGer interface {
	Fer
	Ger
}

type Fer interface {
	F() ff
}

type Ger interface {
	G() gg
}

type ff int
type gg int

type tt struct {
	*ff
	*gg
}

func New() FGer {
	f:=ff(123)
	g:=gg(456)
	return &tt{&f, &g}
}

func (f *ff) F() ff {
	return *f
}

func (g *gg) G() gg {
	return *g
}