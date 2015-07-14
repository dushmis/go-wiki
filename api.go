package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown"
)

func DiffHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	file := vars["file"] + ".md"

	diff, err := Diff(file, vars["hash"])
	if err != nil {
		log.Println("ERROR", "Failed to get commit hash", vars["hash"])
	}

	// XXX: This could probably be done in a nicer way
	wrapped_diff := []byte("```diff\n" + string(diff) + "```")
	// md := blackfriday.MarkdownCommon(wrapped_diff)
	md := github_flavored_markdown.Markdown(wrapped_diff)

	w.Header().Set("Content-Type", "text/html")
	w.Write(md)
}
