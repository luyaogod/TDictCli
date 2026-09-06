---
title: "BDL 1.30 new features"
source: "fgl-topics/fgl_whatsnew_130.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 1.30 new features"
type: "topic"
---

# BDL 1.30 new features

> Features added in 1.30 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 1.30 upgrade guide](0331-bdl-1-30-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 1.30.").

Prior new features guide: [BDL 1.20 new features](0082-bdl-1-20-new-features.md "Features added in 1.20 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| First version of integrated preprocessor using `#` hash syntax for macros. Version 1.32 uses `&` instead | See [Source preprocessor](../13_programming-tools/2562-source-preprocessor.md "A typical preprocessor like in the C language."). |
| Localization support (multibyte character sets). | See [Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules."). |
| The `fglcomp` compiler now adds build information in 42m modules. Compiler version of a 42m module can be checked on site by using the `fglrun` with the `-b` option (line break added for documentation only):$ fglrun -b module.42m 2004-05-17 10:42:05 1.30.2a-620.10 /devel/tests/module.4gl | See [42m module information](../13_programming-tools/2539-42m-module-information.md "Describes how to handle module information in .42m p-code files."). |
| The `fglmkmsg` tool now has the same behavior as other tools like `fglcomp` and `fglform`: If you give only the source file, the message compiler uses the same filename for the compiled output file, adding the .iem extension. | See [Compiling message files (.msg)](../11_user-interface/1596-compiling-message-files-msg.md "The .msg message files must be compiled to .iem binary files, in order to be loaded by the runtime system."). |
| New `BREAKPOINT` instruction to stop a program at a given position when using the debugger. It is ignored when not running in debug mode. | See [Defining a breakpoint in the code](../13_programming-tools/2582-defining-a-breakpoint-in-the-code.md "Set a breakpoint in the program source code with the BREAKPOINT instruction."). |
| New assignment operator `:=` has been added to the language. Assign variables directly within expressions: `IF (i:=(j+1))==2 THEN` | See [Assignment (:=)](../08_language-basics/0651-assignment.md "The := operator assigns a variable with an expression and returns the result."). |
| New `fglcomp` compiler option to detect non-standard SQL syntax: `fglcomp -W stdsql module.4gl` | See [SQL portability](../10_sql-support/1000-sql-portability.md "Writing portable SQL is mandatory, to support different kind of database servers."). |
| New method `base.StringBuffer.replace()`, to replace a sub-string in a string:CALL s.replace("old","new",2)Replaces two occurrences of "old" with "new"... | See [base.StringBuffer.replace](../15_library-reference/3056-base-stringbuffer-replace.md "Replace one string with another."). |
| New methods to read/write complete lines in `base.Channel` built-in class: `readLine()` and `writeLine()`. | See [Read and write text lines](../15_library-reference/3002-read-and-write-text-lines.md). |
| The FGLLDPATH variable is now used during program linking. | See [Compiling source files](../13_programming-tools/2530-compiling-source-files.md "Describes how to build the runtime files from source files."). |
| The linker option `-O` (optimize) is de-supported (was ignored before). You now get a warning if you use this option. | See [Linking programs](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file."). |
| The `[]` array sub-script operator now returns the sub-array:DEFINE a2 DYNAMIC ARRAY WITH DIMENSION 2 OF INTEGER LET a2[5,10] = 123 DISPLAY a2.getLength() -- displays 5 DISPLAY a2[5].getLength() -- displays 10 | See [Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements."). |

| Overview | Reference |
| --- | --- |
| New layout rules and form item attributes provide better control of form design. | See [Form rendering](../11_user-interface/1538-form-rendering.md "The section explains the layout rules to render forms on graphical front-ends."). |
| Decoration attribute can be defined in a presentation style file to set fonts and colors. | See [Presentation styles](../11_user-interface/1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements."). |
| Action defaults can be specified in forms in the `ACTION DEFAULTS` section. | See [ACTION DEFAULTS section](../11_user-interface/1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements."). |
| New `ui.Dialog` built-in class to provide better control over interactive instructions. | See [The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction."). |
| `COMBOBOX` fields now support `UPSHIFT` and `DOWNSHIFT` attributes, to force character case when `QUERYEDITABLE` is used. | See [QUERYEDITABLE attribute](../11_user-interface/1811-queryeditable-attribute.md "The QUERYEDITABLE attribute makes a COMBOBOX field editable during a CONSTRUCT statement."). |
| New presentation style attribute `highlightCurrentRow` for Tables, to indicate if the current row must be highlighted in a specific mode. By default, the current row is highlighted during a `DISPLAY ARRAY`. | See [Table style attributes](../11_user-interface/1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container."). |
| New method `appendElement()` for `ARRAYs`, to append an element at the end of a dynamic array. | See [Array methods](../08_language-basics/0737-array-methods.md "Native BDL arrays and Java arrays can be used to invoke built-in methods."). |
| New assignment operator `:=` has been added to the language. Assign variables directly within expressions: `IF (i:=(j+1))==2 THEN` | See [Assignment (:=)](../08_language-basics/0651-assignment.md "The := operator assigns a variable with an expression and returns the result."). |
| The new method `ui.Dialog.setCellAttributes()` allows you to define colors for each cell of a table. | See [Cell color attributes](../11_user-interface/2313-cell-color-attributes.md "List controllers can display every cell in a specific color."). |
| The `ui.Window` class provides new methods to create or get a form object. | See [ui.Window methods](../15_library-reference/3129-ui-window-methods.md "Methods of the ui.Window class."). |
| When using a dynamic array in `INPUT ARRAY` or `DISPLAY ARRAY`, the number of rows is defined by the size of the dynamic array. The `SET_COUNT()` or `COUNT` attributes are ignored. | See [Controlling the number of rows](../11_user-interface/2302-controlling-the-number-of-rows.md "Methods are provided to set and get the total number of rows in a read-only or editable list of records."). |
| The new form field attribute `TITLE` can be used to specify a table column label with a localized string. | See [TITLE attribute](../11_user-interface/1829-title-attribute.md "The TITLE attribute defines the title of a form item."). |
| New class method `ui.Dialog.setDefaultUnbuffered()` to set the default for the `UNBUFFERED` mode. | See [The buffered and unbuffered modes](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields."). |
| Action defaults are now applied at element creation by the runtime system. In previous versions this was done dynamically by the front-end. Now, changing an action default node at runtime has no effect on existing elements. | See [Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes."). |
| The `DATEEDIT` field type now supports DBDATE/CENTURY settings and the `FORMAT` attribute. | See [FORMAT attribute](../11_user-interface/1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display."). |
| New default action 'close' to control window closing: ON ACTION close | See [Implementing the close action](../11_user-interface/2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). |
| `INPUT ARRAY` using `TABLE` container now needs `FIELD ORDER FORM` attribute to keep tabbing order consistent with visual order of columns. | See [Defining the tabbing order](../11_user-interface/2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute."). |
| New instructions `ACCEPT INPUT` / `ACCEPT CONSTRUCT` / `ACCEPT DISPLAY` to validate a dialog by program.ON ACTION doit ACCEPT INPUT | See [ACCEPT INPUT instruction](../11_user-interface/1951-accept-input-instruction.md), [ACCEPT DISPLAY instruction](../11_user-interface/1998-accept-display-instruction.md), [ACCEPT CONSTRUCT instruction](../11_user-interface/2068-accept-construct-instruction.md). |
| New dialog attribute `ACCEPT` / `CANCEL` to avoid creation of default actions 'accept' and 'cancel'. | See [INPUT instruction configuration](../11_user-interface/1936-input-instruction-configuration.md). |
| New default action 'append' in `INPUT ARRAY`. Allows you to add a row at the end of the list. | See [Default actions in INPUT ARRAY](../11_user-interface/2012-default-actions-in-input-array.md). |
| New method `ui.Window.createForm()` to create an empty form object in order to build forms from scratch at runtime. | See [ui.Window.createForm](../15_library-reference/3131-ui-window-createform.md "Create a new empty form in a window."). |
| `TOPMENU` definition in forms now allows attributes in parenthesis. | See [TOPMENU section](../11_user-interface/1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions."). |
| The form layout syntax now allows you to specify the real width of form items by using a hyphen '-' in the layout tag. | See [Widget width inside hbox tags](../11_user-interface/1559-widget-width-inside-hbox-tags.md). |
| Important remark: Before build 530 the `MENU` has attached the window when returning from the `BEFORE MENU` actions. Since build 530 the `WINDOW` must exist before the `MENU` statement. So now the `Menu` AUI tree node is available in the `BEFORE MENU` block, but a window opened or made current in the `BEFORE MENU` block will NOT be used. |  |
| Layout `GRID` now accepts HBox tags to group items horizontally. | See [Hbox tags](../11_user-interface/1680-hbox-tags.md "Hbox tags group several item tags within the same horizontal layout box, inside a grid-based container (GRID)."). |
| Form `VERSION` attribute to distinguish form revisions. | See [VERSION attribute](../11_user-interface/1846-version-attribute.md "The VERSION attribute is used to specify a user version string for an element."). |
| Form layout `SPACING` attribute to define space between widgets. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| Form `DEFAULT SAMPLE` instruction to define a default sample attribute for all form fields. | See [INSTRUCTIONS section](../11_user-interface/1751-instructions-section.md "The INSTRUCTIONS section is used to define screen arrays, non-default screen records, and global form properties."). |
| New form item attributes: `SAMPLE`, `JUSTIFY`, `SIZEPOLICY` ... | See [SAMPLE attribute](../11_user-interface/1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget."), [JUSTIFY attribute](../11_user-interface/1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers."), [SIZEPOLICY attribute](../11_user-interface/1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") |
| To hide form elements by default, that can be shown by the end user by option, use `HIDDEN=USER` as 'hidden to the user by default'. | See [HIDDEN attribute](../11_user-interface/1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."). |
| Individual table columns now have new attribute `UNMOVABLE` to avoid moving. | See [UNMOVABLE attribute](../11_user-interface/1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table."). |
| `WANTCOLUMNSANCHORED` replaced by `UNMOVABLECOLUMN` and `WANTCOLUMNSVISIBLE` replaced by `UNHIDABLECOLUMNS`. | See [UNMOVABLECOLUMNS attribute](../11_user-interface/1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table."), [UNHIDABLECOLUMNS attribute](../11_user-interface/1831-unhidablecolumns-attribute.md "The UNHIDABLECOLUMNS attribute indicates that the columns of the table cannot be hidden or shown by the user with the context menu."). |
| Tables now accept a `WIDTH` and `HEIGHT` attribute to specify a size. | See [WIDTH attribute](../11_user-interface/1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element."), [HEIGHT attribute](../11_user-interface/1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element."). |
| New `RADIOGROUP` attribute to define the orientation of the radio buttons: `ORIENTATION`. | See [ORIENTATION attribute](../11_user-interface/1805-orientation-attribute.md "The ORIENTATION attribute defines whether an element displays vertically or horizontally."). |
| The `MENU COMMAND` clause now generates action names in lowercase. This means, when you define `COMMAND "Open"`, it will bind to all actions views defined with the name 'open'. | See [COMMAND [KEY()] "option" block](../11_user-interface/1916-command-key-option-block.md). |
| New `ui.Interface.loadTopMenu()` method to load a global `TOPMENU`. | See [ui.Interface.loadTopMenu](../15_library-reference/3116-ui-interface-loadtopmenu.md "Load a default/global topmenu file for all forms of the program."). |
| The `ON CHANGE` block is now invoked when the user clicks on a `CHECKBOX`, `RADIOGROUP`, or changes the item in a `COMBOBOX`. | See [ON CHANGE block](../11_user-interface/1943-on-change-block.md). |
| New `DIALOG` keyword to reference the current dialog as a `ui.Dialog` object. This can be used for example to enable/disable fields during the dialog execution. | See [The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction."). |
| The `ui.Form` built-in class has new methods to handle form elements. The hidden attribute is now also managed at the model level, this allows you to hide form fields by name, instead of using the decoration node.CALL myform.setElementHidden("formonly.field1",2) CALL myform.setFieldHidden("field1",2) -- prefix is optional | See [The Form class](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction."). |
| New methods are provided in `ui.Interface` to control the MDI children. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| In `INPUT ARRAY`, `CANCEL INSERT` now supported in `AFTER INSERT`, to remove the new added line when needed. | See [CANCEL INSERT instruction](../11_user-interface/2037-cancel-insert-instruction.md). |
| `TOOLBAR` and `TOPMENU` elements now have the hidden attribute so you can create them and hide the options the user is not supposed to see.**Important:**Hiding a toolbar or topmenu option does not prevent the use of the accelerator of the action. Use `ui.Dialog.setActionActive()` to disable an action. | See [ui.Form.setElementHidden](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements."). |
| New option `NEXT FIELD CURRENT` to gives control back to the dialog instruction without moving to another field. | See [Giving the focus to a form element](../11_user-interface/2245-giving-the-focus-to-a-form-element.md "How to force the focus to move or stay in a specific form element using program code."). |

| Overview | Reference |
| --- | --- |
| Support for PostgreSQL 7.4 with parameterized queries. | See [PostgreSQL](../10_sql-support/1415-postgresql.md). |
| A MySQL 3.23 driver is now provided for Windows™ platforms (was previously only provided on Linux®). | See [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| The `fglcomp` compiler now converts static SQL updates like:UPDATE tab SET (c1,c2)=(v1,c2) ...to a standard syntax:UPDATE tab SET c1=v1, c2=v2 ... | See [UPDATE](../10_sql-support/1121-update.md "Modifies rows of a database table."). |
| On Windows platforms only, the ix drivers automatically set standard Informix® environment variables with ifx\_putenv(). Values are taken from the console environment with getenv(). Additional variables can be specified with:dbi.stdifx.environment.count = n dbi.stdifx.environment.xx = "variable" | No more available in recent versions. |
