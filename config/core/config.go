/*
Copyright 2021 Upbound Inc.
*/

package core

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("oci_core_vcn", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.Version = "v1alpha1"
	})
}
