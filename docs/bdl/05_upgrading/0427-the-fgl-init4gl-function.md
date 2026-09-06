---
title: "The fgl_init4gl() function"
source: "fgl-topics/c_fgl_Mig0000_026.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > 4GL Programming topics > The fgl_init4gl() function"
type: "concept"
---

# The fgl_init4gl() function

> The fgl_init4js() function has no effect in Genero BDL.

Four Js Business Development Suite (BDS) provided a few utility functions in the
libfgl4js.42x library. This library had to be initialized with a call to
`fgl_init4js()`:

```
MAIN
   ...
  CALL fgl_init4js()
   ...
END MAIN
```

Genero Business Development Language still supports the `fgl_init4js()` function,
but only for backward compatibility. Calling this function has no effect in Genero.
