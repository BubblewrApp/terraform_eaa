package eaaprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]func() (*schema.Provider, error)

var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()

	testAccProviders = map[string]func() (*schema.Provider, error){
		"eaa": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}
