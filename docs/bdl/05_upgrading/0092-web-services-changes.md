---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_600_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 6.00 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 6.00.

## REST functions must include a `RETURNS()` clause

Starting from FGLGWS 6.00.01, all REST functions must declare a `RETURNS()`
clause, even when the function does not return a value. This enforces consistent function signatures
and prevents ambiguous web service definitions.

**Action needed**:

- You must add a `RETURNS()` clause to every REST function declaration.
- Update your existing REST functions to include a `RETURNS()` declaration — either
  an explicit return type or an empty `RETURNS()`.

Omitting the `RETURNS` clause will raise error [-9158](../15_library-reference/4483-genero-bdl-errors.md)

For example:

**Before (no return):** 

```
FUNCTION notifyUser(userId INTEGER)
    ATTRIBUTE(WSPost)
   # ...
END FUNCTION
```

**Before (returns
data)**:

```
FUNCTION notifyUser(userId INTEGER)
    ATTRIBUTE(WSPost)
  # ...
  RETURN "data"
END FUNCTION
```

**Now (no return value — explicit empty `RETURNS()` must be
added):**

```
FUNCTION notifyUser(userId INTEGER)
    ATTRIBUTE(WSPost)
# The web service function 'notifyUser' requires a RETURNS() clause.
# See error number -9158.
END FUNCTION
```

**Now (returns data — explicit
`RETURNS(return-type)` must be
added):**

```
FUNCTION notifyUser(userId INTEGER)
    ATTRIBUTE(WSPost)
# The web service function 'notifyUser' requires a RETURNS() clause.
# See error number -9158.
  RETURN "data"
END FUNCTION
```

For more information about using `RETURN` clauses in RESTful
functions, go to [Return value attributes](../16_web-services/4801-return-value-attributes.md "A RESTful web service function can have return values.")

## JSONEnum: separator changed from pipe to comma

Starting from FGLGWS 6.00.01, the `JSONEnum` attribute uses **commas** instead
of pipes to separate values in its list (for example,
`"pending","approved","rejected"`). Update declarations that use the old pipe format
to the new comma-separated format to avoid validation errors.

An invalid separator will raise error [-9159](../15_library-reference/4483-genero-bdl-errors.md)

Example:

**Before**:

```
TYPE ok_string STRING
   ATTRIBUTE(JSONEnum = `"pending"|"approved"|"rejected"`)
```

**Now**:

```
TYPE ok_string STRING
   ATTRIBUTE(JSONEnum = `"pending","approved","rejected"`)
```

For more details, go to [JSONEnum](../15_library-reference/3680-jsonenum.md "Defines an explicit, typed list of acceptable values for a field and maps to the enum keyword in JSON Schema.")

## Implicit type conversions in `json.Serializer` have been removed

Starting from FGLGWS 6.00.00, implicit type conversions in [`json.Serializer`](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") have been removed.

For example, previously you could deserialize a JSON string like `"true"` to a FGL
`BOOLEAN` (`1`), or `"123"` to a FGL
`INTEGER` (`123`).

Now, these assignments will result in errors unless explicitly authorized.

- If your application relied on automatic type conversions during JSON
  serialization/deserialization with `json.Serializer`, you must now handle these
  conversions explicitly or use [`util.json`](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") if you require the old behavior.
- The [`allowImplicitConversion`](../15_library-reference/3685-json-serializer-options.md) setting does not restore the previous implicit
  conversion behavior in `json.Serializer`.

Ensure that JSON payloads sent to your 4GL applications use the correct types, as mismatches will
now result in errors rather than silent conversions.

For more details, go to [Implicit and explicit JSON conversion for primitive types](../15_library-reference/3664-primitive-type-conversion.md "This topic describes how primitive 4GL values are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.").

## Change in allowImplicitConversion behavior

The handling of implicit conversions for BDL types has also changed. Previously, primitive types
were always allowed to convert implicitly, regardless of the `allowImplicitConversion` setting. Now, implicit conversions for all
BDL types, including primitive types, are controlled by this option. If it is not set to
`1`, implicit conversions are no longer permitted.

For more details about the `json.Serializer`
`allowImplicitConversion` option, go to [allowImplicitConversion](../15_library-reference/3685-json-serializer-options.md).

## Restored HandleRequest() behavior for root-level query-only requests

The behavior of `com.WebServiceEngine.HandleRequest()` for root‑level, query‑only requests has been restored to match versions prior to 5.01.02 and 4.01.09. The sequence of changes is as follows:

| Version | Behavior |
| --- | --- |
| Before v5.01.02 and v4.01.09 | Query‑only requests to the REST service root path (for example, `/service?abc`) were treated as non‑REST and fell back to low‑level handling (status `1`). |
| v5.01.02 and v4.01.09 | These requests were interpreted as REST requests. Because no operation exists on the service root path, the engine returned error `−35` instead of using the low‑level handler. |
| Now with v4.01.11, v5.01.06, and v6.00.02 | The previous fallback behavior has been restored. Root‑level query‑only requests again fall back to low‑level handling when no REST operation matches. |

This restoration ensures compatibility with existing applications that depend on low‑level handling of root‑level query‑only requests.

For more details and examples, go to [com.WebServiceEngine.HandleRequest](../15_library-reference/3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.")

## Enhanced control of HTTP error messages in preprocessing and postprocessing

Starting from FGLGWS 6.00.02, you can set the "`HTTPProcessingErrorMsg`" context
key in [WSPreProcessing](../16_web-services/4811-wspreprocessing.md "Specify functions in a REST web service module that are used as preprocessing request handlers.") and [WSPostProcessing](../16_web-services/4812-wspostprocessing.md "Specify functions in a REST web service module that are used as postprocessing request handlers.") callbacks to supply a custom
HTTP reason phrase for non‑zero status returns. This update extends the existing [WSContext](../16_web-services/4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.") capabilities by giving you more control
over the error message sent to the client. For more details, go to [Using preprocessing and postprocessing callbacks](../16_web-services/4751-using-preprocessing-and-postprocessing-callbacks.md "Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.")

## Global tags generation

Starting in version 6.00.02, all tag names defined with the `WSTags` attribute are now collected and generated once in the OpenAPI document’s global `tags` section.

The `tags` keyword is case-sensitive. For more details about tag behavior, go to [WSTags](../16_web-services/4826-wstags.md "Assign one or more tags to a REST operation. The value corresponds to the \"tags\" keyword in the OpenAPI specification and is used for grouping and documentation.").

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

## CHAR(n) fields now include maxLength in the OpenAPI schema

Starting from FGLGWS 5.01.07 and 6.00.03, the OpenAPI specification generated for a REST service
includes `maxLength: n` for a `CHAR(n)` field, aligning it with
`VARCHAR(n)`; a client stub regenerated with fglrestful maps the
field to `VARCHAR(n)` instead of `STRING`. Where previously
`CHAR(n)` produced only `type: string`, you must now review stub files
generated for services that expose `CHAR(n)` fields, or pass
`--ignore-restrictions` to retain the `STRING` mapping. See
[OpenAPI types mapping: BDL](../16_web-services/4847-openapi-types-mapping-bdl.md "Conversion mapping for Genero BDL data types in OpenAPI documentation.").

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

## Improved communication protocol between GAS and GWS

The communication protocol between the Genero Application Server (GAS)
and the Genero Web Service (GWS) processes has been updated to improve reliability, specifically to
ensure that no connections are lost when the GWS process needs extra time to stop or restart. Prior
to this update, if the GWS process needed time to stop, connections could be lost or requests could
be sent before the service was ready. The protocol was enhanced so that `gwsproxy`
now ensures `fglrun` is ready to receive the next request before proceeding.

To ensure you are using this improved communication protocol between the GAS and
your GWS applications, ensure you are using the following version pairs (or greater):

- FGLGWS 6.00.02 and GAS 6.00.03

## fglrestful mapping of OpenAPI duration properties

Starting from GWS 6.00.03, fglrestful generates
`INTERVAL DAY(9) TO FRACTION(5)` instead of `STRING` for OpenAPI
properties declared with `format: duration`. Use
`--interval STRING` to restore the previous `STRING` mapping. See [Date-time and duration options](../13_programming-tools/2524-fglrestful.md).

## Changes to the OpenAPI format mapping for DATETIME and INTERVAL

Starting from FGLGWS 6.00.03, the GWS generates qualifier-aware `format` keywords
for [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") in the OpenAPI
document: `DATETIME YEAR TO DAY` is now `format: date` (previously
`format: date-time`), and [`INTERVAL`](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") is now
`format: duration` (previously a plain string).
Where previously every `DATETIME` mapped to `format: date-time`
and `INTERVAL` was a plain string, you must now review your code to handle
the qualifier-aware format mapping. See [OpenAPI types mapping: BDL](../16_web-services/4847-openapi-types-mapping-bdl.md).

## DSA signature key

Starting from FGLGWS 6.00.03, the `xml.CryptoKey` class supports the
`dsa-sha256` signature key for DSA digital signatures. SHA-1 is deprecated;
where previously you used `http://www.w3.org/2000/09/xmldsig#dsa-sha1`, you
must now update your code to use
`http://www.w3.org/2009/xmldsig11#dsa-sha256`. See [Supported kind of keys](../15_library-reference/4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.").

## REST API catalog endpoint

Starting from FGLGWS 6.00.03, the GWS engine intercepts
`GET /.well-known/api-catalog` before any user request handler and returns an
RFC 9727 API catalog of the application's REST services. An application that previously handled
this path itself will now see the engine answer it instead. See
[REST API catalog](../16_web-services/4758-rest-api-catalog.md "A Genero application that publishes REST services automatically provides an RFC 9727 API catalog at /.well-known/api-catalog, listing each service and its OpenAPI description.").

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Web
services changes in BDL 5.01](0103-web-services-changes.md "There are changes in support of web services in Genero 5.01.").

Notable changes introduced in maintenance releases:

- [Stricter handling for nulls in serialization of primitive types, objects, and arrays](0103-web-services-changes.md)
  introduced in FGLGWS 5.01.00.
- [json.Serializer.getLastErrorDescription](https://4js.com/online_documentation/fjs-fgl-manual-html/fgl-topics/c_gws_serializer_getLastErrorDescription.html) method to provide a
  more complete message than `sqlca.sqlerrm` for a generated serialization error
  message introduced in FGLGWS 5.01.01.
- [Register web services and resources from a package](0103-web-services-changes.md) feature introduced in FGLGWS
  5.01.02.
- [New json.Serializer option for handling implicit conversion](0103-web-services-changes.md), `allowImplicitConversion`, introduced in FGLGWS 5.01.03.
- [Support for Online Certificate Status Protocol (OCSP)](0155-web-services-changes.md)  was discontinued for FGLGWS versions
  3.21.02, 4.01.07, and 5.01.03.
- [fglrestful now recognizes `+json` media types (RFC 6839)](0103-web-services-changes.md)
  introduced in FGLGWS 5.01.06 and 6.00.03.
- [maxLength generated for CHAR fields in OpenAPI](0103-web-services-changes.md) introduced in FGLGWS 5.01.07 and
  6.00.03.
- [fglrestful now detects `text/event-stream` (SSE) responses](0103-web-services-changes.md)
  introduced in FGLGWS 5.01.07 and 6.00.03.

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
