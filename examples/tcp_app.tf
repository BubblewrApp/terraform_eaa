terraform {
    required_providers {
    eaa = {
      source  = "terraform.eaaprovider.dev/eaaprovider/eaa"
      version = "1.0.0"
    }
  }
}

provider "eaa" {
  contractid       = "1-3CV382"
  edgerc           = ".edgerc"
}


resource "eaa_application" "tfapptcp7" {
  provider = eaa

  name = "tfapptcp7"
  description = "app created using terraform"
  host = "github.com"

  app_profile = "tcp"
  app_type = "tunnel"
  client_app_mode = "tcp"

  domain = "wapp"
  auth_enabled = "true"

  agents = ["terraform-test-connector"]
  app_category ="test"

  servers {
    orig_tls        = true
    origin_protocol = "tcp"
    origin_port     = 443
    origin_host     = "192.168.2.1"
  }


  advanced_settings {
      is_ssl_verification_enabled = "false"
      ignore_cname_resolution = "true"
      g2o_enabled = "true"
      ip_access_allow = "false"
      x_wapp_read_timeout = "300"
      internal_host_port = "300"
      internal_hostname = "myhost999.com"
	}
  popregion = "us-west-1"

  app_authentication {
    app_idp = "kiran-sqa2-okta"
    app_directories {
      name = "Cloud Directory"
      app_groups {
        name = "demo_group"
      }
    }
  }
}
