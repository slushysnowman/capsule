// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"fmt"
)

const (
	AvailableIngressClassesAnnotation             = "capsule.clastix.io/ingress-classes"
	AvailableIngressClassesRegexpAnnotation       = "capsule.clastix.io/ingress-classes-regexp"
	AvailableStorageClassesAnnotation             = "capsule.clastix.io/storage-classes"
	AvailableStorageClassesRegexpAnnotation       = "capsule.clastix.io/storage-classes-regexp"
	AllowedRegistriesAnnotation                   = "capsule.clastix.io/allowed-registries"
	AllowedRegistriesRegexpAnnotation             = "capsule.clastix.io/allowed-registries-regexp"
	ForbiddenNamespaceLabelsAnnotation            = "capsule.clastix.io/forbidden-namespace-labels"
	ForbiddenNamespaceLabelsRegexpAnnotation      = "capsule.clastix.io/forbidden-namespace-labels-regexp"
	ForbiddenNamespaceAnnotationsAnnotation       = "capsule.clastix.io/forbidden-namespace-annotations"
	ForbiddenNamespaceAnnotationsRegexpAnnotation = "capsule.clastix.io/forbidden-namespace-annotations-regexp"
)

func UsedQuotaFor(resource fmt.Stringer) string {
	return "quota.capsule.clastix.io/used-" + resource.String()
}

func HardQuotaFor(resource fmt.Stringer) string {
	return "quota.capsule.clastix.io/hard-" + resource.String()
}
