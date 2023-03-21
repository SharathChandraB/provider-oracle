/*
Copyright 2021 Upbound Inc.
*/

package objectstorage

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("oci_objectstorage_bucket", func(r *ujconfig.Resource) {
		r.Version = "v1alpha1"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
}
