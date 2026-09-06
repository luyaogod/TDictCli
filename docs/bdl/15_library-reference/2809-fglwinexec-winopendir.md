---
title: "fglwinexec.winopendir()"
source: "fgl-topics/c_fgl_utility_functions_WINOPENDIR.html"
breadcrumb: "Library reference > Utility modules > fglwinexec: Front-end dialogs module > fglwinexec.winopendir()"
type: "concept"
---

# fglwinexec.winopendir()

> Opens a dialog window to get a directory path on the front-end workstation.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION winopendir(
   dirname STRING,
   caption STRING )
  RETURNS STRING
```

1. dirname is the default path to be displayed in the dialog window.
2. caption is the label to be displayed.

## Usage

This function opens a dialog window to let the user select a directory path on the front-end
workstation file system.

The function returns the directory path on success.

The function returns [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.")
if a problem has occurred or if the user canceled the dialog.

The function must be called after the front-end connection has been
established. Do not use this function when using another front-end as the Genero Desktop Client
(GDC).

## Related links

**Related concepts**  

[standard.openDir](3408-standard-opendir.md "Displays a file dialog window to get a directory path on the local file system.")
