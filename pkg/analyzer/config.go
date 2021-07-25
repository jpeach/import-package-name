package analyzer

import (
	"flag"
	"fmt"
	"strings"
)

type AliasNames struct {
	index map[string]string
}

func (a *AliasNames) Add(alias string, path string) {
	if a.index == nil {
		a.index = make(map[string]string)
	}

	switch alias {
	case "", ".", "_":
	default:
		a.index[path] = alias
	}
}

func (a *AliasNames) AliasForPath(path string) string {
	return a.index[path]
}

var Config = AliasNames{}

// Flag is a type that can be used with flag.Var to set Config entries.
type Flag struct{}

var _ flag.Value = &Flag{}

// String ...
func (f Flag) String() string {
	var values []string

	for name, path := range Config.index {
		values = append(values, name+"="+path)
	}

	return strings.Join(values, ",")
}

func (f Flag) Set(value string) error {
	for _, spec := range strings.Split(value, ",") {
		parts := strings.Split(spec, "=")
		if len(parts) != 2 {
			return fmt.Errorf("%q must be of the form PackageName=ImportPath", spec)
		}

		Config.Add(parts[0], parts[1])
	}

	return nil
}
