---
title: "BDL 4.01 new features"
source: "fgl-topics/fgl_whatsnew_401.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 4.01 new features"
type: "topic"
---

# BDL 4.01 new features

> Features added in 4.01 releases of the Genero Business Development Language.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 4.01 upgrade guide](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.").

Previous new features page: [BDL 4.00 new features](0060-bdl-4-00-new-features.md "Features added in 4.00 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| The `ARRAY[n]`, `DYNAMIC ARRAY` and `DICTIONARY` definitions accept now another array or dictionary, as type for the `OF` clause. | See [`ARRAY` syntax](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.") and [`DICTIONARY` syntax](../08_language-basics/0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements."). |
| **Starting at 4.01.02** |  |
| Module aliases with `IMPORT FGL path AS alias-name` | See [IMPORT FGL](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols."). |
| **Starting at 4.01.03** |  |
| New `os.Path` methods `getAccessTime()/setAccessTime()`, `getModificationTime()/setModificationTime()`, `isSameFile()`, `describeLastError()` | See [os.Path methods](../15_library-reference/3698-os-path-methods.md). |
| New `ARRAY` method `sortByComparisonFunction()`. | See [`ARRAY.sortByComparisonFunction()`](../15_library-reference/2954-dynamic-array-sortbycomparisonfunction.md "Sorts the rows in the array by using a comparison function."). |
| New `util.String` methods to compare strings. | See [`util.Strings.collate()`](../15_library-reference/3557-util-strings-collate.md "Compares two strings using locale collation rules."), [`util.Strings.collateNumeric()`](../15_library-reference/3558-util-strings-collatenumeric.md "Compares two strings using locale collation rules and sequences of numerical digits."). |
| `LOAD/UNLOAD` and `base.Channel` support semicolon separator with `"DSV=;"` | See [`LOAD`](../10_sql-support/1176-load.md "Inserts data from a file into an existing database table."), [`UNLOAD`](../10_sql-support/1177-unload.md "Copies data from the database tables into a file."), [base.Channel.setDelimiter](../15_library-reference/2996-base-channel-setdelimiter.md "Define the value delimiter for a channel."). |
| `util.Channels` utility class to wait for activity on a set of TCP sockets from `base.Channel` objects. | See [util.Channels methods](../15_library-reference/3496-util-channels-methods.md "Methods for the util.ChannelsDate class."). |
| The `STRING.getMultibyteLength()` method counts the number of bytes of a string. | See [STRING.getMultibyteLength](../15_library-reference/2924-string-getmultibytelength.md "Counts the number of bytes in the string."). |
| The `STRING.expandTabs()` method converts TAB characters to spaces. | See [STRING.expandTabs](../15_library-reference/2920-string-expandtabs.md "Converts TAB characters to a 8 space characters."). |
| **Starting at 4.01.06** |  |
| `util.Channels.selectWithTimeout` method to wait for activity on a set of TCP sockets from `base.Channel` objects, with timeout option. | See [util.Channels.selectWithTimeout](../15_library-reference/3500-util-channels-selectwithtimeout.md "Waits for activity on a set of TCP socket listening base.Channel objects and returns after a given period of inactivity.") |
| New `base.Channel.flush()` method to flush output of a bidirectional channel. | See [base.Channel.flush](../15_library-reference/2986-base-channel-flush.md "Flushes the channel."). |
| Support for `BYTE.getLength()` method like for `TEXT` type. | See [BYTE.getLength](../15_library-reference/2912-byte-getlength.md "Returns the length of BYTE content.") |
| **Starting at 4.01.08** |  |
| The `reflect.Value.hasKey` and `reflect.Value.removeDictionaryElement` methods. | See [reflect.Value.hasKey](../15_library-reference/4372-reflect-value-haskey.md "Checks if existence of an element in a dictionary."), [reflect.Value.removeDictionaryElement](../15_library-reference/4376-reflect-value-removedictionaryelement.md "Deletes an element of a dictionary.") |
| The `base.Channel.closeOut()` method, to close then writing stream of the channel (EOF equivalent) | See [base.Channel.closeOut](../15_library-reference/2984-base-channel-closeout.md "Closes the writing stream of the channel."). |
| [***Related upgrade notes***](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.") |  |

| Overview | Reference |
| --- | --- |
| The `KEYBOARDHINT` form item attribute gets new options `DECIMAL`, `NUMERIC`, `TEXT`, `SEARCH` and `NONE`, for better control of the virtual keyboard on mobile devices. | See [KEYBOARDHINT attribute](../11_user-interface/1798-keyboardhint-attribute.md "The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly."). |
| New rendering for `ROWBOUND` actions with three-dots button in `TABLE`, `TREE` and `SCROLLGRID`. | See [Actions bound to the current row](../11_user-interface/2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row."), [Binding action views to a table row](../11_user-interface/2322-binding-action-views-to-a-table-row.md "A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering."), [Binding action views to a treeview row](../11_user-interface/2360-binding-action-views-to-a-treeview-row.md "A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering."), [Binding action views to a scrollgrid row](../11_user-interface/2341-binding-action-views-to-a-scrollgrid-row.md "A rowbound action specifies an action to apply to the selected row. Rowbound actions get specific rendering."). |
| **Starting at 4.01.03** |  |
| New `ui.Dialog` method `setColumnComparisonFunction()` to implement custom sort in list dialogs. | See [`ui.Dialog.setColumnComparisonFunction`](../15_library-reference/3222-ui-dialog-setcolumncomparisonfunction.md "Associates a comparison function to a form field of a list dialog."). |
| New `ui.Dialog` method `setGroupBy()` to define the grouping column in list dialogs. | See [`ui.Dialog.setGroupBy`](../15_library-reference/3229-ui-dialog-setgroupby.md "Defines the grouping column of a list dialog."). |
| `STRING`-typed parameters in fglsvgcanvas API. | See [fglsvgcanvas: STRING-typed parameters](0129-web-components-changes.md). |
| New fglsvgcanvas API function to define symbols. | See [fglsvgcanvas.symbol()](../15_library-reference/2885-fglsvgcanvas-symbol.md "Produces an SVG \"symbol\" element."). |
| New fglsvgcanvas API function for @media styling. | See [fglsvgcanvas.mediaQuery()](../15_library-reference/2868-fglsvgcanvas-mediaquery.md "Builds an SVG \"@media\" query for CSS."). |
| Form and combobox default initializer as function reference. | See [ui.Form.setDefaultInitializerFunction](../15_library-reference/3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms."), [ui.ComboBox.setDefaultInitializerFunction](../15_library-reference/3250-ui-combobox-setdefaultinitializerfunction.md "Define the default initializer for combobox form items."). |
| [***Related upgrade notes***](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.") |  |
| Upgrade notes for presentation style. | See [Presentation styles changes](0127-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |

| Overview | Reference |
| --- | --- |
| The `standard.playSound` front call gets a new "wait" option, to define if the front call must immediately return while the sound continues playing. | See [standard.playSound](../15_library-reference/3411-standard-playsound.md "Plays the sound file passed as parameter on the front-end platform."). |
| New presentation style attributes `topmenuDesktopRendering` and `topmenuMobileRendering` for Window. | See [New Window.topmenuDesktopRendering and Window.topmenuMobileRendering style attributes](0127-presentation-styles-changes.md) |
| The `defaultTTFColor` style attribute applies to any kind of form element. | See [defaultTTFColor applies to any type of form element](0127-presentation-styles-changes.md) |
| The `packed` style attribute for `HBOX` and `VBOX`. | See [packed attribute for HBOX and VBOX](0127-presentation-styles-changes.md). |
| The `modalOnLargeScreen` value for the `windowType` style. | See [New windowType value modalOnLargeScreen](0127-presentation-styles-changes.md). |
| The `standard.feInfo` front call supports a new property `"colorScheme"` to get the brightness mode. | See [Front calls changes](0128-front-calls-changes.md "Modifications to consider when using front calls."), [`standard.feInfo: colorScheme`](../15_library-reference/3399-standard-feinfo.md). |
| The predefined actions `enterforeground` / `enterbackground` are triggered when using GBC front end in a browser (GAS) | See [Foreground and background modes](../09_advanced-features/0829-executing-programs.md). |
| The `rowAspect` style attribute for flipped `TABLE` and resizable `SCROLLGRID`. | See [New rowAspect attribute](0127-presentation-styles-changes.md). |
| The `position` style attribute for ToolBar elements. | See [New position attribute for ToolBars](0127-presentation-styles-changes.md). |
| New front calls in `"standard"` module. | See [New front calls](0128-front-calls-changes.md). |
| [***Related upgrade notes***](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.") |  |
| Upgrade notes for presentation style. | See [Presentation styles changes](0127-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |

| Overview | Reference |
| --- | --- |
| Support for Microsoft® ODBC driver 18 for SQL Server with `dbmsnc_18` ODI driver. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| **Starting at 4.01.02** |  |
| The fgldbsch tool can extract PostgreSQL OID columns as BIGINT. | See [fgldbsch type conversion option](../09_advanced-features/0802-data-type-conversion-control.md). |
| **Starting at 4.01.03** |  |
| Support for Microsoft SQL Server 2022. | See [Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md). |
| Support for PostgreSQL 15. | See [Database drivers changes](0126-database-drivers-changes.md "New and desupported database drivers."), [PostgreSQL](../10_sql-support/1415-postgresql.md). |
| **Starting at 4.01.06** |  |
| Support for PostgreSQL 16. | See [Database drivers changes](0126-database-drivers-changes.md "New and desupported database drivers."), [PostgreSQL](../10_sql-support/1415-postgresql.md). |
| Support for MySQL 8.2 Innovation Release (with new `dbmmys_8_2` driver) | See [Database drivers changes](0126-database-drivers-changes.md "New and desupported database drivers."), [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| **Starting at 4.01.07** |  |
| New FGLPROFILE entry to define PostgreSQL client trace file. | See [dbi.database.dsname.pgs.trace.file](../10_sql-support/1084-postgresql-specific-fglprofile-parameters.md). |
| [***Related upgrade notes***](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.") |  |
| Upgrade notes for database drivers. | See [Database drivers changes](0126-database-drivers-changes.md "New and desupported database drivers."). |

| Overview | Reference |
| --- | --- |
| fglcomp -W unused option warns for `PREPARE` and `DECLARE CURSOR` statements with no corresponding usage, and for unused imported modules. | See [fglcomp -W option](../13_programming-tools/2516-fglcomp.md). |
| fglform -k / --keep-going option continues compilation of forms after error. | See [fglform -k option](../13_programming-tools/2514-fglform.md). |
| fglcomp -k / --keep-going option continues compilation of sources after error. | See [fglcomp -k option](../13_programming-tools/2516-fglcomp.md). |
| fglcomp -W form-field-name option to produce warnings for invalid field names in `INPUT BY NAME` and `CONSTRUCT BY NAME` dialogs. | See [fglcomp -W option](../13_programming-tools/2516-fglcomp.md). |
| **Starting at 4.01.02** |  |
| Code completer can find modules from FGLLDPATH, to complete `IMPORT FGL`. | See [Source code completion](../13_programming-tools/2543-source-code-completion.md). |
| **Starting at 4.01.03** |  |
| Microsoft Visual Studio Code (VS Code) extension for Genero BDL. | See [Visual Studio Code extension](../13_programming-tools/2545-visual-studio-code-extension.md), [Debugging with VS Code](../13_programming-tools/2584-debugging-with-vs-code.md "The Genero BDL VS Code extension can be used to debug programs with VS Code debugger capabilities."). |
| Command-line debugger improvements (expression evaluation, in-line editing, completion, command history, list range of lines on stop) | See [Understanding the debugger](../13_programming-tools/2574-understanding-the-debugger.md "This is an introduction to the integrated debugger."), [Expressions in debugger commands](../13_programming-tools/2583-expressions-in-debugger-commands.md "A limited expression syntax can be used in debugger commands."). |
| VIM edition enhancements: Go to definition (`gd`) and BDL language tags (`CTRL-]`). | See [Configure VIM for Genero BDL](../13_programming-tools/2544-configure-vim-for-genero-bdl.md). |
| fglgar run --pid-file option writes fglgar process-ids to file. | See [fglgar](../13_programming-tools/2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS)."). |
| [***Related upgrade notes***](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.") |  |

| Overview | Reference |
| --- | --- |
| There is support for the HTTP specification allowing a message body with the HTTP verb DELETE. A message body is supported on client and server Web services using the REST low-level framework. | See [com.HttpRequest methods](../15_library-reference/3848-httprequest-methods.md "Methods for the com.HttpRequest class.") and [com.HttpServiceRequest methods](../15_library-reference/3807-httpservicerequest-methods.md "Methods of the com.HttpServiceRequest class.") |
| For REST Web services using the high-level REST framework, the `WSParam` and `WSQuery` attributes support complex types, such as records and arrays in operation parameters. This enhancement is based on the OpenAPI specification for serialization. It supports default methods for delimiting values and handling items of arrays or records in paths and queries. | See [WSParam](../16_web-services/4834-wsparam.md "Maps a parameter to a value in the resource path template.") and [WSQuery](../16_web-services/4836-wsquery.md "Maps a parameter to a query string value in the request URL.") |
| The `fglrestful` tool has options to allow HTTP and proxy authentication when requesting OpenAPI documentation. | See [fglrestful](../13_programming-tools/2524-fglrestful.md) |
| The `fglwsdl -xmlname` option is added to generate variables named with [XMLName](../16_web-services/5037-xmlname.md) attributes in stubs. | See [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") |
| The REST engine `WSContext` variable has a `RequestURL` entry that you can use to get the URL (base URL, resource path, and query) from the client app. | See [WSContext](../16_web-services/4806-wscontext.md) |
| The REST engine `WSAttachment` attribute has been enhanced with an optional regular expression (regex) type value to validate the names of incoming files. | See [WSAttachment](../16_web-services/4839-wsattachment.md "Defines file attachments in the REST message.") |
| **Starting at 4.01.02** |  |
| The GWS OpenID Connect service has enhancements to how scopes are exchanged. The Genero Identity Provider (GIP) follows the standard RFC 8693 as the default method when creating OAuth ID and access tokens with the scope parameter. | See [Support for RFC 8693 in the Genero Identity Provider (GIP) creation of OAuth ID and access tokens with scopes](0125-web-services-changes.md) |
| **Starting at 4.01.04** |  |
| The GWS OpenID Connect service configuration has a new entry (`oidc.accesstoken.decode`) added to decode roles and scopes in the access token. | See [New option oidc.accesstoken.decode for decoding access tokens with roles and scopes](0125-web-services-changes.md) |
| **Starting at 4.01.06** |  |
| The FGLPROFILE has a new entry, `security.global.options`, added to allow legacy OpenSSL 1 options. | See [New security.global.option in FGLPROFILE to allow legacy OpenSSL 1 options](0125-web-services-changes.md) |
| New fglwsdl option `-SSLOptions` added to support legacy OpenSSL 1 options. | See [fglwsdl option -SSLOptions to support legacy OpenSSL 1 options](0125-web-services-changes.md) and [fglwsdl](../13_programming-tools/2523-fglwsdl.md). |
| `STRING`-typed return parameter in OAUTH API function. | See [Change to OAuthAPI.GetIDSubject returns](0155-web-services-changes.md). |
| **Starting at 4.01.07** |  |
| The FGLPROFILE has a new entry, `security.global.certificate.selfsigned.preload`, added to preload the global certificate and key at the start of the application instead of at the first HTTPS connection. | See [New security.global.certificate.selfsigned.preload entry in FGLPROFILE](0125-web-services-changes.md) |
| **Starting at 4.01.11** |  |
| The function `com.WebServiceEngine.HandleRequest` now restores the previous behavior for query‑only requests targeting the REST service root path, allowing them to fall back to low‑level handling when no REST operation matches. | See [Restored HandleRequest() behavior for root-level query-only requests](0125-web-services-changes.md) and [com.WebServiceEngine.HandleRequest](../15_library-reference/3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.") |
| [***Related upgrade notes***](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.") |  |
| Upgrade notes for web services. | See [Web Services changes](0125-web-services-changes.md "There are changes in support of web services in Genero 4.01.") |

> **Important:**
>
> This version of Genero Mobile (GMA/GMI) is desupported, use latest version of the product.

| Overview | Reference |
| --- | --- |
| **Starting at GMA-4.01.02** |  |
| New gmabuildtool option to enable data collection for analytics, when using Firebase Cloud Messaging. | See [Controlling Firebase data collection for analytics](0159-genero-mobile-for-android-gma-1-40-changes.md) |
| The `standard.playSound` front call supports http/https URLs. | See [GMA 4.01 changes topic](0130-genero-mobile-for-android-gma-4-01-changes.md). |
| **Starting at GMA-4.01.04** |  |
| Default keystore type generated by gmabuildtool is now of type PKCS12. | See [GMA 4.01 changes topic](0130-genero-mobile-for-android-gma-4-01-changes.md). |
| New `mobile.isEmulator` front call. | See [New front calls](0128-front-calls-changes.md). |
| Support GBC theming for app colors. | See [Using GBC theme to define app colors](0130-genero-mobile-for-android-gma-4-01-changes.md). |
| **Starting at GMA-4.01.06** |  |
| Local notification management. | See [mobile.createNotification](../15_library-reference/3449-mobile-createnotification.md "Creates or updates a local notification to be displayed on the mobile device."). |
| New gmabuildtool options to control the manifest file. | See [The app manifest file](../17_mobile-applications/5105-building-android-apps-with-genero.md). |
| New gmabuildtool options to set the path to the program files: `--main-app-path` and `--root-path`. | See [Changes to gmabuildtool command arguments](0130-genero-mobile-for-android-gma-4-01-changes.md) |
| gmabuildtool supports `@argfile` syntax to specify a list of options in one or several files. | See [Changes to gmabuildtool command arguments](0130-genero-mobile-for-android-gma-4-01-changes.md) |
| Permissions summary in gmabuildtool output. | See [Permissions summary in gmabuildtool output](0130-genero-mobile-for-android-gma-4-01-changes.md) |
| **Starting at GMA-4.01.07** |  |
| Support for Firebase Cloud Messaging API V1. | See [Push notifications using FCM API V1](0130-genero-mobile-for-android-gma-4-01-changes.md), [JSON data from mobile.getRemoteNotifications](0128-front-calls-changes.md). |
| **Starting at GMA-4.01.08** |  |
| New gmabuildtool option to define the type of package (AAB or APK), to speed the build. | See [Controlling the production of AAB/APK package files](../17_mobile-applications/5105-building-android-apps-with-genero.md). |
| **Starting at GMI-4.01.04** |  |
| New gmibuildtool option to update the app installed on a simulator. | See [New option --update for gmibuildtool](0160-genero-mobile-for-ios-gmi-1-40-changes.md). |
| Local notification management. | See [mobile.createNotification](../15_library-reference/3449-mobile-createnotification.md "Creates or updates a local notification to be displayed on the mobile device."). |
| New `mobile.isEmulator` front call. | See [New front calls](0128-front-calls-changes.md). |
| **Starting at GMI-4.01.05** |  |
| Notification IDs for remote notifications. | See [JSON data from mobile.getRemoteNotifications](0128-front-calls-changes.md). |
| **Starting at GBC-4.01.20** |  |
| Mobile front calls implemented in GBC JavaScript code. | See [GBC supports mobile front calls](0128-front-calls-changes.md). |
| [***Related upgrade notes***](0123-bdl-4-01-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 4.01.") |  |
| Upgrade notes for GMA. | See [Genero Mobile for Android (GMA) 4.01 changes](0130-genero-mobile-for-android-gma-4-01-changes.md "Modifications to consider when using the Genero Mobile for Android."). |
| Upgrade notes for GMI. | See [Genero Mobile for iOS (GMI) 4.01 changes](0131-genero-mobile-for-ios-gmi-4-01-changes.md "Modifications to consider when using Genero Mobile for iOS."). |
