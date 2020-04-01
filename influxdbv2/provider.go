package influxdbv2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/influxdata/influxdb-client-go"
	"github.com/lancey-energy-storage/influxdb-client-go"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.EnvDefaultFunc(
					"INFLUXDB_URL", "http://localhost:9999/"),
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	influx, error := influxdb.New(d.Get("url").(string), "")
	if error != nil {
		panic(error)
	}

	err := influx.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error pinging server: %s", err)
	}
	return influx, nil
}
