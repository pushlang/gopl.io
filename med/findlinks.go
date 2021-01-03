package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopl.io/med/links"
	"gopl.io/med/testsrv"
)

func saveWebpage(link links.Link) error {
	log.Println("enter saveWebpage:", link.Url)
	defer log.Println("exit saveWebpage", link.Url)

	req, err := http.NewRequest("GET", link.Url, nil)

	/*hd := req.Header
	hd.Set("Host", " medium.com")
	hd.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0")
	hd.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*\/*;q=0.8")
	hd.Set("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3")
	//hd.Set("Accept-Encoding", "gzip, deflate, br")
	hd.Set("DNT", "1")
	hd.Set("Connection", "keep-alive")
	hd.Set("Cookie", "__cfduid=da165d5b344a72b3405b7d99817e868791609619300; uid=lo_8785774c852d; sid=1:jEiQ1tYc45EZ+p3kLVXz9T2vRmn1Py6YEiCQd5xeJPJfWfjpdyrrY3oURCdCHSLM; optimizelyEndUserId=lo_8785774c852d; __cfruid=4ace85550add54750ced638baa9e977145708d13-1609619301; lightstep_guid/lite-web=332832e57f53a547; lightstep_session_id=28d1533b0cec4500; _parsely_session={%22sid%22:1%2C%22surl%22:%22https://medium.com/rungo/building-rpc-remote-procedure-call-network-in-go-5bfebe90f7e9%22%2C%22sref%22:%22%22%2C%22sts%22:1609619355694%2C%22slts%22:0}; _parsely_visitor={%22id%22:%22pid=13d6597594069214e6f2007e5d0e51c1%22%2C%22session_count%22:1%2C%22last_session_ts%22:1609619355694}")
	hd.Set("Upgrade-Insecure-Requests", "1")
	hd.Set("Cache-Control", "max-age=0")*/

	if err != nil {
		return errors.New(fmt.Sprintf("Error - Request not sent: %s => %v", link.Url, err))
	}

	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	now := time.Now()
	for ; err != nil; time.Sleep(10 * time.Millisecond) {
		if time.Since(now).Milliseconds() > 1000 {
			break
		}
		resp, err = http.DefaultClient.Do(req)
	}
	if err != nil {
		return errors.New(fmt.Sprintf("Error - Request timeout: %s => %v", link.Url, err))
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Error - Status code: %s => %v", link.Url, err))
	}
	fn := filepath.Base(link.Url)
	fd, err := os.Create("downloads/" + fn + ".html")
	if err != nil {
		return errors.New(fmt.Sprintf("Error - File not created: %s => %v", fn, err))

	}
	cleanedBody, err := links.CleanHTMLbody(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error - Html not cleaned: %s => %v", fn, err))
	}

	_, err = io.Copy(fd, cleanedBody)
	if err != nil {
		return errors.New(fmt.Sprintf("Error - Html not saved: %s => %v", fn, err))
	}
	fd.Close()

	return nil
}

func findLinks(wl links.Extractor, done chan struct{}) {
	worklist, err := wl.Extract()
	log.Printf("enter findLinks: %d", len(worklist))
	defer log.Printf("exit findLinks: %d", len(worklist))

	if err != nil {
		log.Printf("Error - Extract: %v", err)
	}

	for _, w := range worklist {
		if err := saveWebpage(w); err != nil {
			log.Printf("Link didn't saved: %s(%v)\n", w.Url, err)
		} else {
			log.Printf("Link saved: %s\n", w.Url)
		}
		//fmt.Printf("<a href=\"%s\">%s</a><br>\n", w.Url, w.Name)
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	if len(os.Args) == 1 {
		log.Fatal("No arguments")
	}

	var ex links.Extractor

	if os.Args[1] == "web" {
		log.Println("Web extractor")
		if len(os.Args) <= 2 {
			ex = links.Url("http://127.0.0.1:8000")
			log.Println("Using default http://127.0.0.1:8000")
		} else {
			ex = links.Url(os.Args[2])
			log.Println("using " + os.Args[2])
		}
	} else if os.Args[1] == "file" {
		log.Println("File extractor")
		ex = links.FileName(os.Args[2])
	} else {
		log.Fatal("Unknown action")
	}
	go findLinks(ex, done)

	if os.Args[1] == "web" && len(os.Args) <= 2 {
		testsrv.Run()
	}
	<-done
}
