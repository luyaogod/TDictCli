---
title: "fglwinexec.winsavefile()"
source: "fgl-topics/c_fgl_utility_functions_WINSAVEFILE.html"
breadcrumb: "Library reference > Utility modules > fglwinexec: Front-end dialogs module > fglwinexec.winsavefile()"
type: "concept"
---

# fglwinexec.winsavefile()

> Opens a dialog window to get a file path to save data on the front-end workstation.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION winsavefile(
   dirname STRING,
   typename STRING,
   extlist STRING,
   caption STRING )
  RETURNS STRING
```

1. dirname is the default path to be displayed in the dialog
   window.
2. typename is the name of the file type to be saved.
3. extlist is a blank separated list of file extensions
   defining the file type.
4. caption is the label to be saved.

## Usage

This function opens a dialog window to let the user select a file path on the front-end
workstation file system, in order to save the file.

The function returns the file path on success.

The function returns [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.")
if a problem has occurred or if the user canceled the dialog.

The function must be called after the front-end connection has been
established. Do not use this function when using another front-end as the Genero Desktop Client
(GDC).

## Related links

**Related concepts**  

[standard.saveFile](3413-standard-savefile.md "Displays a file dialog window to get a path to save a file on the local file system.")
