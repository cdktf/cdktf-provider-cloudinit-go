// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudinitconfig


type DataCloudinitConfigPart struct {
	// Body content for the part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.2/docs/data-sources/config#content DataCloudinitConfig#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// A MIME-style content type to report in the header for the part. Defaults to `text/plain`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.2/docs/data-sources/config#content_type DataCloudinitConfig#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A filename to report in the header for the part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.2/docs/data-sources/config#filename DataCloudinitConfig#filename}
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// A value for the `X-Merge-Type` header of the part, to control [cloud-init merging behavior](https://cloudinit.readthedocs.io/en/latest/reference/merging.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.2/docs/data-sources/config#merge_type DataCloudinitConfig#merge_type}
	MergeType *string `field:"optional" json:"mergeType" yaml:"mergeType"`
}

