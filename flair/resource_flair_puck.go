package flair

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePuck() *schema.Resource {
	return &schema.Resource{
		Create: resourcePuckCreate,
		Read:   resourcePuckRead,
		Update: resourcePuckUpdate,
		Delete: resourcePuckDelete,

		Schema: map[string]*schema.Schema{
			"room": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePuckCreate(d *schema.ResourceData, m interface{}) error {
	room := d.Get("room").(string)
	d.SetId(room)
	return resourcePuckRead(d, m)
}

func resourcePuckRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePuckUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePuckRead(d, m)
}

func resourcePuckDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
