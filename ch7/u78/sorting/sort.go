package sorting

import (
	"sort"
	"strconv"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var fields = map[string]bool{"title": true, "artist": true, "album": true, "year": true, "length": true}

type TrackList []*Track

func (t *TrackList) Add(tracks [][]string) {
	for i, track := range tracks {
		*t = append(*t, &Track{})
		for j, field := range track {
			switch j {
			case 0:
				(*t)[i].Title = field
			case 1:
				(*t)[i].Artist = field
			case 2:
				(*t)[i].Album = field
			case 3:
				(*t)[i].Year, _ = strconv.Atoi(field)
			case 4:
				(*t)[i].Length = parseLength(field)
			}
		}
	}
}

func (t *TrackList) Sort(sortby []string) {
	var less func(x, y *Track, sortby []string) bool
	less = sortByMultiple

	sb := make([]string, 0)
	for _, by := range sortby {
		if !fields[by] {
			panic(by)
		}
		sb = append(sb, by)
	}
	sort.Sort(custom{t, less, sb})
	printTracks(t)
}

func sortByMultiple(x, y *Track, sortby []string) bool {
	for _, by := range sortby {
		switch by {
		case "title":
			if x.Title != y.Title {
				return x.Title < y.Title
			}
		case "artist":
			if x.Year != y.Year {
				return x.Year < y.Year
			}
		case "album":
			if x.Album != y.Album {
				return x.Album < y.Album
			}
		case "year":
			if x.Year != y.Year {
				return x.Year < y.Year
			}
		case "length":
			if x.Length != y.Length {
				return x.Length < y.Length
			}
		}
	}
	return false
}
