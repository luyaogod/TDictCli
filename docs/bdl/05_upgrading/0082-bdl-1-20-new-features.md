---
title: "BDL 1.20 new features"
source: "fgl-topics/fgl_whatsnew_120.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 1.20 new features"
type: "topic"
---

# BDL 1.20 new features

> Features added in 1.20 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 1.20 upgrade guide](0339-bdl-1-20-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 1.20.").

Prior new features guide: [BDL 1.10 new features](0083-bdl-1-10-new-features.md "Features added in 1.10 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| Integrated debugger with gdb syntax to interface with graphical tools like ddd. | See [Integrated debugger](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs."). |
| The program profiler can be used to generate statistics of program execution, to find the bottlenecks in the source code. | See [Program profiler](../13_programming-tools/2622-program-profiler.md "Find out what function is causing the bottleneck in your program."). |
| Internationalize your application in different languages with localized strings, by using the %"string" notation. | See [Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site."). |
| The `TERMINATE REPORT` and `EXIT REPORT` can be used in reports to respectively stop a report from outside of the `REPORT` routine, or stop the report from inside the `REPORT` routine. | See [TERMINATE REPORT](../12_reports/2473-terminate-report.md "The TERMINATE REPORT instruction cancels a report execution."), [EXIT REPORT](../12_reports/2490-exit-report.md "Cancels the report processing.") |
| The `fgl_getversion()` function returns the version number of the runtime system. | See [fgl\_getversion()](../15_library-reference/2766-fgl-getversion.md "Returns the product version number of Genero."). |
| Static arrays can be passed as parameters: all elements are expanded. | See [Static arrays](../08_language-basics/0732-static-arrays.md "Static arrays have a predefined and limited size."). |
| New methods for `StringBuffer` class: `base.StringBuffer.replaceAt()` and `base.StringBuffer.insertAt()`. | See [The StringBuffer class](../15_library-reference/3045-the-stringbuffer-class.md "The base.StringBuffer class is a built-in class designed to manipulate character strings."). |
| Operators equal (= or ==) and not equal (<> or !=) now can be used with records: All record members will be compared. If two members are `NULL` the result of this member comparison results in `TRUE`. | See [RECORD](../08_language-basics/0717-record.md "The RECORD keyword defines a structured type or variable."). |
| New `-W` option for fglform to show warnings. | See [fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs."). |
| `LSTR()` operator, to get a localized string by name. Useful when the localized string identifier is known at runtime only. | See [LSTR() [function]](../08_language-basics/0638-lstr-function.md "The LSTR() operator returns a localized string."). |
| `SFMT()` operator, to format strings with parameter placeholders. Useful to localize application messages with parameters. | See [SFMT() [function]](../08_language-basics/0639-sfmt-function.md "The SFMT() operator replaces place holders in a string with values."). |
| The `base.StringTokenizer` class can be used to parse strings for tokens. | See [The StringTokenizer class](../15_library-reference/3071-the-stringtokenizer-class.md "The base.StringTokenizer class is designed to parse a string to extract tokens based on delimiters."). |
| `CONSTANT` language elements can now be defined as `GLOBALs`. | See [Constants](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values."). |
| The `base.Application` class provides an interface to the program properties. | See [The Application class](../15_library-reference/2968-the-application-class.md "The base.Application class provides a set of utility functions related to the program environment."). |
| Review of the definition of `base.Channel` class, now based on objects. | See [The Channel class](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions."). |

| Overview | Reference |
| --- | --- |
| Interactive instructions support the `UNBUFFERED` mode, to synchronise data model and view automatically: When you set a variable, the value is automatically displayed to the field, and when the user fires and action, the field value is automatically assigned to the corresponding program variable. | See [The buffered and unbuffered modes](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields."). |
| `DISPLAY ARRAY` can now work in paged mode, to avoid loading a large array of rows, with the `ON FILL BUFFER` clause. | See [Paged mode of DISPLAY ARRAY](../11_user-interface/2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY."). |
| Centralize default attributes for actions in `ACTION DEFAULTS`. | See [Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes."). |
| Client side settings can now be saved by application name, with a specific API. By default it is the name of the program. | See [ui.Interface.setName](../15_library-reference/3120-ui-interface-setname.md "Define the name of the current program for the front-end."). |
| New attribute `APPEND ROW = TRUE/FALSE` attribute for the `INPUT ARRAY` instruction, to control the creation of the default `append` action. | See [INPUT ARRAY row modifications](../11_user-interface/2310-input-array-row-modifications.md "Controlling row creation and deletion in an editable record list."). |
| New attribute `KEEP CURRENT ROW = TRUE/FALSE` for the `DISPLAY ARRAY` and `INPUT ARRAY` instructions, to defines if the current row must remain highlighted when leaving the dialog. The default is `FALSE`. | See [Handling the current row](../11_user-interface/2303-handling-the-current-row.md "Query and control the current row in a read-only or editable list of records."). |
| You can now define a `TOOLBAR` in form specification files. | See [TOOLBAR section](../11_user-interface/1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions."). |
| You can now define a `TOPMENU` in form specification files. | See [TOPMENU section](../11_user-interface/1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions."). |
| The `fgl_gethelp()` function returns the help text for the given help number. | See [fgl\_gethelp()](../15_library-reference/2760-fgl-gethelp.md "Reads the current help file, returning help text based on the help identifier."). |
| The `fgl_set_arr_curr()` function changes the current row in `DISPLAY ARRAY` or `INPUT ARRAY`. | See [Handling the current row](../11_user-interface/2303-handling-the-current-row.md "Query and control the current row in a read-only or editable list of records."). |
| Users can now send an interruption event to the program,to stop long running SQL queries, processing loops and reports. | See [User interruption handling](../11_user-interface/2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure."). |
| The `statusBarType` window style attribute to define the statusbar layout. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| The new `FIELD ORDER FORM` option can be used to follow the new `TABINDEX` attribute, to define the field tabbing order. `FIELD ORDER FORM` can also be used at the dialog level as dialog attribute. | See [Defining the tabbing order](../11_user-interface/2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute."). |
| For `COMBOBOX` form items, a default `ITEMS` list is created by fglform when an `INCLUDE` attribute is used. | See [COMBOBOX item type](../11_user-interface/1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values."). |
| The `ON IDLE` clause can be used to execute a block of instructions after a timeout. | See [Get program control if user is inactive](../11_user-interface/2226-get-program-control-if-user-is-inactive.md "Execute some code after a given number of seconds, when the user does not interact with the program."). |
| New logical order of execution for `INPUT ARRAY` triggers:`BEFORE INPUT``BEFORE ROW``BEFORE INSERT``BEFORE FIELD` | See [Editable record list (INPUT ARRAY)](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."). |
| New `ui.ComboBox` class to configure `COMBOBOX` fields at runtime. | See [The ComboBox class](../15_library-reference/3247-the-combobox-class.md "The ui.ComboBox class provides an interface to the COMBOBOX form field view in the abstract user interface tree."). |
| `DISPLAY ARRAY` and `INPUT ARRAY` instructions now automatically use two predefined actions `nextrow` and `prevrow`, which allow binding action views for navigation. | See [Predefined actions](../11_user-interface/2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions."). |
| `ON CHANGE` field trigger can be used to detect field modification. Useful for fields such as `CHECKBOX` and `COMBOBOX`. | See [Reacting to field value changes](../11_user-interface/2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block."). |
| Program icon definition with `ui.Interface.setImage()`. | See [ui.Interface.setImage](../15_library-reference/3119-ui-interface-setimage.md "Defines the icon image of the program."). |
| `LABEL` fields can now have a `FORMAT` attribute. | See [LABEL item type](../11_user-interface/1696-label-item-type.md "Defines a simple text area to display a read-only value."). |
| Front-end function calls allow to execut code on the front-end side with the `ui.Interface.frontCall()` method. | See [Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed."). |
| New ui.Form built-in class to handle forms. | See [The Form class](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction."). |
| New `ON ROW CHANGE` clause in `INPUT ARRAY`, executed when if at least one value in the row has been modified, and the user moves to another row or validates the dialog. The `ON ROW CHANGE` block is executed before the `AFTER ROW` block. | See [ON ROW CHANGE block](../11_user-interface/2018-on-row-change-block.md). |
| `MENU` instruction now supports `ON ACTION` clause, to write abstract menus as simple action handlers. | See [Ring menus (MENU)](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from."). |
| New 'help' predefined action, to start help viewer for `HELP` clauses in dialog instructions. | See [Predefined actions](../11_user-interface/2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions."). |

| Overview | Reference |
| --- | --- |
| SQL Server driver now supports the TINYINT data type. | See [Numeric data types](../10_sql-support/1274-numeric-data-types.md). |
| The fglcomp compiler supports now ANSI outer join syntax in SQL statements (`LEFT OUTER JOIN`), to replace the Informix specific `OUTER()` syntax. | See [SELECT](../10_sql-support/1123-select.md "Produces a result set from a query on database tables."). |
| `FOREACH` that raises an error no longer loops infinitely. | See [FOREACH (result set cursor)](../10_sql-support/1155-foreach-result-set-cursor.md "Processes a series of data rows returned from a database cursor."). |
| New `SQLSTATE` and `SQLERRMESSAGE` registers, to give SQL execution information. | See [SQL error identification](../10_sql-support/0989-sql-error-identification.md "Identify SQL exceptions in your programs with sqlca.sqlcode."). |
