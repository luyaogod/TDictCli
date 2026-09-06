---
title: "BDL 2.00 new features"
source: "fgl-topics/fgl_whatsnew_200.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 2.00 new features"
type: "topic"
---

# BDL 2.00 new features

> Features added in 2.00 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 2.00 upgrade guide](0306-bdl-2-00-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.00.").

Prior new features guide: [BDL 1.33 new features](0078-bdl-1-33-new-features.md "Features added in 1.33 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| The runtime system (`fglrun`) now uses shared libraries for database drivers; there is no need to link anymore. | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| The `TYPE` instruction allows to define your own data type structures. | See [Types](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."). |
| File management function library provided as loadable extension. | See [The os.Path class](../15_library-reference/3697-the-os-path-class.md "The os.Path class provides functions to manipulate files and directories on the machine where the program executes."). |
| Mathematical function library provided as loadable extension. | See [The util.Math class](../15_library-reference/3561-the-util-math-class.md "The util.Math class provides basic mathematical functions based on floating point numbers (FLOAT)."). |
| C extension support has been extended with Informix-like C API functions. | No longer applicable as of Genero 2.51. |
| The runtime system now shares several static elements among all processes, reducing the memory usage. The shared elements are: Data type definitions, string constants and debug information. For example, when a program defines a string containing a long SQL statement, all fglrun processes will share the same string, which is allocated only once. | See [Runtime system basics](../09_advanced-features/0967-runtime-system-basics.md "This section contains topics about Genero BDL runtime system basic concepts."). |
| The `IMPORT` instruction allows to declare a C extension module. | See [IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module."). |
| New debugger commands (`call`, `ignore`). | See [Debugger commands](../13_programming-tools/2586-debugger-commands.md "This topic lists all debugger commands."). |
| The `base.Channel` class now has an `isEof()` method to detect end of file. | See [Read and write text lines](../15_library-reference/3002-read-and-write-text-lines.md). |
| Ignoring the CTRL\_LOGOFF\_EVENT events on Microsoft™ Windows® platforms. | See [Responding to CTRL\_LOGOFF\_EVENT](../09_advanced-features/0936-responding-to-ctrl-logoff-event.md "FGLPROFILE fglrun.ignoreLogoffEvent controls program behavior in case of logoff events on Windows platforms."). |
| New built-in function to set an environment variable: `FGL_SETENV()`. | See [fgl\_setenv()](../15_library-reference/2781-fgl-setenv.md "Sets the value of an environment variable."). |
| The XML reader and writer classes have been extended to properly support markup language entities (like HTML's `&nbsp;` ). | See [The XmlReader class](../15_library-reference/3367-the-xmlreader-class.md "The om.XmlReader class provides methods to read and process a file written in XML format."), [The XmlWriter class](../15_library-reference/3377-the-xmlwriter-class.md "The om.XmlWriter class implements methods to write XML to a stream."). |

| Overview | Reference |
| --- | --- |
| New form item types (i.e. widgets): `SLIDER`, `SPINEDIT`, `TIMEEDIT`. | See [ATTRIBUTES section](../11_user-interface/1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form."). |
| The `WIDTH` and `HEIGHT` attributes can be used for `IMAGE` form items, as a replacement for `PIXELWIDTH`/`PIXELHEIGHT`. | See [HEIGHT attribute](../11_user-interface/1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element."), [WIDTH attribute](../11_user-interface/1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element."). |
| New debugger commands (`call`, `ignore`). | See [Debugger commands](../13_programming-tools/2586-debugger-commands.md "This topic lists all debugger commands."). |
| Presentation styles support now pseudo selectors such as `focus`, `active`, `inactive`, `input`, `display` for fields and `odd` / `even` states for table rows. | See [Pseudo selectors](../11_user-interface/1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled."). |
| New presentation style attributes were added:`'errorMessagePosition`' can be used for windows to define how the ERROR message must be displayed;`'highlightTextColor`' for tables allows you to change the color of the selected line;`'border`' allows you to remove the border of some widgets like button, images;`'firstDayOfWeek`' can be used for DateEdit widget to specify the first day of the week in the calendar;The auto-selection behavior for ComboBoxes and RadioGroup can be changed using '`autoSelectionStart'`. | See [Style attributes reference](../11_user-interface/1628-style-attributes-reference.md "A presentation style attribute may be a common attribute that can be applied to any graphical element. Most presentation style attributes apply only to a specific graphical element."). |
| With X11 or Windows TSE environments, you can now automatically start up the front-end with FGLPROFILE entries. | See [Automatic front-end startup](../11_user-interface/1530-automatic-front-end-startup.md). |
| Up to fourth accelerators can now be defined for an action in actions defaults files or in the `ACTION DEFAULTS` section of form files. | See [Defining keyboard accelerators for actions](../11_user-interface/2265-defining-keyboard-accelerators-for-actions.md). |
| Specify TTY attributes (`COLOR`, `REVERSE`) and conditional TTY attributes (`COLOR WHERE`) for all type of fields. | See [COLOR attribute](../11_user-interface/1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element."), [REVERSE attribute](../11_user-interface/1813-reverse-attribute.md "The REVERSE attribute displays any value in the field in reverse video (dark characters in a bright field)."), [COLOR WHERE Attribute](../11_user-interface/1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically."). |

| Overview | Reference |
| --- | --- |
| Database schema files have been extended to centralize form field definition with the new `FIELD` item type. | **Important:**This feature is deprecated in 2.51 and +. |
| Call database stored procedures with output parameters with the new `IN`/`OUT` keywords. | See [EXECUTE (SQL statement)](../10_sql-support/1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared."), [Stored procedures](../10_sql-support/1047-stored-procedures.md "Executing stored procedures with different database engine types."). |
| Primary key, foreign key and check constraints can be specified in static SQL CREATE TABLE statements:CREATE TABLE t1 ( col1 INTEGER PRIMARY KEY, col2 CHAR(2), col3 DATE, FOREIGN KEY (col2) REFERENCES t2(col1) ) | See [CREATE TABLE](../10_sql-support/1128-create-table.md "Creates a new table object in the database."). |
| The `fgldbsch` tool can now extract database tables with LVARCHAR columns. The LVARCHAR type is converted to VARCHAR2(n>255) in the .sch file. | See [Data type conversion control](../09_advanced-features/0802-data-type-conversion-control.md). |
| Upgrade notes for database drivers. | See [Database drivers changes](0310-database-drivers-changes.md "Desupported database drivers."). |

| Overview | Reference |
| --- | --- |
| You can now choose to use Document Style Service (Doc/Literal) or RPC Literal Style Service (RPC/Literal) with Genero Web Services (GWS), for .NET compatibility and WS-I compatibility (standards defined by the Web Services Interoperability organization).Document Style Service allows you to exchange complex data structures, such as database tables or word processing documents (MS.Net default)RPC Literal Style Service is usually used to execute a function, such as a service that returns a stock option**Note:**RPC/Encoded Style Service (Traditional SOAP section 5) is available for backward compatibility. | See [Choosing a web services style](../16_web-services/4673-choosing-a-web-services-style.md "Genero Web Services contains style options for creating SOAP Web services. Your choice is dependent on the type of service, (Document or RPC), and the encoding mechanism (literal or encoded) required.") and [Writing a Web server application](../16_web-services/4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. ."). |
| Genero Web Services now provides a tool, `fglwsdl`, to allow a Genero application that is accessing a Web Service to obtain the WSDL information for the service. It does not matter what language the Web Service is written in. The `fglwsdl` tool is installed in Genero as part of the Genero Web Services package. | See [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD)."). |
| You no longer need to create a runner that includes the Genero Web Services package. Instead, your applications import the Genero Web Services library named `com`. This library provides classes and methods that allow you to perform tasks associated with creating GWS Servers and Clients, and managing the Web Services. | See [The com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services."). |
| GWS now supports SOAP header management through the `CreateHeader` method in the Web Service class that is part of the Web Services library (`com`). | See [The WebService class](../15_library-reference/3754-the-webservice-class.md "The com.WebService class provides an interface to create and manage Genero Web Services."). |
| HTTPS support has been added on the client side. GWS supports secure communications through the use of encryption and standard X.509 certificates. Based on the OpenSSL engine, new security features allow a Web Services client to communicate with any secured server over HTTP or HTTPS.A new tool is provided, `fglpass`, allowing you to encrypt a password from a standard X.509 certificate, and to decrypt a password you previously encrypted with a certificate.Entries in the FGLPROFILE file are used to define the configuration for client security. | See [fglpass](../13_programming-tools/2525-fglpass.md "The fglpass tool allows you to encrypt passwords."), [Encryption, BASE64 and password agent with fglpass tool](../16_web-services/4557-encryption-base64-and-password-agent-with-fglpass-tool.md "Genero Web Services supports password encryption with fglpass as password agent."), and [The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files"). |
| You can configure a GWS Client to connect via an HTTP proxy by adding an entry in the FGLPROFILE file. | See [Configure a WS client to connect via an HTTP Proxy](../16_web-services/4634-configure-a-ws-client-to-connect-via-an-http-proxy.md "Configuration steps to connect via a HTTP proxy."). |
| You can define multiple Web Services in a single Genero DVM. When you start the Web Services engine, all registered Web Services are started. | See [The WebServiceEngine class](../15_library-reference/3785-the-webserviceengine-class.md "The com.WebServiceEngine class provides an interface to manage the Web Services engine."). |
| You can remap the location of Genero Web Services using entries in the FGLPROFILE file, depending on the network configuration and the access rights management of the deployment site. | See [Use logical names for service locations](../16_web-services/4629-use-logical-names-for-service-locations.md "Using a logical reference for the Web service, instead of the real URL, in your client application URL binding has advantages for working with and deploying applications."). |
| Serializing Genero data types: you can add optional attributes to the definition of data types. You can use these attributes to map the BDL data types in a Genero Web Services Client or Server application to their corresponding XML data types. | See [XML serialization rules and customization](../16_web-services/4958-xml-serialization-rules-and-customization.md). |
| The WSHelper.42m library included in the $FGLDIR/lib directory of the Genero Web Services package file contains internal BDL functions to handle SOAP requests and errors.It is recommended that it is linked into every Genero Web Services Server or Client program. | See [Compile the client application](../16_web-services/4606-compile-the-client-application.md "Compiling the client and the stub file.") and [Compiling GWS server applications](../16_web-services/4668-compiling-gws-server-applications.md "When compiling, remember to include the WSHelper library."). |
| Upgrade notes for web services. | See [Web Services changes](0307-web-services-changes.md "There are changes in support of Web services in Genero 2.00."). |
