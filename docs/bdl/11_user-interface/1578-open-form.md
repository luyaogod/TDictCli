---
title: "OPEN FORM"
source: "fgl-topics/c_fgl_windows_and_forms_OPEN_FORM.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > OPEN FORM"
type: "concept"
---

# OPEN FORM

> Declares a compiled form in the program.

## Syntax

```
OPEN FORM identifier FROM form-file
```

1. identifier is an identifier that defines the
   name of the form object.
2. form-file defines the name of the compiled form file, without
   .42f extension.

## Usage

In order to use a .42f compiled version of a form specification file, the
programs must declare the form with the `OPEN FORM` instruction and then display the
form in the current window by using the `DISPLAY FORM` instruction.

`OPEN FORM` / `DISPLAY FORM` are typically used at the beginning of
programs to display the main form in the default `SCREEN`
window:

```
OPEN FORM custform FROM "customer"
DISPLAY FORM custform
```

The form identifier does not need to match the name of the form specification files, but it must
be unique among form names in the program. Its scope of reference is the entire program.

> **Important:**
>
> The form filename identifies the .42f compiled [form file](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.") to be loaded. The filename may use a
> .42f extension, but this is not recommended.
>
> The filename can be a simple filename, a relative file path, or an absolute file path.
>
> - When using a simple filename or a relative path, form files are found relative to several
>   directories in a given order, as described in [the FGLRESOURCEPATH reference topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").
> - When specifying an absolute path, FGLRESOURCEPATH (or DBPATH) is not used. On Windows®, an absolute filename must start with a drive letter
>   ( C: ), a backslash ( \ ) or a slash ( / ).
>   If FGLRESOURCEPATH contains a driver letter ( C: ), a form file specified as
>   "/foo/bar" will only be found in "C:/foo/bar", if
>   C: is the current drive.

If you execute an `OPEN FORM` with the name of an open form, the runtime system
first closes the existing form before opening the new form.

The scope of reference of form identifier is the entire program.

When the window is dedicated to the form, use the `OPEN WINDOW WITH FORM` instruction to create the
window and the form object in one statement.

In TUI mode, the form is displayed in the current window at the position defined by the
`FORM LINE` attribute that can be specified in the `ATTRIBUTE` clause
of `OPEN WINDOW` or as default with the
`OPTIONS` instruction.

After the form is loaded, you can activate the form by executing a `CONSTRUCT`,
`DISPLAY ARRAY`, `INPUT`, `INPUT ARRAY`, or
`DIALOG` statement. When the runtime system executes the `OPEN FORM`
instruction, it allocates resources and loads the form into memory. To release the allocated
resources when the form is no longer needed, the program must execute the `CLOSE FORM` instruction. This is a memory-management
feature to recover memory from forms that the program no longer displays on the screen. If the form
was loaded with a window by using the `WITH FORM` clause, it is automatically closed
when the window is closed with a `CLOSE WINDOW` instruction.

## Example

```
MAIN
  OPEN FORM f1 FROM "customer"
  DISPLAY FORM f1
  CALL input_customer()
  CLOSE FORM f1
  OPEN FORM f2 FROM "custlist"
  DISPLAY FORM f2
  CALL input_custlist()
  CLOSE FORM f2
END MAIN
```

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")

[Defining the position of reserved lines](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.")
