package provider


type CloudinitProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.2/docs#alias CloudinitProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

