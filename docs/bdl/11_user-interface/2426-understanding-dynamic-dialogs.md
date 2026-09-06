---
title: "Understanding dynamic dialogs"
source: "fgl-topics/c_fgl_dynamic_dialogs_intro.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Understanding dynamic dialogs"
type: "concept"
---

# Understanding dynamic dialogs

> This section provides basics about dynamic dialogs.

The `ui.Dialog` class can create dialog objects at runtime, to implement generic
code controlling forms that are created at runtime, when the data structure is not known at compile
time.

Unlike static dialog instructions, dynamic dialogs do not require a data model (that is a program
variables containing the values for fields). Dynamic dialogs are created with a list for field
definitions that is built at runtime, no static (RECORD) structure is required to instantiate a
dynamic dialog.

> **Important:**
>
> **Dynamic dialogs are provided to resolve specific needs, like implementing a generic zoom
> window to select a record in a list, and control forms generated at runtime. This feature is not a
> replacement for regular "static" [dialog
> instructions](1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation."), used to control the forms defined in [form
> specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.").**

Dynamic dialogs use by default the `UNBUFFERED` mode and `WITHOUT
DEFAULTS`. For more details, see [Controlling field values](2431-controlling-field-values.md "Fields values in dynamic dialogs can be manipulated dynamically.").

Dynamic dialogs can be used in conjunction with [`base.SqlHandle`](../15_library-reference/3015-the-sqlhandle-class.md "The base.SqlHandle class is a built-in class providing an API to execute parameterized SQL statements, with or without result sets.") objects, to get database table column information in order
to build forms dynamically and define the field list.
