package predefined

import (
	"github.com/fpluchorg/pango/objs/app"
	"github.com/fpluchorg/pango/objs/srvc"
	dlpft "github.com/fpluchorg/pango/predefined/dlp/filetype"
	tdbft "github.com/fpluchorg/pango/predefined/tdb/filetype"
	"github.com/fpluchorg/pango/predefined/threat"
	"github.com/fpluchorg/pango/util"
)

type Firewall struct {
	Application *app.Predefined
	DlpFileType *dlpft.Firewall
	Services    *srvc.Predefined
	TdbFileType *tdbft.Firewall
	Threat      *threat.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		Application: app.PredefinedNamespace(x),
		DlpFileType: dlpft.FirewallNamespace(x),
		Services:    srvc.PredefinedNamespace(x),
		TdbFileType: tdbft.FirewallNamespace(x),
		Threat:      threat.FirewallNamespace(x),
	}
}
