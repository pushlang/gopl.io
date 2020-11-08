package sorting

import (
	_ "sort"
)

// sort.Interface implementing

type custom struct {
	t      *TrackList
	less   func(x, y *Track, s []string) bool
	sortby []string
}

func (x custom) Len() int           { return len(*x.t) }
func (x custom) Less(i, j int) bool { return x.less((*x.t)[i], (*x.t)[j], x.sortby) }
func (x custom) Swap(i, j int)      { (*x.t)[i], (*x.t)[j] = (*x.t)[j], (*x.t)[i] }
