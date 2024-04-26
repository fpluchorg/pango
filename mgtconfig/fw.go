package mgtconfig

import (
	"github.com/fpluchorg/pango/mgtconfig/passwordcomplexity"
	"github.com/fpluchorg/pango/util"
)

// Firewall is the client.MGTConfig namespace.
type Firewall struct {
	PasswordComplexity *passwordcomplexity.Panorama
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		PasswordComplexity: passwordcomplexity.PanoramaNamespace(x),
	}
}
