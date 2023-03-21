/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	vcn "github.com/SharathChandraB/provider-oracle/internal/controller/core/vcn"
	bucket "github.com/SharathChandraB/provider-oracle/internal/controller/objectstorage/bucket"
	providerconfig "github.com/SharathChandraB/provider-oracle/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vcn.Setup,
		bucket.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
