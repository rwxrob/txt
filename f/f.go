package F

import "html/template"

var DefaultFuncMap = template.FuncMap{
	`ping`: Ping,
}

func Ping() string { return `pong` }
