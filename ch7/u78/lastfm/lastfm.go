package lastfm

const URL = "http://ws.audioscrobbler.com/2.0/"
const Key = "967a950503428759d47a78db62bbf6fd"
const Limit = "25"

type method struct {
	name   string
	params map[string]string
}

var trackGetInfo = method{
	name:   "track.getInfo",
	params: map[string]string{},
}
var artistGettoptracks = method{
	name:   "artist.gettoptracks",
	params: map[string]string{},
}

//art.gtt//////////////////////////////////////////////////
type ArtistTopTracks struct {
	Toptracks ArtistTopTracksList
}
type ArtistTopTracksList struct {
	Track []ArtistTopTrack
	Attr  ArtistTopTracksListAttr
}
type ArtistTopTracksListAttr struct {
	Artist     string
	Page       string
	PerPage    string
	TotalPages string
	Total      string
}
type ArtistTopTrack struct {
	Mbid       string
	Name       string
	Artist     Artists
	Streamable string
	Attr       ArtistTopTrackAttr `json:"@attr"`
}
type ArtistTopTrackAttr struct {
	Rank string
}

//trk.gi//////////////////////////////////////////////////
type ArtistTrack struct {
	Track Tracks
}
type Tracks struct {
	Mbid            string
	Name            string
	Url             string
	Duration        string
	Streamable      Streamables
	Listeners       string
	Playcount       string
	Artist          Artists
	Album           Albums
	Toptags         Toptagss
	Wiki            Wikis
	Image           []Images
	artistTrackRank string
}

func (t *ArtistTrack) SetRank(r string) {
	t.Track.artistTrackRank = r
}
func (t *ArtistTrack) GetRank() string {
	return t.Track.artistTrackRank
}

type Streamables struct {
	Text      string `json:"#text"`
	Fulltrack string
}
type Artists struct {
	Mbid string
	Name string
	Url  string
}
type Albums struct {
	Mbid   string
	Artist string
	Title  string
	Url    string
	Image  []Images
}
type Toptagss struct {
	Tag []Tags
}
type Tags struct {
	Name string
	Url  string
}
type Wikis struct {
	Published string
	Summary   string
	Content   string
}
type Images struct {
	Text string `json:"#text"`
	Size string
}
