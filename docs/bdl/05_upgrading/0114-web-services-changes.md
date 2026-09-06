---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_500_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.00 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 5.00.

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

## New Genero Web Services JSON library for working with JSON

Starting with FGLGWS 5.00, the Genero Web Services JSON library (`json`) has been
added. This library provides classes and methods to perform:

- JSON manipulation with a streaming API for JSON
- Serialization of BDL variables in JSON

The `json` library provides an API with three classes: `JSONWriter`, `JSONReader`, and `Serializer` to stream JSON over HTTP.

Where before you may have used the JSON utility API (`util.JSON,util.JSONObject,util.JSONArray`)
classes to perform JSON manipulation, the new `json` package provides an alternative.
By contrast with the JSON utility API, the new `json` library does not load a JSON
file into memory, but instead streams and serializes data over the network on the fly. This
streaming method improves communication and performance, particularly when large JSON files need to
be handled.

Tools that process JSON also use the new `json` class. The
fglrestful tool has been enhanced to take advantage of the streaming methods of
the new `json` class to generate code in the stub file.

For more information, go to [The json package](../15_library-reference/3616-the-json-package.md "The Genero Web Services JSON package provides classes and methods to process JSON documents.").

## Changes to how WSRetCode attribute handles return status

Starting with FGLGWS 5.00, the high-level REST API attribute `WSRetCode` can be
set to support the Swagger and OpenAPI specification for the "2XX" value, which is a specific code
that can be set in the [WSRetCode](../16_web-services/4830-wsretcode.md "Sets the HTTP success status code returned by the REST function.") attribute.
This specific code means that the REST web service can return any HTTP response status value from
200 to 299 dynamically at runtime.

In earlier versions you could only set an explicit return code in the function declaration. From
version 5.00 onwards, the GWS supports both the older method and the new method using the 2XX value.
For more information and examples using the attribute, see [SetRestStatus](../15_library-reference/3801-com-webserviceengine-setreststatus.md "Manages HTTP response status codes (200 - 299) for a REST high-level web service function.") and [WSRetCode](../16_web-services/4830-wsretcode.md "Sets the HTTP success status code returned by the REST function.").

## WSDescription attribute can be used on JSON schemas

Starting with FGLGWS 5.00, the high-level REST API attribute `WSDescription` can
be used in user-defined BDL types in addition to REST function input and output parameters, added to
support the use of JSON schemas in the Swagger and OpenAPI specification.

In earlier versions you could only set this attribute on REST function input and output
parameters. From version 5.00 onwards, you can use the attribute in user-defined [type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.") with data types that use an
`ATTRIBUTES()` clause. You can set the `WSDescription` attribute on
type members to specify metadata information. For more information and examples using the attribute,
see [WSDescription](../16_web-services/4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")

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

## New `security.global.verifyserver` entry in FGLPROFILE to support certificate validation process

Starting from 5.00.00, it is now possible to turn off certificate validation for requests for
applications or services once the server has been validated.

This feature is intended for development purposes only; it allows you to disable the mechanism
that checks the chain of trust for a certificate. To do this, set the FGLPROFILE
`security.global.verifyserver` property to `FALSE`, which turns off
certificate validation when requests are made for applications or services run with the HTTPS
protocol.

For details, go to [Security Configuration FGLPROFILE entries](../16_web-services/4916-fglprofile-entries-for-web-services.md).

## fglrestful oauth option changes

Starting with FGLGWS 5.00, the fglrestful tool `--oauth` 
option has just two settings ("yes" or "no") to specify if OAuth specification should be generated
or not.

For more details, go to [fglrestful](../13_programming-tools/2524-fglrestful.md).

## New method `GetIdRoles()` for retrieving authorization roles

Starting from FGLGWS 5.00, the OAUTH API has a new method called
`OAuthAPI.GetIdRoles()` to explicitly retrieve authorization roles.

Where previously you used `OAuthAPI.GetIdScopes()` to retrieve both the list of
scopes and roles from an ID or access token provided by an Identity provider using OAuth2 Single
sign-on, you now need to use the dedicated method to retrieve roles.

For details, go to [OAuthAPI.GetIdRoles()](../16_web-services/4945-oauthapi-getidroles.md "Get OAuth ID Token authorization roles.").

## OpenIDConnect service supports `OIDC_ROLES`

Starting from FGLGWS 5.00, the Genero OpenIDConnect service now decodes ID tokens containing
roles instead of scopes, and creates a new environment variable called `OIDC_ROLES`
containing the list of roles.

Where previously you used `OIDC_SCOPES` to retrieve both the list of scopes and
roles from an ID or access token provided by an Identity provider using OpenID Connect/OAuth2 Single
sign-on, you now need to use the dedicated variable to retrieve roles.

For an example using `OIDC_ROLES`, refer to Retrieve roles and scopes
in the Single Sign-On User Guide.

## New security.global.certificate.selfsigned.preload entry in FGLPROFILE

Starting from FGLGWS 3.21.02, 4.01.07, and 5.00.02, there is an option to preload the global
self-signed certificate and private key used for HTTPS connections. Typically, the certificate and
key is loaded at the first HTTPS request. If you find the GWS computation of the certificate and key
takes too long, you can speed things up by setting the
`security.global.certificate.selfsigned.preload = TRUE` to preload the certificate
and key at the start of the application instead of at the first HTTPS connection.

For more information on web service security configuration, go to [Security Configuration FGLPROFILE entries](../16_web-services/4916-fglprofile-entries-for-web-services.md).

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

## The `GeneroAccessService` supports scopes set in configuration file

Starting from FGLGWS 5.00.02, the `GeneroAccessService` service can now secure web
services from scopes set in the configuration file (xcf).

No action needs to be taken on your part, but if you want to secure legacy REST web services that
do not define any `WSScope` attributes or that are written with the REST low-level
API, then you can secure them by configuring the `SCOPE` element in the configuration
file. The `GeneroAccessService` will verify if the access token provided by the
client application has all the necessary scopes specified in the `SCOPE` element.

For more information about Single sign-on and the GIP, refer to the Single Sign-On User Guide. For more
information about configuring scopes, refer to the SCOPE (for service) topic in the Genero Application Server User Guide.

## fglwsdl oauth option added

Starting from FGLGWS 5.00.03, the fglwsdl tool `--oauth` 
option can be used to generate OAuth specification in the SOAP service call. The stub is generated
with the `IMPORT FGL OAuthAPI` and requests use the
`OAuthAPI.CreateHTTPAuthorizationRequest()` method to take the access token into
account.

For more details, go to [fglwsdl](../13_programming-tools/2523-fglwsdl.md).

## New OAuthAPI methods to use with SSO authentication for apps accessing web services

Starting from FGLGWS 5.00.03, the OAuthAPI introduces two new dedicated methods designed
specifically for applications that do not function behind a GAS, enabling access to web services
across all applications:

1. First, call the `OAuthAPI.RetrievePasswordTokenForNativeApp()` method, using the
   client app’s client and secret credentials along with the user credentials to obtain an access
   token, a refresh token, and the access token expiration date from the Identity Provider. These
   tokens must be stored in the [OpenIdCResponseType](../16_web-services/4939-oauthapi-openidcresponsetype.md "The OpenIdCResponseType record stores the access token, refresh token, and token expiry date retrieved in a request to the IdP.") record.
2. Next, call the `OAuthAPI.InitNativeApp()` method to initiate access to the web
   service. Pass the `OpenIdCResponseType` record to refresh the token when needed,
   allowing for seamless access without requiring a restart of the application.

Previously, you would use `OAuthAPI.InitService()` before calling any of the
OAuthAPI methods, such as `CreateHTTPAuthorizationRequest`, to perform requests to
the service. Now, you will use the `OAuthAPI.RetrievePasswordTokenForNativeApp()` and
`OAuthAPI.InitNativeApp()` methods instead.

For more details, go to [OAuthAPI.RetrievePasswordTokenForNativeApp()](../16_web-services/4954-oauthapi-retrievepasswordtokenfornativeapp.md "Returns the OAuth service access token via user credentials (username/password) and client credentials (client_id/secret_id). A refresh token allows the access token to be refreshed when it expires.") and
[OAuthAPI.InitNativeApp()](../16_web-services/4942-oauthapi-initnativeapp.md "To be called in a Genero application accessing a secure RESTful web service directly (not behind a Genero Application Server).").

## Using an identity provider's registration endpoint

Starting from FGLGWS 5.00.03, there are updates to how you manage and use the registration
endpoint of an identity provider using the OpenID Connect/OAuth2 Single sign-on protocol.

If you have an application that is protected by an identity provider via
`OpenIDConnectServiceProvider` and a registration endpoint has been set:

- You can view the identity provider's registration endpoint using the
  ImportOAuth tool's `--show` option.
- At runtime, an environment variable called OIDC\_REGISTRATION\_ENDPOINT will be available, and you
  can load it in your Genero application via
  `fgl_getenv("OIDC_REGISTRATION_ENDPOINT")`. For example, you might want to register
  or unregister an application with the identity provider. A valid access token is required to perform
  this procedure.

For more information about Single sign-on and managing registration endpoints, refer to the
ImportOAuth page in the Single Sign-On User Guide.

## New `json.Serializer` method for handling null values

Starting from FGLGWS 5.00.03, `json.Serializer` has a global option, [allowNullAsDefault](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class."), to allow `NULL`
values to be accepted during deserialization even if the JSON type is not explicitly specified. By
default, `allowNullAsDefault` is `0`.

Where previously you did not have an option to allow null values, you now can use the two new
methods `json.Serializer.setOption` and `json.Serializer.getOption`
to set options and retrieve values.

For details, go to [json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.") and [json.Serializer methods](../15_library-reference/3671-json-serializer-methods.md "Methods for the json.Serializer class.").

## Using the `json_null="null"` attribute with fglrestful

Starting from FGLGWS 5.00.03, the GWS uses [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") class for serialization and deserialization.
`json.Serializer` introduces changes to how nulls are handled that comply with
OpenAPI 3.0.

Previously, when "null" values ​​were passed in an array, record, or on one of their primitive
elements, the GWS sent a `null` value. Now a serialization error is thrown instead.
Where previously you did not have an option to allow null values, you must now explicitly allow a
null value; you must set the `json_null="null` attribute on the relevant field, or
set the [allowNullAsDefault](../15_library-reference/3685-json-serializer-options.md) option for the [The json.Serializer class](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") to allow nulls
globally.

The JSON OpenAPI documentation generated by fglrestful will represent
`json_null="null"` using the `nullable: true` property.

For conversion examples, go to [Null value conversions](../16_web-services/4849-null-value-conversions.md "These examples show how null values are handled during data type conversion.").

To learn more about the OpenAPI specification, refer to the [OpenAPI](http://swagger.io/docs/specification/about/)
documentation.

## Summary of changes in JSON streaming from GWS v4 to v5

The classes `util.JSON` and [GWS JSON API](../15_library-reference/3617-the-streaming-api-for-json-classes.md "The streaming API for JSON uses streaming to process JSON documents. These topics cover the classes for the API.") serve different purposes.

The [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") API is designed to convert program
variable values to/from JSON data, similar to JSON handling in Python or JavaScript.
`util.json` is more flexible and focuses on producing JSON output without strict
adherence to schema rules.

The [GWS JSON API](../15_library-reference/3617-the-streaming-api-for-json-classes.md "The streaming API for JSON uses streaming to process JSON documents. These topics cover the classes for the API.") is designed for streaming. Its [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") class follows the rules of JSON Schema, and
more specifically, the OpenAPI specification. Therefore, this API must ensure that the serialization
of data is compliant with these standards, and that deserialization also follows the defined rules.

Table 1 summarizes the changes to
the OpenAPI specification that may affect how you handle null values when migrating from version
FGLGWS v4.01 to version FGLGWS v5.00 of the streaming class.

| Feature | FGLGWS v4.01 | FGLGWS v5.00 | From FGLGWS 5.00.03 |
| --- | --- | --- | --- |
| Serialization/deserialization of null values | GWS used [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") for serialization `util.JSON.parse()` accepts `null` values for all types. | The GWS uses [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa."), which conforms to how nulls are handled in OpenAPI 3.0.`json.Serializer` will generate a deserialization error if "null" values ​​are passed without the `json_null="null"` attribute being explicitly set. This applies to:complex types (object and array) from v5.00.00primitive types (boolean, string, integer, and so on) from v5.00.03 | [json.Serializer.setOption](../15_library-reference/3675-json-serializer-setoption.md "Sets an option on the JSON serializer.") and [json.Serializer.getOption](../15_library-reference/3674-json-serializer-getoption.md "Returns the current value of a JSON serializer option.") are added to change the serialization/deserialization rules (making it more relaxed)[allowNullAsDefault](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.") option is added to accept `NULL` values ​​by default even when no `json_null="null"` attribute is set.fglrestful is enhanced to take `json_null` attribute and generate the "`nullable`" keyword in the OpenAPI documentation.fglrestful is enhanced to generate `json_null="null"` when detecting the `nullable : true` keyword in JSON schemas. |

Where before the GWS used [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") for
serialization, now it uses [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa."). Therefore,
you may need to review your code, setting the `json_null="null"` attribute on BDL
variables to allow null values. Alternatively, you can set the [allowNullAsDefault](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.") option to specify the default
behavior for nulls even when there is no `json_null="null"` attribute set.

## Specific exception -15582 for status code not in the range 200 - 299

Starting from FGLGWS 5.00.00 and 5.01.02, the GWS method [com.WebServiceEngine.SetRestStatus](../15_library-reference/3801-com-webserviceengine-setreststatus.md "Manages HTTP response status codes (200 - 299) for a REST high-level web service function.") raises a specific error code [-15582](../15_library-reference/4483-genero-bdl-errors.md), when the status code is not
in the range 200 - 299.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Web
services changes in BDL 4.01](0125-web-services-changes.md "There are changes in support of web services in Genero 4.01.").

Notable changes introduced in maintenance releases:

- No particular change to consider.

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
