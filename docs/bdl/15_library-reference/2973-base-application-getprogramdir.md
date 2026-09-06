---
title: "base.Application.getProgramDir"
source: "fgl-topics/c_fgl_ClassApplication_getProgramDir.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods > base.Application.getProgramDir"
type: "concept"
---

# base.Application.getProgramDir

> Returns the directory path of the current program.

## Syntax

```
base.Application.getProgramDir()
  RETURNS STRING
```

## Usage

This method provides the directory path where the program file used by
fglrun is located.

The directory path is system-dependent.

For example, when starting fglrun myapp.42m in the
/opt/app/bin directory, or when starting the program with fglrun
/tmp/app/bin/myapp.42m, the method will return `"/tmp/app/bin/"`.

## Related links

**Related concepts**  

[FGLAPPDIR](../07_configuration/0519-fglappdir.md "Contains the path to the application directory when executing on a mobile device.")
