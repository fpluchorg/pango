package settingmanagement

import (
	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

func versioning() (normalizer, func(Config) interface{}) {
	return &container_v1{}, specify_v1
}

func specifier(e Config) []namespace.Specifier {
	return []namespace.Specifier{e}
}

func container(v version.Number) normalizer {
	r, _ := versioning()
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
