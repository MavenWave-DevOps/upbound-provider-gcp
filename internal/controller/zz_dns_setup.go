/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	managedzone "github.com/upbound/provider-gcp/internal/controller/dns/managedzone"
	managedzoneiammember "github.com/upbound/provider-gcp/internal/controller/dns/managedzoneiammember"
	policy "github.com/upbound/provider-gcp/internal/controller/dns/policy"
	recordset "github.com/upbound/provider-gcp/internal/controller/dns/recordset"
)

// Setup_dns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managedzone.Setup,
		managedzoneiammember.Setup,
		policy.Setup,
		recordset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}