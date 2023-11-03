package device

import (
	"fmt"

	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
)

// Panorama is the client.Panorama.LogCollectorGroupDevice namespace.
type Panorama struct {
	ns *namespace.Standard
}

// GetList performs GET to retrieve a list of all objects.
func (c *Panorama) GetList(logcollectorgroup string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Get, c.pather(logcollectorgroup), ans)
}

// ShowList performs a SHOW to retrieve a list of all objects.
func (c *Panorama) ShowList(logcollectorgroup string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Show, c.pather(logcollectorgroup), ans)
}

// Get performs GET to retrieve configuration for the given object.
func (c *Panorama) Get(logcollectorgroup string, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(logcollectorgroup), name, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve configuration for the given object.
func (c *Panorama) Show(logcollectorgroup string, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(logcollectorgroup), name, ans)
	return first(ans, err)
}

// GetAll performs GET to retrieve all objects configured.
func (c *Panorama) GetAll(logcollectorgroup string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Get, c.pather(logcollectorgroup), ans)
	return all(ans, err)
}

// ShowAll performs SHOW to retrieve all objects configured.
func (c *Panorama) ShowAll(logcollectorgroup string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Show, c.pather(logcollectorgroup), ans)
	return all(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Panorama) Set(logcollectorgroup string, e ...Entry) error {
	return c.ns.Set(c.pather(logcollectorgroup), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Panorama) Edit(logcollectorgroup string, e Entry) error {
	return c.ns.Edit(c.pather(logcollectorgroup), e)
}

// Delete performs DELETE to remove the specified objects.
//
// Objects can be either a string or an Entry object.
func (c *Panorama) Delete(logcollectorgroup string, e ...interface{}) error {
	names, nErr := toNames(e)
	return c.ns.Delete(c.pather(logcollectorgroup), names, nErr)
}

func (c *Panorama) pather(logcollectorgroup string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(logcollectorgroup, v)
	}
}

func (c *Panorama) xpath(logcollectorgroup string, vals []string) ([]string, error) {
	if logcollectorgroup == "" {
		return nil, fmt.Errorf("logcollectorgroup must be specified")
	}

	ans := make([]string, 0, 7)
	ans = append(ans, util.LogCollectorGroupXpathPrefix(logcollectorgroup)...)
	ans = append(ans,
		"logfwd-setting",
		"devices",
		util.AsEntryXpath(vals),
	)

	return ans, nil
}

func (c *Panorama) container() normalizer {
	return container(c.ns.Client.Versioning())
}
