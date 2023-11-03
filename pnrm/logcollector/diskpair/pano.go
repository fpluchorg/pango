package diskpair

import (
	"fmt"

	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
)

// Panorama is the client.Panorama.LogCollectorDiskPair namespace.
type Panorama struct {
	ns *namespace.Standard
}

// GetList performs GET to retrieve a list of all objects.
func (c *Panorama) GetList(logcollector string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Get, c.pather(logcollector), ans)
}

// ShowList performs a SHOW to retrieve a list of all objects.
func (c *Panorama) ShowList(logcollector string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Show, c.pather(logcollector), ans)
}

// Get performs GET to retrieve configuration for the given object.
func (c *Panorama) Get(logcollector string, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(logcollector), name, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve configuration for the given object.
func (c *Panorama) Show(logcollector string, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(logcollector), name, ans)
	return first(ans, err)
}

// GetAll performs GET to retrieve all objects configured.
func (c *Panorama) GetAll(logcollector string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Get, c.pather(logcollector), ans)
	return all(ans, err)
}

// ShowAll performs SHOW to retrieve all objects configured.
func (c *Panorama) ShowAll(logcollector string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Show, c.pather(logcollector), ans)
	return all(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Panorama) Set(logcollector string, e ...Entry) error {
	return c.ns.Set(c.pather(logcollector), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Panorama) Edit(logcollector string, e Entry) error {
	return c.ns.Edit(c.pather(logcollector), e)
}

// Delete performs DELETE to remove the specified objects.
//
// Objects can be either a string or an Entry object.
func (c *Panorama) Delete(logcollector string, e ...interface{}) error {
	names, nErr := toNames(e)
	return c.ns.Delete(c.pather(logcollector), names, nErr)
}

func (c *Panorama) pather(logcollector string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(logcollector, v)
	}
}

func (c *Panorama) xpath(logcollector string, vals []string) ([]string, error) {
	if logcollector == "" {
		return nil, fmt.Errorf("logcollector must be specified")
	}

	ans := make([]string, 0, 7)
	ans = append(ans, util.LogCollectorXpathPrefix(logcollector)...)
	ans = append(ans,
		"disk-settings",
		"disk-pair",
		util.AsEntryXpath(vals),
	)

	return ans, nil
}

func (c *Panorama) container() normalizer {
	return container(c.ns.Client.Versioning())
}
