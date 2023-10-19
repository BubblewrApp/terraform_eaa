package main

import (
	"git.source.akamai.com/terraform-provider-eaa/pkg/eaaprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return eaaprovider.Provider()
		},
	})
}
