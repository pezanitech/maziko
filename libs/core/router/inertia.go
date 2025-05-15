package router

import (
	"net/http"

	inertia "github.com/romsar/gonertia"
)

type Inertia = *inertia.Inertia
type Props = inertia.Props
type InertiaHandler func(Inertia, http.ResponseWriter, *http.Request)

// Renders an Inertia page with the component and props
func RenderInertiaPage(i Inertia, w http.ResponseWriter, r *http.Request, component string, props Props) error {
	return i.Render(w, r, component, props)
}

// Creates a new Inertia instance with the given options
func NewInertia(templateString string, options ...inertia.Option) (Inertia, error) {
	return inertia.New(templateString, options...)
}

// Holds the available options from the original library
var InertiaOptions struct {
	WithVersion         func(version string) inertia.Option
	WithVersionFromFile func(filename string) inertia.Option
	WithSSR             func(url ...string) inertia.Option
}

func init() {
	// initialize the options struct with options
	InertiaOptions.WithVersion = inertia.WithVersion
	InertiaOptions.WithVersionFromFile = inertia.WithVersionFromFile
	InertiaOptions.WithSSR = inertia.WithSSR
}
