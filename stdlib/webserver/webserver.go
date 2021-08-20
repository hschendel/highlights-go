package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

var rootTmpl *template.Template

func main() {
	rootTmpl = mustLoadTemplates()
	mux := http.NewServeMux()
	mux.HandleFunc("/check", checkHandler)
	mux.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", mux)) // not a graceful shutdown!
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	checkData := struct {
		Environment string `json:"environment"`
		Version     string `json:"version"`
	}{
		Environment: "development",
		Version:     "0.0.1",
	}
	enc := json.NewEncoder(w)
	err := enc.Encode(checkData)
	if err != nil {
		log.Printf("while JSON encoding check data: %s", err)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var data timePageData
	data.Parse(r)
	if data.LocationError != "" {
		renderTemplate(w, "index.html", &data, http.StatusBadRequest)
	}
	if data.Location != "" {
		timeUTC := time.Now().UTC()
		location, _ := time.LoadLocation(data.Location) // location was checked in Parse
		timeLocal := timeUTC.In(location)
		data.TimeUTC = timeUTC.Format(time.RFC850)
		data.TimeLocal = timeLocal.Format(time.RFC850)
		data.HasOutput = true
	}
	renderTemplate(w, "index.html", data, http.StatusOK)
}

type timeForm struct {
	Location      string
	LocationError string
}

func (f *timeForm) Parse(r *http.Request) {
	f.Location = ""
	f.LocationError = ""
	if err := r.ParseForm(); err != nil {
		f.LocationError = "form parsing error"
		return
	}
	f.Location = strings.TrimSpace(r.Form.Get("location"))
	if f.Location != "" {
		if _, err := time.LoadLocation(f.Location); err != nil {
			f.LocationError = "unknown location"
		}
	}
}

type timePageData struct {
	timeForm  // composition
	HasOutput bool
	TimeUTC   string
	TimeLocal string
}

func mustLoadTemplates() *template.Template {
	r := template.New("root")
	if _, err := r.ParseGlob("_tmpl/*.html"); err != nil {
		log.Fatal(err)
	}
	return r
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}, statusCode int) {
	tmpl := rootTmpl.Lookup(name)
	if tmpl == nil {
		log.Printf("template %q not found", name)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("template rendering error for %q: %s", name, err)
	}
}
