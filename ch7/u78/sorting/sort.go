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

func (t *TrackList) Sort(sortby string) {
	var less func(x, y *Track, sortby []string) bool
	less = sortByField

	switch sortby {
	case "custom":
	case "artist":
	case "year":
	default:
		panic(sortby)
	}
	sort.Sort(custom{t, less, []string{sortby}})
	printTracks(t)
}

func (t *TrackList) SortMultiple(sortby []string) {
	var less func(x, y *Track, sortby []string) bool
	less = sortByMultiple
	
	sb := make([]string, 0)
	for _, by := range sortby {
		switch by {
		case "title":
			sb = append(sb, "title")
		case "year":
			sb = append(sb, "year")
		default:
			panic(by)
		}
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
		case "year":
			if x.Year != y.Year {
				return x.Year < y.Year
			}
		}
	}
	return false
}

func sortByCustom(x, y *Track, sortby []string) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	if x.Year != y.Year {
		return x.Year < y.Year
	}
	if x.Length != y.Length {
		return x.Length < y.Length
	}
	return false
}

func sortByField(x, y *Track, sortby []string) bool {
	switch sortby[0] {
	case "artist":
		return x.Artist < y.Artist
	case "year":
		return x.Year < y.Year
	case "title":
		return x.Title < y.Title
	default:
		panic(sortby[0])
	}
}
