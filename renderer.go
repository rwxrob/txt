package txt

import (
	"fmt"
	"html/template"
	"log"
)

// Renderer encapsulates a template.FuncMap to provide strings from templates
// rendered from default or customized domain-specific template
// grammars. See the f subpackage for available template grammar
// possibilities.
type Renderer struct {
	FuncMap template.FuncMap
}

// Render returns a rendered template (t) as a string.
func (r *Renderer) Render(t string) string {
	// TODO
	return ""
}

// Print prints the rendered template (t) to standard output (without
// a new line).
func (r *Renderer) Print(t string) { fmt.Print(r.Render(t)) }

// Log logs the rendered template (t) to standard error (without
// a new line).
func (r *Renderer) Log(t string) { log.Print(r.Render(t)) }
