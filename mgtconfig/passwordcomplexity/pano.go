package passwordcomplexity

import (
	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
)

// Panorama is the client.Panorama.Template namespace.
type Panorama struct {
	ns *namespace.Standard
}

// Get performs GET to retrieve information for the given object.
func (c *Panorama) Get(tmpl string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(tmpl), util.EmptyString, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve information for the given object.
func (c *Panorama) Show(tmpl string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(tmpl), util.EmptyString, ans)
	return first(ans, err)
}

// GetAll performs GET to retrieve all objects configured.
func (c *Panorama) GetAll(tmpl string) (Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Get, c.pather(tmpl), ans)
	return all(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Panorama) Set(tmpl string, e ...Entry) error {
	return c.ns.Set(c.pather(tmpl), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Panorama) Edit(tmpl string, e Entry) error {
	return c.ns.Edit(c.pather(tmpl), e)
}

// Delete performs DELETE to remove the specified objects.
// Objects can be either a string or an Entry object.
func (c *Panorama) Delete(tmpl string, e ...interface{}) error {
	return c.ns.Delete(c.pather(tmpl), nil, nil)
}

func (c *Panorama) pather(tmpl string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(tmpl, v)
	}
}

func (c *Panorama) xpath(tmpl string, vals []string) ([]string, error) {
	var ans []string

	if tmpl != "" {
		ans = make([]string, 0, 12)
		ans = append(ans, util.TemplateXpathPrefix(tmpl, util.EmptyString)...)
	} else {
		ans = make([]string, 0, 4)
	}

	ans = append(ans,
		"config",
		"mgt-config",
		"password-complexity",
		util.AsEntryXpath(vals),
	)

	return ans, nil
}

func (c *Panorama) container() normalizer {
	return container(c.ns.Client.Versioning())
}
