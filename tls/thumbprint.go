package tls

import (
	"bytes"
	"crypto/md5"
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
			"md5": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The MD5 hash of the certificate.",
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
	for x, cert := range conn.ConnectionState().PeerCertificates {
		fmt.Println(x)
		fmt.Println(cert.Subject)
	}
	cert := conn.ConnectionState().PeerCertificates[0]
	sha1_fingerprint := sha1.Sum(cert.Raw)
	md5_fingerprint := md5.Sum(cert.Raw)

	var buf bytes.Buffer
	for i, f := range sha1_fingerprint {
		if i > 0 {
			_, _ = fmt.Fprintf(&buf, ":")
		}
		_, _ = fmt.Fprintf(&buf, "%02X", f)
	}

	var md5buf bytes.Buffer
	for i, f := range md5_fingerprint {
		if i > 0 {
			_, _ = fmt.Fprintf(&md5buf, ":")
		}
		_, _ = fmt.Fprintf(&md5buf, "%02X", f)
	}

	d.SetId(buf.String())
	d.Set("sha1", strings.ToLower(strings.ReplaceAll(buf.String(), ":", "")))
	d.Set("md5", strings.ToLower(strings.ReplaceAll(md5buf.String(), ":", "")))
	return nil
}
