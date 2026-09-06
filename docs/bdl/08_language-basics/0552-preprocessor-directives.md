---
title: "Preprocessor directives"
source: "fgl-topics/c_fgl_language_features_preprocessor_directives.html"
breadcrumb: "Language basics > Syntax features > Preprocessor directives"
type: "concept"
---

# Preprocessor directives

> Preprocessor directives can be used in Genero BDL sources.

Genero BDL supports preprocessing instructions, which allow you to write macros and conditional
compilation rules as in the following example:

```
&include "myheader.4gl"
FUNCTION debug( msg )
  DEFINE msg STRING
&ifdef DEBUG
  DISPLAY msg 
&endif
END FUNCTION
```

Use the preprocessor with care, and only when there is no native language solution. Too many
preprocessing directives make the code unreadable and unmaintainable.

## Related links

**Related concepts**  

[Source preprocessor](../13_programming-tools/2562-source-preprocessor.md "A typical preprocessor like in the C language.")
