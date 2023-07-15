package txt

import (
	"fmt"
	"log"

	F "github.com/rwxrob/txt/f"
)

// Render returns the rendered template (t) by calling
// DefaultRenderer.Render.
func Render(t string) string { return DefaultRenderer.Render(t) }

// Print uses the DefaultRenderer to render and print the template
// string argument (t) to standard output (without adding a new line).
func Print(t string) { fmt.Print(Render(t)) }

// Log uses the DefaultRenderer to render and log the template
// string argument (t) to standard error (without adding a new line).
func Log(t string) { log.Print(Render(t)) }

// DefaultRenderer encapsulates all the available template grammar
// possibilities by wrapping F.DefaultFuncMap (see f subpackage). To
// customize the template grammar first create a Renderer and explicitly
// set its FuncMap.
var DefaultRenderer = &Renderer{
	FuncMap: F.DefaultFuncMap,
}
