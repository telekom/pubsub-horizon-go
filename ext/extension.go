package ext

import (
	"fmt"
	"plugin"
)

type Extension interface {
	Info() *Info
	Configure()
	Enable()
	Disable()
}

type Info struct {
	Name string
}

func open(extFile string) (Extension, error) {
	pl, err := plugin.Open(extFile)
	if err != nil {
		return nil, err
	}

	sym, err := pl.Lookup("Extension")
	if err != nil {
		return nil, err
	}

	ex, ok := sym.(Extension)
	if !ok {
		return nil, fmt.Errorf("the 'Extension' variable in %s does not implement the Extension interface", extFile)
	}

	return ex, nil
}
