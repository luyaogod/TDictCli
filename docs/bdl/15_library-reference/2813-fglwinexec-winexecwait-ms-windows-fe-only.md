---
title: "fglwinexec.winexecwait() MS Windows FE Only!"
source: "fgl-topics/c_fgl_utility_functions_WINEXECWAIT.html"
breadcrumb: "Library reference > Utility modules > fglwinexec: Front-end dialogs module > fglwinexec.winexecwait() MS Windows® FE Only!"
type: "concept"
description: "Executes a program on the workstation where the Windows front-end runs and waits for termination."
---

# fglwinexec.winexecwait() MS Windows FE Only!

> Executes a program on the workstation where the Windows® front-end runs and waits for termination.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION winexecwait(
   command STRING )
  RETURNS INTEGER
```

1. command is the command to be executed on the front-end.

## Usage

The function executes the program on the Windows front
end and waits for its termination.

The function must be called after the front-end connection has been
established. Do not use this function when using another front-end as the Genero Desktop Client
(GDC).
