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

resource "eaa_application" "tftunnelapp" {
  provider = eaa

  name = "tfapptunnel1"
  description = "app created using terraform"
  host = "github.com"

  app_profile = "tcp"
  app_type = "tunnel"
  client_app_mode = "tunnel"


  domain = "wapp"
  auth_enabled = "true"

  popregion = "us-west-1"

  agents = ["terraform-test-connector"]
  app_category ="test"

  tunnel_internal_hosts {
      proto_type= 1
      port_range=     "22"
      host=     "192.168.2.1"
    }

     tunnel_internal_hosts {
      proto_type= 1
      port_range=     "22"
      host=     "192.168.2.2"
    }

  advanced_settings {
      is_ssl_verification_enabled = "false"
      ignore_cname_resolution = "true"
      g2o_enabled = "true"
      ip_access_allow = "false"
      x_wapp_read_timeout = "300"
	}


  app_authentication {
    app_idp = "kiran-sqa2-okta"
    app_directories {
      name = "Cloud Directory"
      app_groups {
        name = "Admins"
      }
    }
  }
}

