---
title: "Preprocessor directive syntax changed"
source: "fgl-topics/c_fgl_Migrate_to_200_preprocessor_char.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Preprocessor directive syntax changed"
type: "concept"
---

# Preprocessor directive syntax changed

> The preprocessor directives use an ampersand character (&) instead of a hash (#) character.

Before version 2.00, the preprocessor directives start with a (`#`) hash
character, to be compliant with standard preprocessors (like **cpp**). This caused too many
conflicts with standard language comments that use the same character:

```
#include "myheader.4gl"
# This is a comment
```

Starting with version 2.00, the preprocessor directives use an ampersand character
(&):

```
&include "myheader.4gl"
FUNCTION debug( msg )
  DEFINE msg STRING
&ifdef DEBUG
  DISPLAY msg 
&endif
END FUNCTION
```

The preprocessor is now integrated in the compiler, to achieve faster compilation.

> **Important:**
>
> To simplify the migration, the `#` hash character is still
> supported when using the `-p` fglpp option of the compiler. However, it is
> recommended that you review your source code and use the & character instead; `#`
> hash will be desupported in a future version.

## Related links

**Related concepts**  

[Source preprocessor](../13_programming-tools/2562-source-preprocessor.md "A typical preprocessor like in the C language.")
