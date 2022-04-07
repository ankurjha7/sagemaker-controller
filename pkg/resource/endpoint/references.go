// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package endpoint

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/client"

	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// ResolveReferences finds if there are any Reference field(s) present
// inside AWSResource passed in the parameter and attempts to resolve
// those reference field(s) into target field(s).
// It returns an AWSResource with resolved reference(s), and an error if the
// passed AWSResource's reference field(s) cannot be resolved.
// This method also adds/updates the ConditionTypeReferencesResolved for the
// AWSResource.
func (rm *resourceManager) ResolveReferences(
	ctx context.Context,
	apiReader client.Reader,
	res acktypes.AWSResource,
) (acktypes.AWSResource, error) {
	return res, nil
}

// validateReferenceFields validates the reference field and corresponding
// identifier field.
func validateReferenceFields(ko *svcapitypes.Endpoint) error {
	return nil
}

// hasNonNilReferences returns true if resource contains a reference to another
// resource
func hasNonNilReferences(ko *svcapitypes.Endpoint) bool {
	return false
}
