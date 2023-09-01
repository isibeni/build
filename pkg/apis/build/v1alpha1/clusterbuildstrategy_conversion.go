// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/shipwright-io/build/pkg/webhook"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ensure v1alpha1 implements the Conversion interface
var _ webhook.Conversion = (*ClusterBuildStrategy)(nil)

func (src *ClusterBuildStrategy) ConvertTo(_ context.Context, _ *unstructured.Unstructured) error {
	return fmt.Errorf("v1alpha1 is the current storage version, nothing to convert to")
}

func (src *ClusterBuildStrategy) ConvertFrom(_ context.Context, _ *unstructured.Unstructured) error {
	return fmt.Errorf("v1alpha1 is the current storage version, nothing to convert from")
}
