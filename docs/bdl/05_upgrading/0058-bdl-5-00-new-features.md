---
title: "BDL 5.00 new features"
source: "fgl-topics/fgl_whatsnew_500v.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 5.00 new features"
type: "topic"
---

# BDL 5.00 new features

> Features added in 5.00 releases of the Genero Business Development Language.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 5.00 upgrade guide](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.").

Previous new features page: [BDL 4.01 new features](0059-bdl-4-01-new-features.md "Features added in 4.01 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| **Starting at 5.00.03** |  |
| New `base.Channel` method to get the exit status of the pipe child process. | See [base.Channel.getExitStatus](../15_library-reference/2987-base-channel-getexitstatus.md "Returns the exit status of the child process of a pipe channel."). |
| [***Related upgrade notes***](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.") |  |

| Overview | Reference |
| --- | --- |
| Support for front-ends of different major version number. | See [VM and front-end compatibility](0088-vm-and-front-end-compatibility.md "Always combine Genero Virtual Machine with the latest compatible Front-End.") |
| Proper action default attributes in FGLDIR/lib/default.4ad, for special actions such as `enterforeground`/`enterbackground`. | See [List of predefined actions](../11_user-interface/2257-list-of-predefined-actions.md), [Using action defaults files](../11_user-interface/1604-using-action-defaults-files.md "To use action default files, you must understand how they work and how to structure the code."). |
| Form file syntax allows now string literal for `ACCELERATOR*` attributes. | See [ACCELERATOR attribute](../11_user-interface/1754-accelerator-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action."). |
| **Starting at 5.00.03** |  |
| Declarative dialogs support parameters (`SUBDIALOG`). | See [Declarative dialogs (DIALOG - at module level)](../11_user-interface/2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs."). |
| [***Related upgrade notes***](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.") |  |

| Overview | Reference |
| --- | --- |
| **Starting at 5.00.03** |  |
| New method `UseGbcThemeVariables()` in gICAPI web components, to use GBC theme variables in your code. | See [gICAPI.UseGbcThemeVariables()](../11_user-interface/2402-gicapi-usegbcthemevariables.md "The gICAPI.UseGbcThemeVariables() function can be used to declare a set of GBC theme variables to be referenced in the CSS of a web component.") |
| [***Related upgrade notes***](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.") |  |

| Overview | Reference |
| --- | --- |
| Support for Dameng® database 8. | See [Dameng database server](../10_sql-support/1214-dameng-database-server.md). |
| Support for MariaDB 11. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md) and [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| Support for Oracle® database 23c, with support for native `BOOLEAN` and `JSON` types. | See [Oracle Database](../10_sql-support/1359-oracle-database.md), [ORACLE BOOLEAN type](0122-oracle-boolean-type.md "Support for native Oracle SQL BOOLEAN type."), [Oracle JSON data type](../10_sql-support/1384-oracle-json-data-type.md). |
| **Starting at 5.00.02** |  |
| New FGLPROFILE entry to define PostgreSQL client trace file. | See [dbi.database.dsname.pgs.trace.file](../10_sql-support/1084-postgresql-specific-fglprofile-parameters.md). |
| **Starting at 5.00.03** |  |
| Support for PostgreSQL 17. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md) and [PostgreSQL](../10_sql-support/1415-postgresql.md). |
| [***Related upgrade notes***](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.") |  |

| Overview | Reference |
| --- | --- |
| The `--web-content` option of the fglgar war command to create a war file is now mandatory. It must specify the path to the Java Web Content directory of the JGAS installation. | See [Changes to fglgar](0121-command-tools-changes.md) and [options table for fglgar war](../13_programming-tools/2526-fglgar.md). |
| The fgldbsch tool has a new `-sl` ("sloppy") option, to skip detection of invalid table or column names. | See [Command tools changes](0121-command-tools-changes.md "Modifications to consider regarding command line tools.") |
| [***Related upgrade notes***](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.") |  |

| Overview | Reference |
| --- | --- |
| The Genero Web Services JSON library (`json`) has been added. This library provides classes and methods to perform:JSON manipulation with a Streaming API for JSONSerialization of BDL variables in JSON | See[The json package](../15_library-reference/3616-the-json-package.md "The Genero Web Services JSON package provides classes and methods to process JSON documents.")[New Genero Web Services JSON library for working with JSON](0114-web-services-changes.md) |
| Serializer functions have been added to the `com.HttpRequest` and `com.HttpResponse` classes to start and end HTTP requests and responses streaming as `JSONWriter` and `JSONReader` objects. | Go to[com.HttpRequest.beginJSONRequest](../15_library-reference/3850-com-httprequest-beginjsonrequest.md "Starts a HTTP request streaming JSON.")[com.HttpRequest.endJSONRequest](../15_library-reference/3864-com-httprequest-endjsonrequest.md "Terminates a streaming HTTP request.")[com.HttpResponse.beginJSONResponse](../15_library-reference/3890-com-httpresponse-beginjsonresponse.md "Starts a HTTP response streaming JSON.")[com.HttpResponse.endJSONResponse](../15_library-reference/3892-com-httpresponse-endjsonresponse.md "Performs the HTTP request.") |
| The fglrestful tool now supports the new JSON API (`json`) and generates code with this new API. However, it is possible to use the legacy JSON API (`util.json`) using the option `--legacyJSONApi`. | See [fglrestful](../13_programming-tools/2524-fglrestful.md) |
| The `JSONOneOf` and `JSONSelector` attributes are added to support the Swagger and OpenAPI `oneOf` JSON schema property. | See[JSONOneOf](../15_library-reference/3681-jsononeof.md "Defines a record that matches exactly one of several possible JSON schemas during serialization or deserialization.")[JSONSelector](../15_library-reference/3684-jsonselector.md "Identifies the variant member of a JSONOneOf record to use during serialization or deserialization.") |
| The `JSONRequired` attribute is added to support the Swagger and OpenAPI `required` JSON schema property. | See [JSONRequired](../15_library-reference/3683-jsonrequired.md "Specify properties that are required in a JSON schema.") |
| The `JSONAdditionalProperties` attribute is added to support the Swagger and OpenAPI `AdditionalProperties` JSON schema property. | See [JSONAdditionalProperties](../15_library-reference/3678-jsonadditionalproperties.md "Allows a record to accept JSON properties not explicitly listed in its schema definition.") |
| The FGLPROFILE has a new entry, `security.global.options`, added to support legacy OpenSSL 1 options. | See [New security.global.option in FGLPROFILE to allow legacy OpenSSL 1 options](0114-web-services-changes.md) |
| The FGLPROFILE has a new entry, `security.global.verifyserver`, added to support certificate validation. | See [New security.global.verifyserver entry in FGLPROFILE to support certificate validation process](0114-web-services-changes.md) |
| New fglwsdl option `-SSLOptions` added to support legacy OpenSSL 1 options. | See [fglwsdl option -SSLOptions to support legacy OpenSSL 1 options](0114-web-services-changes.md) and [fglwsdl](../13_programming-tools/2523-fglwsdl.md). |
| Set the HTTP response status code (200 - 299) of a web service operation with `com.WebServicesEngine.SetRestStatus`. | See [com.WebServiceEngine.SetRestStatus](../15_library-reference/3801-com-webserviceengine-setreststatus.md "Manages HTTP response status codes (200 - 299) for a REST high-level web service function.") |
| A new method allows the return of a specific header in an HTTP error response. | See [WSErrorHeader](../16_web-services/4831-wserrorheader.md "Defines a custom HTTP header to include in an error response."). |
| The `WSRetCode` attribute can be set with the "2XX" value to allow setting the response status with any code from 200 to 299 dynamically at runtime, added to support the Swagger and OpenAPI specification. | See[WSRetCode](../16_web-services/4830-wsretcode.md "Sets the HTTP success status code returned by the REST function.")[Changes to how WSRetCode attribute handles return status](0114-web-services-changes.md) |
| The `WSTypeDescription` high-level REST attribute is added to support the Swagger and OpenAPI `description` JSON schema property. | See [WSTypeDescription](../16_web-services/4810-wstypedescription.md "Provides a description for a user‑defined type in the REST service.") |
| The `WSDescription` attribute can be used on JSON schemas in addition to REST function input and output parameters, added to support the Swagger and OpenAPI specification. | See[WSDescription](../16_web-services/4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSDescription attribute can be used on JSON schemas](0114-web-services-changes.md) |
| The `WSPreProcesssing` and `WSPostProcessing` attributes are added to support preprocessing and postprocessing for REST callback functions. | See[Using preprocessing and postprocessing callbacks](../16_web-services/4751-using-preprocessing-and-postprocessing-callbacks.md "Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.")[WSPreProcessing](../16_web-services/4811-wspreprocessing.md "Specify functions in a REST web service module that are used as preprocessing request handlers.")[WSPostProcessing](../16_web-services/4812-wspostprocessing.md "Specify functions in a REST web service module that are used as postprocessing request handlers.") |
| The REST engine `WSContext` variable has three new entries: `RequestPATH`, `RequestVERB`, and `RequestOPERATION` to support preprocessing and postprocessing for REST callback functions. | See [WSContext](../16_web-services/4806-wscontext.md) |
| `STRING`-typed return parameter in OAUTH API function. | See [Change to OAuthAPI.GetIDSubject returns](0155-web-services-changes.md). |
| The OAuthAPI has a new method called `OAuthAPI.GetIdRoles()` to explicitly retrieve authorization roles from ID tokens. | See [New method GetIdRoles() for retrieving authorization roles](0114-web-services-changes.md) and [OAuthAPI.GetIdRoles()](../16_web-services/4945-oauthapi-getidroles.md "Get OAuth ID Token authorization roles."). |
| The Genero OpenIDConnect service now decodes ID tokens containing roles instead of scopes, and creates a new environment variable called `OIDC_ROLES` containing the list of roles. | See [OpenIDConnect service supports OIDC\_ROLES](0114-web-services-changes.md) |
| **Starting at 5.00.02** |  |
| The FGLPROFILE has a new entry, `security.global.certificate.selfsigned.preload`, added to preload the global certificate and key at the start of the application instead of at the first HTTPS connection. | See [New security.global.certificate.selfsigned.preload entry in FGLPROFILE](0114-web-services-changes.md) |
| The fglrestful tool now supports adding a prefix in front of named types in a JSON schema when generating the GWS client stub. The `--typePrefix` option is added. | See [fglrestful](../13_programming-tools/2524-fglrestful.md) |
| The GWS `GeneroAccessService` service has enhancements to how scopes are managed. The Genero Identity Provider can now secure a web service from scopes set in the configuration file (xcf). | See [The GeneroAccessService supports scopes set in configuration file](0114-web-services-changes.md) |
| **Starting at 5.00.03** |  |
| New fglwsdl option `-oauth` added to support OAuth. | See [fglwsdl oauth option added](0114-web-services-changes.md) |
| The OAuthAPI has a two new method to use with Single sign-on authentication for mobile apps accessing web services:`OAuthAPI.RetrievePasswordTokenForNativeApp()``OAuthAPI.InitNativeApp()` | See [New OAuthAPI methods to use with SSO authentication for apps accessing web services](0114-web-services-changes.md) |
| The ImportOAuth tool, which registers endpoints provided by an identity provider, has an update to its `--show` option to display the identity provider's registration endpoint. | See [Using an identity provider's registration endpoint](0114-web-services-changes.md) |
| The `json.Serializer.setOption()` and `json.Serializer.getOption()` methods are added to support the global configuration of JSON serializer options such as `allowNullAsDefault`. | See:[json.Serializer.getOption](../15_library-reference/3674-json-serializer-getoption.md "Returns the current value of a JSON serializer option.")[json.Serializer.setOption](../15_library-reference/3675-json-serializer-setoption.md "Sets an option on the JSON serializer.") |
| The `allowNullAsDefault` option is added to `json.Serializer` to support null values during deserialization even if the JSON type is not explicitly specified. By default `allowNullAsDefault` is `FALSE`. You can set this option with the `json.Serializer.setOption()` method. | See:[New json.Serializer method for handling null values](0114-web-services-changes.md)[allowNullAsDefault](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.") |
| The fglrestful tool has enhancements:It detects the `nullable: true` keyword set for OpenAPI 3.0 and generates `json_null="null"` attribute in BDL variables to allow nulls.It generates the `nullable: true` keyword in the OpenAPI document when GWS detects `json_null="null"` on a BDL type. | See [Using the json\_null="null" attribute with fglrestful](0114-web-services-changes.md) |
| [***Related upgrade notes***](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.") |  |
| Upgrade notes for web services. | See [Web Services changes](0114-web-services-changes.md "There are changes in support of web services in Genero 5.00.") |

| Overview | Reference |
| --- | --- |
| New front calls introduced in V4.01 front end maintenance releases. | See [New front calls](0128-front-calls-changes.md). |
| **Starting at GMA-5.00.02** |  |
| New gmabuildtool options `--custom-foreground-service-type` and `--background-explanation`. | See [Specifying the Android foreground service type](0119-genero-mobile-for-android-gma-5-00-changes.md). |
| [***Related upgrade notes***](0112-bdl-5-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.00.") |  |
| Upgrade notes for GMA. | See [Genero Mobile for Android (GMA) 5.00 changes](0119-genero-mobile-for-android-gma-5-00-changes.md "Modifications to consider when using the Genero Mobile for Android."). |
| Upgrade notes for GMI. | See [Genero Mobile for iOS (GMI) 5.00 changes](0120-genero-mobile-for-ios-gmi-5-00-changes.md "Modifications to consider when using Genero Mobile for iOS."). |
