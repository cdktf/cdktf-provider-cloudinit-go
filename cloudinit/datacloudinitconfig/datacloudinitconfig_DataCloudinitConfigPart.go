package datacloudinitconfig


type DataCloudinitConfigPart struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudinit/d/config#content DataCloudinitConfig#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudinit/d/config#content_type DataCloudinitConfig#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudinit/d/config#filename DataCloudinitConfig#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudinit/d/config#merge_type DataCloudinitConfig#merge_type}.
	MergeType *string `field:"optional" json:"mergeType" yaml:"mergeType"`
}

