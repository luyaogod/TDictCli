---
title: "The OpenSSL tool"
source: "fgl-topics/c_gws_ssl_openssl_003.html"
breadcrumb: "Web services > Security > Certificates in practice > The OpenSSL tool"
type: "concept"
---

# The OpenSSL tool

> The openssl command line tool creates certificates for the configuration of secured communications.

OpenSSL is provided in the FGLGWS package. To use it, you need to add its directory,
$FGLDIR/web\_utilities/tools, to your PATH environment variable.

The openssl tool has a default configuration file,
openssl.cnf. It looks for the openssl.cnf file in the
directory where it is executed; it stops if the file is not present. To use the
openssl tool from any directory, set the `OPENSSL_CONF`
environment variable to specify the location of the configuration file.

For information on how the openssl tool works, refer to the [openssl](https://www.openssl.org/) (external link)
documentation.

## Related links

**Related concepts**  

[OpenSSL requirements](4567-openssl-requirements.md "FGLGWS uses OpenSSL 3 libraries for security and encryption.")
