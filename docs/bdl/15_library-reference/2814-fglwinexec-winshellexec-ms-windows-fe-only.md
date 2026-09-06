---
title: "fglwinexec.winshellexec() MS Windows FE Only!"
source: "fgl-topics/c_fgl_utility_functions_WINSHELLEXEC.html"
breadcrumb: "Library reference > Utility modules > fglwinexec: Front-end dialogs module > fglwinexec.winshellexec() MS Windows® FE Only!"
type: "concept"
description: "Opens a document on the workstation where the Windows front-end runs."
---

# fglwinexec.winshellexec() MS Windows FE Only!

> Opens a document on the workstation where the Windows® front-end runs.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION winshellexec(
   filename STRING )
  RETURNS INTEGER
```

1. filename is the file to be opened on the front-end.

## Usage

The function opens a document on the Windows front-end
without waiting.

The function must be called after the front-end connection has been
established. Do not use this function when using another front-end as the Genero Desktop Client
(GDC).
