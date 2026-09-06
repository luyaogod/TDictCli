---
title: "WITH FORM clause"
source: "fgl-topics/c_fgl_windows_and_forms_009.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > OPEN WINDOW > WITH FORM clause"
type: "concept"
---

# WITH FORM clause

> Creating a window object with a form.

The `WITH FORM` clause can be used to specify the name of a compiled [form file](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.") to be used to create a window. A window object is
automatically opened and sized to the screen layout of the form.

```
OPEN WINDOW w1 WITH FORM "custlist"
```

When using the [TUI mode](1517-text-mode-rendering-tui-mode.md), the width of the window is
from the left-most character on the screen form (including leading blank spaces) to the right-most
character on the screen form (truncating trailing blank spaces). The length of the window is
calculated as (form line) + (form length).

It is recommended that you use the `WITH FORM` clause, especially in the default
[GUI mode](1518-graphical-mode-rendering-gui-mode.md), because the window is created in accordance
with the form. If you use this clause, you do not need the [`OPEN FORM`](1578-open-form.md "Declares a compiled form in the program."), [`DISPLAY FORM`](1579-display-form.md "Displays and associates a form with the current window."), or
`CLOSE FORM` statement to
open and close the form. The `CLOSE WINDOW` statement closes the window and the form.

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
