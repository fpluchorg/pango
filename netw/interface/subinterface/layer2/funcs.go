package layer2

import (
	"fmt"

	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

func versioning(v version.Number) (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func specifier(e ...Entry) []namespace.ImportSpecifier {
	ans := make([]namespace.ImportSpecifier, 0, len(e))

	var val namespace.ImportSpecifier
	for _, x := range e {
		val = x
		ans = append(ans, val)
	}

	return ans
}

func container(v version.Number) normalizer {
	r, _ := versioning(v)
	return r
}

func first(ans normalizer, err error) (Entry, error) {
	if err != nil {
		return Entry{}, err
	}

	return ans.Normalize()[0], nil
}

func all(ans normalizer, err error) ([]Entry, error) {
	if err != nil {
		return nil, err
	}

	return ans.Normalize(), nil
}

func toNames(e []interface{}) ([]string, error) {
	ans := make([]string, len(e))
	for i := range e {
		switch v := e[i].(type) {
		case string:
			ans[i] = v
		case Entry:
			ans[i] = v.Name
		default:
			return nil, fmt.Errorf("invalid type: %s", v)
		}
	}

	return ans, nil
}

// FirewallNamespace returns an initialized namespace.
func FirewallNamespace(client util.XapiClient) *Firewall {
	return &Firewall{
		ns: &namespace.Importable{
			ImportPath: util.InterfaceImport,
			Common: namespace.Common{
				Singular: singular,
				Plural:   plural,
				Client:   client,
			},
		},
	}
}

// PanoramaNamespace returns an initialized namespace.
func PanoramaNamespace(client util.XapiClient) *Panorama {
	return &Panorama{
		ns: &namespace.Importable{
			ImportPath: util.InterfaceImport,
			Common: namespace.Common{
				Singular: singular,
				Plural:   plural,
				Client:   client,
			},
		},
	}
}
