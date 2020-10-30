package artree

// ART implements The Adaptive Radix Tree.
type ART struct {

}

func (t *ART) Insert() {

}

func (t *ART) Search() {

}

func (t *ART) Delete() {

}

type Item interface {
	Less(than Item) bool
}

type ItemIterator func(i Item) bool

func (t *ART) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {

}
