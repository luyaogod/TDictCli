---
title: "OpenSSL requirements"
source: "fgl-topics/c_gws_ssl_openssl_002.html"
breadcrumb: "Web services > Security > OpenSSL requirements"
type: "concept"
---

# OpenSSL requirements

> FGLGWS uses OpenSSL 3 libraries for security and encryption.

## Genero Web Services and OpenSSL

For security and encryption, GWS need OpenSSL libraries and tools. If the operating system does
not have OpenSSL libraries, the FGLGWS installer will provide the most recent OpenSSL version in
FGLDIR.

## Security note: OpenSSL 3.0 LTS support

Starting with FGLGWS 3.21.01, 4.01.05 and 5.00.00, OpenSSL 3.0 LTS is required for
encryption and security.

Because [OpenSSL 1.1.1 goes EOL in September 2023](https://www.openssl.org/blog/blog/2023/03/28/1.1.1-EOL/) (external link), it is now
mandatory to use OpenSSL 3.0 LTS to get the latest security fixes.

When installing an FGLGWS package, OpenSSL 3.0 libs will be provided in FGLDIR, if no OpenSSL 3.0
exists on the system.

Starting with OpenSSL 3.0, the SHA-1 digest algorithm is no longer supported by default. The
OpenSSL 3.0 libs provided in FGLDIR still have SHA-1 digest activated by default. If you want to
enable SHA-1 with the system OpenSSL 3.0 libs, use a command such as update-crypto-policies
--set DEFAULT:SHA1 in order to use SHA-1. However, the SHA-1 digest algorithm is no longer
recommended, because it is increasingly vulnerable as computers become more and more powerful. If
you are using SHA-1 with GWS crypto APIs, consider moving to SHA-256 or to a stronger secure hash
algorithm.

See [GWS Security](4556-security.md "These topics cover security for Genero Web Services.") for more details about security
and encryption with GWS.

## Get OpenSSL from third-party vendors for Windows

OpenSSL libraries are provided in the FGLGWS package; however, starting with FGLGWS 3.21.03,
4.01.08 and 5.00.03 if you want to install more recent OpenSSL libraries from third-party vendors,
you must do the following on your Windows® system:

- Go to an official OpenSSL library vendor. OpenSSL recommends [Shining Light
  Productions](https://slproweb.com/products/Win32OpenSSL.html) (external link) for OpenSSL libraries for Windows.
- Download the OpenSSL libraries from there.
- By default, GWS will look for the OpenSSL libraries in the $FGLDIR\bin
  directory unless you have also specified OPENSSL\_MODULES to look in another path. Copy the following
  libraries into your $FGLDIR\bin directory.
  - libcrypto-3-x64.dll
  - libssl-3-x64.dll
  - legacy.dll
- Verify that the installation has worked by running the command fglpass -Vssl;
  it should return the version of the OpenSSL libraries.

See [GWS Security](4556-security.md "These topics cover security for Genero Web Services.") for more details about security
and encryption with GWS.
