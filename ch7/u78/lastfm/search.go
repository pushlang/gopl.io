package lastfm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetTrackInfo(trackMbid, artistName, trackName string) (*ArtistTrack, error) {
	var meth = trackGetInfo

	meth.params["artist"] = artistName
	meth.params["track"] = trackName
	meth.params["mbid"] = trackMbid

	var result ArtistTrack
	if err := json.NewDecoder(mustQuery(&meth)).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
func GetArtistTopTracks(artistName, limit string) (*ArtistTopTracks, error) {
	var meth = artistGettoptracks

	meth.params["artist"] = artistName
	meth.params["limit"] = limit

	var result ArtistTopTracks
	if err := json.NewDecoder(mustQuery(&meth)).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
func mustQuery(m *method) io.Reader {
	q := fmt.Sprintf("%s?method=%s&api_key=%s&format=json", URL, m.name, Key)
	for key, value := range m.params {
		q += fmt.Sprintf("&%s=%s", url.QueryEscape(key), url.QueryEscape(value))
	}
	resp, err := http.Get(q)
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			panic("PANIC: search query failed: " + resp.Status)
		}
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
		panic("PANIC: http.Get(q) -> err!=nil")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic("PANIC: ioutil.ReadAll(resp.Body) -> err != nil")
	}
	return bytes.NewReader(body)
}
