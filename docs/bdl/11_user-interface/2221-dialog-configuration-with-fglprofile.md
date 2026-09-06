---
title: "Dialog configuration with FGLPROFILE"
source: "fgl-topics/c_fgl_prog_dialogs_fglprofile.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > Dialog configuration with FGLPROFILE"
type: "concept"
---

# Dialog configuration with FGLPROFILE

> FGLPROFILE parameters can be used to configure dialog behavior.

By setting global parameters in FGLPROFILE, you can control the behavior of all dialogs of the
program. These options are provided as global parameters to define a common pattern for all
dialogs of your application. A complete description is available in the runtime configuration
section.

List of FGLPROFILE entries affecting the behavior of dialogs:

1. `Dialog.fieldOrder` (only used by singular dialogs like INPUT)
2. `Dialog.currentRowVisibleAfterSort`

## The Dialog.fieldOrder entry

```
Dialog.fieldOrder = {true|false}
```

The `Dialog.fieldOrder` FGLPROFILE entry defines the execution of `BEFORE
FIELD` and `AFTER FIELD` triggers of intermediate fields.

When this parameter is set to **true**, as the end user moves to a new field with a mouse
click, the runtime system executes the `BEFORE FIELD` and `AFTER FIELD`
dialog control blocks of the input fields between the source field and the destination field. When
the parameter is set to **false**, intermediate field triggers are not executed.

The `Dialog.fieldOrder` configuration parameter is ignored by the
`DIALOG` multiple-dialog instruction or when using the `FIELD ORDER FORM`
option in singular dialogs such as `INPUT`.

Do not use this feature for new developments: GUI applications allow users to jump from one field
to any other field of the form by using the mouse. Therefore, it makes no sense to execute the
`BEFORE FIELD` and `AFTER FIELD` triggers of intermediate fields in a
graphical application.

> **Important:**
>
> The default setting for the runtime system is **false**; while the default
> setting in FGLPROFILE for `Dialog.fieldOrder` is **true**. As a result, the overall
> setting after installation is **true**. To modify the behavior of intermediate field trigger
> execution, change the setting of `Dialog.fieldOrder` in FGLPROFILE to **false**, or
> use the `FIELD ORDER FORM` program option.

## The Dialog.currentRowVisibleAfterSort entry

```
Dialog.currentRowVisibleAfterSort = {true|false}
```

The `Dialog.currentRowVisibleAfterSort` FGLPROFILE entry controls the visibility
of the current row after a sort in tables

When this parameter is set to **true**, the offset of table page is automatically adapted to
show the current row after a sort. By default, the offset is not changed and current row may not
be visible after sorting the rows of a table. Changing this parameter has no impact on existing
code, it is just an indicator to force the dialog to shift to the page of rows having the current
row, as if the end-user had used the scrollbar. You can use this parameter to get the same behavior
as well known e-mail readers.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
