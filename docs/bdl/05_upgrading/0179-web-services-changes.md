---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_310_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 3.10.

## The XMLElementNillable, XMLOptional, and XMLNillable attributes

The management of XML serialization when serializing nillable has changed. These attributes
define how a NULL value is interpreted in XML.

When using the [XMLOptional](../16_web-services/5022-xmloptional.md) attribute, the behavior of the
XML serialization has changed; it no longer handles XML nillable.

The new [XMLNillable](../16_web-services/5023-xmlnillable.md) attribute can be used to specify that
the XML representation of a NULL value must be `xsi:nil="true"`.

In order to get the same result as when using `XMLOptional` in prior versions, set
both `XMLOptional` and `XMLNillable`.

Instead of specifying each element individually with `XMLNillable`, the [XMLElementNillable](../16_web-services/5025-xmlelementnillable.md) attribute can be used with a BDL
`RECORD` defined as `TYPE` or `DEFINE` to specify the
XML representation for nillable in all elements in the record. This eases migration to the new
`XMLNillable` behavior, because now you can get the same behavior as before defining
database records by adding the `XMLElementNillable` attribute.

> **Important:**
>
> If in versions prior to GWS 3.10 you had the following expected XML
> serialization behavior, you need to take action to avoid serializer errors:
>
> - A variable that was expected to serialize an `xsi:nil` value, you must set the
>   XMLNillable attribute on that variable.
> - A record defined with `LIKE` for fields in a database table which may allow null
>   values, you must set the XMLElementNillable attribute as in the
>   example.
>
>   ```
>   DEFINE rec RECORD ATTRIBUTE(XMLElementNillable) LIKE customer.*
>   ```

It is therefore recommended to do the following:

- Recompile the Genero Web Service server to create a new WSDL that supports the XML nillable
  feature.
- Regenerate all Genero Web Service client stubs from the newly-generated WSDL to get the support
  of XML nillable. Regenerate client stubs using the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool.

## Better support for Web Services API on iOS/GMI

Most of the Web Services APIs are not available on iOS mobile devices.

There are still some exceptions, related to iOS restrictions. For details see [Web Services on mobile devices](../17_mobile-applications/5100-web-services-on-mobile-devices.md "Web Services can be used within mobile applications.").

> **Note:**
>
> In versions before 3.10, the iOS app displayed a pop-up dialog to cancel a long running HTTP
> request. Starting with 3.10, the pop-up dialog is no longer used; only the activity indicator
> displays, and the app goes into background mode, if the HTTP request fails to complete after several
> minutes. The program gets a runtime error -15553 if the user taps the app icons, to bring the app
> back to foreground mode.

## OpenID Single sign-on (SSO) protocol not supported

Support of the OpenID Single sign-on (SSO) protocol has been dropped. This service based on
Genero REST is no longer delivered in the Genero Web Services package under
$FGLDIR/web\_utilities/services/openid.

If you have previously used `OpenID` to authenticate users launching applications,
you must now use an alternate method, such as OpenID Connect. See the OpenID Connect
SSO topics in Genero Application Server User Guide.

## Define server configuration based on regular expressions

Starting with GWS 3.10.09, instead of using the
`ws.myident.url` FGLPROFILE entry, you can now use a regular
expression to identify several server URLs, by using the
`ws.ident.regex.url` entry.

For more details, see [Web Services FGLPROFILE:
Server configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## fglwsdl option -fRPCNamespace to support namespaces in RPC parameters

Starting with GWS 3.10.09, the fglwsdl tool supports the new option
`-fRPCNamespace` to produce BDL code supporting the `namespace`
attribute for RPC parameters.

For more details, see [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).").

## fglpass option -gid to allow agent authentication through a UNIX user group

Starting with GWS 3.10.10, the fglpass tool can be used with the
`-gid` to allow agent authentication for all users which belong to the group of the
current user executing the fglpass command.

For more details, see [fglpass](../13_programming-tools/2525-fglpass.md "The fglpass tool allows you to encrypt passwords."), [Using the password agent](../16_web-services/4561-using-the-password-agent.md "Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand.").

## Changes to security.global.protocol in FGLPROFILE

Starting with FGLGWS 3.10.22 GWS secured communication is based on the [OpenSSL 1.1](http://www.openssl.org) engine. This version
of OpenSSL always selects the security protocol. It no longer allows you to specify a specific
Transport Layer Security (TLS) or Secure Sockets Layer (SSL).

The `security.global.protocol` entry in the fglprofile file
is therefore no longer supported.

For instance, if you have set `security.global.protocol = "TLsv1.2"` to configure
OpenSSL to use TLSv1.2 for HTTPS for earlier versions, you may encounter the following error message
in your Web service:

```
OpenSSL 1.1 doesn't support specific protocol anymore
```

It is
therefore recommended to remove the `security.global.protocol` entry from your
fglprofile file. For more information on Web service security configuration,
see [HTTPS and password encryption](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## Changes to how GWS handles cookies

When upgrading your Web services, you need to be aware of the requirement for sameSite when
setting HTTP cookies. Review the changes in how the Genero Web Services (GWS) handles cookies:

- [WSHelper.WSServerCookiesType](../16_web-services/4925-wshelper-wsquerytype.md "The WSQueryType defines a dynamic array of key-value pairs that stores the query string of a URL.") has a new member (`sameSite`) added to the
  record to specify how cookies are sent to other sites. Three values are possibles for sending
  cookies: "Strict", "Lax", or "None". The default is "Lax". If `NULL`, the default
  "Lax" is used.
- The [com.HttpServiceRequest.setResponseCookies](../15_library-reference/3841-com-httpservicerequest-setresponsecookies.md "Allows the server to return cookies to be set on the client application sending the request.") method taking `WSHelper.WSServerCookiesType` as parameter may raise a runtime
  error if the `sameSite` attribute is not set correctly. For instance, you can not set
  it to "None" if secure is not set to `TRUE`. This raises the error
  `SameSite None value requires Secure to be set`.

If you use cookies on the server side, we recommend you make the changes necessary for sending
cookies, and that you recompile your sources. For instance, make sure that cookies you use to secure
your application are set with sameSite "Strict". Setting `SameSite="Strict"` on the
secure cookie ensures that only the exact URL the cookie has been set with will return the cookie to
the server. If not set, the GWS sets each cookie's sameSite attribute as "Lax".

To have SameSite available, you must recompile your sources with fglcomp 3.10.24 or greater.

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
