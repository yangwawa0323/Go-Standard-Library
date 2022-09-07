package texttemplate

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"text/template"
)

var templateDirListing, templateNotFound *template.Template
var basePath string

type listing struct {
	Path   string
	Parent string
	Files  []*listingEntry
}

type listingEntry struct {
	IsDir    bool
	IsLink   bool
	Fullname string
	Basename string
	Size     string
	Modified string
}

const (
	KB = 1024 << (10 * iota)
	MB
	GB
	TB
)

func formatFilesize(n int64) string {
	switch {
	case n < KB:
		return fmt.Sprintf("%d B", n)
	case n < MB:
		return fmt.Sprintf("%.2f KB", float64(n)/KB)
	case n < GB:
		return fmt.Sprintf("%.2f MB", float64(n)/MB)
	case n < TB:
		return fmt.Sprintf("%.2f GB", float64(n)/GB)
	}
	return fmt.Sprintf("%.2f TB", float64(n)/TB)
}

func getParentDir(requestPath string) (parentDir string) {
	if requestPath != "/" {
		// eat string up to the last forward slash
		// unless it's the last character in the path
		i := len(requestPath) - 1
		index := i
		for {
			if requestPath[i] == '/' && i < index {
				// include the slash, too
				index = i + 1
				break
			}
			i--
		}
		parentDir = requestPath[0:index]
	}
	return
}

func errorNotFound(w http.ResponseWriter, path string) {
	templateNotFound.Execute(w, path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//	if referer, ok := r.Header["Referer"]; ok {
	//
	//	}
	requestPath := r.URL.Path
	completePath := basePath + requestPath
	fmt.Fprintf(os.Stdout, "Access the path: %s\n\n", completePath)
	file, err := os.Stat(completePath)
	if err != nil {
		errorNotFound(w, requestPath)
		return
	}

	if file.IsDir() {
		// make sure all of our requestPaths have a / at the end
		if requestPath[len(requestPath)-1] != '/' {
			requestPath += "/"
			completePath += "/"
		}

		// we don't want to show the actual root, the apparent root will actually be in the given basePath
		dir, err := ioutil.ReadDir(completePath)
		if err != nil {
			errorNotFound(w, requestPath)
			return
		}

		entries := make([]*listingEntry, len(dir))
		for id, this := range dir {
			thisFullPath := completePath + this.Name()

			// if there exists an index.html in the current directory,
			// serve it up instead of the listing
			// TODO: Make this configurable?
			if this.Name() == "index.html" {
				http.ServeFile(w, r, thisFullPath)
				return
			}

			isLink := this.Mode()&os.ModeSymlink != 0
			isDir := this.IsDir()
			// someone please look at this and help me DRY it up!
			var size string
			if isDir {
				thisDirInfo, _ := ioutil.ReadDir(thisFullPath)
				size = fmt.Sprintf("%v items", len(thisDirInfo))
			} else if !isLink {
				size = formatFilesize(this.Size())
			} else {
				thisLink, err := os.Stat(thisFullPath)
				if err != nil {
					size = "???"
				} else if thisLink.IsDir() {
					isDir = true
					thisLinkInfo, _ := ioutil.ReadDir(thisFullPath)
					size = fmt.Sprintf("%v items", len(thisLinkInfo))
				} else {
					size = formatFilesize(thisLink.Size())
				}
			}

			entries[id] = &listingEntry{
				IsDir:    isDir,
				IsLink:   isLink,
				Fullname: strings.Replace(requestPath+this.Name(), " ", "%20", -1),
				Basename: this.Name(),
				Size:     size,
				Modified: this.ModTime().Format("02 Jan 2006 15:04"),
			}
		}

		templateDirListing.Execute(w, &listing{requestPath, getParentDir(requestPath), entries})
	} else {
		http.ServeFile(w, r, completePath)
	}
}

func Test_HTTP_ListDir(t *testing.T) {
	HOME := os.Getenv("HOME")
	if HOME == "" {
		HOME = `e:\\software`
	}
	HOME = `e:\\Software`
	var listenPort string
	flag.StringVar(&listenPort, "listen", "8080", "Port to listen on")
	flag.StringVar(&basePath, "root", HOME, "Base path to serve files from '/'")
	flag.Parse()

	t.Log("Base Path: ", basePath)

	_, err := os.Open(basePath)
	if err != nil {
		t.Fatal(err)
	}

	templateDirListing, _ = template.ParseFiles("dir_listing.html.got")
	templateNotFound, _ = template.ParseFiles("404.html.got")

	log.Println("Listening on port", listenPort)
	http.HandleFunc("/", handler)
	// for some reason can't get http.FileServer to work properly
	// this probably isn't the best way to do this - if you have a folder called 'static'
	// in basePath then the static handler will take precedence
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+r.URL.Path)
	})
	panic(http.ListenAndServe(":"+listenPort, nil))
}
