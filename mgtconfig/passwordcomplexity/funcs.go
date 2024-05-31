package passwordcomplexity

import (
	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

func versioning(v version.Number) (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func specifier(e ...Entry) []namespace.Specifier {
	ans := make([]namespace.Specifier, 0, len(e))

	var val namespace.Specifier
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

func all(ans normalizer, err error) (Entry, error) {
	if err != nil {
		return Entry{}, err
	}

	return ans.Normalize()[0], nil
}

// PanoramaNamespace returns an initialized namespace.
func PanoramaNamespace(client util.XapiClient) *Panorama {
	return &Panorama{
		ns: &namespace.Standard{
			Common: namespace.Common{
				Singular: singular,
				Plural:   plural,
				Client:   client,
			},
		},
	}
}

// FirewallNamespace returns an initialized namespace.
func FirewallNamespace(client util.XapiClient) *Firewall {
	return &Firewall{
		ns: &namespace.Standard{
			Common: namespace.Common{
				Singular: singular,
				Plural:   plural,
				Client:   client,
			},
		},
	}
}
