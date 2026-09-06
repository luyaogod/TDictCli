---
title: "BDL 3.20 new features"
source: "fgl-topics/fgl_whatsnew_320.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 3.20 new features"
type: "topic"
---

# BDL 3.20 new features

> Features added in 3.20 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 3.20 upgrade guide](0157-bdl-3-20-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 3.20.").

Prior new features guide: [BDL 3.10 new features](0063-bdl-3-10-new-features.md "Features added in 3.10 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| Support for circular dependency of modules with `IMPORT FGL`. | See [Circular dependency with IMPORT FGL](0173-circular-dependency-with-import-fgl.md "The compiler allows that two modules reference each other with IMPORT FGL."). |
| Method declaration, to define functions acting on a user-defined type. | See [Methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), [FUNCTION syntax](../08_language-basics/0763-function-definitions.md). |
| `INTERFACE` structure, to define a group of methods for a type. | See [Interfaces](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type."). |
| Passing records by reference (`INOUT`). | See [Passing records by reference (with INOUT)](../09_advanced-features/0835-passing-records-as-parameter.md). |
| Variable initializers: The `DEFINE` instruction supports a clause to initialize the variable. | See [Variable initializers](../08_language-basics/0691-variable-initializers.md "Variables can be initialized in their definition."), [DEFINE](../08_language-basics/0688-define.md "The DEFINE instruction declares a program variable with a given type."). |
| Named parameters in function calls: `CALL func(p1:value)`. | See [Naming parameters in a function call](../08_language-basics/0766-calling-functions.md). |
| Function attributes (for RESTful Web Services). A `FUNCTION` definition can specify function attributes, parameter attributes and return values attributes. | See [Function attributes](../08_language-basics/0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values."), [FUNCTION definitions](../08_language-basics/0763-function-definitions.md "A FUNCTION definition defines a named procedure with a set of statements."). |
| New `STRING` and `base.StringBuffer` methods to remove all kinds of whitespace characters in a string. | See the `*WhiteSpace` methods in [STRING data type methods](../15_library-reference/2916-string-data-type-methods.md) and [base.StringBuffer methods](../15_library-reference/3046-base-stringbuffer-methods.md). |
| `FUNCTION` definition syntax allows the `RETURNS ()` clause, to specify an empty return list and enforce compilation verifications to avoid `RETURN` instruction misusage. | See [FUNCTION syntax](../08_language-basics/0763-function-definitions.md), [Returning values](../08_language-basics/0768-returning-values.md "A function can return values with the RETURN instruction."). |
| Arrays can be assigned without the `.*` notation. | See [Dynamic array assignment with .\* notation](0171-dynamic-array-assignment-with-notation.md "The .* notation to assign dynamic arrays is discouraged."). |
| Records can be assigned without the `.*` notation. | See [Record copy without .\* notation](0172-record-copy-without-notation.md "The .* notation to assign records is discouraged."). |
| FGLPROFILE `fglrun.floatToCharScale2` and `fglrun.floatToCharScale2.print` to control `FLOAT/SMALLFLOAT` to string conversion. | See [FLOAT/SMALLFLOAT to string conversion](0175-float-smallfloat-to-string-conversion.md "New FGLPROFILE entry fglrun.floatToCharScale2 for FLOAT/SMALLFLOAT types."). |
| GetOpt utility type and methods, to process command line arguments in the getopt style (`--option[=value]`). | See [getopt: Command line options module](../15_library-reference/2893-getopt-command-line-options-module.md). |
| New `os.Path.glob()` method to get a list of files from a pattern. | See [os.Path.glob](../15_library-reference/3721-os-path-glob.md "Returns a list of files matching the specified pattern."). |
| `TEXT` and `BYTE` variables are now automatically located in memory by JSON parsing methods. | See [JSON to BDL data conversion](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md). |

| Overview | Reference |
| --- | --- |
| The new `gui.rendering` FGLPROFILE entry, to enable Universal Rendering.**Important:**This feature is deprecated, its use is discouraged although not prohibited. | See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| FGLGBCDIR environment variable, to define the directory of the GBC component to be used for Universal Rendering. | See [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode."), [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| `KEYBOARDHINT` form attribute allows now the `URL` value to define an `EDIT` field for URL input. | See [KEYBOARDHINT attribute](../11_user-interface/1798-keyboardhint-attribute.md "The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly."). |
| GBC supports the `"chrome"` value for `actionPanelPosition`, `ringMenuPosition` and `toolBarPosition` window style attributes. | See [Action views rendering in GBC chromebar](0161-presentation-styles-changes.md). |
| `ui.Interface` class methods to identify the Universal Rendering client and version.**Important:**This feature is deprecated, its use is discouraged although not prohibited. | See [ui.Interface.getUniversalClientName](../15_library-reference/3110-ui-interface-getuniversalclientname.md "Returns the name of the front-end used for Universal Rendering."), [ui.Interface.getUniversalClientVersion](../15_library-reference/3111-ui-interface-getuniversalclientversion.md "Returns the version of the front-end used for Universal Rendering."). |
| Form elements of type `GROUP` can be defined as collapsible with new style attributes (`collapsible` / `initiallyCollapsed`). | See [Collapsible groups with GBC](0161-presentation-styles-changes.md). |
| Collapsible groups and folders can be configured to display the collapser icon at a specific position with the collapserPosition style attribute. | See [New Folder.collapserPosition (GBC)](0161-presentation-styles-changes.md), [Collapsible groups with GBC](0161-presentation-styles-changes.md). |
| `ON FILL BUFFER`, `ON IDLE` and `ON TIMER` triggers for display array dynamic dialog. | See [ui.Dialog.addTrigger](../15_library-reference/3184-ui-dialog-addtrigger.md "Adds an event trigger to the dynamic dialog"). |
| Built-in classes methods to control form elements are now case-insensitive. | See [Case insensitive names with UI methods](0174-case-insensitive-names-with-ui-methods.md "Methods of built-in classes using user interface object names are now case insensitive."). |
| Use the `itemsAlignment` style attribute to control the alignment of elements inside a scrollgrid. | See [Controlling element alignment inside a scrollgrid](../11_user-interface/2339-controlling-scrollgrid-rendering.md). |
| New `enableCalendar` style attribute for `DateTimeEdit` with GDC. | See [New DateTimeEdit.enableCalendar style attribute (GDC)](0161-presentation-styles-changes.md). |
| New `qtStyle` style attribute for `ComboBox` with GDC. | See [New ComboBox.qtStyle style attribute (GDC)](0161-presentation-styles-changes.md). |
| Set the `sanitize` style attribute to `"no"` to disable HTML content cleanup (Warning: allows Stored XSS attacks!) | See [New sanitize style attribute (GBC)](0161-presentation-styles-changes.md). |
| The `INITIALIZER` attribute of `COMBOBOX` accepts module specification as function prefix (the module is then automatically loaded) | See [INITIALIZER attribute](../11_user-interface/1791-initializer-attribute.md "The INITIALIZER attribute allows you to specify an initialization function that will be automatically called by the runtime system to set up the form item."). |
| Setting the comment of a form element or form field. | See [ui.Form.setElementComment](../15_library-reference/3155-ui-form-setelementcomment.md "Set the comment/hint of form elements."), [ui.Form.setFieldComment](../15_library-reference/3160-ui-form-setfieldcomment.md "Set the comment/hint of a form field."). |
| New FGLPROFILE entry `gui.programStoppedMessage` to define a generic message to be displayed to the end user when a program stops because of a runtime error. | See [Default exception handling](../09_advanced-features/0855-default-exception-handling.md "Default exception handling must be adapted to your programming pattern."), [Program stop error message box](0203-program-stop-error-message-box.md "In GUI mode, runtime errors stopping the program are now displayed to the end user."). |
| `AUTONEXT` can be used on `EDIT` or `BUTTONEDIT` with the `COMPLETER` attribute, and on `COMBOBOX`, to jump to the next field when a value is selected from the list. | See [AUTONEXT attribute](../11_user-interface/1763-autonext-attribute.md "The AUTONEXT attribute forces the focus to automatically leave the current field when completed."). |
| The `lateRendering` style attribute for `FOLDER` containers can be used to optimize the layout processing. | See [New Folder.lateRendering style attribute (GBC)](0161-presentation-styles-changes.md). |
| Upgrade notes for presentation styles. | See [Presentation styles changes](0161-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |
| Upgrade notes for web components. | See [Web components changes](0163-web-components-changes.md "Modifications to consider when using web components."). |
| Upgrade notes for front calls. | See [Front calls changes](0162-front-calls-changes.md "Modifications to consider when using front calls."). |

| Overview | Reference |
| --- | --- |
| Support for Oracle® database 18c and 19c with the `dbmora_18` driver. | See [Database drivers changes](0164-database-drivers-changes.md "New and desupported database drivers."). |
| Using Oracle 18c private temporary tables to emulate Informix® temporary tables. | See [Using the private temporary table emulation](../10_sql-support/1393-using-the-private-temporary-table-emulation.md). |
| Support for PostgreSQL 11, 12 with the `dbmpgs_9` driver (PostgreSQL 13 and 14 backported from FGL 4.01) | See [Database drivers changes](0164-database-drivers-changes.md "New and desupported database drivers."), [PostgreSQL 12 notes](0169-postgresql-12-notes.md "This topics contains notes about PostgreSQL 12 changes that affect Genero applications."). |
| Support for SQL Server 2019 (v15) with the `dbmsnc_17` driver. | See [Database drivers changes](0164-database-drivers-changes.md "New and desupported database drivers."). |
| New `dbmmys_5_6` driver for MySQL 5.6 as replacement of `dbmmys_5_5` no longer available because of MySQL 5.5 desupport. | See [Database drivers changes](0164-database-drivers-changes.md "New and desupported database drivers."). |
| New `ifxemul.datatype.serial.sqlerrd2` FGPROFILE entry to disable automatic serial retrieval for `sqlca.sqlerrd[2]`. | See [FGLPROFILE entries for core language](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL."), [Disabling automatic serial retrieval for sqlca.sqlerrd[2]](../10_sql-support/1377-serial-and-bigserial-data-types.md). |
| Utility function `db_get_last_serial()` to retrieve the last generated serial, mandatory for `BIGINT` incremented columns. | See [fgldbutl.db\_get\_last\_serial()](../15_library-reference/2803-fgldbutl-db-get-last-serial.md "Retrieves the last generated serial for a given serial emulation and database table."). |
| FGLPROFILE entry `dbi.database.dbname.pgs.schema` to define the schema search path for PostgreSQL. | See [PostgreSQL specific FGLPROFILE parameters](../10_sql-support/1084-postgresql-specific-fglprofile-parameters.md). |
| FGLSQLDEBUG ouput enhancement to identify when `SQL INTERRUPT ON/OFF` is used. | See note in [Performances with SQL interruption](0168-performances-with-sql-interruption.md "With some database drivers, performances can be impacted when using OPTIONS SQL INTERRUPT ON."). |
| Support for IBM® Informix trusted connections, with the new `TRUSTED` keyword for `CONNECT TO`, and the new `SET SESSION AUTHORIZATION` statement. | See [CONNECT TO](../10_sql-support/1099-connect-to.md "Opens a new database session in multi-session mode."), [SET SESSION AUTHORIZATION](../10_sql-support/1101-set-session-authorization.md "Select the user under which database operations are performed in the current connection."). |
| Informix SQL `SELECT` statement row limiting clause `SKIP n FIRST m` can be converted to native equivalent. | See [Row limiting clause (SELECT)](../10_sql-support/1051-row-limiting-clause-select.md "How to use the right clause to limit the number of rows produced by a SELECT statement?"). |
| With Microsoft SQL Server, in DDL statements such as `CREATE TABLE`, character data types can be converted to the equivalent native national character types, with FGLPROFILE entry `dbi.database.dsname.ifxemul.nationalchars=true`. | See [character types with SQL Server](../10_sql-support/1273-char-and-varchar-data-types.md), [ifxemul FGLPROFILE settings](../10_sql-support/1079-ibm-informix-emulation-parameters-in-fglprofile.md). |
| To avoid Microsoft SQL Server 468 collation conflict errors with temporary tables, the ODI drivers for SQL Server add now a `COLLATE DATABASE_DEFAULT` clause after character types in the native statement generated from a `CREATE TEMP TABLE`. This is not needed for `SELECT INTO TEMP`. | See [Temporary tables](../10_sql-support/1291-temporary-tables.md). |
| Oracle database rowid (in base 64 format) can be found in `sqlca.sqlerrm`. | See [ORACLE rowid in sqlca.sqlerrm](0176-oracle-rowid-in-sqlca-sqlerrm.md "With ORACLE, the rowid of the last affected row is available in sqlca.sqlerrm."). |
| Upgrade notes for database drivers. | See [Database drivers changes](0164-database-drivers-changes.md "New and desupported database drivers."). |

| Overview | Reference |
| --- | --- |
| Source code formatting tool: fglcomp --format, including new command line tools fglgitformat, fglformatdiff. | See [Source code beautifier](../13_programming-tools/2636-source-code-beautifier.md "Reformat the source code for better readability."), [fglgitformat](../13_programming-tools/2527-fglgitformat.md "The fglgitformat tool reformats the source lines changed in a GIT history."), [fglformatdiff](../13_programming-tools/2528-fglformatdiff.md "The fglformatdiff tool reformats source lines of a unified diff text."). |
| Qualifying imported symbols: fglcomp --qualify-imports. | See [Qualifying imports](../13_programming-tools/2642-qualifying-imported-symbols.md "The fglcomp compiler can automatically qualify imported symbols."). |
| Use lowercase keywords in VIM with `fgl_lowercase_keywords=1`. | See [Configure VIM for Genero BDL](../13_programming-tools/2544-configure-vim-for-genero-bdl.md). |
| fglrun -l/fgllink supports `@argfile` to specify the list of files to link. | See [Providing the files to link in an arguments file](../13_programming-tools/2537-linking-programs.md), [Providing the files to link in an arguments file](../13_programming-tools/2536-linking-libraries.md). |
| fglcomp can now process several .4gl files on the command line. | See [Compiling several .4gl sources in a single command](../13_programming-tools/2534-compiling-program-code-files-4gl.md), [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks."). |
| Commands accepting a list of files like fglcomp or fgllink, can do pathname expansion (fglcomp \*.4gl), on Unix and Windows platforms. | See [Compiling several .4gl sources in a single command](../13_programming-tools/2534-compiling-program-code-files-4gl.md), [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks."). |
| To reduce compilation time of large projects, fglcomp supports the `--make` option, to compile the provided .4gl sources only when the .42m is not up-to-date. The `--make` option can be used with the `--simulate` option, to simulate the process without compiling. | See [Compiling in make mode](../13_programming-tools/2534-compiling-program-code-files-4gl.md). |
| The fglcomp `--dependencies` option can be used to produce makefile-style dependency rules for imported modules. | See [Producing make-style dependency rules](../13_programming-tools/2534-compiling-program-code-files-4gl.md). |
| fglcomp supports `@argfile` to specify the list of source files in an input file. | See [Providing the source files in an arguments file](../13_programming-tools/2534-compiling-program-code-files-4gl.md). |
| fglcomp supports the value "`auto`" for `-p` preprocessor type option. | See [Preprocessing style option](../13_programming-tools/2564-compilers-command-line-options.md). |
| fglrun provides the `--print-missing-imports` option to show only missing `IMPORT FGL` (`--print-imports` option prints required imports even if `IMPORT FGL` is used) | See [Identifying modules to be imported](../09_advanced-features/0821-identifying-modules-to-be-imported.md "Use the --print-missing-imports and --print-imports options to identify missing IMPORT FGL instructions."). |
| fglcomp supports the new warning argument `-W unused-parameter`, that prints a message when a function parameter is not used. The negative form `-W no-unused-parameter` can be combined with `-W unused`, to ignore parameters that are unused. | See [fglcomp option `-W` option](../13_programming-tools/2516-fglcomp.md). |
| **Starting at 3.20.16** |  |
| fglgar run --pid-file option writes fglgar process-ids to file. | See [fglgar](../13_programming-tools/2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS)."). |

| Overview | Reference |
| --- | --- |
| The Genero Web Service engine has been enhanced with a mechanism to provide a high level REST web service. You can implement REST Web services using function attributes (identified with `WS*` prefix) that you specify in `ATTRIBUTES()` clauses of functions. | See:[Function attributes](../08_language-basics/0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values.")[High-level RESTful Web service attributes](../16_web-services/4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")[RESTful Web services (high-level framework)](../16_web-services/4705-restful-web-services-high-level-framework.md "These topics give you the information you need to begin working with RESTful Web services applications using BDL function with support for attributes.") |
| The fglrestful tool takes an OpenAPI documentation file (JSON format) and generates the client stub to interact with the REST service. | See [fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification."). |
| Descriptions using the `WSDescription` attribute are allowed in input and output parameters. These descriptions can be generated in the REST stub file using the fglrestful tool's `--comment` option. | See [fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.") and [WSDescription](../16_web-services/4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types."). |
| `com.WebServiceEngine` option `server_restdefaultformat` added to define runtime support for MIME types in REST operations. | See [WebServiceEngine options](../15_library-reference/3804-webserviceengine-options.md). |
| Method added to publish high-level RESTful Web services in the GWS engine. | See [com.WebServiceEngine.RegisterRestService](../15_library-reference/3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine."). |
| Method added to provide error management in the high-level RESTful Web service. | See [com.WebServiceEngine.SetRestError](../15_library-reference/3800-com-webserviceengine-setresterror.md "Manages error handling for a REST high-level Web Service function."). |
| Empty HTTP POST and PUT requests. | See [Support for empty HTTP POST or PUT requests](0158-web-services-changes.md). |
| The FGLPROFILE entry `http.global.request.date` can be set to `true`, in order to force sending HTTP Date header for GET, HEAD and DELETE requests. | See [Control HTTP Date header for GET, HEAD and DELETE requests](0158-web-services-changes.md). |
| The `com.HttpRequest.setAuthentication` has an enhancement to support the NTLM authentication protocol of Windows® server that requires the connection to be kept open. | See [com.HttpRequest.setAuthentication](../15_library-reference/3869-com-httprequest-setauthentication.md "Defines the user login and password to authenticate to the server.") or [com.HttpRequest.setKeepConnection](../15_library-reference/3878-com-httprequest-setkeepconnection.md "Defines whether a connection is kept open if a new request occurs."). |
| API added to provide management of OAuth authentication for RESTful Web service secured by OpenID Connect. | See [OAuthAPI library](../16_web-services/4930-oauthapi-library.md "The OAuthAPI library provides a set of functions for working with access to web services using the OAuth protocol."). |
| **Starting at 3.20.13** |  |
| The `WSHelper.WSServerCookiesType` record has a member added for setting the `sameSite` security value. The `com.HttpServiceRequest.setResponseCookies()` method, which uses the record, may raise a runtime error if the `SameSite` attribute is not set correctly. | See: [Changes to how GWS handles cookies](0158-web-services-changes.md), [Single sign-on (OpenID Connect, SAML, and GIP) sameSite security](0158-web-services-changes.md). |
| **Starting at 3.20.14** |  |
| The `fglwsdl -xmlname` option is added to generate variables named with [XMLName](../16_web-services/5037-xmlname.md) attributes in stubs. | See [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") |
| **Starting at 3.20.15** |  |
| The REST engine `WSAttachment` attribute has been enhanced with an optional regular expression (regex) type value to validate the names of incoming files. | See [WSAttachment](../16_web-services/4839-wsattachment.md "Defines file attachments in the REST message.") |
| **Starting at 3.20.17** |  |
| The GWS OpenID Connect service has enhancements to how scopes are exchanged. The Genero Identity Provider (GIP) follows the standard RFC 8693 as the default method when creating OAuth ID and access tokens with the scope parameter. | See [Support for RFC 8693 in the Genero Identity Provider (GIP) creation of OAuth ID and access tokens with scopes](0158-web-services-changes.md) |
| [***Related upgrade notes***](0157-bdl-3-20-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 3.20.") |  |
| Upgrade notes for web services. | See [Web Services changes](0158-web-services-changes.md "There are changes in support of web services in Genero 3.20."). |

> **Important:**
>
> This version of Genero Mobile (GMA/GMI) is desupported, use latest version of the product.

| Overview | Reference |
| --- | --- |
| gmabuildtool option `-bgr / --build-gbc-runtime`, to define the GBC to be used to build the app. | See [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices."). |
| gmibuildtool option `--gbc`, to define the GBC to be used to build the app. | See [Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices."). |
| GMA supports now the `alignment` style attribute for `BUTTON`. | See [Button.alignment style attribute (GMA)](0161-presentation-styles-changes.md), [Button style attributes](../11_user-interface/1631-button-style-attributes.md "Button presentation style attributes apply to BUTTON elements."). |
| Cordova plugin wrapper libraries now available on the FOURJS Cordova Github. | See [Cordova plugins for GMA](0159-genero-mobile-for-android-gma-1-40-changes.md), [Cordova plugins for GMI](0160-genero-mobile-for-ios-gmi-1-40-changes.md). |
| New Cordova plugin front calls to get plugin information. | See [New Cordova front calls in GMA](0159-genero-mobile-for-android-gma-1-40-changes.md), [New Cordova front calls in GMI](0160-genero-mobile-for-ios-gmi-1-40-changes.md). |
| GMA and GMI support `standard.openFile/openFiles` front calls. | See [standard.openFile](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system."), [standard.openFiles](../15_library-reference/3410-standard-openfiles.md "Displays a file dialog window to let the user select a list of file paths on the local file system."), [File management](../17_mobile-applications/5089-file-management.md "Specific APIs are available to manipulate file resources in mobile apps."). |
| gmabuildtool produces Android App Bundle .aab packages. | See [Android App Bundle (.aab) packages](0159-genero-mobile-for-android-gma-1-40-changes.md) |
| GMA build tool command `updatesdk` has a `--no-interaction` option to handle user interaction during installation of extras. It allows you to silently accept these installations during updates of Android SDK. | See [gmabuildtool updatesdk --no-interactions](../17_mobile-applications/5106-gmabuildtool.md) and [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices."). |
| **Starting at GMA-1.40.22** |  |
| New gmabuildtool option to enable data collection for analytics, when using Firebase Cloud Messaging. | See [Controlling Firebase data collection for analytics](0159-genero-mobile-for-android-gma-1-40-changes.md) |
| **Starting at GMI-1.40.14** |  |
| New gmibuildtool option to update the app installed on a simulator. | See [New option --update for gmibuildtool](0160-genero-mobile-for-ios-gmi-1-40-changes.md) |
| [***Related upgrade notes***](0157-bdl-3-20-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 3.20.") |  |
| Upgrade notes for GMA. | See [Genero Mobile for Android (GMA) 1.40 changes](0159-genero-mobile-for-android-gma-1-40-changes.md "Modifications to consider when using the Genero Mobile for Android."). |
| Upgrade notes for GMI. | See [Genero Mobile for iOS (GMI) 1.40 changes](0160-genero-mobile-for-ios-gmi-1-40-changes.md "Modifications to consider when using Genero Mobile for iOS."). |

| Overview | Reference |
| --- | --- |
| New parameter for monitor.update front call to control administrator permission elevation prompt. | See [Front calls changes](0162-front-calls-changes.md "Modifications to consider when using front calls."), [monitor.update front call](../15_library-reference/3427-monitor-update.md "Starts the GDC update."). |
