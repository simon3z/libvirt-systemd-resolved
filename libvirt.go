package main

import (
	"encoding/xml"
	"io"
	"os"
)

type LibvirtNetwork struct {
	Name   string `xml:"name"`
	Bridge struct {
		Name string `xml:"name,attr"`
	} `xml:"bridge"`
	Domain struct {
		Name string `xml:"name,attr"`
	} `xml:"domain"`
	IP struct {
		Address string `xml:"address,attr"`
	} `xml:"ip"`
}

type HookData struct {
	Network LibvirtNetwork `xml:"network"`
}

func ReadLibvirtNetwork() (*LibvirtNetwork, error) {
	x, _ := io.ReadAll(os.Stdin)

	var d HookData

	err := xml.Unmarshal([]byte(x), &d)

	if err != nil {
		return nil, err
	}

	return &d.Network, nil
}
