package general

import (
	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

func versioning(v version.Number) (normalizer, func(Config) interface{}) {
	if v.Gte(version.Number{10, 0, 0, ""}) {
		return &container_v3{}, specify_v3
	} else if v.Gte(version.Number{9, 0, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func specifier(e Config) []namespace.Specifier {
	return []namespace.Specifier{e}
}

func container(v version.Number) normalizer {
	r, _ := versioning(v)
	return r
}

func first(ans normalizer, err error) (Config, error) {
	if err != nil {
		return Config{}, err
	}

	return ans.Normalize()[0], nil
}

// FirewallNamespace returns an initialized namespace.
func FirewallNamespace(client util.XapiClient) *Firewall {
	return &Firewall{
		ns: &namespace.Standard{
			Common: namespace.Common{
				Singular: singular,
				Client:   client,
			},
		},
	}
}

// PanoramaNamespace returns an initialized namespace.
func PanoramaNamespace(client util.XapiClient) *Panorama {
	return &Panorama{
		ns: &namespace.Standard{
			Common: namespace.Common{
				Singular: singular,
				Client:   client,
			},
		},
	}
}
