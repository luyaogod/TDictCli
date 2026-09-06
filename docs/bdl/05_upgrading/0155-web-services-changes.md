---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_321_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.21 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 3.21.

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

See [GWS Security](../16_web-services/4556-security.md "These topics cover security for Genero Web Services.") for more details about security
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

See [GWS Security](../16_web-services/4556-security.md "These topics cover security for Genero Web Services.") for more details about security
and encryption with GWS.

## New security.global.options entry in FGLPROFILE to allow legacy OpenSSL 1 options

Starting from FGLGWS 3.21.02, 4.01.06, and 5.00.00, it is now mandatory to use OpenSSL 3.0 LTS to
get the latest security fixes. This change is due to [OpenSSL 1.1.1 going EOL in September 2023](https://www.openssl.org/blog/blog/2023/03/28/1.1.1-EOL/) (external link).

To ease your migration from OpenSSL 1 to OpenSSL 3, the FGLPROFILE option
`security.global.options` can be used to set OpenSSL 1 options to connect to a legacy
server.

For details, go to [Security Configuration FGLPROFILE entries](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## New fglwsdl option -SSLOptions to support legacy OpenSSL 1 options

Starting from FGLGWS 3.21.02, 4.01.06, and 5.00.00, it is now mandatory to use OpenSSL 3.0 LTS to
get the latest security fixes. This change is due to [OpenSSL 1.1.1 going EOL in September 2023](https://www.openssl.org/blog/blog/2023/03/28/1.1.1-EOL/) (external link).

The fglwsdl tool supports the option (`-SSLOptions`) to set
OpenSSL 1 options when connecting to a legacy server.

For more details, go to [fglwsdl](../13_programming-tools/2523-fglwsdl.md).

## New security.global.certificate.selfsigned.preload entry in FGLPROFILE

Starting from FGLGWS 3.21.02, 4.01.07, and 5.00.02, there is an option to preload the global
self-signed certificate and private key used for HTTPS connections. Typically, the certificate and
key is loaded at the first HTTPS request. If you find the GWS computation of the certificate and key
takes too long, you can speed things up by setting the
`security.global.certificate.selfsigned.preload = TRUE` to preload the certificate
and key at the start of the application instead of at the first HTTPS connection.

For more information on web service security configuration, go to [Security Configuration FGLPROFILE entries](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## Change to OAuthAPI.GetIDSubject returns

Starting from FGLGWS 3.21.02 and 4.01.06, the `OAuthAPI.GetIDSubject` function
returns the subject identifier of an ID token in a string instead of an integer.

If you have previously used the function, review your code and ensure that the variable that
affects the return value is of type `STRING`.

For details, go to [OAuthAPI.GetIDSubject()](../16_web-services/4947-oauthapi-getidsubject.md "Get OAuth subject identifier of ID Token.").

## Changes to the OpenIDConnect service configuration

Starting from FGLGWS 3.21.02, 4.01.07, and 5.00.02, two parameters of the Genero OpenIDConnect
service configuration
($FGLDIR\web\_utilities\services\openid-connect\res\configuration) have
changes:

- The `oidc.logout.id_token_hint` parameter, used in the logout request sent to the
  provider, has been replaced by `oidc.logout.identifier`. The new parameter supports
  the values "`id_token_hint`", "`client_id`", or an
  `empty` (" ") value. The default value is "`id_token_hint`".
- The `oidc.logout.post_redirect` parameter, used to send the post redirect uri in
  the logout request, now needs a string value instead of the boolean value true/false on previous
  versions. The default value is now "`post_logout_redirect_uri`".

No action needs to be taken on your part, but if you have previously used a custom OpenIDConnect
configuration file and you want to use it when upgrading FGLGWS version, ensure that you review your
configuration for these parameters.

For more information on OpenID Connect Single sign-on, refer to the Single Sign-On User Guide.

## Discontinued support for OCSP

We are discontinuing support
for the Online Certificate Status Protocol (OCSP) due to the decision by several Certificate
Authorities to phase out support for this protocol in web services. However, please note that the
OCSP functionality will remain intact within our product, allowing you to continue using this
feature as part of your secure communications.

For more information on OCSP, go to [Enable OCSP](../16_web-services/4574-enable-ocsp.md "To enable Online Certificate Status Protocol (OCSP), set the security.global.ocsp.enable and security.global.ocsp.url entries in FGLPROFILE.").

## Improved communication protocol between GAS and GWS

The communication protocol between the Genero Application Server (GAS)
and the Genero Web Service (GWS) processes has been updated to improve reliability, specifically to
ensure that no connections are lost when the GWS process needs extra time to stop or restart. Prior
to this update, if the GWS process needed time to stop, connections could be lost or requests could
be sent before the service was ready. The protocol was enhanced so that `gwsproxy`
now ensures `fglrun` is ready to receive the next request before proceeding.

To ensure you are using this improved communication protocol between the GAS and
your GWS applications, ensure you are using the following version pairs (or greater):

- FGLGWS 3.21.05 and GAS 3.21.05

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Web
services changes in BDL 3.20](0158-web-services-changes.md "There are changes in support of web services in Genero 3.20.").

Notable changes introduced in maintenance releases:

- [Support for validating filenames in WSAttachments](0158-web-services-changes.md). The high-level REST `WSAttachment` attribute has an option to verify filenames in
  received files using a regular expression pattern, also available in FGLGWS 4.01.00.
- [Changes to default IP version used by a GWS client](0158-web-services-changes.md). The default IP version is now IPv4, also available in FGLGWS 4.01.00.
- [fglwsdl -xmlname option added to generate variables named with XMLName](0158-web-services-changes.md), also available in
  FGLGWS 4.01.00.
- [Dynamic loading of zlib library for data compression](0158-web-services-changes.md), also available in FGLGWS 4.01.03.

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
