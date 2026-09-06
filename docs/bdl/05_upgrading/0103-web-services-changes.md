---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_501_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.01 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 5.01.

## The handling of NULL values for primitive types during serialization has changed

Starting from FGLGWS 5.01.00, there are changes to how nulls are handled in primitive types.
Previously, when `null` values ​​were passed, the GWS sent `null`. Now
[error-15807](../15_library-reference/4483-genero-bdl-errors.md) will be raised if
a `null` value is sent. Where previously you did not have an option to allow null
values, you must now review your code to explicitly allow nulls. The following is a summary of the
behavioral changes for handling nulls in primitive types, objects, and arrays:

- Primitive type: If the value is `NULL`:
  - **Before:** It was serialized to `NULL`.
  - **Now:** If the `json_null="null"` attribute is not set or the [serializeNullAsDefault](../15_library-reference/3685-json-serializer-options.md) option is not set to 1, a serialization error is generated.
- Object: If a primitive type field is `NULL`, it is omitted unless there is a
  `JSONRequired` attribute set:
  - **Before:** It was serialized to `NULL`.
  - **Now:** If the [json\_null="null"](../16_web-services/4847-openapi-types-mapping-bdl.md) attribute is not set or the [serializeNullAsDefault](../15_library-reference/3685-json-serializer-options.md) option is not set to 1, a serialization error is generated.
- Array: If an index is `NULL`:
  - **Before:** It was serialized to `NULL`.
  - **Now:** If the `json_null="null"` attribute is not set or the [serializeNullAsDefault](../15_library-reference/3685-json-serializer-options.md) option is not set to 1, a serialization error is generated.

For more information about serialization, go to [json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.").

## Using the `json_null="null"` attribute with `JSONRequired`

Starting from FGLGWS 5.01.00, [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") class
introduces changes to how nulls are handled with `JSONRequired`.

Previously, when `null` values ​​were passed in a record, or on one of its
primitive elements, marked with `JSONRequired`, the GWS sent `null`.
Now serialization [error-15807](../15_library-reference/4483-genero-bdl-errors.md)
will be raised if a `null` value is sent. Where previously you did not have an option
to allow null values, you must now review your code to explicitly allow null values.

To explicitly allow a null value, you must set the
[json\_null="null"](../16_web-services/4847-openapi-types-mapping-bdl.md) attribute on the relevant field, or set the [serializeNullAsDefault](../15_library-reference/3685-json-serializer-options.md) option for the [The json.Serializer class](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") to allow
nulls.

For more information about `JSONRequired` attribute, go to [JSONRequired](../15_library-reference/3683-jsonrequired.md "Specify properties that are required in a JSON schema.").

## New HttpRequest methods dynamically override FGLPROFILE entries

Starting from FGLGWS 5.01.00, new methods in the `HttpRequest` class have been
implemented to allow you to override configurations in FGLPROFILE dynamically at runtime.

Calling [setCertificateAndKey()](../15_library-reference/3873-com-httprequest-setcertificateandkey.md "Specifies the certificate and key to use for the HttpRequest request.")
overrides configuration for certificates that might be set in [FGLPROFILE](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography."). Calling the [setCipher()](../15_library-reference/3875-com-httprequest-setcipher.md "Defines the type of cipher to use for encryption and decryption.") and [setVerifyServer()](../15_library-reference/3885-com-httprequest-setverifyserver.md "Defines if certificates for applications or services are validated on each request.") override other parameters
that may be set in FGLPROFILE.

While clearing a parameter using [clearCertificateAndKey()](../15_library-reference/3853-com-httprequest-clearcertificateandkey.md "Removes the client certificate and key set by setCertificateAndKey()."), [clearCipher()](../15_library-reference/3854-com-httprequest-clearcipher.md "Removes the cipher set by setCipher()."), and [clearVerifyServer()](../15_library-reference/3856-com-httprequest-clearverifyserver.md "Removes the value set by setVerifyServer.") resets values to the ones configured in FGLPROFILE.

For details, go to [com.HttpRequest methods: connection options](../15_library-reference/3848-httprequest-methods.md).

## New xml.keystore.cadir entry in FGLPROFILE for XML signature verification

Starting from FGLGWS 5.01.00, an `xml.keystore.cadir` entry can be used to specify
a directory for all trusted X509 certificates to be used during XML signature verification.
Previously, you could only specify a priority list of certificates with the
`xml.keystore.calist` entry.

Now, `xml.keystore.cadir` allows you to provide all the CA certificates in a
directory. You can continue to use the `xml.keystore.calist` as before; as files in
the "cadir" directory will be appended to those in `xml.keystore.calist`.

For more details, go to [FGLPROFILE entries for XML configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## New `json.Serializer.getLastErrorDescription()` method has improved error messaging

Starting from FGLGWS 5.01.01, a new method in the `json.Serializer` class has been
implemented to allow you return last generated serialization error message.

Calling [json.Serializer.getLastErrorDescription](https://4js.com/online_documentation/fjs-fgl-manual-html/fgl-topics/c_gws_serializer_getLastErrorDescription.html) returns a more
complete message than `sqlca.sqlerrm`. It will return the name of the class to
identify the source of the error, the name of the fields, when known, along with a detailed
description of the error.

## Specific exception -15582 for status code not in the range 200 - 299

Starting from FGLGWS 5.00.00 and 5.01.02, the GWS method [com.WebServiceEngine.SetRestStatus](../15_library-reference/3801-com-webserviceengine-setreststatus.md "Manages HTTP response status codes (200 - 299) for a REST high-level web service function.") raises a specific error code [-15582](../15_library-reference/4483-genero-bdl-errors.md), when the status code is not
in the range 200 - 299.

## Specific exception -15583 when setting a LANG value not recognized by iconv

Starting from FGLGWS 5.01.02, the GWS raises a specific error code [-15583](../15_library-reference/4483-genero-bdl-errors.md) when the
`LANG` value is not a codeset recognized by iconv.

## Register web services and resources from a package

Starting from FGLGWS 5.01.02, the GWS supports package-style modules in the
`com.WebServiceEngine.RegisterRestService()` and
`com.WebServiceEngine.RegisterRestResources()` methods.

No action needs to be taken on your part, but if you have previously wanted to use a Genero BDL
[package](../09_advanced-features/0816-package.md "Defines the package the module belongs to.") when registering web services, this is now
possible.

For more information on how to use this feature, go to [com.WebServiceEngine.RegisterRestService](../15_library-reference/3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine.") or [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine.").

## New `json.Serializer` option for handling implicit conversion

To implement strict JSON schema validation while allowing flexibility when needed, we are
introducing an option for implicit conversions during JSON to BDL deserialization.

Starting with FGLGWS 5.01.03, the global option `allowImplicitConversion` in
`json.Serializer` permits implicit type conversion for
`INTEGER`/`NUMBER(DECIMAL)`, `BOOLEAN`, and
`STRING` values in an array. For example, a boolean value can be represented as
`1`, `0`, `true`, `false`,
`"true"`, `"false"`, or `"1"`,
`"0"`.

You can now use `CALL json.Serializer.setOption("allowImplicitConversion", 1)` in
your server module to enable this option and avoid deserialization errors related to type
mismatches.

For details, go to [json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.").

## fglrestful now recognizes `+json` media types (RFC 6839)

Starting from FGLGWS 5.01.06 and 6.00.03, the fglrestful tool recognizes
all media types derived from JSON via the `+json` structured syntax suffix defined
by [RFC 6839](https://www.rfc-editor.org/rfc/rfc6839) (external link). Previously, only `application/json` and `text/json`
were recognized; other `+json` media types in OpenAPI content sections were handled
as binary, causing the generated client stub code to expect binary data rather than JSON.

No action is required. If your OpenAPI specification uses `+json` media types
such as `application/problem+json` or `application/vnd.api+json`,
fglrestful now generates the corresponding client stub code from the associated
OpenAPI schema, with JSON serialization of the request body and JSON parsing of the
response.

For more details about supported media types, go to [fglrestful](../13_programming-tools/2524-fglrestful.md).

## fglrestful now detects `text/event-stream` (SSE) responses

Starting from FGLGWS 5.01.07 and 6.00.03, fglrestful detects the
`text/event-stream` (Server-Sent Events, SSE) response content type but does not
generate response-handling code for it, because the Genero HTTP client does not provide an API to
consume SSE streams. Previously, fglrestful could generate an incorrect response
handler for SSE responses.

No action is required if your OpenAPI specification uses `text/event-stream`
alongside another response content type. If `text/event-stream` is the only declared
response content type, fglrestful skips the operation, issues a warning, and
writes the operation as a comment in the generated .4gl file (run fglrestful
with `-W yes` to list skipped operations). You must implement SSE handling manually
if you need this operation. For more about SSE, see [Server-Sent Events on MDN](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events).

For more details about supported media types, go to [fglrestful](../13_programming-tools/2524-fglrestful.md).

## CHAR(n) fields now include maxLength in the OpenAPI schema

Starting from FGLGWS 5.01.07 and 6.00.03, the OpenAPI specification generated for a REST service
includes `maxLength: n` for a `CHAR(n)` field, aligning it with
`VARCHAR(n)`; a client stub regenerated with fglrestful maps the
field to `VARCHAR(n)` instead of `STRING`. Where previously
`CHAR(n)` produced only `type: string`, you must now review stub files
generated for services that expose `CHAR(n)` fields, or pass
`--ignore-restrictions` to retain the `STRING` mapping. See
[OpenAPI types mapping: BDL](../16_web-services/4847-openapi-types-mapping-bdl.md "Conversion mapping for Genero BDL data types in OpenAPI documentation.").

## Discontinued support for OCSP

We are discontinuing support
for the Online Certificate Status Protocol (OCSP) due to the decision by several Certificate
Authorities to phase out support for this protocol in web services. However, please note that the
OCSP functionality will remain intact within our product, allowing you to continue using this
feature as part of your secure communications.

For more information on OCSP, go to [Enable OCSP](../16_web-services/4574-enable-ocsp.md "To enable Online Certificate Status Protocol (OCSP), set the security.global.ocsp.enable and security.global.ocsp.url entries in FGLPROFILE.").

## Restored HandleRequest() behavior for root-level query-only requests

The behavior of `com.WebServiceEngine.HandleRequest()` for root‑level, query‑only requests has been restored to match versions prior to 5.01.02 and 4.01.09. The sequence of changes is as follows:

| Version | Behavior |
| --- | --- |
| Before v5.01.02 and v4.01.09 | Query‑only requests to the REST service root path (for example, `/service?abc`) were treated as non‑REST and fell back to low‑level handling (status `1`). |
| v5.01.02 and v4.01.09 | These requests were interpreted as REST requests. Because no operation exists on the service root path, the engine returned error `−35` instead of using the low‑level handler. |
| Now with v4.01.11, v5.01.06, and v6.00.02 | The previous fallback behavior has been restored. Root‑level query‑only requests again fall back to low‑level handling when no REST operation matches. |

This restoration ensures compatibility with existing applications that depend on low‑level handling of root‑level query‑only requests.

For more details and examples, go to [com.WebServiceEngine.HandleRequest](../15_library-reference/3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.")

## Improved communication protocol between GAS and GWS

The communication protocol between the Genero Application Server (GAS)
and the Genero Web Service (GWS) processes has been updated to improve reliability, specifically to
ensure that no connections are lost when the GWS process needs extra time to stop or restart. Prior
to this update, if the GWS process needed time to stop, connections could be lost or requests could
be sent before the service was ready. The protocol was enhanced so that `gwsproxy`
now ensures `fglrun` is ready to receive the next request before proceeding.

To ensure you are using this improved communication protocol between the GAS and
your GWS applications, ensure you are using the following version pairs (or greater):

- FGLGWS 5.01.06 and GAS 5.01.05

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
