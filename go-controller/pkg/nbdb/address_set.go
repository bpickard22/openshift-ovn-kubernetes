// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

import "github.com/ovn-org/libovsdb/model"

// AddressSet defines an object in Address_Set table
type AddressSet struct {
	UUID        string            `ovsdb:"_uuid"`
	Addresses   []string          `ovsdb:"addresses"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	Name        string            `ovsdb:"name"`
}

func copyAddressSetAddresses(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalAddressSetAddresses(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func copyAddressSetExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalAddressSetExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *AddressSet) DeepCopyInto(b *AddressSet) {
	*b = *a
	b.Addresses = copyAddressSetAddresses(a.Addresses)
	b.ExternalIDs = copyAddressSetExternalIDs(a.ExternalIDs)
}

func (a *AddressSet) DeepCopy() *AddressSet {
	b := new(AddressSet)
	a.DeepCopyInto(b)
	return b
}

func (a *AddressSet) CloneModelInto(b model.Model) {
	c := b.(*AddressSet)
	a.DeepCopyInto(c)
}

func (a *AddressSet) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *AddressSet) Equals(b *AddressSet) bool {
	return a.UUID == b.UUID &&
		equalAddressSetAddresses(a.Addresses, b.Addresses) &&
		equalAddressSetExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		a.Name == b.Name
}

func (a *AddressSet) EqualsModel(b model.Model) bool {
	c := b.(*AddressSet)
	return a.Equals(c)
}

var _ model.CloneableModel = &AddressSet{}
var _ model.ComparableModel = &AddressSet{}
