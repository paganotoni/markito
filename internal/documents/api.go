package documents

import (
	"cmp"
	"encoding/json"
	"net/http"
	"os"

	"github.com/lithammer/shortuuid/v4"
)

type documentResponse struct {
	URL      string `json:"url"`
	Markdown string `json:"markdown_url"`
	HTML     string `json:"html_url"`
}

func baseURL() string {
	return cmp.Or(os.Getenv("BASE_URL"), "http://localhost:3000")
}

func APICreate(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	content := r.FormValue("content")
	if content == "" {
		http.Error(w, `{"error":"content is required"}`, http.StatusBadRequest)
		return
	}

	id := shortuuid.New()
	err := documents.Save(id, content)
	if err != nil {
		http.Error(w, `{"error":"failed to save document"}`, http.StatusInternalServerError)
		return
	}

	base := baseURL()
	resp := documentResponse{
		URL:      base + "/" + id,
		Markdown: base + "/raw/" + id + "/md",
		HTML:     base + "/raw/" + id + "/html",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func APIUpdate(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, `{"error":"id is required"}`, http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		http.Error(w, `{"error":"content is required"}`, http.StatusBadRequest)
		return
	}

	err := documents.Save(id, content)
	if err != nil {
		http.Error(w, `{"error":"failed to save document"}`, http.StatusInternalServerError)
		return
	}

	base := baseURL()
	resp := documentResponse{
		URL:      base + "/" + id,
		Markdown: base + "/raw/" + id + "/md",
		HTML:     base + "/raw/" + id + "/html",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
