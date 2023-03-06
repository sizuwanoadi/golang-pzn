package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Page{})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title": "Comparator Bilek",
		"Value": 62,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Range",
		"Hobbies": []string{
			"Gaming", "Coding", "Reading Manga",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(w, "address.gohtml", map[string]interface{}{
		"Title": "Template With",
		"Name":  "Adi Siswanto",
		"Address": map[string]interface{}{
			"Street": "Jl. Raya Dago",
			"City":   "Bogor",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
