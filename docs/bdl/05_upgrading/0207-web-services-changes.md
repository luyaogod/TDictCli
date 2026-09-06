---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_300_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 3.00.

## Migration to 3.00 on client side

If migrating from a version 2.xx of a GWS client application to version 3.00, you need to
regenerate all client stubs in your application using the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool.

> **Important:**
>
> It is mandatory to regenerate the client stubs, to support fault response
> with HTTP error code of 200. For more information see SOAP fault handling in client stub.

See also [Change client
behavior at runtime](../16_web-services/4607-change-ws-client-behavior-at-runtime.md "Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.").

## Default SSL/TLS protocol on server side

The default for the FGLPROFILE entry `security.global.protocol` is now
`SSLv23`, enabling all supported SSL/TLS protocols, including
`TLSv1.2` as required by the Federal Law of USA. In prior versions, the default was
TLSv1 (v1.0). It is up to the web server administrator to restrict the SSL/TLS protocol to
TLSv1.2.

> **Important:**
>
> Starting with FGLGWS version 3.00.24 GWS secured communication is based on the [OpenSSL 1.1](http://www.openssl.org) engine. This version
> of OpenSSL always selects the security protocol. It no longer allows you to specify a specific
> Transport Layer Security (TLS) or Secure Sockets Layer (SSL).
>
> The `security.global.protocol` entry in the fglprofile file
> is therefore not supported.
>
> For instance, if you have set `security.global.protocol = "TLsv1.2"` to configure
> OpenSSL to use TLSv1.2 for HTTPS for earlier versions, you may encounter the following error message
> in your Web service:
>
> ```
> OpenSSL 1.1 doesn't support specific protocol anymore
> ```
>
> It is
> therefore recommended to remove the `security.global.protocol` entry from your
> fglprofile file.

For more details, see [HTTPS and password encryption](../16_web-services/4916-fglprofile-entries-for-web-services.md)

## Server socket read/write timeout on server side

Before version 3.00, when a WS client did not send all the HTTP body (for instance, after
connection has been accepted), by default the WS server would wait indefinitely, and this could end
up in a denial of service.

The `com.WebServiceEngine` class supports now a new option called
`server_readwritetimeout`, to define the server socket read/write timeout: If a
timeout occurs, the WS server program will raise the BDL exception [-15553](../15_library-reference/4483-genero-bdl-errors.md). By default this timeout is
defined as 5 seconds.

For more details, see [WebServiceEngine options](../15_library-reference/3804-webserviceengine-options.md).

## HTTPPart header default settings with com.HTTPPart.CreateAttachment()

The `com.HTTPPart.CreateAttachment()` method now by default creates header fields
based on the filename and file extension.

For more details, see [com.HttpPart.CreateAttachment](../15_library-reference/3915-com-httppart-createattachment.md "Creates a new HttpPart object based on a given filename located on disk.").

## File path returned by com.HTTPPart.getAttachment()

Before version 3.00, the `com.HTTPPart.getAttachment()` method returned the path
to a temporary file. Starting with Version 3.00, this method will now return the absolute path
location of the received part filename, based on the "Content-Disposition" header.

For more details, see [com.HttpPart.getAttachment](../15_library-reference/3917-com-httppart-getattachment.md "Returns the absolute path to the HTTP part.").

## XForms characters in com.HttpServiceRequest.readFormEncodedRequest()

Starting with version 3.00, if the result string of the HTTP request contains
`&` or `=` XForms special characters, these are escaped by
doubling them.

For more details, see [com.HttpServiceRequest.readFormEncodedRequest](../15_library-reference/3832-com-httpservicerequest-readformencodedrequest.md "Returns the string of a GET request with UTF-8 conversion option.").

## Specific exception -15575 when GAS disconnects web service server

The GWS methods listed below raise an exception with a specific error code [-15575](../15_library-reference/4483-genero-bdl-errors.md), when the GAS disconnects
properly from the web service server. Before version 3.00, the generic error [-15565](../15_library-reference/4483-genero-bdl-errors.md) was raised. A specific error
code allows you to distinguish fully a normal disconnection from other errors, in a
`TRY/CATCH` block. See code examples in method reference pages:

- [com.WebServiceEngine.GetHTTPServiceRequest](../15_library-reference/3788-com-webserviceengine-gethttpservicerequest.md "Get a handle for an incoming HTTP service request.")
- [com.WebServiceEngine.HandleRequest](../15_library-reference/3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.")

## SOAP fault handling in client stub

Web Services client stub generation has been changed to support fault response with HTTP error
code of 200.

The generated code supports SOAP fault with HTTP error code of 200 and 500. To enable this new
feature in your client stub code, regenerate the stubs with the fglwsdl
tool.

For more details, see [Client side SOAP fault handling](../16_web-services/4516-client-side.md "A Genero Web Services client can receive a SOAP fault number in the operation status and act accordingly.").

## Optional multipart handling in client stub

In the generated client stub code, all functions handling the SOAP request with multipart get
an additional input parameter and/or return parameter as a `DYNAMIC ARRAY OF com.HTTPPart`,
to pass and return optional parts.

When generating client stubs managing multipart, you get an extra input and/or output variable
called "`AnyInputParts`" and "`AnyOutputParts`" that is a
`DYNAMIC ARRAY` of `com.HTTPPart` objects. Those variables may contain
additional input and/or output HTTP parts not specified in the WSDL. You will have to adapt your
client program by handling those dynamic arrays in any Genero functions calling such stubs.

Request example prior to 3.00:

```
FUNCTION xxx_g(InputHttpPart_1, ..., InputHttpPart_n)
  DEFINE InputHttpPart_1 com.HTTPPart
  ...
  DEFINE InputHttpPart_n com.HTTPPart
  ...
  RETURN wsstatus
END FUNCTION
```

Request example 3.00 and greater, with extra input variable `AnyInputParts`:

```
FUNCTION xxx_g(InputHttpPart_1, ..., InputHttpPart_n)
  DEFINE InputHttpPart_1 com.HTTPPart
  ...
  DEFINE InputHttpPart_n com.HTTPPart
  DEFINE AnyInputParts DYNAMIC ARRAY OF com.HTTPPart
  ...
  RETURN wsstatus
END FUNCTION
```

> **Note:**
>
> This change has also been backported in 2.50.25.

For more details, see [Multipart in the client stub](../16_web-services/4626-multipart-in-the-client-stub.md "You can generate a client stub for a Web service that has multiple parts.").

## Removal of FGLWSNOINFO environment variable

Before version 3.00, the GWS library displayed by default a message about certificates used by the
program:

```
—
WS-INFO (Certificate authority) | Loading from Windows keystore
—
```

To avoid this message, it was possible to set the FGLWSNOINFO environment variable to TRUE.

Starting with version 3.00, this message is no longer displayed by the GWS library, and the
FGLWSNOINFO is no longer required.

## Desupported Web Services APIs

The methods listed in the following table are desupported in Genero 3.00.

| Methods desupported as of 3.00 | Alternative method to use |
| --- | --- |
| `com.Util.CreateDigestString` | `security.Digest.CreateDigestString` |
| `com.Util.CreateRandomString` | `security.RandomGenerator.CreateRandomString` |
| `com.Util.CreateUUIDString` | `security.RandomGenerator.CreateUUIDString` |
