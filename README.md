# Libvirt Systemd Resolved

Libvirt Systemd Resolved is a Golang project to automatically update the systemd resolved configuration with libvirt networks dns.

## Installation

Use the golang build command to build from local sources:

```bash
go build -o 01-libvirt-systemd-resolved .
```

## Usage

Copy the `01-libvirt-systemd-resolved` binary in the libvirt network hooks directory:

```bash
mkdir -p /etc/libvirt/hooks/network.d
cp 01-libvirt-systemd-resolved /etc/libvirt/hooks/network.d/01-libvirt-systemd-resolved
```

## Libvirt Networks Requirements

The Libvirt network definitions must include a `domain` `name` and an `ip` `address` in order to properly configure the systemd resolved service:

```xml
<network>
  <name>default</name>
  <...>
  <domain name='default.example.com' .../>
  <ip address='192.168.122.1' ...>
    <...>
  </ip>
</network>
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)
