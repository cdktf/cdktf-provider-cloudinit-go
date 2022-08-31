//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudinit Provider for Terraform CDK (cdktf)
package cloudinit

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigPartList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigPartList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigPartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigPartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigPartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigPartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigPartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

