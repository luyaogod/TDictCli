---
title: "FGL_LENGTH_SEMANTICS"
source: "fgl-topics/c_fgl_EnvVariables_FGL_LENGTH_SEMANTICS.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGL_LENGTH_SEMANTICS"
type: "concept"
---

# FGL_LENGTH_SEMANTICS

> Defines the length semantics to be used in programs.

Define the FGL\_LENGTH\_SEMANTICS environment variable to specify byte or character length
semantics, by setting the value to BYTE or CHAR, respectively.

If the variable is not set, byte length semantics will be used by default.

When using a single-byte character set such as ISO-8859-1, use byte length semantics (the
default). If the application character set is UTF-8, it is recommended that you use char length semantics.

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")
