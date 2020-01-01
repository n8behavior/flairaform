package flair

import (
	"fmt"
	"github.com/google/jsonapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"net/http"
	"net/url"
	"os"
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

type Auth struct {
	Token   string `json:"access_token"`
	Expires int    `json:"expires_in"`
	Type    string `json:"token_type"`
	Scope   string `json:"scope"`
}

func authorize() (*Auth, error) {
	auth := new(Auth)
	endpoint, err := url.Parse("https://api.flair.co/oauth/token")
	if err != nil {
		return auth, fmt.Errorf("Endpoint URL: %s", err)
	}
	q := endpoint.Query()
	q.Set("client_id", os.Getenv("CLIENT_ID"))
	q.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	q.Set("scope", "thermostats.view+structures.view+structures.edit")
	q.Set("grant_type", "client_credentials")
	endpoint.RawQuery = q.Encode()

	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint.String(), nil)

	if err != nil {
		return auth, fmt.Errorf("Request: %s", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	defer res.Body.Close()
	jsonapi.UnmarshalPayload(res.Body, auth)

	return auth, nil
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
