package collector

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
func (c *Panorama) GetList(logcollectorgroup, device string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Get, c.pather(logcollectorgroup, device), ans)
}

// ShowList performs a SHOW to retrieve a list of all objects.
func (c *Panorama) ShowList(logcollectorgroup, device string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Show, c.pather(logcollectorgroup, device), ans)
}

// Get performs GET to retrieve configuration for the given object.
func (c *Panorama) Get(logcollectorgroup, device string, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(logcollectorgroup, device), name, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve configuration for the given object.
func (c *Panorama) Show(logcollectorgroup, device string, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(logcollectorgroup, device), name, ans)
	return first(ans, err)
}

// GetAll performs GET to retrieve all objects configured.
func (c *Panorama) GetAll(logcollectorgroup, device string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Get, c.pather(logcollectorgroup, device), ans)
	return all(ans, err)
}

// ShowAll performs SHOW to retrieve all objects configured.
func (c *Panorama) ShowAll(logcollectorgroup, device string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Show, c.pather(logcollectorgroup, device), ans)
	return all(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Panorama) Set(logcollectorgroup, device string, e ...Entry) error {
	return c.ns.Set(c.pather(logcollectorgroup, device), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Panorama) Edit(logcollectorgroup, device string, e Entry) error {
	return c.ns.Edit(c.pather(logcollectorgroup, device), e)
}

// Delete performs DELETE to remove the specified objects.
//
// Objects can be either a string or an Entry object.
func (c *Panorama) Delete(logcollectorgroup, device string, e ...interface{}) error {
	names, nErr := toNames(e)
	return c.ns.Delete(c.pather(logcollectorgroup, device), names, nErr)
}

func (c *Panorama) pather(logcollectorgroup, device string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(logcollectorgroup, device, v)
	}
}

func (c *Panorama) xpath(logcollectorgroup, device string, vals []string) ([]string, error) {
	if logcollectorgroup == "" || device == "" {
		return nil, fmt.Errorf("logcollectorgroup and device must be specified")
	}

	ans := make([]string, 0, 7)
	ans = append(ans, util.LogCollectorGroupDeviceXpathPrefix(logcollectorgroup, device)...)
	ans = append(ans,
		"collectors",
		util.AsEntryXpath(vals),
	)

	return ans, nil
}

func (c *Panorama) container() normalizer {
	return container(c.ns.Client.Versioning())
}
