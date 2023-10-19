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

resource "eaa_application" "tfapp3" {
  provider = eaa

  name        = "tfapp3"
  description = "app created using terraform"
  host        = "terraformapp300" /* Application Name */

  app_profile = "http"
  app_type    = "enterprise"

  client_app_mode = "tcp"

  domain = "wapp"

  auth_enabled = "true"

  agents = ["DND_Shekhar_Grafana_RP_Access"]

  servers {
    orig_tls        = true
    origin_protocol = "https"
    origin_port     = 443
    origin_host     = "origin-perftest.akamaidemo.net"
  }

  advanced_settings {
      is_ssl_verification_enabled = "false"
      ignore_cname_resolution = "true"
      g2o_enabled = "true"
	}

  popregion = "us-east-1"

  app_authentication {
    app_idp = "kiran-sqa2-okta"
    
    app_directories {
      name = "Cloud Directory"
      app_groups {
        name = "Admins"
      }
      app_groups {
        name = "demo_group"
      }
    }

  }
}


