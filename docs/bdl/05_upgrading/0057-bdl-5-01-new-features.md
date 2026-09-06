---
title: "BDL 5.01 new features"
source: "fgl-topics/fgl_whatsnew_501.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 5.01 new features"
type: "topic"
---

# BDL 5.01 new features

> Features added in 5.01 releases of the Genero Business Development Language.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 5.01 upgrade guide](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.").

Previous new features page: [BDL 5.00 new features](0058-bdl-5-00-new-features.md "Features added in 5.00 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| Support for GWA (Genero Web Applications) | See [Genero Web applications](../18_genero-web-applications/5123-genero-web-applications.md "These topics cover programming subjects about Genero Web applications"). |
| **Starting at 5.01.02** |  |
| UTF-8 locale support is now set on Windows® with `LANG=.utf8` (`LANG=.fglutf8` is still supported for backward compatibility) | See [Language and character set settings](../09_advanced-features/0880-language-and-character-set-settings.md). |
| **Starting at 5.01.03** |  |
| New `base.SqlHandle` methods to open the SQL cursor with hold option. | See [base.SqlHandle.openCursorWithHold](../15_library-reference/3032-base-sqlhandle-opencursorwithhold.md "Opens the SQL handle with holdable option.") and [base.SqlHandle.openScrollCursorWithHold](../15_library-reference/3034-base-sqlhandle-openscrollcursorwithhold.md "Opens the SQL handle with scrollable and holdable option."). |
| **Starting at 5.01.06** |  |
| Server TCP socket APIs and tools support hostname in addition to IPv4 addresses. | See [base.Channel.openServerSocket](../15_library-reference/2992-base-channel-openserversocket.md "Open a TCP server socket channel.") and [Syntax 3: Starting the debug-server](../13_programming-tools/2513-fglrun.md). |
| **Starting at 5.01.07** |  |
| String to `INTERVAL` and `util.JSON` parser accepts ISO 8601 formatted durations. | See [Character string to date time types](../08_language-basics/0578-data-type-conversion-reference.md), [JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.") |
| `util.JSON` serialization of `DATETIME` as RFC 3339. | See [util.JSON.setDatetimeSerializationMode](../15_library-reference/3584-util-json-setdatetimeserializationmode.md "Defines the JSON formatting mode for DATETIME values."). |
| `util.JSON` serialization of `INTERVAL` as ISO 8601. | See [util.JSON.setIntervalSerializationMode](../15_library-reference/3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values."). |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |

| Overview | Reference |
| --- | --- |
| **Starting at 5.01.06** |  |
| Define the initial window container size with GDC. | See [ui.Interface.setSize](../15_library-reference/3121-ui-interface-setsize.md "Specify the initial size of the window container."), [Initial window size (GDC/desktop)](../11_user-interface/1563-containers-for-program-windows.md). |
| **Starting at 5.01.07** |  |
| `CONSTRUCT` allows combination of `\|` (pipe) with `*` and `?` for multiple pattern-based character type search. | See [Query operators in CONSTRUCT](../11_user-interface/2052-query-operators-in-construct.md). |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |

| Overview | Reference |
| --- | --- |
| New style attribute `headerPosition` for table column titles. | See [headerPosition attribute](0105-presentation-styles-changes.md) |
| New style attribute `showVirtualKeyboard` on mobile. | See [showVirtualKeyboard attribute](0105-presentation-styles-changes.md) |
| **Starting at 5.01.03** |  |
| New front calls to control table column layout. | See [New front calls](0106-front-calls-changes.md) |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |

| Overview | Reference |
| --- | --- |
| Support for Oracle® MySQL 8.4 with the `dbmmys_8_4` ODI driver. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md) and [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| Support for Informix® 15.0, using CSDK 4.50 with the `dbmifx_9` ODI driver. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| **Starting at 5.01.02** |  |
| Support for Oracle 23ai and the `VECTOR` data type. | See [Oracle VECTOR data type](../10_sql-support/1385-oracle-vector-data-type.md). |
| Support for PostgreSQL `JSON` and `JSONB` data types. | See [PostgreSQL JSON/JSONB data types](../10_sql-support/1439-postgresql-json-jsonb-data-types.md). |
| **Starting at 5.01.06** |  |
| Support for PostgreSQL 18. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [PostgreSQL](../10_sql-support/1415-postgresql.md). |
| Support for MariaDB 12. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| Support for SQL Server 2025. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md) |
| Support for SQL Server 2025 `JSON` data type. | See [SQL Server JSON data type](../10_sql-support/1285-sql-server-json-data-type.md). |
| Support for SQL Server 2025 `VECTOR` data type. | See [SQL Server VECTOR data type](../10_sql-support/1286-sql-server-vector-data-type.md) |
| **Starting at 5.01.07** |  |
| Support for MySQL 9.7 LTS. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md), [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| Support for MySQL 9 `VECTOR` data type. | See [MySQL/MariaDB VECTOR data type](../10_sql-support/1336-mysql-mariadb-vector-data-type.md). |
| PostgreSQL `TSVECTOR` and `TSQUERY` types for full text search. | See [PostgreSQL TSVECTOR/TSQUERY data types](../10_sql-support/1440-postgresql-tsvector-tsquery-data-types.md). |
| PostgreSQL `pgvector` extension support for vector similarity search. | See [PostgreSQL pgvector extension data types](../10_sql-support/1441-postgresql-pgvector-extension-data-types.md). |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |

| Overview | Reference |
| --- | --- |
| The fglgar tool has a `gwa` verb for packaging Genero Web Applications in a Genero archive (gwa) file. | See [options table for fglgar gwa](../13_programming-tools/2526-fglgar.md) and [Packaging gwa files](../13_programming-tools/2650-packaging-gwa-files.md "Using the fglgar tool to build a Genero Web Application (gwa) file allows you to deploy applications that are ready to run in a browser."). |
| VS Code extension uses now the Language Server technology. This implies fast responses on VS Code front-end requests, go to references is aware of all files in the workspace, diagnostic is now sensitive for inter-module dependencies. | See [Visual Studio Code extension](../13_programming-tools/2545-visual-studio-code-extension.md), [Debugging with VS Code](../13_programming-tools/2584-debugging-with-vs-code.md "The Genero BDL VS Code extension can be used to debug programs with VS Code debugger capabilities."). |
| **Starting at 5.01.06** |  |
| New VS Code extension version 0.0.28. Install the new extension file genero-fgl-0.0.28.vsix into VS Code. | See [Installing Genero BDL extension into VS Code](../13_programming-tools/2545-visual-studio-code-extension.md). |
| Debugger-server (debuggee proxy) for VS Code extension. | See [Using the debug-server](../13_programming-tools/2585-using-the-debug-server.md "The Genero BDL debug-server is a proxy for fglrun processes which can not be accessed directly by the debugger."), [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs."). |
| No-configuration debugging with VS Code extension. | See [No-configuration debugging](../13_programming-tools/2584-debugging-with-vs-code.md). |
| Automatic TCP port selection with debug-server. | See [Automatic port allocation (--da-listen 0)](../13_programming-tools/2585-using-the-debug-server.md). |
| New fglcomp -W keywords warning option. | See [fglcomp -W keywords warning option](0111-command-tools-changes.md) |
| New fglcomp -W shadow warning option. | See [fglcomp -W shadow warning option](0111-command-tools-changes.md) |
| VS Code extension: Parameter name inlay hints. | See [Visual Studio Code extension](../13_programming-tools/2545-visual-studio-code-extension.md). |
| **Starting at 5.01.07** |  |
| New VS Code extension version 0.0.28. Install the new extension file genero-fgl-0.0.29.vsix into VS Code. | See [Installing Genero BDL extension into VS Code](../13_programming-tools/2545-visual-studio-code-extension.md). |
|  |  |
| VS Code extension supports Rename Symbol (F2), to rename a variable, function in current file and all files referencing this symbol. | See [Visual Studio Code extension](../13_programming-tools/2545-visual-studio-code-extension.md), [Debugging with VS Code](../13_programming-tools/2584-debugging-with-vs-code.md "The Genero BDL VS Code extension can be used to debug programs with VS Code debugger capabilities."). |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |

| Overview | Reference |
| --- | --- |
| There are new options for handling `NULL` values for primitive types during serialization. | See [The handling of NULL values for primitive types during serialization has changed](0103-web-services-changes.md) |
| [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") class for serialization and deserialization introduces changes to how nulls are handled with `JSONRequired`. [error-15807](../15_library-reference/4483-genero-bdl-errors.md) is raised when a `null` value is sent to a variable with `JSONRequired` unless you also use `json_null="null"`. | See [Using the json\_null="null" attribute with JSONRequired](0103-web-services-changes.md) |
| The `serializeNullAsDefault` option is added to `json.Serializer` to support null values during BDL/JSON serialization. By default `serializeNullAsDefault` is `0`. You can set this option with the `json.Serializer.setOption()` method. | See:[serializeNullAsDefault](../15_library-reference/3685-json-serializer-options.md)[The handling of NULL values for primitive types during serialization has changed](0103-web-services-changes.md) |
| New methods allow you to set HTTPS client certificates at runtime and dynamically override some FGLPROFILE entries for secured communication. | See [New HttpRequest methods dynamically override FGLPROFILE entries](0103-web-services-changes.md) |
| FGLPROFILE has a new entry, `xml.keystore.cadir`, to specify a directory for all trusted X509 certificates to be used during XML signature verification. | See [New xml.keystore.cadir entry in FGLPROFILE for XML signature verification](0103-web-services-changes.md) |
| **Starting at 5.01.01** |  |
| The `getLastErrorDescription()` method is introduced in the `json.Serializer` class. This method retrieves the most recent error generated during serialization. | See[New json.Serializer.getLastErrorDescription() method has improved error messaging](0103-web-services-changes.md)[json.Serializer.getLastErrorDescription](../15_library-reference/3676-json-serializer-getlasterrordescription.md "Retrieves the most recent error generated during serialization.") |
| **Starting at 5.01.02** |  |
| The `com.WebServiceEngine.RegisterRestService()` and `com.WebServiceEngine.RegisterRestResources()` methods support package-style module names. | See:[Register web services and resources from a package](0103-web-services-changes.md)[com.WebServiceEngine.RegisterRestService](../15_library-reference/3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine.") or [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine.") |
| **Starting at 5.01.03** |  |
| [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") class introduces a new option, `allowImplicitConversion`, to support implicit conversion in JSON deserialization. | See[New json.Serializer option for handling implicit conversion](0103-web-services-changes.md)[json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.") |
| **Starting at 5.01.06** |  |
| The function `com.WebServiceEngine.HandleRequest` now restores the previous behavior for query‑only requests targeting the REST service root path, allowing them to fall back to low‑level handling when no REST operation matches. | See [Restored HandleRequest() behavior for root-level query-only requests](0125-web-services-changes.md) and [com.WebServiceEngine.HandleRequest](../15_library-reference/3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.") |
| The fglrestful tool now recognizes all `+json` media types in OpenAPI content sections (RFC 6839). | See [fglrestful now recognizes +json media types (RFC 6839)](0103-web-services-changes.md) |
| **Starting at 5.01.07** |  |
| A `CHAR(n)` field exposed by a REST service now outputs `maxLength: n` in the generated OpenAPI specification, like `VARCHAR(n)`. | See [OpenAPI types mapping: BDL](../16_web-services/4847-openapi-types-mapping-bdl.md "Conversion mapping for Genero BDL data types in OpenAPI documentation.") and [CHAR(n) fields now include maxLength in the OpenAPI schema](0103-web-services-changes.md). |
| The fglrestful tool now detects the `text/event-stream` (Server-Sent Events) response content type and excludes it from response-body type selection and from the generated `Accept` header. | See [fglrestful now detects text/event-stream (SSE) responses](0103-web-services-changes.md) |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |
| Upgrade notes for web services. | See [Web Services changes](0103-web-services-changes.md "There are changes in support of web services in Genero 5.01.") |

| Overview | Reference |
| --- | --- |
| **Starting at 5.01.01** |  |
| Extension of the `mobile.scanBarCode` front call to specify the ouput format. | See [mobile.scanBarCode output format specification (GMA)](0106-front-calls-changes.md). |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |
| Upgrade notes for GMA. | See [Genero Mobile for Android (GMA) 5.01 changes](0108-genero-mobile-for-android-gma-5-01-changes.md "Consider these modifications when you use Genero Mobile for Android."). |
| Upgrade notes for GMI. | See [Genero Mobile for iOS (GMI) 5.01 changes](0109-genero-mobile-for-ios-gmi-5-01-changes.md "Modifications to consider when using Genero Mobile for iOS."). |

| Overview | Reference |
| --- | --- |
| [***Related upgrade notes***](0101-bdl-5-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 5.01.") |  |
| **Starting at 5.01.02** |  |
| New gwabuildtool option `--app-dir-is-pwd`. | See [Specifying the working directory as app](0110-genero-for-web-applications-gwa-5-01-changes.md) |
| Upgrade notes for GWA. | See [Genero for Web Applications (GWA) 5.01 changes](0110-genero-for-web-applications-gwa-5-01-changes.md "Modifications to consider when using Genero for Web Applications."). |
