package lastfm

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func GetListFromLastfm() [][]string {
	start := time.Now()
	tl, err := GetArtistTopTracks("queen", Limit)

	ch := make(chan *ArtistTrack)

	if err != nil {
		log.Println(err)
	}
	for _, track := range tl.Toptracks.Track {
		go func(t ArtistTopTrack) {
			ti, err := GetTrackInfo(t.Mbid, t.Artist.Name, t.Name)
			ti.SetRank(t.Attr.Rank)

			if err != nil {
				log.Println(err)
			}
			ch <- ti
		}(track)
	}
	l, err := strconv.Atoi(Limit)
	if err != nil {
		log.Println(err)
	}
	list := [][]string{}
	for i := 0; i < l; i++ {
		ti := <-ch
		n, err := strconv.Atoi(ti.Track.Duration)
		if err != nil {
			log.Println(err)
		}
		d := time.Duration(uint64(n) * 1000000)

		var elem []string
		elem = append(elem, ti.Track.Name, ti.Track.Artist.Name, ti.Track.Album.Title, ti.GetRank(), d.String())
		list = append(list, elem)
	}
	fmt.Printf("%.2f\n", time.Since(start).Seconds())

	return list
}

/*
var list = [][]string{
	{"Go", "Delilah", "From the Roots Up", "2012", "3m38s"},
	{"Go", "Moby", "Moby", "1992", "3m37s"},
	{"Go Ahead", "Alicia Keys", "As I Am", "2007", "4m36s"},
	{"Ready 2 Go", "Martin Solveig", "Smash", "2011", "4m24s"},
	{"Go", "Def Leppard", "Def Leppard", "2008", "5m03s"},
	{"Go", "Def Leppard", "Def Leppard", "2012", "5m03s"},
	{"Go", "Def Leppard", "Leppard Def", "2011", "5m03s"},
}
*/
