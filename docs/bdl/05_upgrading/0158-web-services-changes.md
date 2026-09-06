---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_320_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 3.20.

## Changes to how GWS handles cookies

When upgrading your Web services, you need to be aware of the requirement for sameSite when
setting HTTP cookies. Review the changes in how the Genero Web Services (GWS) handles cookies:

- [WSHelper.WSServerCookiesType](../16_web-services/4925-wshelper-wsquerytype.md "The WSQueryType defines a dynamic array of key-value pairs that stores the query string of a URL.") has a new
  member (`sameSite`) added to the record to specify how cookies are sent
  to other sites. Three values are possibles for sending cookies: "Strict", "Lax", or
  "None". The default is "Lax". If `NULL`, the default "Lax" is used.
- The [com.HttpServiceRequest.setResponseCookies](../15_library-reference/3841-com-httpservicerequest-setresponsecookies.md "Allows the server to return cookies to be set on the client application sending the request.") method taking
  `WSHelper.WSServerCookiesType` as parameter may raise a runtime error
  if the `sameSite` attribute is not set correctly. For instance, you can
  not set it to "None" if secure is not set to `TRUE`.
  This raises the error `SameSite None value requires Secure to be
  set`.

If you use cookies on the server side, we recommend you make the changes necessary for
sending cookies, and that you recompile your sources. For instance, make sure that cookies
you use to secure your application are set with sameSite "Strict". Setting
`SameSite="Strict"` on the secure cookie ensures that only the exact URL
the cookie has been set with will return the cookie to the server. If not set, the GWS sets
each cookie's sameSite attribute as "Lax".

To have SameSite available, you must recompile your sources with fglcomp 3.20.13 or
greater.

## Single sign-on (OpenID Connect, SAML, and GIP) sameSite security

Single sign-on (SSO) services delivered in the Genero Web Services package under
$FGLDIR/web\_utilities/services have enhancements in FGLGWS 3.20.13 to
set "Strict" on the sameSite attribute of security cookies.

If you have previously used these protocols to authenticate users launching applications,
ensure you use the services delivered in FGLGWS 3.20.13 and higher to take advantage of the
SameSite security feature.

For more information about SSO, see the Genero Application Server User Guide.

## sameSite security cookie set to lax for OpenID Connect and SAML

From FGLGWS 3.20.14 onwards, the sameSite attribute of security cookies used in OpenID
Connect and SAML Single sign-on (SSO) services delivered in the Genero Web Services package
is set to "Lax" instead of "Strict". As the security cookie is used to perform URL
redirection to the Identity Provider, it is used only once to authenticate the user log in;
therefore, there is no need to restrict its use.

> **Note:**
>
> **Genero Identity Provider (GIP) SSO**
>
> SSO functionality is
> applied differently in the Genero Identity Provider (GIP). The GIP retains the cookie.
> If the authenticated user selects the option to "stay connected" at log in, any further
> authentication requests coming from the same browser are silently authorized by the GIP
> without requesting the user to log in again. The cookie forwarded must identify the user
> to the GIP, so its sameSite attribute is set to "Strict" to ensure it comes from the
> same domain, and is not sent via a third-party site.

No action needs to be taken on your part, but if you have previously used these protocols
to authenticate users launching applications, ensure you use the services delivered in the
Genero Web Services package under $FGLDIR/web\_utilities/services of
FGLGWS 3.20.14 and higher.

For more information about SSO, see the Genero Application Server User Guide.

## XML serializer is case sensitive

From FGLGWS 3.20 on, the XML serializer is case sensitive, same as the JSON serializer.
This means that the serializer uses the case of the variable name as defined in the 4GL
file.

What impact has this on my existing Web services? It is recommended to:

1. Generate a new client and server stub via the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool.
   - The tool has an enhancement which does not generate `XMLName`
     attributes, unless they are really needed.
   - If you do not generate the stub, the `XMLName` attribute is still
     taken into account, so it should not have any impact.
2. Check on server sides that all published SOAP web services input and output records are
   in lower case. This ensures compatibility with older versions. If backward compatibility
   is not an issue, then you must query the server WSDL file and generate the client stub
   from it.

For more details, see [XML serialization attributes](../16_web-services/5021-xml-serialization-attributes.md).

## Control HTTP Date header for GET, HEAD and DELETE requests

Starting with FGLGWS 3.20.00, the FGLPROFILE entry
`http.global.request.date` can be used to control whether GWS HTTP GET,
HEAD and DELETE requests must send the HTTP Date header.

For more details, see [HTTP configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## Support for empty HTTP POST or PUT requests

FGLGWS 3.20.01 allows now to POST or PUT empty requests and read them on server
side:

```
DEFINE req com.HttpRequest
LET req = com.HttpRequest.Create("http://tempuri.org")
CALL req.setMethod("POST") # or PUT
CALL req.doRequest()
```

Prior to GWS 3.20.01, the `doRequest()` call returned error code -15555 and
the error message "Unsupported request-response feature".

The server side can now read empty POST or PUT requests as well. For instance in the
following code sample, the return value of `readTextRequest()` will be
`NULL`:

```
DEFINE ser com.HttpServiceRequest
DEFINE txt STRING
LET ser = com.WebServiceEngine.getHttpServiceRequest(-1)
IF ser.getMethod() == "POST" THEN
  LET txt = ser.readTextRequest()
  IF txt IS NULL THEN
    DISPLAY "No body"
  END IF
END IF
```

Notice that the Content-Type is still checked and may raise an error, if it doesn't match.
So using `readXmlRequest()` must have a Content-Type of XML format, even if
the body is empty. In practice, it is recommended to check the Content-Type before doing any
read operation.

For more details, see [com.HttpRequest.doRequest](../15_library-reference/3861-com-httprequest-dorequest.md "Performs the HTTP request.").

## New high-level RESTful framework

Genero BDL 3.20 provides a new framework for RESTful Web services programming.

See [RESTful Web services (high-level framework)](../16_web-services/4705-restful-web-services-high-level-framework.md "These topics give you the information you need to begin working with RESTful Web services applications using BDL function with support for attributes.").

## Changes to security.global.protocol in FGLPROFILE

Starting with FGLGWS 3.20.07, GWS secured communication is based on the [OpenSSL 1.1](http://www.openssl.org) engine.
This version of OpenSSL always selects the security protocol. It no longer allows you to
specify a specific Transport Layer Security (TLS) or Secure Sockets Layer (SSL).

The `security.global.protocol` entry in the fglprofile
file is therefore no longer supported.

For instance, if you have set `security.global.protocol = "TLSv1.2"` to
configure OpenSSL to use TLSv1.2 for HTTPS for earlier versions, you may encounter the
following error message in your Web
service:

```
OpenSSL 1.1 doesn't support specific protocol anymore
```

It is
therefore recommended to remove the `security.global.protocol` entry from
your fglprofile file. For more information on Web service security
configuration see [Web services FGLPROFILE configuration](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.").

## XMLChoice = "nested" attribute option

Starting with FGLGWS 3.20, the fglwsdl tool now supports the attribute
`XMLChoice="nested"` option to produce an XML representation supporting a
`substitutionGroup`. A substitutionGroup in an XML schema is where one XML
element can be replaced by another element that has a substitutionGroup with the same name.

See [Example 3: Using XMLChoice="nested" for substitutionGroup](../16_web-services/5029-xmlchoice.md).

This change applies if you have previously used [XMLChoice](../16_web-services/5029-xmlchoice.md) for a WSDL containing a substitutionGroup.

> **Important:**
>
> If in versions prior to FGLGWS 3.20 you had a variable that was expected to serialize an
> `XMLChoice` value, you can set `XMLChoice="nested"`
> attribute on that variable to improve the serialization options and to avoid serializer
> errors.

It is therefore recommended to do the following:

- Recompile the Genero Web Service server to create a new WSDL that supports the XML
  `XMLChoice="nested"` feature.
- Regenerate all Genero Web Service client stubs from the newly-generated WSDL to get
  the support of `XMLChoice="nested"`. Regenerate client stubs using the
  [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool.

## fglwsdl -xmlname option

Starting with FGLGWS 3.20.14, the `fglwsdl -xmlname`
option is added to generate variables named with [XMLName](../16_web-services/5037-xmlname.md) attributes in stubs. The option may be used if
you need to serialize sub records when generating stubs.

For more details, see [fglwsdl](../13_programming-tools/2523-fglwsdl.md).

## Changes to default IP version used by a GWS client

Starting with FGLGWS 3.20.15, the default IP protocol version used by a GWS client is now
IPv4.

The `ip.global.version` entry in the fglprofile can be
used to set the IP version. For instance, if you want to use IPv6, set
`ip.global.version="IPv6"`.

The previous default was to use IPv6 if available, and fall back on IPv4 if not. In some
cases, the fallback led to performance issues. To use this previous behavior, set
`ip.global.version="undefined"`.

For more information on Web service IP configuration see [IPv6 configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## Support for validating filenames in WSAttachments

Starting with FGLGWS 3.20.15, the high-level REST `WSAttachment` attribute
supports validating file names in received files. There is an option to specify that the
file name must match a regular expression pattern on the server side.

```
WSAttachment = "[a-zA-Z0-9_]*\.[a-zA-Z0-9_]*"
```

Prior to FGLGWS 3.20.15, file names in file attachments were not validated. It is
recommended to use a regexp pattern to limit file names to a set of allowed characters, to
avoid the risk of code injection. For more information, see [WSAttachment](../16_web-services/4839-wsattachment.md "Defines file attachments in the REST message.").

## Support for RFC 8693 in the Genero Identity Provider (GIP) creation of OAuth ID and access tokens with scopes

Starting with FGLGWS 3.20.17 (and FGLGWS 4.01.02), the GIP follows the standard RFC 8693 as the
default method when creating OAuth ID and access tokens with the scope parameter.

Prior to 3.20.17, GIP created a JSON Web Token (JWT) with a "scopes" element defined as a JSON
array for the list of scopes. Now, according to the RFC 8693 standard, the JWT has a "scope" element
defined as a string with the scopes in a space-separated list.

No action needs to be taken on your part, but if you have previously used the GIP to authenticate
users launching applications and you want to use the new scope member, ensure that the
OpenIDConnectServiceProvider.xcf and the
GeneroAccessService.xcf delivered in the Genero Web Services package under
$FGLDIR/web\_utilities/services use FGLGWS 3.20.17 or higher. The
OpenIDConnectService and GeneroAccessService services have
been enhanced to handle both the old and new methods for exchanging scopes.

For more information about GIP, see the Genero Application Server User Guide.

## Dynamic zlib loading (to use system zlib)

Starting with FGLGWS 3.20.17 (and FGLGWS 4.01.03), the GWS modules requiring compression will
search for the zlib library installed on the operating system, and fallback to a zlib library
provided in FGLDIR/lib/wse.

In previous versions, a specific zlib was linked statically in
libfglutils.so, that could enter into conflict with other components, such as
JVM using the zlib of the system.

For more details, see [GWS uses zlib to compress data](../16_web-services/4546-gws-uses-zlib-to-compress-data.md "GWS support data compression by using the zlib library.").

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
