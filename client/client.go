package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/sapcc/ontap-restapi-client/swagger/client"
)

type NetAppClientConfig struct {
	Host             string
	Username         string
	Password         string
	Debug            bool
	Insecure         bool
	TargetServerName string
}

type ONTAPAPI struct {
	Client *apiclient.ONTAPRESTAPIOnlineReference
}

// CreateHTTPClient creates and configures an HTTP client with TLS settings.
func CreateHTTPClient(insecure bool, targetServerName string) (*http.Client, error) {
	tlsOptions := httptransport.TLSClientOptions{}
	if insecure {
		tlsOptions.InsecureSkipVerify = true
	} else if targetServerName != "" {
		tlsOptions.ServerName = targetServerName
	}
	client, err := httptransport.TLSClient(tlsOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewClient(cfg NetAppClientConfig, client *http.Client) (*ONTAPAPI, error) {
	var err error
	if client == nil {
		client, err = CreateHTTPClient(cfg.Insecure, cfg.TargetServerName)
		if err != nil {
			return nil, err
		}
	}

	rt := httptransport.NewWithClient(cfg.Host, apiclient.DefaultBasePath, apiclient.DefaultSchemes, client)
	rt.Debug = cfg.Debug
	rt.DefaultAuthentication = cfg.BasicAuth()
	rt.Consumers["application/hal+json"] = runtime.JSONConsumer()
	rt.Consumers["text/html"] = runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		if bytes, err := io.ReadAll(reader); err == nil {
			return fmt.Errorf("error: %v", string(bytes))
		} else {
			return err
		}
	})
	return &ONTAPAPI{Client: apiclient.New(rt, strfmt.Default)}, nil
}

func (c NetAppClientConfig) BasicAuth() runtime.ClientAuthInfoWriter {
	return httptransport.BasicAuth(c.Username, c.Password)
}
