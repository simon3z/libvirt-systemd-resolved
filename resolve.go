package main

// cspell:ignore resolvectl

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	ErrVirNetMissingDomainName       = errors.New("missing libvirt network domain name")
	ErrVirNetMissingBridge           = errors.New("missing libvirt network bridge")
	ErrVirNetMissingDNSServerAddress = errors.New("missing libvirt network dns server address")
)

func UpdateNetworkDNS(name string, network *LibvirtNetwork) error {
	if network.Domain.Name == "" {
		return ErrVirNetMissingDomainName
	}

	if network.Bridge.Name == "" {
		return ErrVirNetMissingBridge
	}

	if network.IP.Address == "" {
		return ErrVirNetMissingDNSServerAddress
	}

	c := exec.Command("resolvectl", "dns", network.Bridge.Name, network.IP.Address)

	if err := c.Run(); err != nil {
		return err
	}

	c = exec.Command("resolvectl", "domain", network.Bridge.Name, fmt.Sprintf("~%s", network.Domain.Name))

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
