package passwordcomplexity

import (
	"github.com/fpluchorg/pango/namespace"
	"github.com/fpluchorg/pango/util"
)

// Firewall is the client.Firewall.Template namespace.
type Firewall struct {
	ns *namespace.Standard
}

// Get performs GET to retrieve information for the given object.
func (c *Firewall) Get() (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(), util.EmptyString, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve information for the given object.
func (c *Firewall) Show() (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(), util.EmptyString, ans)
	return first(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Firewall) Set(e ...Entry) error {
	return c.ns.Set(c.pather(), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Firewall) Edit(e Entry) error {
	return c.ns.Edit(c.pather(), e)
}

// Delete performs DELETE to remove the specified objects.
// Objects can be either a string or an Entry object.
func (c *Firewall) Delete(e ...interface{}) error {
	return c.ns.Delete(c.pather(), nil, nil)
}

func (c *Firewall) pather() namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(v)
	}
}

func (c *Firewall) xpath(vals []string) ([]string, error) {
	var ans []string

	ans = append(ans,
		"config",
		"mgt-config",
		"password-complexity",
	)

	return ans, nil
}

func (c *Firewall) container() normalizer {
	return container(c.ns.Client.Versioning())
}
