---
title: "BDL 2.20 new features"
source: "fgl-topics/fgl_whatsnew_220.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 2.20 new features"
type: "topic"
---

# BDL 2.20 new features

> Features added in 2.20 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 2.20 upgrade guide](0280-bdl-2-20-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.20.").

Prior new features guide: [BDL 2.11 new features](0073-bdl-2-11-new-features.md "Features added in 2.11 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| The Java Interface allows your programs to use the Java library. | See [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs."). |
| New TINYINT, BIGINT and BOOLEAN data types. | See [TINYINT](../08_language-basics/0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers."), [BIGINT](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers."), [BOOLEAN](../08_language-basics/0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE."). |
| Private functions: It is now possible to hide a function (or report) to the other modules with the new PRIVATE keyword. | See [Understanding functions](../08_language-basics/0762-understanding-functions.md "This is an introduction to functions."). |
| Automatic source documentation generator. | See [Source documentation](../13_programming-tools/2546-source-documentation.md "Explains how to automatically generate documentation from your sources."). |
| The fglcomp compiler has been extended with a new option (`--timestamp`) to write the compilation timestamp to the generated 42m p-code module. If present, the timestamp will be printed when using `fglrun -b`. Use compilation timestamps only if really needed; every new compiled .42m module will be different, even if the source code has not changed. | See [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks."). |
| The FGLRESOURCEPATH environment variable to define search paths for program resource files like forms. | See [FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files."). |
| New precision math built-in functions for DECIMAL data. | See [fgl\_decimal\_truncate()](../15_library-reference/2738-fgl-decimal-truncate.md "Returns a decimal truncated to the precision passed as parameter."), [fgl\_decimal\_sqrt()](../15_library-reference/2739-fgl-decimal-sqrt.md "Computes the square root of the decimal passed as parameter."), [fgl\_decimal\_exp()](../15_library-reference/2740-fgl-decimal-exp.md "Returns the value of Euler's constant (e) raised to the power of the decimal passed as parameter."), [fgl\_decimal\_logn()](../15_library-reference/2741-fgl-decimal-logn.md "Returns the natural logarithm of the decimal passed as parameter."), [fgl\_decimal\_power()](../15_library-reference/2742-fgl-decimal-power.md "Raises decimal to the power of the real exponent."). |
| Automatic Code Completion with VIM: If you have Vim 7 installed, you can now use .per and .4gl code completion. | See [Source code edition](../13_programming-tools/2540-source-code-edition.md "Simple helper to better render sources in configurable text editors."). |

| Overview | Reference |
| --- | --- |
| The START REPORT instruction now allows to specify the XML SAX Document Handler to process XML output with the TO XML HANDLER clause. | See [TO XML HANDLER syntax](../12_reports/2463-xml-output-for-reports.md "For better integration with external tools based on XML standards, reports can produce XML output."). |
| Report definition file generation with fglcomp --build-rdd option. | See See [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks."). |

| Overview | Reference |
| --- | --- |
| Support for typical Tree-View widgets with the new TREE container. | See [Tree views](../11_user-interface/2352-tree-views.md "Describes how to implement tree views."). |
| The traditional user interface mode: To simplify migration from Informix® 4GL or Four Js BDS, you can now run applications in traditional mode to render windows as simple boxes, as in the WTK front-end. | See [Graphical mode with Traditional Display](../11_user-interface/1519-graphical-mode-with-traditional-display.md). |
| Phantom form fields can be used to define the screen-record or screen-array, but are not used in the LAYOUT section of the form. Phantom fields are especially useful when implementing a TREE container. | See [Phantom fields](../11_user-interface/1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field)."). |
| Multi-row selection allows end users to highlight several rows in a list of records. | See [Syntax of DISPLAY ARRAY instruction](../11_user-interface/1961-syntax-of-display-array-instruction.md "The DISPLAY ARRAY instruction controls the display of a program array on the screen."). |
| Built-in sort works now in INPUT ARRAY. | See [Sorting rows in a list](../11_user-interface/2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required."). |
| New `contextMenu` action default attribute to allow you to specify whether the menu option is visible in the default context menu. The default value is "yes" - the option is visible whenever the action is visible. | See [Action defaults files](../11_user-interface/1600-action-defaults-files.md "Action defaults files allow to centralize action configuration parameters such as text, icon, accelerators and behavior options in XML format."). |
| New `integratedSearch` presentation style attribute for TEXTEDIT fields to enable text search. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| FOLDER elements can now use a "`position`" style attribute to define the position (top, left, right, bottom) of folder tabs. | See [Folder style attributes](../11_user-interface/1638-folder-style-attributes.md "Folder presentation style attributes apply to FOLDER tab elements."). |
| BUTTON form items get a new "`buttonType`" attribute to define the rendering of the button. | See [Button style attributes](../11_user-interface/1631-button-style-attributes.md "Button presentation style attributes apply to BUTTON elements."). |
| MENU object created with the popup option can be placed with the "`position`" style attribute. | Desupported in V4, see [Presentation style changes in V4.00](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |
| Window Menu and Action panel decoration can be customized using the new "`ringMenuDecoration`", "`actionPanelDecoration`" style attributes. | See [Window style attributes: Ring Menu](../11_user-interface/1658-window-style-attributes-ring-menu.md "Presentation style attributes that apply to a window ring menu."), [Window style attributes: Action Panel](../11_user-interface/1657-window-style-attributes-action-panel.md "Presentation style attributes that apply to the window action panel."). |
| The new "`tabbedContainer`", "`tabbedContainerCloseMethod`" style attributes can be used to turn on and customize tabbed WCI containers. | Deprecated in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| TABLE elements can use the new "`tableType`" attribute to render data in different ways. The new "`resizeFillsEmptySpace`" attribute can be used to define how the last column is resized when the table is resized. | See [Table style attributes](../11_user-interface/1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container."). |
| All items with an `IMAGE` attribute can use the new "`imageCache`" attribute to define if the picture can be cached locally on the front-end. | Desupported in GDC 3.20. See [Presentation style changes in BDL 3.20](0161-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |
| New Front-End Functions "`getWindowId`", "`feInfo`", "`launchURL`". | See [Standard front calls](../15_library-reference/3387-standard-front-calls.md "Standard front call functions provide common utility APIs to control the front-end."). |
| Front-End protocol compression can now be disabled with a new FGLPROFILE entry. This is especially useful in fast networks to save processor time. | This feature is desupported, see [Presentation styles changes in V4.00](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |
| New built-in functions are now available to control the part of the text that is selected in the current field. | See [fgl\_dialog\_getselectionend()](../15_library-reference/2754-fgl-dialog-getselectionend.md "Returns the position of the last selected character in the current field."), [fgl\_dialog\_setselection()](../15_library-reference/2756-fgl-dialog-setselection.md "Selects the text in the current field."). |
| New IMAGE attribute in form LAYOUT element: The LAYOUT section of a form definition can now use the IMAGE attribute to define the icon to be used for the parent Window. This is especially useful in a Container-based application, to distinguish child programs inside the WCI container. | See [LAYOUT section](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers."). |
| Use the new INFIELD clause in ON ACTION interactive block to automatically enable/disable the action when entering/leaving the specified field. | See [Field-specific actions (INFIELD clause)](../11_user-interface/2285-field-specific-actions-infield-clause.md "Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus."). |
| Getting the current active dialog with ui.Dialog.getCurrent(). | See [ui.Dialog.getCurrent](../15_library-reference/3176-ui-dialog-getcurrent.md "Returns the current dialog object."). |

| Overview | Reference |
| --- | --- |
| New database drivers. | List of new database drivers:dbmsqt3xx for an [SQLite 3 library](../10_sql-support/1466-sqlite.md) (2.20.01) |
| MySQL Driver supports TEXT/BYTE data types. | See [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| To work around conflicts with the Informix database path specification in DBPATH, use the FGLRESOURCEPATH environment variable. | See [FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files."). |
| Database user authentication callback function can be used to specify a database user and password when the DATABASE instruction cannot be replaced by CONNECT TO. | See [User authentication callback function](../10_sql-support/1093-user-authentication-callback-function.md). |
| FGLSQLDEBUG output is improved to display and SQL command header with SQL command name and source/line information before executing the underlying ODI driver code. If the driver code crashes or stops the process with an assertion, you can easily identify the last SQL instruction that was executed. | See [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions."). |

| Overview | Reference |
| --- | --- |
| The Genero Web Services XML Library has been improved to support the XML-Signature and XML-Encryption specifications defined by the W3C (also known as XML-Security).The library enables BDL applications to handle public, private, symmetric or hmac keys and X509 certificates in order to sign XML documents or document fragments, and verify a XML signature against a certificate or key. It also enables the applications to encrypt XML nodes using symmetric keys, and decrypt them back using DOM manipulation. Combined with the COM library, any BDL application can now exchange any XML documents over the Internet in a completely secured manner.The library provides classes for:Manipulating cryptography keysHandling X509 certificates for identificationEncrypting and decrypting XML documents, document fragments, or symmetric keysSigning XML documents, document fragments, or any kind of data, and validating them againstXML signatures | See [XML security classes](../15_library-reference/4190-xml-security-classes.md "XML Security classes handle encryption and signature of XML documents entirely in memory with keys and certificates."). |
| The Genero Web Services XML library provides APIs to encrypt and decrypt strings with symmetric or RSA public/private keys. These APIs can be used to encrypt/decrypt passwords directly in BDL applications. | See [The Encryption class](../15_library-reference/4309-the-encryption-class.md "The xml.Encryption class provides methods to encrypt and decrypt XML documents, nodes or symmetric keys.") and [fglpass](../13_programming-tools/2525-fglpass.md "The fglpass tool allows you to encrypt passwords."). |
| The Genero Web Services provides support for the new `BOOLEAN`, `TINYINT` and `BIGINT` data types.You can use these data types when writing your web service or to customize your BDL RECORDs for XML serialization. The `fglwsdl` tool has been enhanced to generate these new data types automatically when encountered in WSDL files or XML schemas.**Note:**For compatibility issues, the `fglwsdl` tool allows code generation without these new data types by using the option '`-legacyTypes`'. | See [XML serialization rules and customization](../16_web-services/4958-xml-serialization-rules-and-customization.md) and [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD)."). |
