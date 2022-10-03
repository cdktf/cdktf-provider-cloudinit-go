package provider


type CloudinitProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudinit#alias CloudinitProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

