---
title: "fgldialog.fgl_winprompt()"
source: "fgl-topics/c_fgl_utility_functions_FGL_WINPROMPT.html"
breadcrumb: "Library reference > Utility modules > fgldialog: Common dialog functions > fgldialog.fgl_winprompt()"
type: "concept"
---

# fgldialog.fgl_winprompt()

> Displays a dialog box containing a field that accepts a value.

## Syntax

```
FUNCTION fgl_winprompt(
   x INTEGER,
   y INTEGER,
   text STRING,
   default STRING,
   length INTEGER,
   type INTEGER )
  RETURNS STRING
```

1. x is the column position in characters.
2. y is the line position in characters.
3. text is the message shown in the box.
4. default is the default value.
5. length is the maximum length of the input value.
6. type is the data type of the return value.

## Usage

The `fgl_winprompt()` function allows the end user
to enter a value.

This function is provided for backward compatibility, you can also
use your own input dialog with a customized
[form](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.") to get a value
from the user.
Or use the standard [`PROMPT`](../11_user-interface/1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.")
instruction.

Possible values for the type parameter are: `0=CHAR,
1=SMALLINT, 2=INTEGER, 7=DATE, 255=invisible`

Avoid passing `NULL` values.

## Example

```
IMPORT FGL fgldialog
MAIN
  DEFINE answer DATE
  LET answer = fgl_winprompt( 10, 10, "Today", DATE, 10, 7 )
  DISPLAY "Today is " || answer 
END MAIN
```
