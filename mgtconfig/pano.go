/*
Package mgtconfig is the client.MGTConfig namespace.
*/
package mgtconfig

import (
	"github.com/fpluchorg/pango/mgtconfig/user"
	"github.com/fpluchorg/pango/util"
)

// Panorama is the panorama.Panorama namespace.
type Panorama struct {
	User *user.Panorama
}

func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		User: user.PanoramaNamespace(x),
	}
}
