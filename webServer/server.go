package main

import (
	"contype"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/cognifly/cognilog"
)

// Name of Web Server
const serverName = "reVres"

// HTTP verbs defined
const get = "GET"
const post = "POST"

// Page keeps request-scoped values
type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/home", root)
	http.HandleFunc("/about", about)
	http.HandleFunc("/about/company", company)
	http.HandleFunc("/", serveRes)

	host := flag.String("host", "localhost", "The host or ip address E.G. localhost")
	port := flag.Int("port", 8080, "port E.G. 8080")
	flag.Parse()
	cognilog.LogINFO("green", "reVres", fmt.Sprintf("http://%s:%d", *host, *port))
	cognilog.Log("red", http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil))
}

// handleErr deals with all non http errors
func handleErr(err error) {
	if err != nil {
		cognilog.Log("red", fmt.Sprintf("%v", err))
	}
}

// render static html files
func renderTmp(w http.ResponseWriter, t string, p Page) {
	if t == "404" {
		tmp, err := template.ParseFiles(path.Join("tmp", "404.html"))
		handleErr(err)
		err = tmp.Execute(w, p)
		handleErr(err)
	} else {
		tmp, err := template.ParseFiles(t)
		handleErr(err)
		err = tmp.Execute(w, p)
		handleErr(err)
		return
	}
}

// The root of the web server.
func root(w http.ResponseWriter, r *http.Request) {
	var page Page
	page.Title = "Cognifly"
	renderTmp(w, path.Join("tmp", "index.html"), page)
}

// The root of the web server.
func about(w http.ResponseWriter, r *http.Request) {
	var page Page
	page.Title = "About Cognifly"
	renderTmp(w, path.Join("tmp", "about.html"), page)
}

// The root of the web server.
func company(w http.ResponseWriter, r *http.Request) {
	var page Page
	page.Title = "Company Information"
	renderTmp(w, path.Join("tmp", "company.html"), page)
}

// maping the resource got from fileRead
type resConfig map[string]string

// prepare the resource link
func filePath(s, r string) string {
	num := strings.Index(r, s)
	str := r[num:]
	return str
}

// serve resources to the client or serve 404 page if its a bad request.
func serveRes(w http.ResponseWriter, r *http.Request) {
	var page Page
	if r.URL.Path == "/" {
		page.Title = "Cognifly"
		renderTmp(w, path.Join("tmp", "index.html"), page)
		return
	}
	jsn, err := ioutil.ReadFile(path.Join("etc", "static.json"))
	handleErr(err)
	var resource resConfig
	err = json.Unmarshal(jsn, &resource)
	handleErr(err)

	checkDir := strings.Contains(r.URL.Path, resource["dir"])
	checkfiles := strings.Contains(r.URL.Path, resource["files"])
	if checkDir == true || checkfiles == true {
		var state string
		if checkDir {
			state = "res/"
		} else if checkfiles {
			state = "docs/"
		}
		fsPath := filePath(state, r.URL.Path)
		data, err := ioutil.ReadFile(fsPath)
		// if resource path in url is false but contains res serve as 404.
		if err != nil {
			renderTmp(w, "404", page)
			return
		}
		cont := contype.FileType(r.URL.Path)
		w.Header().Set("Content-Type", cont)
		w.Header().Set("Server", serverName)
		w.Write(data)
	} else {
		switch r.Method {
		case get:
			renderTmp(w, "404", page)
		case post:
			renderTmp(w, "404", page)
		}
	}
}
