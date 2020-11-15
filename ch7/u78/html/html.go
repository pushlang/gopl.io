package html

import (
	"fmt"
	s "gopl.io/ch7/u78/sorting"
	"html/template"
	"log"
	"net/http"
)

var tracklist *template.Template

var tracks *s.TrackList

func init() {
	tracklist = template.Must(template.New("issuelist").Parse(`
<h1>Track List</h1>
<table>
<tr style='text-align: left'>
  <th><a href='http://localhost:8000/sort?title=true'>Title</a></th>
  <th><a href='http://localhost:8000/sort?artist=true'>Artist</a></th>
  <th><a href='http://localhost:8000/sort?album=true'>Album</a></th>
  <th><a href='http://localhost:8000/sort?year=true'>Year</a></th>
  <th><a href='http://localhost:8000/sort?length=true'>Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

	var list = [][]string{
		{"Go", "Delilah", "From the Roots Up", "2012", "3m38s"},
		{"Go", "Moby", "Moby", "1992", "3m37s"},
		{"Go Ahead", "Alicia Keys", "As I Am", "2007", "4m36s"},
		{"Ready 2 Go", "Martin Solveig", "Smash", "2011", "4m24s"},
		{"Go", "Def Leppard", "Def Leppard", "2008", "5m03s"},
		{"Go", "Def Leppard", "Def Leppard", "2012", "5m03s"},
		{"Go", "Def Leppard", "Leppard Def", "2011", "5m03s"},
	}

	tracks = new(s.TrackList)

	tracks.Add(list)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	args := []string{}
	for k, _ := range r.Form {
		args = append(args, k)
	}

	tracks.Sort(args)

	if err := tracklist.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, args[0])
}
