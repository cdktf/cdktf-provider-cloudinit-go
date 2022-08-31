//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudinit Provider for Terraform CDK (cdktf)
package cloudinit

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

func validateNewCloudinitProviderParameters(scope constructs.Construct, id *string, config *CloudinitProviderConfig) error {
	return nil
}

