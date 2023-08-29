// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudinitProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CloudinitProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCloudinitProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudinitProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCloudinitProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewCloudinitProviderParameters(scope constructs.Construct, id *string, config *CloudinitProviderConfig) error {
	return nil
}

