package engine

import (
	"strings"
	"unicode"

	"github.com/leonhfr/search-go/pkg/bloom"
)

type Entry struct {
	Title  string
	URL    string
	Filter bloom.Filter
}

func New(title, url, body string) Entry {
	words := Clean(body)
	filter := bloom.New(len(words))
	for _, w := range words {
		filter.Add(w)
	}
	return Entry{Title: title, URL: url, Filter: filter}
}

func (e *Entry) isResult(queries []string) bool {
	for _, query := range queries {
		if !e.Filter.Query(query) {
			return false
		}
	}
	return true
}

type Entries []Entry

type Result struct {
	Title string
	URL   string
}

func (e *Entries) Search(query string) []Result {
	var results []Result
	queries := Clean(query)

	for _, entry := range *e {
		if entry.isResult(queries) {
			results = append(results, Result{entry.Title, entry.URL})
		}
	}

	return results
}

func Clean(s string) []string {
	dict := make(map[string]struct{})

	for _, line := range strings.Split(s, "\n") {
		clean := strings.FieldsFunc(
			line,
			func(r rune) bool { return !unicode.IsLetter(r) },
		)

		for _, c := range clean {
			word := strings.ToLower(c)
			if !isStopWord(word) {
				dict[word] = struct{}{}
			}
		}
	}

	var words []string
	for word := range dict {
		words = append(words, word)
	}

	return words
}
