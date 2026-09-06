---
title: "fglwinexec.winexec() MS Windows FE Only!"
source: "fgl-topics/c_fgl_utility_functions_WINEXEC.html"
breadcrumb: "Library reference > Utility modules > fglwinexec: Front-end dialogs module > fglwinexec.winexec() MS Windows® FE Only!"
type: "concept"
description: "Executes a program on the workstation where the Windows front-end runs and returns immediately."
---

# fglwinexec.winexec() MS Windows FE Only!

> Executes a program on the workstation where the Windows® front-end runs and returns immediately.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION winexec(
   command STRING )
  RETURNS INTEGER
```

1. command is the command to be executed on the front-end.

## Usage

The function executes the program on the Windows front
end and returns the control to the program without waiting.

The function must be called after the front-end connection has been
established. Do not use this function when using another front-end as the Genero Desktop Client
(GDC).
