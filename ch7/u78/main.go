// rm u78; goimports -v -w *.go; gofmt -w *.go; go build; ./u78
package main

import (
	s "gopl.io/ch7/u78/sorting"
)

func main() {
	var list = [][]string{
		{"Go", "Delilah", "From the Roots Up", "2012", "3m38s"},
		{"Go", "Moby", "Moby", "1992", "3m37s"},
		{"Go Ahead", "Alicia Keys", "As I Am", "2007", "4m36s"},
		{"Ready 2 Go", "Martin Solveig", "Smash", "2011", "4m24s"},
		{"Go", "Def Leppard", "Def Leppard", "2008", "5m03s"},
		{"Go", "Def Leppard", "Def Leppard", "2012", "5m03s"},
		{"Go", "Def Leppard", "Leppard Def", "2011", "5m03s"},
	}

	tracks := new(s.TrackList)

	tracks.Add(list)

	tracks.Sort([]string{"title", "album", "year"}) // artist, year, custom
}
