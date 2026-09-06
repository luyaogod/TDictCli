---
title: "BDL 6.00 new features"
source: "fgl-topics/fgl_whatsnew_600v.html"
breadcrumb: "What's new in 6.00"
type: "topic"
---

# BDL 6.00 new features

> Features added in 6.00 releases of the Genero Business Development Language.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 6.00 upgrade guide](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.").

Previous new features page: [BDL 5.01 new features](../05_upgrading/0057-bdl-5-01-new-features.md "Features added in 5.01 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| The `sqlca.sqlerrd[1-6]` base type is now `BIGINT` for Informix 15 support. | See [sqlca.sqlerrd[1-6] as BIGINT](../05_upgrading/0091-core-language-changes.md). |
| Server TCP socket APIs and tools support hostname in addition to IPv4 addresses. | See [base.Channel.openServerSocket](../15_library-reference/2992-base-channel-openserversocket.md "Open a TCP server socket channel.") and [Syntax 3: Starting the debug-server](../13_programming-tools/2513-fglrun.md). |
| **Starting at 6.00.03** |  |
| String to `INTERVAL` and `util.JSON` parser accepts ISO 8601 formatted durations. | See [Character string to date time types](../08_language-basics/0578-data-type-conversion-reference.md), [JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.") |
| `util.JSON` serialization of `DATETIME` as RFC 3339. | See [util.JSON.setDatetimeSerializationMode](../15_library-reference/3584-util-json-setdatetimeserializationmode.md "Defines the JSON formatting mode for DATETIME values."). |
| `util.JSON` serialization of `INTERVAL` as ISO 8601. | See [util.JSON.setIntervalSerializationMode](../15_library-reference/3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values."). |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |

| Overview | Reference |
| --- | --- |
| Multiple column sorting in list dialogs, with new APIs to get sort specification. | See [Sorting rows in a list](../11_user-interface/2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required.").New `ui.Dialog` methods:[`addGroupBy()`](../15_library-reference/3181-ui-dialog-addgroupby.md "Appends a grouping column of a list dialog.")[`setGroupByDesc()`](../15_library-reference/3230-ui-dialog-setgroupbydesc.md "Defines sort order for a grouping column of a list dialog.")[`getSortKeyAt()`](../15_library-reference/3203-ui-dialog-getsortkeyat.md "Returns the name of field used as grouping or sorting column, for a given sort column position.")[`isSortKeyReverseAt()`](../15_library-reference/3208-ui-dialog-issortreverseat.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending), for a given sort column position.") |
| New method to reset the sort specification in list dialogs. | See [`resetSort()`](../15_library-reference/3211-ui-dialog-resetsort.md "Resets the sort columns of a list dialog."). |
| **Starting at 6.00.02** |  |
| Define the initial window container size with GDC. | See [ui.Interface.setSize](../15_library-reference/3121-ui-interface-setsize.md "Specify the initial size of the window container."), [Initial window size (GDC/desktop)](../11_user-interface/1563-containers-for-program-windows.md). |
| **Starting at 6.00.03** |  |
| `CONSTRUCT` allows combination of `\|` (pipe) with `*` and `?` for multiple pattern-based character type search. | See [Query operators in CONSTRUCT](../11_user-interface/2052-query-operators-in-construct.md). |
| Simplified dialog error message translation with .str localized strings. | See [Runtime system messages](../09_advanced-features/0888-runtime-system-messages.md "This section describes how to translate default English runtime system message files in a different language."). |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |

| Overview | Reference |
| --- | --- |
| New UX for window containers (SideBarRail, SideBarDrawer, Application List, etc) | See [Containers for program windows](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."). |
| Multiple column sorting with ALT-CLICK in table column headers. | See [Sorting rows in a list](../11_user-interface/2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required."). |
| Back button in chromebar, when an action with the name `back` is defined and active. | See [Implementing the back action](../11_user-interface/2290-implementing-the-back-action.md "The back action is a predefined action dedicated to move back in the stack of windows/forms."). |
| The chromebar is now hidden by default on GDC/UR desktop. It can be enabled with a GBC theme variable. | See [Containers for program windows](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.") and GBC documentation. |
| Phone number edit field with the new `customWidget` `phoneEdit` style attribute. | See [New customWidget value phoneEdit](../05_upgrading/0094-presentation-styles-changes.md), [Phone number fields](../11_user-interface/2249-phone-number-fields.md "Phone number fields allow to display and input formatted telephone numbers, with international call prefix selection."). |
| Tag edit field with the new `customWidget` `tagEdit` style attribute. | See [New customWidget value tagEdit](../05_upgrading/0094-presentation-styles-changes.md), [Multi-valued fields](../11_user-interface/2248-multi-valued-fields.md "Multi-valued form fields allow to display and input a set of values in the same character string field."). |
| Front calls to control the browser URL `#` anchors, with new `applicationstatechanged` special action. | See [Controlling web application state (#anchor)](../11_user-interface/2229-controlling-web-application-state-anchor.md "With GBC in a web browser, the context/state of a program can be managed with URL # anchors."), [Browser front calls](../15_library-reference/3432-browser-front-calls.md "This section describes browser handling front calls."), [Front calls changes](../05_upgrading/0095-front-calls-changes.md "Modifications to consider when using front calls."). |
| Front calls to manage local notifications on non-mobile front end hosts. | See [standard.createNotification](../15_library-reference/3397-standard-createnotification.md "Creates a new local notification to be displayed on the host system of the front end."), [Front calls changes](../05_upgrading/0095-front-calls-changes.md "Modifications to consider when using front calls."). |
| Presentation style attribute for rounder border on images: `imageBorderRadius`. | See [New imageBorderRadius style attribute](../05_upgrading/0094-presentation-styles-changes.md) |
| The `applicationListPosition` style attribute for `UserInterface`, to display the application list in the ChromeBar. | See [New applicationListPosition style attribute](../05_upgrading/0094-presentation-styles-changes.md). |
| AI writing assistant with new `aiWritingAssistant` style attribute for `TEXTEDIT` fields. | See [New aiWritingAssistant style attribute](../05_upgrading/0094-presentation-styles-changes.md). |
| New `customWidget` style attribute for `COMBOBOX`. | See [New customWidget style attribute for COMBOBOX](../05_upgrading/0094-presentation-styles-changes.md) |
| `DATEEDIT` and `DATETIMEEDIT` field support `showCurrentMonthOnly` style attribute. | See [showCurrentMonthOnly for DATEEDIT/DATETIMEEDIT](../05_upgrading/0094-presentation-styles-changes.md) |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |

| Overview | Reference |
| --- | --- |
| Support for Informix® 15.0, using CSDK 15 with the `dbmifx_15` ODI driver. | See [New ODI driver for Informix CSDK 15](../05_upgrading/0093-database-drivers-changes.md). |
| Support for PostgreSQL 18. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [PostgreSQL](../10_sql-support/1415-postgresql.md). |
| **Starting at 6.00.02** |  |
| Support for MariaDB 12. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| Support for SQL Server 2025. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md) |
| Support for SQL Server 2025 `JSON` data type. | See [SQL Server JSON data type](../10_sql-support/1285-sql-server-json-data-type.md). |
| Support for SQL Server 2025 `VECTOR` data type. | See [SQL Server VECTOR data type](../10_sql-support/1286-sql-server-vector-data-type.md) |
| **Starting at 6.00.03** |  |
| Support for MySQL 9.7 LTS. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| Support for MySQL 9 `VECTOR` data type. | See [MySQL/MariaDB VECTOR data type](../10_sql-support/1336-mysql-mariadb-vector-data-type.md). |
| PostgreSQL `TSVECTOR` and `TSQUERY` types for full text search. | See [PostgreSQL TSVECTOR/TSQUERY data types](../10_sql-support/1440-postgresql-tsvector-tsquery-data-types.md). |
| PostgreSQL `pgvector` extension support for vector similarity search. | See [PostgreSQL pgvector extension data types](../10_sql-support/1441-postgresql-pgvector-extension-data-types.md). |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |

| Overview | Reference |
| --- | --- |
| The Prometheus library has been added. This library provides classes and methods for collecting and storing metrics for integration with a monitoring platform such as Prometheus. | See [The prometheus package](../15_library-reference/4462-the-prometheus-package.md "These topics cover the classes for the prometheus package."). |
| Debugger-server (debuggee proxy) for VS Code extension. | See [Using the debug-server](../13_programming-tools/2585-using-the-debug-server.md "The Genero BDL debug-server is a proxy for fglrun processes which can not be accessed directly by the debugger."), [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs."). |
| New fglcomp -W keywords warning option. | See [fglcomp -W keywords warning option](../05_upgrading/0111-command-tools-changes.md) |
| **Starting at 6.00.02** |  |
| New VS Code extension version 0.0.28. Install the new extension file genero-fgl-0.0.28.vsix into VS Code. | See [Installing Genero BDL extension into VS Code](../13_programming-tools/2545-visual-studio-code-extension.md). |
| No-configuration debugging with VS Code extension. | See [No-configuration debugging](../13_programming-tools/2584-debugging-with-vs-code.md). |
| Automatic TCP port selection with debug-server. | See [Automatic port allocation (--da-listen 0)](../13_programming-tools/2585-using-the-debug-server.md). |
| New fglcomp -W shadow warning option. | See [fglcomp -W shadow warning option](../05_upgrading/0111-command-tools-changes.md) |
| VS Code extension: Parameter name inlay hints. | See [Visual Studio Code extension](../13_programming-tools/2545-visual-studio-code-extension.md). |
| **Starting at 6.00.03** |  |
| `[DESCRIBE] DATABASE` schema specification detected by `fglcomp -W stdsql`. | See [[DESCRIBE] DATABASE detected by -W stdsql](../05_upgrading/0100-command-tools-changes.md). |
| New `fglcomp -W fragile-cursor` warning option to detect risky SQL cursor usage. | See [Compiler warning for risky SQL cursor usage](../05_upgrading/0100-command-tools-changes.md). |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |

| Overview | Reference |
| --- | --- |
| The `JSONAllOf` attribute defines a type to merge the properties of several record types. | See [JSONAllOf](../15_library-reference/3679-jsonallof.md "Combine multiple record schemas into a single type using the JSONAllOf attribute in Genero BDL.") |
| The `JSONEnum` defines a typed list of acceptable values in a field. It maps to the `enum` keyword in the JSON Schema specification. | See [JSONEnum](../15_library-reference/3680-jsonenum.md "Defines an explicit, typed list of acceptable values for a field and maps to the enum keyword in JSON Schema.") |
| The `JSONPattern` attribute defines regular expressions that string values must match. | See [JSONPattern](../15_library-reference/3682-jsonpattern.md "Specify a regular expression pattern for string values in a JSON schema.") |
| The `--no-merge-allof` option of `fglrestful` disables the merging of properties in schemas where the `JSONAllOf` attribute is applied. | See [fglrestful](../13_programming-tools/2524-fglrestful.md) and [JSONAllOf](../15_library-reference/3679-jsonallof.md "Combine multiple record schemas into a single type using the JSONAllOf attribute in Genero BDL.") |
| The `xml.CryptoKey` class now supports Elliptic Curve Digital Signature Algorithm (ECDSA) keys for creating digital signatures. | See [Supported kind of keys](../15_library-reference/4222-supported-kind-of-keys.md) |
| New methods in `xml.CryptoKey` for ECDSA key management:`getEllipticCurveName()``generateEllipticCurveKey()``loadEllipticCurve()` | See:[xml.CryptoKey.getEllipticCurveName](../15_library-reference/4201-xml-cryptokey-getellipticcurvename.md "Returns the name of the elliptic curve for an ECDSA key, if available.")[xml.CryptoKey.generateEllipticCurveKey](../15_library-reference/4199-xml-cryptokey-generateellipticcurvekey.md "Generates a new ECDSA key for the specified named elliptic curve.")[xml.CryptoKey.loadEllipticCurve](../15_library-reference/4209-xml-cryptokey-loadellipticcurve.md "Loads an ECDSA key using a specific public point.") |
| **Starting at 6.00.01** |  |
| The `WSTags` attribute maps to the `tags` keyword in the OpenAPI specification and categorizes and groups REST operations. | See [WSTags](../16_web-services/4826-wstags.md "Assign one or more tags to a REST operation. The value corresponds to the \"tags\" keyword in the OpenAPI specification and is used for grouping and documentation.") |
| The `WSSummary` attribute maps to the `summary` keyword in the OpenAPI specification and provides a short, human-readable summary for a REST operation. | See [WSSummary](../16_web-services/4843-wssummary.md "Specifies a short, human-readable summary of a REST operation. It corresponds to the summary keyword in the OpenAPI specification.") |
| REST functions now require a `RETURNS()` clause; use an empty `RETURNS()` to indicate no return value (previously `RETURNS` clauses were optional). | See [REST functions must include a RETURNS() clause](../05_upgrading/0092-web-services-changes.md) |
| **Starting at 6.00.02** |  |
| `WSTags` values are now collected and generated once in the OpenAPI document’s global `tags` section. | See [Global tags generation](../05_upgrading/0092-web-services-changes.md). |
| Retrieve all available versions of the service via the `?version.json` / `?version.yaml` query parameter. | See [Get available versions of a service](../16_web-services/4778-get-available-versions-of-a-service.md "Retrieve the list of available versions of a REST service and identify the default version if defined.") |
| The function `com.WebServiceEngine.HandleRequest` now restores the previous behavior for query‑only requests targeting the REST service root path, allowing them to fall back to low‑level handling when no REST operation matches. | See [Restored HandleRequest() behavior for root-level query-only requests](../05_upgrading/0125-web-services-changes.md) and [com.WebServiceEngine.HandleRequest](../15_library-reference/3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.") |
| The `WSOperationId` attribute maps to the `operationId` keyword in the OpenAPI specification and defines a unique identifier for a REST operation. | See [WSOperationId](../16_web-services/4827-wsoperationid.md "Specifies a unique string to identify the operation. It corresponds to the operationId keyword in the OpenAPI specification.") |
| The new `RegisterRestOpenAPIHandler()` method registers a callback function that is invoked when a client requests the OpenAPI documentation for a REST service, allowing dynamic modification of the OpenAPI JSON before it is returned to the client. | See [com.WebServiceEngine.RegisterRestOpenAPIHandler](../15_library-reference/3794-com-webserviceengine-registerrestopenapihandler.md "Registers a callback function that dynamically modifies OpenAPI documentation for a REST service.") |
| A new context key, "`HTTPProcessingErrorMsg`", lets you provide a custom HTTP error message in WSPreprocessing and WSPostProcessing callbacks. | See [Enhanced control of HTTP error messages in preprocessing and postprocessing](../05_upgrading/0092-web-services-changes.md) |
| **Starting at 6.00.03** |  |
| The fglrestful tool now recognizes all `+json` media types in OpenAPI content sections (RFC 6839). | See [fglrestful now recognizes +json media types (RFC 6839)](../05_upgrading/0103-web-services-changes.md) |
| A `CHAR(n)` field exposed by a REST service now outputs `maxLength: n` in the generated OpenAPI specification, like `VARCHAR(n)`. | See [OpenAPI types mapping: BDL](../16_web-services/4847-openapi-types-mapping-bdl.md "Conversion mapping for Genero BDL data types in OpenAPI documentation.") and [CHAR(n) fields now include maxLength in the OpenAPI schema](../05_upgrading/0103-web-services-changes.md). |
| The fglrestful tool now detects the `text/event-stream` (Server-Sent Events) response content type and excludes it from response-body type selection and from the generated `Accept` header. | See [fglrestful now detects text/event-stream (SSE) responses](../05_upgrading/0103-web-services-changes.md) |
| The `xml.CryptoKey` class now supports the `sha256-rsa-MGF1` key type for RSA signatures requiring SHA-256 and RSA-PSS padding. | See [Supported kind of keys](../15_library-reference/4222-supported-kind-of-keys.md). |
| `json.Serializer` and `json.JSONWriter` serialization of `DATETIME` as RFC 3339. | See [datetimeSerializationMode](../15_library-reference/3688-datetimeserializationmode.md "Controls the JSON output format for DATETIME values during serialization."). |
| `json.Serializer` and `json.JSONWriter` serialization of `INTERVAL` as ISO 8601. | See [intervalSerializationMode](../15_library-reference/3689-intervalserializationmode.md "Controls the JSON output format for INTERVAL values during serialization."). |
| fglrestful generates `INTERVAL` variables for OpenAPI duration properties (`format: duration`); the new `--interval` option overrides the qualifier or restores the previous `STRING` mapping. | See [Date-time and duration options](../13_programming-tools/2524-fglrestful.md). |
| fglrestful generates client stubs that can send `DATETIME` in RFC 3339 and `INTERVAL` as ISO 8601 durations in JSON request bodies, with the new `--datetime-interval-format` option. | See [fglrestful options](../13_programming-tools/2524-fglrestful.md). |
| The GWS generates a qualifier-aware `format` keyword for `DATETIME` and `INTERVAL` in the OpenAPI document. | See [OpenAPI types mapping: BDL](../16_web-services/4847-openapi-types-mapping-bdl.md). |
| The `xml.CryptoKey` class now supports the DSA-SHA256 signature key (`http://www.w3.org/2009/xmldsig11#dsa-sha256`). SHA-1 is deprecated. | See [Supported kind of keys](../15_library-reference/4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class."). |
| High-level REST services are now listed in an RFC 9727 API catalog at `GET /.well-known/api-catalog`, with a link to each service's OpenAPI description. | See [REST API catalog](../16_web-services/4758-rest-api-catalog.md "A Genero application that publishes REST services automatically provides an RFC 9727 API catalog at /.well-known/api-catalog, listing each service and its OpenAPI description."). |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |
| Upgrade notes for web services. | See [Web Services changes](../05_upgrading/0092-web-services-changes.md "There are changes in support of web services in Genero 6.00.") |

| Overview | Reference |
| --- | --- |
| The gmibuildtool option `--deploy` can be used to install an app .ipa file on a device. | See [Installing the app on a device](../17_mobile-applications/5109-building-ios-apps-with-genero.md). |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |
| Upgrade notes for GMA. | See [Genero Mobile for Android (GMA) 6.00 changes](../05_upgrading/0097-genero-mobile-for-android-gma-6-00-changes.md "Modifications to consider when using the Genero Mobile for Android."). |
| Upgrade notes for GMI. | See [Genero Mobile for iOS (GMI) 6.00 changes](../05_upgrading/0098-genero-mobile-for-ios-gmi-6-00-changes.md "Modifications to consider when using Genero Mobile for iOS."). |

| Overview | Reference |
| --- | --- |
| Internationalize your application in different languages with localized strings. | See [Internationalize your app](../18_genero-web-applications/5135-internationalize-your-app.md "GWA applications can be translated and modified for international markets."). |
| The GWA\_TARGET environment variable defines the target platform or environment for the GWA app. | See [GWA\_TARGET](../18_genero-web-applications/5155-gwa-environment-variables.md). |
| [***Related upgrade notes***](../05_upgrading/0089-bdl-6-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 6.00.") |  |
| Upgrade notes for GWA. | See [Genero for Web Applications (GWA) 6.00 changes](../05_upgrading/0099-genero-for-web-applications-gwa-6-00-changes.md "Modifications to consider when using Genero for Web Applications.") |
