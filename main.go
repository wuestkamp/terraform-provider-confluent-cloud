package main

import (
	c "github.com/Mongey/terraform-provider-confluent-cloud/ccloud"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: c.Provider})
}
