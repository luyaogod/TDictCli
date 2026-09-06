---
title: "BDL 2.30 new features"
source: "fgl-topics/fgl_whatsnew_230.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 2.30 new features"
type: "topic"
---

# BDL 2.30 new features

> Features added in 2.30 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 2.30 upgrade guide](0261-bdl-2-30-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.30.").

Prior new features guide: [BDL 2.21 new features](0071-bdl-2-21-new-features.md "Features added in 2.21 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| Genero is now available on Mac OS X™. You need at least Mac OS X version 10.5. The Operating System code for Mac OS X 10.5 64-bit is m64x105. | See [Supported operating systems](../04_installation/0034-supported-operating-systems.md "Details of the supported operating systems for the Genero Business Development Language."). |
| Platform identifier is now displayed when using the -V option with command-line tools. | See [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs."). |
| The FGLPROFILE environment variable now accepts multiple file specification with an operating-system-specific path separator. | See [The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files"). |
| The LOAD, UNLOAD and base.Channel class support the "CSV" delimiter specification to read/write files in Comma Separated Value format. | See [LOAD](../10_sql-support/1176-load.md "Inserts data from a file into an existing database table."), [UNLOAD](../10_sql-support/1177-unload.md "Copies data from the database tables into a file.") and [The Channel class](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions."). |
| Version 2.30.04 supports now the `fglrun.arrayIgnoreRangeError` entry which can be set to true to force the runtime system to return the first element of an array when the array index is out of bounds. | See [Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements."). |
| The version 2.30.04 introduces the new `fglrun.mapAnyErrorToError` FGLPROFILE entry. This configuration parameter can be set to true to map the default action of the WHENEVER ANY ERROR exceptions to the action defined for the WHENEVER ERROR exception type. | See [Exceptions](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs."). |

| Overview | Reference |
| --- | --- |
| Drag & Drop support in DISPLAY ARRAY for tables or tree views. | See [The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.") |
| A new form item type called WEBCOMPONENT is provided to integrate external Java-Script-based widgets in your forms. | See [WEBCOMPONENT item type](../11_user-interface/1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component."). |
| New ui.Form class method to make a specific form field visible, showing the parent containers automatically.This method can also be used to bring a given folder page to the front, even if the field is not active (i.e. not driven by a dialog). | See [ui.Form.ensureFieldVisible](../15_library-reference/3149-ui-form-ensurefieldvisible.md "Ensure visibility of a form field.") and [ui.Form.ensureElementVisible](../15_library-reference/3148-ui-form-ensureelementvisible.md "Ensure the visibility of a form element."). |
| The ERROR and MESSAGE instructions get an additional STYLE attribute, to reference a presentation style and define the rendering with font, color, and position. | See [MESSAGE](../11_user-interface/1885-message.md "The MESSAGE instruction displays a message to the user."). |
| New style for TOOLBAR and TOPMENU elements. See Front-End documentation for more details about possible decoration attributes. | See [Toolbars](../11_user-interface/1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.") and [Topmenus](../11_user-interface/1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms."). |
| As with COMBOBOX, the items of a RADIOGROUP are now filled with the values of the INCLUDE attribute, if specified. | See [RADIOGROUP item definition](../11_user-interface/1742-radiogroup-item-definition.md "Defines attributes for a mutually-exclusive set of option fields."). |
| Identify the last clicked CANVAS item with the drawGetClickedItemId() function of fgldraw.4gl. | See [Example 1: Simple canvas](../11_user-interface/2445-example-1-simple-canvas.md). |
| The FIELD\_TOUCHED() operator and ui.Dialog.getFieldTouched() method accept now a simple star as parameter, in order to check all fields used by the dialog. | See [FIELD\_TOUCHED() [function]](../08_language-basics/0673-field-touched-function.md "The FIELD_TOUCHED() operator checks if fields were modified during the dialog execution.") and [ui.Dialog.getFieldTouched](../15_library-reference/3198-ui-dialog-getfieldtouched.md "Returns the modification flag for a field."). |
| The JUSTIFY attribute is now supported for all form item types, in order to let you specify both the data justification in the field/cell and the alignment of the table column header. | See [JUSTIFY attribute](../11_user-interface/1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."). |
| The ui.Dialog.setFieldActive() method takes now a list of fields as parameter, with the "dot-asterisk" notation, like the setFieldTouched() method. | See [ui.Dialog.setFieldActive](../15_library-reference/3226-ui-dialog-setfieldactive.md "Enable and disable form fields."). |
| This new feature is part of the fix for bug #18224.When modifying a tree during the dialog execution (for example, when implementing dynamic trees with ON EXPAND / ON COLLAPSE triggers), if you use the ui.Dialog.insertRow(), ui.Dialog.deleteRow() or ui.Dialog.deleteAllRows() methods to modify the node list, the internal tree structure was corrupted. The program array can be safely modified directly with array methods, but multi-range selection flags and cell attributes are not synchronized when doing this. Starting with 2.30.02, you can now use the ui.Dialog.insertNode(), ui.Dialog.appendNode() and ui.Dialog.deleteNode() methods to manipulate the node list and get additional data like row selection flags and cell attributes synchronized. | See [The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction."). |
| Upgrade notes for presentations styles. | See [Presentation styles changes](0271-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |

| Overview | Reference |
| --- | --- |
| New database drivers | List of new database drivers:dbmase0Fx for SAP ASE 15.x (2.30.01) Warning: SAP ASE is now desupported in all BDL versions.dbmmys55x for a [Mysql 5.5.x client](../10_sql-support/1314-oracle-mysql-mariadb.md) (2.30.01)dbmpgs90x for a [PostgreSQL 9.0.x client](../10_sql-support/1415-postgresql.md) (2.30.02) |
| Informix® SMALLFLOAT and FLOAT can now be stored in Oracle native BINARY\_FLOAT / BINARY\_DOUBLE types. | See [Oracle Database](../10_sql-support/1359-oracle-database.md). |
| The LOAD, UNLOAD and base.Channel class support the "CSV" delimiter specification to read/write files in Comma Separated Value format. | See [LOAD](../10_sql-support/1176-load.md "Inserts data from a file into an existing database table."), [UNLOAD](../10_sql-support/1177-unload.md "Copies data from the database tables into a file.") and [The Channel class](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions."). |
| Use the fgl\_db\_driver\_type() built-in function to identify the target database type. | See [fgl\_db\_driver\_type()](../15_library-reference/2737-fgl-db-driver-type.md "Returns the 3-letter identifier/code of the current database driver."). |
| In order to identify the reason why a database driver cannot be loaded, when setting FGLSQLDEBUG you now get an additional debug message that contains the operating system error message (dlerror()) | See [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions."). |
| The fgldbsch tool can now extract database schema from SQLite. However, pay attention to the data types used in SQLite (V 3.6): This database supports some standard type names in the SQL syntax but in reality the types used to store data are very limited. For example, a DATE will be stored as an integer or string (i.e. there is no native DATE type). See SQLite documentation for more details.The fgldbsch tool will extract the schema based on the original type names used to create the table. | See [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database."). |
