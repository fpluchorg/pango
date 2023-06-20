package predefined

import (
	"github.com/fpluchorg/pango/objs/app"
	"github.com/fpluchorg/pango/objs/srvc"
	dlpft "github.com/fpluchorg/pango/predefined/dlp/filetype"
	tdbft "github.com/fpluchorg/pango/predefined/tdb/filetype"
	"github.com/fpluchorg/pango/predefined/threat"
	"github.com/fpluchorg/pango/util"
)

type Panorama struct {
	Application *app.Predefined
	DlpFileType *dlpft.Panorama
	Services    *srvc.Predefined
	TdbFileType *tdbft.Panorama
	Threat      *threat.Panorama
}

func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		Application: app.PredefinedNamespace(x),
		DlpFileType: dlpft.PanoramaNamespace(x),
		Services:    srvc.PredefinedNamespace(x),
		TdbFileType: tdbft.PanoramaNamespace(x),
		Threat:      threat.PanoramaNamespace(x),
	}
}
