//rm u78; goimports -v -w main.go sorting/*.go lastfm/*.go; gofmt -w main.go  sorting/*.go lastfm/*.go; go build; ./u78
//del u78.exe & goimports -v -w main.go ./sorting/sort.go ./sorting/interface.go & gofmt -w main.go ./sorting/sort.go ./sorting/interface.go & go build & u78.exe

package main

import (
	"fmt"
	"log"
	"net/http"

	h "gopl.io/ch7/u78/html"
	//s "gopl.io/ch7/u78/sorting"
)

func main() {
	fmt.Println("Starting webserver...")
	http.HandleFunc("/sort", h.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//test()
}

/*
func test() {
	list := fm.GetListFromLastfm()

	tracks := new(s.TrackListS)

	tracks.Add(list)

	tracks.SortS([]string{"title"})
	tracks.SortS([]string{"artist"})
	tracks.SortS([]string{"album"})

	s.PrintTracks(&tracks.List)
}*/

/*var list = [][]string{
	{"Go", "Delilah", "From the Roots Up", "2012", "3m38s"},
	{"Go", "Moby", "Moby", "1992", "3m37s"},
	{"Go Ahead", "Alicia Keys", "As I Am", "2007", "4m36s"},
	{"Ready 2 Go", "Martin Solveig", "Smash", "2011", "4m24s"},
	{"Go", "Def Leppard", "Def Leppard", "2008", "5m03s"},
	{"Go", "Def Leppard", "Def Leppard", "2012", "5m03s"},
	{"Go", "Def Leppard", "Leppard Def", "2011", "5m03s"},
}*/
