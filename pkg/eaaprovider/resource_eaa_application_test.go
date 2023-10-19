package eaaprovider

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccEaaApplication_basic(t *testing.T) {
	appName1 := fmt.Sprintf("tf-app-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	appName2 := fmt.Sprintf("tf-app-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	host1 := fmt.Sprintf("tfhost%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	host2 := fmt.Sprintf("tfhost%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		IsUnitTest:        false,
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEaaApplicationConfig_basic(appName1, host1, "http", "enterprise"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEaaApplicationExists(fmt.Sprintf("eaa_application.%s", appName1)),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "name", appName1),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "host", host1),
				),
			},
			{
				Config: testAccEaaApplicationConfig_basic(appName2, host2, "http", "enterprise"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEaaApplicationExists(fmt.Sprintf("eaa_application.%s", appName2)),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "name", appName2),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "host", host2),
				),
			},
		},
	})
}

func TestAccEaaApplication_complex(t *testing.T) {
	appName1 := fmt.Sprintf("tf-app-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	appName2 := fmt.Sprintf("tf-app-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	appName3 := fmt.Sprintf("tf-app-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	host1 := fmt.Sprintf("tfhost%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	host2 := fmt.Sprintf("tfhost%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	host3 := fmt.Sprintf("tfhost%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		IsUnitTest:        false,
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEaaApplicationConfig_complex(appName1, host1, "http", "enterprise", "terraform-test-connector", "terraform-idp", "Cloud Directory", "Admins"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEaaApplicationExists(fmt.Sprintf("eaa_application.%s", appName1)),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "name", appName1),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "host", host1),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "agents.#", "1"), // Check the count of agents
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "agents.0", "terraform-test-connector"),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "app_authentication.#", "1"), // Check the count of agents
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "app_authentication.0.app_directories.0.name", "Cloud Directory"),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "app_authentication.0.app_directories.0.app_groups.0.name", "Admins"),
				),
			},
			{
				Config: testAccEaaApplicationConfig_complex(appName2, host2, "http", "enterprise", "terraform-test-connector", "terraform-idp", "Cloud Directory", "demo_group"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEaaApplicationExists(fmt.Sprintf("eaa_application.%s", appName2)),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "name", appName2),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "host", host2),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "agents.#", "1"), // Check the count of agents
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "agents.0", "terraform-test-connector"),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "app_authentication.#", "1"), // Check the count of agents
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "app_authentication.0.app_directories.0.name", "Cloud Directory"),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName2), "app_authentication.0.app_directories.0.app_groups.0.name", "demo_group"),
				),
			},
			{
				Config:      testAccEaaApplicationConfig_complex(appName3, host3, "http", "enterprise", "terraformappnoconnector", "terraform-idp", "Cloud Directory", "demo_group"),
				ExpectError: regexp.MustCompile(`Error: agents assign failed: Action failed - Unable to process request`),
			},
		},
	})
}

func TestAccEaaApplication_G2O(t *testing.T) {
	appName1 := fmt.Sprintf("tf-app-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	host1 := fmt.Sprintf("tfhost%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		IsUnitTest:        false,
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEaaApplicationConfig_G2O(appName1, host1, "http", "enterprise"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEaaApplicationExists(fmt.Sprintf("eaa_application.%s", appName1)),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "name", appName1),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "host", host1),
					resource.TestCheckResourceAttr(fmt.Sprintf("eaa_application.%s", appName1), "advanced_settings.0.g2o_enabled", "true"),

					// Custom check: Verify that the g2o_key attribute is not empty
					func(s *terraform.State) error {

						attr := s.RootModule().Resources[fmt.Sprintf("eaa_application.%s", appName1)].Primary.Attributes

						// Get the g2o_key attribute from the first element
						g2oKey := attr["advanced_settings.0.g2o_key"]
						if g2oKey == "" {
							return fmt.Errorf("Attribute 'g2o_key' is empty")
						}
						g2o_nonce := attr["advanced_settings.0.g2o_nonce"]
						if g2o_nonce == "" {
							return fmt.Errorf("Attribute 'g2o_nonce' is empty")
						}

						return nil
					},
				),
			},
		},
	})
}

func testAccCheckEaaApplicationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		name := rs.Primary.Attributes["name"]
		app_type := rs.Primary.Attributes["app_type"]
		app_profile := rs.Primary.Attributes["app_profile"]

		if name == "" {
			return errors.New("app Name is not set")
		}

		if app_type == "" {
			return errors.New("app_type is not set")
		}

		if app_profile == "" {
			return errors.New("app_profile is not set")
		}

		return nil
	}
}

func testAccEaaApplicationConfig_basic(appName string, host string, appProfile string, appType string) string {

	return fmt.Sprintf(`

	provider "eaa" {
		contractid       = "1-3CV382"
		edgerc           = ".edgerc"
	  }
	  
	  resource "eaa_application" "%s" {
		provider = eaa
	  
		name        = "%s"
		description = "app created using terraform"
		host        = "%s" 
	  
		app_profile = "%s"
		app_type    = "%s"
	  
		client_app_mode = "tcp"
	  
		domain = "wapp"
	  
		advanced_settings {
			is_ssl_verification_enabled = "false"
			ignore_cname_resolution = "true"
			g2o_enabled = "false"
		}
		
		popregion = "us-east-1"

	  }	  
`, appName, appName, host, appProfile, appType)
}

func testAccEaaApplicationConfig_G2O(appName string, host string, appProfile string, appType string) string {

	return fmt.Sprintf(`

	provider "eaa" {
		contractid       = "1-3CV382"
		edgerc           = ".edgerc"
	  }
	  
	  resource "eaa_application" "%s" {
		provider = eaa
	  
		name        = "%s"
		description = "app created using terraform"
		host        = "%s" 
	  
		app_profile = "%s"
		app_type    = "%s"
	  
		client_app_mode = "tcp"
	  
		domain = "wapp"
	  
		advanced_settings {
			is_ssl_verification_enabled = "false"
			ignore_cname_resolution = "true"
			g2o_enabled = "true"
		}
		
		popregion = "us-east-1"

	  }	  
`, appName, appName, host, appProfile, appType)
}

func testAccEaaApplicationConfig_complex(appName, host, appProfile, appType, agent, idp, directory, group string) string {

	return fmt.Sprintf(`

	provider "eaa" {
		contractid       = "1-3CV382"
		edgerc           = ".edgerc"
	  }
	  
	  resource "eaa_application" "%s" {
		provider = eaa
	  
		name        = "%s"
		description = "app created using terraform"
		host        = "%s" 
	  
		app_profile = "%s"
		app_type    = "%s"
	  
		client_app_mode = "tcp"
	  
		domain = "wapp"
	  
		auth_enabled = "true"

  agents = ["%s"]

  servers {
    orig_tls        = true
    origin_protocol = "https"
    origin_port     = 443
    origin_host     = "origin-perftest.akamaidemo.net"
  }

  advanced_settings {
      is_ssl_verification_enabled = "false"
      ignore_cname_resolution = "true"
      g2o_enabled = "false"
			}

  popregion = "us-east-1"

  app_authentication {
    app_idp = "%s"
    
    app_directories {
      name = "%s"
      app_groups {
        name = "%s"
      }
      
    }

  }

	  }	  
`, appName, appName, host, appProfile, appType, agent, idp, directory, group)
}

func testAccPreCheck(_ *testing.T) {

}
