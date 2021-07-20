package sumologic

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSumologicGlobalFolder() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSumologicGlobalFolderRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceSumologicGlobalFolderRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	// Based on https://api.sumologic.com/docs/#operation/getGlobalFolderAsync
	GlobalFolder, err := c.GetFolder("global")
	if err != nil {
		return err
	}

	d.SetId(GlobalFolder.ID)
	d.Set("name", GlobalFolder.Name)
	d.Set("description", GlobalFolder.Description)

	return nil
}
