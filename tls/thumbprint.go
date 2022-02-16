package tls

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHostThumbprint() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHostThumbprintRead,
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address of the host to extract the thumbprint from.",
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "443",
				Description: "The port to connect to on the host.",
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Boolean that can be set to true to disable SSL certificate verification.",
			},
			"sha1": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA1 hash of the certificate.",
			},
		},
	}
}

func dataSourceHostThumbprintRead(d *schema.ResourceData, _ interface{}) error {
	config := &tls.Config{}
	config.InsecureSkipVerify = d.Get("insecure").(bool)
	conn, err := tls.Dial("tcp", d.Get("address").(string)+":"+d.Get("port").(string), config)
	if err != nil {
		return err
	}
	cert := conn.ConnectionState().PeerCertificates[0]
	fingerprint := sha1.Sum(cert.Raw)

	var buf bytes.Buffer
	for i, f := range fingerprint {
		if i > 0 {
			_, _ = fmt.Fprintf(&buf, ":")
		}
		_, _ = fmt.Fprintf(&buf, "%02X", f)
	}
	d.SetId(buf.String())
	d.Set("sha1", strings.ToLower(strings.ReplaceAll(buf.String(), ":", "")))
	return nil
}
