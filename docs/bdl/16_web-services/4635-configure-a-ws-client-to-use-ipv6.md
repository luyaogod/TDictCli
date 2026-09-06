---
title: "Configure a WS client to use IPv6"
source: "fgl-topics/t_gws_client_config_ipv6.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Configure a WS client to use IPv6"
type: "task"
---

# Configure a WS client to use IPv6

> Configuration steps to customize IPv6 for a WS client.

A Web Services client program can access a WS server using IPv6.

URLs that map to IPv6 addresses will be automatically handled by the Web Services library. It is
also possible to specify an IPv6 address directly as URL in your BDL code by enclosing the address
in `[]` square brackets, for example:

```
LET myURL = "http://[fe80::20c:29ff:fe05:9ca3]:80/index.html"
```

By default, the WS library will automatically use IPv4 (default behavior since FGLGWS version
3.20.15). To override the default behavior, you can explicitly specify the IP protocol version.

The platform where WS client programs execute must support IPv6.

1. Force the IP version with the `ip.global.version` entry in your
   fglprofile file, by specifying "IPv6", or "undefined" to use IPv6 addresses if
   available, and fall back to IPv4 if not.

   For example, to force IPv6 (when IPv6 is available):

   `ip.global.version =
   "IPv6"`
2. When using IPv6 for link-local addresses, if several network interfaces exist on the machine,
   you can explicitly specify which interface must be used with the
   `ip.global.v6.interface.name` or `ip.global.v6.interface.id` entry in
   fglprofile.

   In order to specify the IPv6 network interface by name, use:

   `ip.global.v6.interface.name = "eth0"`
   > **Important:**
   >
   > The `ip.global.v6.interface.name` entry is not supported on
   > Microsoft™
   > Windows® platforms.

   In order to specify the IPv6 network interface by id, use:

   `ip.global.v6.interface.id = "2"`

## Related links

**Related reference**  

[IPv6 configuration](4916-fglprofile-entries-for-web-services.md "IPv6 configuration")
