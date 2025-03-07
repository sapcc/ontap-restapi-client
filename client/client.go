package client

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/sapcc/ontap-restapi-client/swagger/client"
)

type NetAppClientConfig struct {
	Host       string
	ServerName string
	Username   string
	Password   string
	Debug      bool
	Insecure   bool
}

func NewClient(cfg NetAppClientConfig) (*apiclient.ONTAPRESTAPIOnlineReference, error) {
	if cfg.ServerName == "" {
		cfg.ServerName = cfg.Host
	}

	tlsOptions := httptransport.TLSClientOptions{}
	if cfg.Insecure {
		tlsOptions.InsecureSkipVerify = true
	} else if cfg.ServerName != "" {
		tlsOptions.ServerName = cfg.ServerName
	}
	tslClient, err := httptransport.TLSClient(tlsOptions)
	if err != nil {
		return nil, err
	}

	// data is of type *models.ErrorResponse
	// maybe we need to manually parse error in some other consumer?
	// rt.Consumers["application/hal+json"] = runtime.JSONConsumer()
	// rt.Producers["application/hal+json"] = runtime.JSONProducer()

	rt := httptransport.NewWithClient(cfg.Host, apiclient.DefaultBasePath, apiclient.DefaultSchemes, tslClient)
	rt.Debug = cfg.Debug
	rt.DefaultAuthentication = cfg.BasicAuth()
	rt.Consumers["text/html"] = runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		if bytes, err := io.ReadAll(reader); err == nil {
			return fmt.Errorf("error: %v", string(bytes))
		} else {
			return err
		}
	})
	return apiclient.New(rt, strfmt.Default), nil
}

func (c NetAppClientConfig) BasicAuth() runtime.ClientAuthInfoWriter {
	return httptransport.BasicAuth(c.Username, c.Password)
}
