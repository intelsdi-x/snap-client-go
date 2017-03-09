package snap

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/intelsdi-x/snap-client-go/client"
	"github.com/intelsdi-x/snap-client-go/client/operations"
)

// ClientParams defines parameters to initiates a Snap client.
type ClientParams struct {
	URL    string
	APIVer string
	Scheme string
}

// Client is an instance of the generated Snap client.
type Client struct {
	*operations.Client
}

// New returns an instance of Snap client.
// It returns the default client if an empty parameter is supplied.
func New(cp ClientParams) *Client {

	if cp.URL == "" {
		cp.URL = client.DefaultHost
	}
	if cp.APIVer == "" {
		cp.APIVer = client.DefaultBasePath
	}
	if cp.Scheme == "" {
		cp.Scheme = client.DefaultSchemes[0]
	}

	transport := httptransport.New(cp.URL, cp.APIVer, []string{cp.Scheme})
	c := client.New(transport, strfmt.Default)

	return &Client{c.Operations}
}
