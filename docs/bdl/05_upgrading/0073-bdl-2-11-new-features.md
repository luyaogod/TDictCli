---
title: "BDL 2.11 new features"
source: "fgl-topics/fgl_whatsnew_211.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 2.11 new features"
type: "topic"
---

# BDL 2.11 new features

> Features added in 2.11 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 2.11 upgrade guide](0298-bdl-2-11-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.11.").

Prior new features guide: [BDL 2.10 new features](0074-bdl-2-10-new-features.md "Features added in 2.10 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| New `-p noln` preprocessor option to remove line number information to get a readable output:fglcomp -E -p noln mymodule.4gl | See [Source preprocessor](../13_programming-tools/2562-source-preprocessor.md "A typical preprocessor like in the C language."). |
| The `-b` option of `fglrun` has been extended to recognize headers of p-code modules compiled with older versions of Genero. | See [42m module information](../13_programming-tools/2539-42m-module-information.md "Describes how to handle module information in .42m p-code files."). |
| The `fglform` compiler now writes build information in the .42f files, to identify on the production site what version was used to compile forms. | See [Compiling form specification files (.per)](../13_programming-tools/2531-compiling-form-specification-files-per.md "The .per form definition files must be compiled to .42f XML files, in order to be loaded by the runtime system."). |

| Overview | Reference |
| --- | --- |
| The `ui.ComboBox` class has been extended with new methods: `getTextOf()` and `getIndexOf()`. | See [The ComboBox class](../15_library-reference/3247-the-combobox-class.md "The ui.ComboBox class provides an interface to the COMBOBOX form field view in the abstract user interface tree."). |
| A new FGLPROFILE entry has been added to force the current row to be shown automatically after a sort in a table:Dialog.currentRowVisibleAfterSort = 1By default, the offset does not change and the current row may disappear from the window. When this new parameter is used, the current row will always be visible. | See [Dialog configuration with FGLPROFILE](../11_user-interface/2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior."). |

| Overview | Reference |
| --- | --- |
| Static SQL syntax now supports derived tables and derived column lists in the FROM clause. For example:SELECT * FROM (SELECT * FROM customer ORDER BY cust_num) AS t(c1,c2,c3,...)See database server documentation for more details about this SQL feature.Informix® 11 does not support the full ANSI SQL 92 specification for derived columns, while other databases do. For this reason, fglcomp allows the ANSI standard syntax. | See [SELECT](../10_sql-support/1123-select.md "Produces a result set from a query on database tables."). |
| The `SET ISOLATION` statement now supports the newInformix 11 clauses for the `COMMITTED READ` option:SET ISOLATION TO COMMITTED READ [LAST COMMITTED] [RETAIN UPDATE LOCKS]When connecting to a non-Informix database, the `LAST COMMITTED` and `RETAIN UPDATE LOCKS` are ignored; other databases do not support these options, and have the same behavior as when these options are used with Informix 11. | See [SET ISOLATION](../10_sql-support/1113-set-isolation.md "Defines the transaction isolation level for the current connection."). |
| The `CAST` operator can now be used in static SQL statements:CAST ( expression AS sql-data-type ) Only Informix data types are supported after the AS keyword. | See [Static SQL statements](../10_sql-support/1115-static-sql-statements.md "Describes static SQL statements supported in the language."). |
| In order to execute database administration tasks, you can now connect to Oracle as SYSDBA or SYSOPER with the `CONNECT` instruction:CONNECT TO "dbname" USER "scott/SYSDBA" USING "tiger" | See [CONNECT TO](../10_sql-support/1099-connect-to.md "Opens a new database session in multi-session mode."). |

| Overview | Reference |
| --- | --- |
| The Genero Web Services com library provides the HttpServiceRequest class to perform low-level XML and TEXT over HTTP communication on the server side. This allows communication at a very low-level layer, to write your own type of web services. | See [The HttpServiceRequest class](../15_library-reference/3806-the-httpservicerequest-class.md "The com.HttpServiceRequest class provides an interface to process incoming XML and TEXT requests over HTTP on the server side, with an access to the HTTP layer and additional XML streaming possibilities."). |
| XML facet constraints attributes: the Genero Web Services XML library provides 12 new XML attributes to map to simple BDL variables. These attributes restrict the acceptable value-space for each variable in different ways such as:a minimum or a maximum number of XML characters or bytes.a strict number of XML characters or bytes.a minimum inclusive or exclusive value depending on the data type.a maximum inclusive or exclusive value depending on the data type,a enumeration of authorized values.a number of digits and fraction digits.how whitespaces have to be handled.a regular expression to match. [(See Section F of XML Schema Part 2)](http://www.w3.org/TR/xmlschema-2/#regexs) | See [XML serialization rules and customization](../16_web-services/4958-xml-serialization-rules-and-customization.md). |
| The fglwsdl tool has been enhanced with the following three new options :`-disk` : to retrieve locally a WSDL or an XSD with all its dependencies from an URL on the disk`-noFacets` : to avoid the generation of the new facet constrain attributes (for compatibility)`-regex` : to validate a value against a regular expression as described in the XML Schema specification | See [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD)."). |

| Overview | Reference |
| --- | --- |
| The Genero Web Services library provides two new methods in the `WebOperation` class to create One-Way operations in services.A One-Way operation means that the server accepts an incoming request, but doesn't return any response back to the client. There is one method called `CreateOneWayRPCStyle` to create an RPC Style operation, and another one called `CreateOneWayDOCStyle` to create a Document Style operation.For instance, a One-Way operation can be used as a logger service, where a client sends a message to the server, but doesn't care about what the server is doing with it. | See [The WebOperation class](../15_library-reference/3770-the-weboperation-class.md "The com.WebOperation class provides an interface to create and manage the operations of a Genero Web Service."). |
| The `fglwsdl` tool has been enhanced with the following new options:`-b`: Generate code from a WSDL using the binding section instead of the service section`-autoNsPrefix`: Determine the prefix for variables and types based on the XML namespace they belong to`-nsPrefix`: Set the prefix for a variable or a type belonging to the given XML namespaceThe following options have been changed:`-o`: If there are several services in one WSDL, they will be generated in the same file with the given base name instead of returning an error`-disk`: Retrieves and displays all dependencies to the current directory but there are no sub directories any longer.`-prefix`: Accepts patterns %s, %f and %p | See [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD)."). |
| The Genero Web Services library has been enhanced to support WSDL with circular references.The Genero language doesn't provide a way to define variables or types that refer to themselves. However, to provide better interoperability and a way to handle such circular data, the `fglwsdl` tool now generates variables or types of `xml.DomDocument` type when circular references are detected during the processing of WSDL files. This gives the user the ability to manipulate the circular data by hand, using the XML DOM API. | See [The xml package](../15_library-reference/3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces."). |
