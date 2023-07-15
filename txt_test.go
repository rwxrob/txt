package txt_test

import (
	"fmt"
	"text/template"

	"github.com/rwxrob/txt"
)

func ExampleRender() {

	str := txt.Render(`some template text`)
	fmt.Println(str)

	// Output:
	// some template text
}

func ExamplePrint() {}

func ExampleLog() {}

func ExampleRenderer() {

	r := txt.Renderer{
		FuncMap: template.FuncMap{
			`hello`: func(a string) string {
				if a == "" {
					a = "there"
				}
				return fmt.Sprintf(`Hello, %v!`, a)
			},
		},
	}

	r.Print(`{{ hello }} Glad you're here.`)
	r.Print(`{{ hello "Rob" }} Glad you're here.`)

	// Output:
	// Hello, there! Glad you're here.
	// Hello, Rob! Glad you're here.
}
