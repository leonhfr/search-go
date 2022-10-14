package gen

import (
	"bufio"
	"bytes"
	"encoding/json"
	"go/format"
	"io"
	"os"
	"text/template"

	"github.com/leonhfr/search-go/pkg/engine"
	"github.com/leonhfr/search-go/tmpl"
)

type entry struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Body  string `json:"body"`
}

func Entries(path string) (engine.Entries, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var jsonEntries []entry
	err = json.Unmarshal(contents, &jsonEntries)
	if err != nil {
		return nil, err
	}

	var entries engine.Entries
	for _, entry := range jsonEntries {
		entries = append(entries, engine.New(entry.Title, entry.URL, entry.Body))
	}

	return entries, nil
}

func SearchCode(entries engine.Entries) ([]byte, error) {
	templates := template.Must(template.New("").Parse(tmpl.Search))

	var code bytes.Buffer
	err := templates.ExecuteTemplate(&code, "", struct {
		Entries engine.Entries
	}{
		Entries: entries,
	})
	if err != nil {
		return nil, err
	}

	formatted, err := format.Source(code.Bytes())
	if err != nil {
		return nil, err
	}

	return formatted, nil
}

func WriteCode(code []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString(string(code))
	if err != nil {
		return err
	}

	return w.Flush()
}
