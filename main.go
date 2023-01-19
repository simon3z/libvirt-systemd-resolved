package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"path"
)

var (
	ErrUnexpectedArgsNum = errors.New("unexpected number of command arguments")
	ErrMissingDashArg    = errors.New("missing dash as last command argument")
)

func main() {
	l, err := syslog.New(syslog.LOG_INFO, path.Base(os.Args[0]))

	if err != nil {
		log.Fatal("unable to open syslog: ", err)
	}

	log.SetOutput(io.MultiWriter(log.Writer(), l))

	if len(os.Args) != 5 {
		log.Fatal(fmt.Errorf("%w: %d", ErrUnexpectedArgsNum, len(os.Args)))
	}

	if os.Args[4] != "-" {
		log.Fatal(ErrMissingDashArg)
	}

	name := os.Args[1]
	operation := os.Args[2]
	status := os.Args[3]

	network, err := ReadLibvirtNetwork()

	if err != nil {
		log.Fatal("unable to read the libvirt network: ", err)
	}

	if operation == "started" && status == "begin" {
		err := UpdateNetworkDNS(name, network)

		if err == ErrVirNetMissingDomainName {
			log.Printf("libvirt network %s domain is not set, systemd resolved service will not be updated", name)
		} else if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("dns server %s for %s on the libvirt network %s with bridge %s has been configured",
				network.IP.Address, network.Domain.Name, network.Name, network.Bridge.Name)
		}
	}
}
