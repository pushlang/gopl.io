package html

import (
	"fmt"
	s "gopl.io/ch7/u78/sorting"
	"html/template"
	"log"
	"net/http"
	fm "gopl.io/ch7/u78/lastfm"
)
var url = "http://localhost:8000/sort?"
var tracklist = template.Must(template.New("tracklist").Parse(`
<h1>Track List</h1>
<table>
<tr style='text-align: left'>

<th>
	<a href='`+ url +`
	{{range .FormValues}}
		{{if ne . "title"}}
			{{.}}=true&
		{{end}}
	{{end}}title=true'>Title</a>
</th>
<th>
<a href='`+ url +`
	{{range .FormValues}}
		{{if ne . "artist"}}
			{{.}}=true&
		{{end}}
	{{end}}artist=true'>Artist</a>
</th>
<th>
	<a href='`+ url +`
	{{range .FormValues}}
		{{if ne . "album"}}
			{{.}}=true&
		{{end}}
	{{end}}album=true'>Album</a>
</th>
<th>
	<a href='`+ url +`
	{{range .FormValues}}
		{{if ne . "year"}}
			{{.}}=true&
		{{end}}
	{{end}}year=true'>Year</a>
</th>
<th>
<a href='`+ url +`
	{{range .FormValues}}
		{{if ne . "length"}}
			{{.}}=true&
		{{end}}
	{{end}}length=true'>Length</a></th>
</tr>
{{range .List}}
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

//var list [][]string

var tracks = new(s.TrackListS)

func init() {
	list := fm.GetListFromLastfm()
	tracks.Add(list)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var tr = new(s.TrackListS)
	*tr = *tracks
	
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	args := []string{}
	for k, _ := range r.Form {
		args = append(args, k)
	}
	tr.SortS(args)

	if err := tracklist.Execute(w, tr); err != nil {
		log.Fatal(err)
	}
	for _, a := range args {
		fmt.Fprintf(w, a+" ")
	}
}
