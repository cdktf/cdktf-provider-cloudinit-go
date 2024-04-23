// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type CloudinitProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.4/docs#alias CloudinitProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

