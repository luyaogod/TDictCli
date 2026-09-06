---
title: "BDL 3.00 new features"
source: "fgl-topics/fgl_whatsnew_300v.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 3.00 new features"
type: "topic"
---

# BDL 3.00 new features

> Features added in 3.00 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 3.00 upgrade guide](0206-bdl-3-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 3.00.").

Prior new features guide: [BDL 2.51 new features](0065-bdl-2-51-new-features.md "Features added in 2.51 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| Attach the debugger to a running program with `fgldb -p process-id`. | See [Attaching to a running program](../13_programming-tools/2577-attaching-to-a-running-program.md "It is possible to start the debugger for a program running on the same computer."). |
| Improved compilation time (fglcomp and fglform) | See [Improved compilation time](0222-improved-compilation-time.md "The fglcomp and fglform compilers have been reviewed to achieve faster compilation."). |
| Datetime-related utility methods. | See [util.Datetime.getCurrentAsUTC](../15_library-reference/3490-util-datetime-getcurrentasutc.md "Returns the current date/time in UTC."), [util.Datetime.format](../15_library-reference/3488-util-datetime-format.md "Formats a date/time value based on a specified format."), [util.Datetime.parse](../15_library-reference/3491-util-datetime-parse.md "Converts a string to a DATETIME value based on a specified format."). |
| Date-related utility methods. | See [util.Date methods](../15_library-reference/3483-util-date-methods.md "Methods for the util.Date class."). |
| Interval-related utility methods. | See [util.Interval methods](../15_library-reference/3527-util-interval-methods.md "Methods for the util.Interval class."). |
| Temporary filename creation with `os.Path.makeTempName()`. | See [os.Path.makeTempName](../15_library-reference/3730-os-path-maketempname.md "Generates a new file path to be used to create a temporary file or directory."). |
| JSON stringification method to omit NULL elements. | See [util.JSON.stringifyOmitNulls](../15_library-reference/3587-util-json-stringifyomitnulls.md "Produces a JSON formatted string from the input, by excluding empty records and empty arrays."). |
| New fglcomp warning for invalid NULL usage in expressions like var==NULL. | See [Compiler warning -6636](../15_library-reference/4483-genero-bdl-errors.md). |
| fglcomp option to avoid source name in the .42m module. | See [42m module information](../13_programming-tools/2539-42m-module-information.md "Describes how to handle module information in .42m p-code files."). |
| C Extension runtime stack introspection (parameter type and actual string value size in bytes). | See [Runtime stack functions](../14_extending-the-language/2711-runtime-stack-functions.md "To pass values between a C function and a program, the C function and the runtime system use the runtime stack."). |
| The fglmkext command line tool can build your C Extension library. | See [fglmkext](../13_programming-tools/2519-fglmkext.md "The fglmkext tool compiles and links a user C Extension."). |

| Overview | Reference |
| --- | --- |
| Dynamic dialog creation (equivalent of INPUT, CONSTRUCT and DISPLAY ARRAY blocks). | See [Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class."), [ui.Dialog.createConstructByName](../15_library-reference/3171-ui-dialog-createconstructbyname.md "Creates an ui.Dialog object to implement a dynamic CONSTRUCT BY NAME."), [ui.Dialog.createInputByName](../15_library-reference/3174-ui-dialog-createinputbyname.md "Creates an ui.Dialog object to implement a dynamic INPUT BY NAME."), [ui.Dialog.createDisplayArrayTo](../15_library-reference/3172-ui-dialog-createdisplayarrayto.md "Creates an ui.Dialog object to implement a dynamic DISPLAY ARRAY TO."). |
| Resizable `SCROLLGRID` containers (`WANTFIXEDPAGESIZE=NO`). | See [WANTFIXEDPAGESIZE attribute](../11_user-interface/1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element."). |
| The `ON SORT` dialog control block can be used to execute code when the record list is re-ordered by the user. | See [Sorting rows in a list](../11_user-interface/2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required."), [Populating a DISPLAY ARRAY](../11_user-interface/2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog."), [ON SORT block](../11_user-interface/1989-on-sort-block.md), [ui.Dialog.getSortKey](../15_library-reference/3202-ui-dialog-getsortkey.md "Returns the name of the sort field selected by the user."), [ui.Dialog.isSortReverse](../15_library-reference/3207-ui-dialog-issortreverse.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending)"). |
| `ON TIMER` trigger in dialogs, to execute a block of code at regular intervals. | See [Get program control on a regular (timed) basis](../11_user-interface/2227-get-program-control-on-a-regular-timed-basis.md "Execute some code after a given number of seconds, with or without user interaction with the program."). |
| Autocompletion in text edit fields with the COMPLETER attribute. | See [Enabling autocompletion](../11_user-interface/2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field."). |
| Centralization of icon definitions with the FGLIMAGEPATH environment variable. | See [Providing the image resource](../11_user-interface/1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program."), [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources."), [Built-in front-end icons desupport](0218-built-in-front-end-icons-desupport.md "Image resources included in front-ends are desupported with Genero 3.00."). |
| Binding structured ARRAYs in DISPLAY ARRAY and INPUT ARRAY. | See [Structured ARRAYs in list dialogs](0225-structured-arrays-in-list-dialogs.md "ARRAYs with sub-records can be used in list dialogs, to simplify array definition based on database tables, requiring additional information at runtime.") |
| Defining an action for `IMAGE` form items (clickable images). | See [Defining action views in forms](../11_user-interface/2278-defining-action-views-in-forms.md "How to define action views that will fire action events."), [Image columns firing actions](../11_user-interface/2321-image-columns-firing-actions.md "Columns in tables displaying images can trigger action events, when the user selects the image."), [IMAGE item type](../11_user-interface/1695-image-item-type.md "Defines an area that can display an image resource."). |
| Detect window resizing or device orientation change with the `windowresized` predefined action. | See [Adapting to viewport changes](../11_user-interface/1548-adapting-to-viewport-changes.md "Application forms and functions can be adapted to the front-end viewport size or mobile device orientation."). |
| Dialog methods to convert the program array row index to the visual index, and the opposite. | See [ui.Dialog.arrayToVisualIndex](../15_library-reference/3187-ui-dialog-arraytovisualindex.md "Converts the program array index to the visual index for a given screen array."), [ui.Dialog.visualToArrayIndex](../15_library-reference/3234-ui-dialog-visualtoarrayindex.md "Converts the visual index to the program array index for a given screen array."). |
| Providing application image resources to Web Components with `ui.Interface.filenameToURI()`. | See [Using image resources with the gICAPI web component](../11_user-interface/2409-using-image-resources-with-the-gicapi-web-component.md "This section explains how to use image resources in a gICAPI web component."), [ui.Interface.filenameToURI](../15_library-reference/3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource."). |
| The `standard.openFile` front call is now supported with GBC. | See [standard front call support matrix](../15_library-reference/3387-standard-front-calls.md "Standard front call functions provide common utility APIs to control the front-end."). |
| The `dictionariesDirectory` parameter for the `standard.feInfo` front call can be used to get the directory where spell checker dictionary files can be uploaded. | See [`standard.feInfo` front call](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties."). |
| The `allowWebSelection` style attribute can used to enable items selection with a simple mouse drag. | See [Table style attributes](../11_user-interface/1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container."). |
| The `browserMultiPage` style can be used to specify whether the `RUN` and `RUN WITHOUT WAITING` instructions will be executed in the same browser tab or in a new browser tab. | See [UserInterface style attributes](../11_user-interface/1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface."). |
| Upgrade notes for presentations styles. | See [Presentation styles changes](0219-presentation-styles-changes.md "Deprecated and renamed presentation style attributes."). |
| Upgrade notes for front calls. | See [Front calls changes](0220-front-calls-changes.md "Describes changes applied to front calls."). |

| Overview | Reference |
| --- | --- |
| Support for PostgreSQL 9.4. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| Support for SAP® ASE 16.x - Warning: SAP ASE is now desupported in all BDL versions. | SAP ASE is desupported in [BDL version 4.00](../10_sql-support/1073-database-driver-specification-driver.md). |
| Support for SQL Server 2008, 2012 and 2014 with FreeTDS driver (using FreeTDS 0.95) | See [FreeTDS driver supports SQL Server 2008, 2012, 2014](0216-freetds-driver-supports-sql-server-2008-2012-2014.md "The FreeTDS driver can now be used for SQL Server versions > 2005."). |
| Maria DB support (V5.5 and V10): Use the dbmmys driver. | See [MariaDB support](0215-mariadb-support.md "The MariaDB database is now supported by Genero 3.00."). |
| Support for Microsoft® SQL Server 2016 with SNC 11, ESM and FTM drivers. | See [Microsoft SQL Server](../10_sql-support/1256-microsoft-sql-server.md). |
| Dynamic cursor built-in class `base.SqlHandle`. | See [The SqlHandle class](../15_library-reference/3015-the-sqlhandle-class.md "The base.SqlHandle class is a built-in class providing an API to execute parameterized SQL statements, with or without result sets."). |
| SQL interruption is now supported with Oracle® MySQL. | See [Using SQL interruption](../10_sql-support/0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data."). |
| MySQL VARCHAR(N) can be used when N is greater than 255. | See [MySQL VARCHAR size limit](0212-mysql-varchar-size-limit.md "MySQL 5 VARCHAR columns can be used to store VARCHAR(N>255) values."). |
| MySQL DATETIME can store fractional seconds. | See [MySQL DATETIME fractional seconds](0213-mysql-datetime-fractional-seconds.md "MySQL 5.6.4 TIME and DATETIME types support fractions of seconds that can be used to store DATETIME HOUR TO FRACTION(N) or DATETIME YEAR TO FRACTION(N)."). |
| Native Oracle NUMBER type (without precision/scale) can be extracted by fgldbsch. | See [Oracle DB NUMBER type](0210-oracle-db-number-type.md "The NUMBER/FLOAT Oracle data type can now be extracted by fgldbsch to create .sch files."). |
| Serial emulation based on triggers and sequences with SQL Server 2012 and +. | See [SERIAL and BIGSERIAL data types](../10_sql-support/1277-serial-and-bigserial-data-types.md). |
| PostgreSQL connection string option specification in the `source` parameter. | See [Database source specification (source)](../10_sql-support/1072-database-source-specification-source.md), [Prepare the runtime environment - connecting to the database](../10_sql-support/1419-prepare-the-runtime-environment-connecting-to-the-database.md). |
| Upgrade notes for database drivers. | See [Database drivers changes](0209-database-drivers-changes.md "Desupported database drivers."). |

| Overview | Reference |
| --- | --- |
| IPv6 support for Web Services clients. | See [Configure a WS client to use IPv6](../16_web-services/4635-configure-a-ws-client-to-use-ipv6.md "Configuration steps to customize IPv6 for a WS client."). |
| Flushing immediately the response of a web service operation with `com.WebServicesEngine.flush`. | See [com.WebServiceEngine.Flush](../15_library-reference/3787-com-webserviceengine-flush.md "Forces the Web Service engine to immediately flush the response of the web service operation."). |
| Base64 / Hexadecimal / Digest methods using a specific character set for string data. | See [security.Base64.FromStringWithCharset](../15_library-reference/4418-security-base64-fromstringwithcharset.md "Encodes the given string in base64, based on a given charset."), [security.Base64.ToStringWithCharset](../15_library-reference/4424-security-base64-tostringwithcharset.md "Decodes the given base64 string, based on a given charset."), [security.HexBinary.FromStringWithCharset](../15_library-reference/4431-security-hexbinary-fromstringwithcharset.md "Encodes a given string in hexadecimal, based on a given charset."), [security.HexBinary.ToStringWithCharset](../15_library-reference/4437-security-hexbinary-tostringwithcharset.md "Decodes an hexadecimal string to a clear, human-readable string, based on a given charset."), [security.Digest.AddStringDataWithCharset](../15_library-reference/4445-security-digest-addstringdatawithcharset.md "Adds a data string to the digest buffer, after converting to the specified character set."). |
| `com.WebServiceEngine` option `server_readwritetimeout` to define a server socket read/write timeout. | See [Web Services changes](0207-web-services-changes.md "There are changes in support of web services in Genero 3.00."), [WebServiceEngine options](../15_library-reference/3804-webserviceengine-options.md). |
| Specific APIs for Apple® Push Notification Service support. | Desupported since V5.00 |
| Methods to perform RESTful requests using files on disk. | See [com.HttpServiceRequest.readFileRequest](../15_library-reference/3831-com-httpservicerequest-readfilerequest.md "Returns the body of the request to a file on disk and returns the filename."), [com.HttpServiceRequest.sendFileResponse](../15_library-reference/3836-com-httpservicerequest-sendfileresponse.md "Sends an HTTP response with the data contained in a file."), [com.HttpRequest.doFileRequest](../15_library-reference/3859-com-httprequest-dofilerequest.md "Performs the request by sending data contained in a file."), [com.HttpResponse.getFileResponse](../15_library-reference/3895-com-httpresponse-getfileresponse.md "Returns the entire HTTP response to a file on the disk."), [com.HttpPart.getAttachment](../15_library-reference/3917-com-httppart-getattachment.md "Returns the absolute path to the HTTP part."), [com.HttpPart.CreateAttachment](../15_library-reference/3915-com-httppart-createattachment.md "Creates a new HttpPart object based on a given filename located on disk."). |
| FGLPROFILE entries to define XML Signature and XML Encrypted data prefix: `xml.signature.prefix` and `xml.encrypted.prefix`. | See [XML configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md). |
| SOAP fault handling works now when HTTP error 200 is returned by the server. | See [SOAP fault handling in client stub](0207-web-services-changes.md). |
| Client stub multipart supports now optional parts. | See [Optional multipart handling in client stub](0207-web-services-changes.md). |
| Upgrade notes for web services. | See [Web Services changes](0207-web-services-changes.md "There are changes in support of web services in Genero 3.00."). |

| Overview | Reference |
| --- | --- |
| Command line tools to build mobile apps. | See [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices."), [Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices."). |
| Starting remote applications from a mobile device with the `runOnServer` front call. | See [Running mobile apps on an application server](../17_mobile-applications/5111-running-mobile-apps-on-an-application-server.md "From the mobile device, programs can be started remotely on an application server, and displayed on the device."). |
| Push notification APIs for Google Cloud Messaging (GMA) and Apple Push Notification Service (GMI), with new predefined actions (`notificationpushed`). | See [Push notifications](../17_mobile-applications/5112-push-notifications.md "This section describes how to implement push notification with Genero."). |
| Extended feInfo front call options for mobile devices (deviceModel, deviceId, freeStorageSpace, iccid, imei, ppi, windowSize, and so on). | See [standard.feInfo](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties."). |
| New `materialFABType` and `materialFABActionList` style attributes for Window class, to control the FAB button on devices following Material Design guidelines. | This feature is desupported since version [4.00 Universal Rendering](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| Front call to display a box controlling debug settings on GMA. | See [android.showSettings](../15_library-reference/3466-android-showsettings.md "Shows the GMA settings box controlling debug options."). |
| Automatic FGLAPPDIR environment variable (defining the path to the appdir), and automatic FGLDIR environment variable, when executing on mobile devices. | See [FGLAPPDIR](../07_configuration/0519-fglappdir.md "Contains the path to the application directory when executing on a mobile device."), [FGLDIR](../07_configuration/0523-fgldir.md "Defines the installation directory of Genero Business Development Language."), [Setting environment variables in FGLPROFILE (mobile)](../07_configuration/0495-setting-environment-variables-in-fglprofile-mobile.md). |
| Front calls to take or choose videos on mobile devices. | See [mobile.chooseVideo](../15_library-reference/3444-mobile-choosevideo.md "Lets the user select a video from the mobile device's video gallery and returns a video identifier."), [mobile.takeVideo](../15_library-reference/3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier.") front calls. |
| Front call to ask user for Android permissions. | See [android.askForPermission](../15_library-reference/3464-android-askforpermission.md "Ask the user to enable a dangerous feature on the Android device.") front call. |
| GMA buildtool `--clean` option to cleanup the scaffold directory in case of interruption or failure in prior build. | See [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices."). |
| GMA buildtool `--no-install-extras` option to avoid installation of extras during Android SDK update. | See [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices."). |
| GMI specific style attribute `iosTabBarUnselectedColor`, to define the color of unselected tab bar elements. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| GMA specific style attribute `androidKeepForeground`, to control the app state and the background state notification. | See [Foreground and background modes](../09_advanced-features/0829-executing-programs.md), [UserInterface style attributes](../11_user-interface/1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface."). |

| Overview | Reference |
| --- | --- |
| Stacked form definition in .per files with the new STACK container, for mobile programming. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
