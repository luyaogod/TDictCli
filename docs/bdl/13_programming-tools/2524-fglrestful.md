---
title: "fglrestful"
source: "fgl-topics/c_gws_tools_fglrestful.html"
breadcrumb: "Programming tools > Command reference > fglrestful"
type: "concept"
---

# fglrestful

> The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.

## Syntax

```
fglrestful [ options ] service
```

1. options are described in fglrestful
   options and in fglrestful network options.
2. service is the OpenAPI service definition file or URL.

## Options

| Options | Description |
| --- | --- |
| `-V` or `--version` | Displays version information. |
| `-h` or `--help` | Displays options for the tool. |
| `-l` or `--list` `{` `path \| resource` `}` | List all the paths or resources of a service. See List path or resource (--list). |
| `-o` or `--output filename` | Specify a name for the output stub file. See Generate the stub file. |
| `-n` or `--name resource` | In a service providing access to one or more resources, specify the creation of a stub file for a given resource. See Generate stub file for resource. |
| `-p` or `--prefix prefix` | Set a prefix for constants and variables. |
| `--typePrefix prefix` | Set a prefix for types. Where prefix allows you to add a prefix in front of the name of types in a JSON schema. See Set types with prefix value (--typePrefix). |
| `-b` or `--binary` `{` `byte \| file` `}` | Specify that binary types are generated as either `BYTE` or `FILE` types in the stub. |
| `-f` or `--format` `{` `xml \| json` `}` | Specify the preferred REST format (xml or json). You must use lowercase xml or json in the command. |
| `-a` or `--oauth` `{` `yes \| no` `}` | Specify if OAuth specification should be generated or not:With option "yes", OAuth support code is generated even if the OpenAPI documentation does not specify it.With option "no", OAuth support code is not generated even if the OpenAPI documentation requires it.If you do not use the `--oauth` argument, fglrestful generates the OAuth support code by default if it is specified in the OpenAPI documentation. |
| `--legacyJSONApi` | Generate stub file with legacy JSON API (`util.JSON`) when used with the --output option. See Usage. |
| `--ignore-restrictions` | Ignore OpenAPI format keywords (`maxLength`, `multipleOf`, `minimum`, `maximum`) when converting OpenAPI types to BDL types in the stub. Affected types: `CHAR(N)`, `VARCHAR(N)`, `DECIMAL(P,S)`, and `DATETIME`. See [Converting OpenAPI types to BDL in client stubs](../16_web-services/4847-openapi-types-mapping-bdl.md). |
| `--datetime-fraction { 1 \| 2 \| 3 \| 4 \|5 }` | Specify `DATETIME` fraction qualifier. The default is `NULL`, meaning `YEAR TO SECOND`. If value is between 1 and 5, fglrestful generates `YEAR TO FRACTION(N)` where N is the value provided. |
| `--interval qualifier` | Specify the `INTERVAL` qualifier used for OpenAPI duration properties (`type: string`, `format: duration`). The default is `DAY(9) TO FRACTION(5)`. Use `STRING` instead of an `INTERVAL` qualifier to generate `STRING` variables for duration properties. See Set the INTERVAL qualifier for durations (--interval). |
| `--datetime-interval-format format` | Specify the JSON serialization format for `DATETIME` and `INTERVAL` values in generated stubs:`native` (default): BDL string representation (for example, `2024-11-05 09:15:58`)`iso8601` (UTC)`iso8601-local` (local timezone offset)This option affects only values sent in requests. Received `DATETIME` and `INTERVAL` values in responses are auto-detected; both native and ISO 8601 formats are accepted.An invalid format value stops fglrestful with `Error: invalid DATETIME/INTERVAL format` and a non-zero exit status.For details, see the example in Serialize DATETIME and INTERVAL in ISO 8601 (--datetime-interval-format). |
| `-t` or `--token token` | Specify the access token value if the service is protected by an access token. |
| `-k` or `--tokenfile filename` | Specify the access token file if the service is protected by an access token. |
| `--comment` | Add comments to the generated stub file. See Output comments (--comment). |
| `-W` or `--Warnings` `{` `yes \| no` `}` | Specify if warnings should be output or not. By default the option value is `yes`, meaning the warnings are reported. See Output warnings (-W). |
| `--no-merge-allof` | Disable merging of properties in '`allof`' schemas. See [JSONAllOf](../15_library-reference/3679-jsonallof.md "Combine multiple record schemas into a single type using the JSONAllOf attribute in Genero BDL.") |

| Options | Description |
| --- | --- |
| `--proxy location` | Connect via proxy where location is `host[:port]` or `ip[:port]`. |
| `--pLogin login` | Proxy authentication login. |
| `--pPass pass` | Proxy authentication password. |
| `--hLogin login` | HTTP authentication login. |
| `--hPass pass` | HTTP authentication password. |

## Usage

Use fglrestful to generate GWS client stub files from an OpenAPI version 3
specification. The typical command form is:

```
fglrestful -o stubname spec
```

where spec is either an OpenAPI file or a service URL. Run
fglrestful -h to list all available options.

If the service uses the Genero REST high-level API, the specification is available by default
at the ?openapi.json query string. For example, to generate a stub from a
service running locally:

```
fglrestful -o myStub http://localhost:8090/MyService?openapi.json
```

If your service uses versioning, request a specific version by appending
`&version=version-name` to the query string. For examples,
see Generate stub file for a version.

fglrestful returns a non-zero exit code on error, making it suitable for
use in scripts and makefiles.

The fglrestful tool recognizes all media types derived from JSON via the
`+json` structured syntax suffix defined by [RFC 6839](https://www.rfc-editor.org/rfc/rfc6839) (external link). The subtype preceding
`+json` must not be empty; the comparison is case-insensitive. All recognized types
are handled like `application/json`: JSON serialization of the request body,
JSON parsing of the response, and generation of content handling code from the associated
OpenAPI schema.

| Media type | Description |
| --- | --- |
| `application/json`, `text/json` | Standard JSON types, supported in all versions. |
| `application/*+json` | Any application subtype with the `+json` suffix. Common examples:`application/ld+json`: JSON Linked Data; enriches JSON with metadata such as `@id` and `@context`.`application/problem+json`: RFC 7807 standard format for returning error details in JSON.`application/vnd.api+json`: JSON:API specification for building consistent REST APIs.`application/hal+json`: Hypertext Application Language (HAL); a JSON format for hypermedia APIs. |
| `text/*+json` | Any text subtype with the `+json` suffix. |
| `text/event-stream` | `text/event-stream` is the media type for [Server-Sent Events (SSE)](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events), a web standard for streaming real-time updates from a server to a client. fglrestful does not generate response-handling code for this type. Behavior depends on the operation:If the operation also declares another response content type, the stub is generated against that type.If `text/event-stream` is the only declared response content type, the operation is skipped; a warning is issued and the skipped operation is written as a comment in the generated .4gl file. Run fglrestful with warnings enabled (`-W yes`) to see skipped operations. You must implement SSE handling for this operation manually. |

## Date-time and duration options

Use `--datetime-fraction` and `--interval` to override the default
BDL types that fglrestful generates for OpenAPI temporal properties. Use
`--datetime-interval-format` to control how `DATETIME` and
`INTERVAL` values serialize to JSON.

| Option | OpenAPI property | Default BDL type | Override example |
| --- | --- | --- | --- |
| `--datetime-fraction` | `type: string`, `format: date-time` | `DATETIME YEAR TO SECOND` | `--datetime-fraction 3` generates `DATETIME YEAR TO FRACTION(3)` |
| `--interval` | `type: string`, `format: duration` | `INTERVAL DAY(9) TO FRACTION(5)` | `--interval "HOUR TO SECOND"` generates `INTERVAL HOUR TO SECOND`; `--interval STRING` generates `STRING` |

For examples, see Set the INTERVAL qualifier for durations (--interval) and
Serialize DATETIME and INTERVAL in ISO 8601 (--datetime-interval-format).

> **Warning:**
>
> BDL [`INTERVAL`](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") qualifiers fall into one of two incompatible classes: year-month
> (`YEAR` to `MONTH`) or day-time (`DAY` to
> `FRACTION`). If the service returns a year-month duration (for example
> `P2Y3M`) and the stub uses a day-time qualifier (the default
> `DAY(9) TO FRACTION(5)`), the runtime raises error -15806. Use `--interval "YEAR TO MONTH"` when the
> service returns only year-month durations, or `--interval STRING` when the
> specification mixes both classes.

> **Note:**
>
> Only qualifiers that end in `FRACTION` preserve fractional seconds:
> `DAY`, `HOUR`, `MINUTE`, or
> `SECOND TO FRACTION(n)`. If the qualifier you pass to
> `--interval` does not end in `FRACTION` — for example
> `HOUR TO SECOND` — the runtime silently drops any fractional seconds in a received
> duration value; `PT1.5S` becomes `PT1S` without error. The default
> qualifier `DAY(9) TO FRACTION(5)` preserves up to five fractional-second
> digits.

In the next sections there are examples of common commands used by REST web service developers.

## Generate the stub file

Use the `-o` / `--output` option to name the output stub file:

- Using the URL of the running
  service:

  ```
  fglrestful -o myStub http://localhost:8090/MyService?openapi.json
  ```
- Or using an OpenAPI documentation file. If you access the OpenAPI information on the browser,
  you can copy the description to a file and save the file with an openapi
  extension, for example, myservice.openapi. Then run the command to generate the
  stub file from this:

  ```
  fglrestful -o myStub myservice.openapi
  ```

  > **Tip:**
  >
  > In Genero Studio, you can place the openapi file in your project and the
  > stub file is generated from it when you compile the file. For more information, see the Genero Studio User Guide.

The generated client stub file, myStub in the example, is a Genero BDL
file (.4gl).

By default, fglrestful uses the [JSON streaming API](../15_library-reference/3616-the-json-package.md "The Genero Web Services JSON package provides classes and methods to process JSON documents.") to generate the stub file code. It is
possible to use the old JSON API (`util.JSON`) instead by specifying the
`--legacyJSONApi`
option:

```
fglrestful -o myStub --legacyJSONApi myservice.openapi
```

## Generate stub code for a `+json` media type

If an OpenAPI specification declares a response using an `application/problem+json`
content type, fglrestful recognizes it and generates the corresponding client
stub code. For example, given this OpenAPI response declaration:

```
responses:
  '400':
    description: Bad request
    content:
      application/problem+json:
        schema:
          $ref: '#/components/schemas/Problem'
```

Running fglrestful generates a stub function containing a
`CASE` block that tests the response `Content-Type` and
retrieves the body:

```
WHEN retCode == 400 # Bad request
  IF contentType MATCHES "*application/problem+json*" THEN
    # Parse FILE response
    LET getproblemjson_400 = resp.getFileResponse()
    RETURN C_GETPROBLEMJSON_400, resp_body.*
  END IF
  LET wsError.code = retCode
  LET wsError.description = "Unexpected Content-Type"
  RETURN -1, resp_body.*
```

## Set types with prefix value (`--typePrefix`)

Use `--typePrefix` to avoid naming conflicts with Genero reserved words or library
names (such as "base" or "util"). In this example, the prefix "foo" is added to all named types:

```
fglrestful -o myStub --typePrefix foo http://localhost:8090/MyService?openapi.json
```

In
the generated stub file, the "Base" type is renamed, and whereever it is referenced, to
"foo\_Base":

```
# components/schemas/foo_Base
PUBLIC TYPE foo_Base RECORD
    id INTEGER,
    path STRING
END RECORD
#...
PUBLIC TYPE useBase RECORD
  id INTEGER,
  country foo_Base,
  postal_code VARCHAR
END RECORD
```

## List path or resource (`--list`)

Use `--list` to output all paths or resources of a service. These commands list the
resources:

- using the URL of the running
  service:

  ```
  fglrestful -l resource http://localhost:8090/MyService?openapi.json
  ```
- or using a specification
  file:

  ```
  fglrestful -l resource myservice.openapi
  ```

## Generate stub file for resource

Use `-n` to generate a stub file for a specific resource. These commands generate a
stub for the "users" resource:

- using the URL of the running
  service:

  ```
  fglrestful -n users -o myUsersStub http://localhost:8090/MyService?openapi.json
  ```
- or using a specification
  file:

  ```
  fglrestful -n users -o myUsersStub myservice.openapi
  ```

## Generate stub file for a version

Use the `&version=version-name` query string to generate a
stub for a specific version. These commands generate a stub for "v2":

- using the URL of the running service (the URL must be enclosed in
  quotes):

  ```
  fglrestful -o myVersion2Stub "http://localhost:8090/MyService?openapi.json&version=v2"
  ```
- or using an OpenAPI documentation file that contains the version 2 operations
  only:

  ```
  fglrestful -o myV2Stub myserviceV2.openapi
  ```

If the requested version does not exist, the 404 error will be returned.

## Output comments (`--comment`)

Use `--comment` to include [WSDescription](../16_web-services/4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") and [WSTypeDescription](../16_web-services/4810-wstypedescription.md "Provides a description for a user‑defined type in the REST service.") text in the generated stub file.
These commands generate a stub with comments:

- using the URL of the running
  service:

  ```
  fglrestful -o myStub --comment http://localhost:8090/MyService?openapi.json
  ```
- or using a specification
  file:

  ```
  fglrestful -o myStub --comment  myservice.openapi
  ```

The descriptive text appears in the generated stub file in curly braces, for
example:

```
PUBLIC TYPE user RECORD {This is a user}
```

Or

```
 p_resourceId INTEGER {this is the customer id}
```

The `WSDescription` set on the `ATTRIBUTES()` clause of the function is not
affected. It is always generated; even if the `--comment` option is not
specified.

## Set the INTERVAL qualifier for durations (`--interval`)

Use `--interval` to override the default `INTERVAL DAY(9) TO
FRACTION(5)` qualifier for duration properties. This command generates
`INTERVAL HOUR TO SECOND` variables:

```
fglrestful --interval "HOUR TO SECOND" -o stub myapi.json
```

To generate `STRING` variables instead (pre-6.00.03 behavior):

```
fglrestful --interval STRING -o stub myapi.json
```

## Serialize DATETIME and INTERVAL in ISO 8601 (`--datetime-interval-format`)

This command generates a stub that sends `DATETIME` values in RFC 3339 UTC and
`INTERVAL` values as ISO 8601 durations:

```
fglrestful --datetime-interval-format iso8601 -o stub myapi.json
```

In the generated stub, the option translates to [`json.Serializer.setOption()`](../15_library-reference/3675-json-serializer-setoption.md "Sets an option on the JSON serializer.") calls
placed before each JSON request body is sent:

```
CALL req.setHeader("Content-Type", "application/json")
CALL json.Serializer.setOption("datetimeSerializationMode", 1)
CALL json.Serializer.setOption("intervalSerializationMode", 1)
VAR writer = req.beginJSONRequest()
CALL writer.startJSON()
CALL json.Serializer.variableToJSON(p_body, writer)
CALL writer.endJSON()
CALL req.endJSONRequest(writer)
```

For details about the `datetimeSerializationMode` and
`intervalSerializationMode` options, see [datetimeSerializationMode](../15_library-reference/3685-json-serializer-options.md) and [intervalSerializationMode](../15_library-reference/3685-json-serializer-options.md).

For the output formats, see [DATE, DATETIME, and INTERVAL conversions](../16_web-services/4852-date-datetime-and-interval-conversions.md "GWS serializes DATE, DATETIME, and INTERVAL values as JSON strings using either the Genero BDL native format or standard RFC 3339 / ISO 8601 formats.").

> **Note:**
>
> This option has no effect when used with `--legacyJSONApi`.
> Free-form and multipart JSON bodies always use the native format. fglrestful
> outputs the warning: `Option --datetime-interval-format has no effect with
> --legacyJSONApi`.

## Output warnings (`-W`)

Use `-W yes` to output warnings when generating the stub. These commands write
warnings to standard output:

- using the URL of the running
  service:

  ```
  fglrestful -W yes -o myStub http://localhost:8090/MyService?openapi.json
  ```
- or using a specification
  file:

  ```
  fglrestful -W yes -o myStub myservice.openapi
  ```

The warning message gives the path in the openapi.json file where you locate
the issue, and gives the
reason.

```
Warning in /paths/Property/GetAllProperties/{applicationID}/{includeWillBeSold}/{update}/get/responses
Reason : Unsupported media='text/plain' on type='WebAPI_CreditAnalysis_Models_PropertyContainer'
```

Errors are reported by default. The error message gives same kind of path, but with `Error
in` instead of `Warning in`.

## Related links

**Related concepts**  

[RESTful Web services (high-level framework)](../16_web-services/4705-restful-web-services-high-level-framework.md "These topics give you the information you need to begin working with RESTful Web services applications using BDL function with support for attributes.")

**Related reference**  

[Error codes of com.WebServicesEngine](../15_library-reference/3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods.")
