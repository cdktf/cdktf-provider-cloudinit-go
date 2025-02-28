// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudinitconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudinitConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Specify whether or not to base64 encode the `rendered` output.
	//
	// Defaults to `true`, and cannot be disabled if gzip is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.6/docs/data-sources/config#base64_encode DataCloudinitConfig#base64_encode}
	Base64Encode interface{} `field:"optional" json:"base64Encode" yaml:"base64Encode"`
	// Specify the Writer's default boundary separator. Defaults to `MIMEBOUNDARY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.6/docs/data-sources/config#boundary DataCloudinitConfig#boundary}
	Boundary *string `field:"optional" json:"boundary" yaml:"boundary"`
	// Specify whether or not to gzip the `rendered` output. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.6/docs/data-sources/config#gzip DataCloudinitConfig#gzip}
	Gzip interface{} `field:"optional" json:"gzip" yaml:"gzip"`
	// part block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.6/docs/data-sources/config#part DataCloudinitConfig#part}
	Part interface{} `field:"optional" json:"part" yaml:"part"`
}

