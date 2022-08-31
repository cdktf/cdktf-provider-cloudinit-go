//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudinit Provider for Terraform CDK (cdktf)
package cloudinit

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudinitConfigPartList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudinitConfigPartList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudinitConfigPartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudinitConfigPartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

