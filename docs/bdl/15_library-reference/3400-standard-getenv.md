---
title: "standard.getEnv"
source: "fgl-topics/c_fgl_frontcall_standard_getenv.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.getEnv"
type: "concept"
---

# standard.getEnv

> Returns an environment variable set in the user session on the front end platform.

## Syntax

```
ui.Interface.frontCall("standard", "getEnv",
 [name], [value])
```

1. name - The name of the environment variable.
2. value - The value of the environment variable.

## Usage

The "`getEnv`" front call returns an environment variable set in the user session
on the front-end platform.

This front call is only supported by the GDC front-end, running on a desktop workstation. When
using the GAS, it is not possible to get the value of an environment variable in the context of a
web browser. On mobile devices, this front call is also invalid.s

## Related links

**Related concepts**  

[Environment variables](../07_configuration/0492-environment-variables.md "Genero BDL related environment variables.")
