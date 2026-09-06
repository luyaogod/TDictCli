---
title: "Form specification file syntax"
source: "fgl-topics/c_fgl_MigI4GL_056.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Form specification file syntax"
type: "concept"
---

# Form specification file syntax

> This topic describes syntax differences between I4GL and FGL in .per form specification file.

## Closing curly brace and END keyword

With IBM® Informix® 4GL, the form4gl compiler allows to specify the
`END` keyword just after the closing curly brace of the `SCREEN`
definition:

```
SCREEN
{
   ...
} END
```

This syntax is not allows by the Genero BDL fglform compiler.

## Related links

**Related concepts**  

[Form specification files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
