---
title: "Default action of WHENEVER ANY ERROR"
source: "fgl-topics/c_fgl_Mig0000_034.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > 4GL Programming topics > Default action of WHENEVER ANY ERROR"
type: "concept"
---

# Default action of WHENEVER ANY ERROR

> By default, the WHENEVER ANY ERROR action is to CONTINUE the program flow.

With old Four Js Business Development Suite (BDS) versions like 2.10, expression evaluation
errors such as a division by zero stop the program with an error message. Genero Business
Development Language behaves like IBM® Informix® 4GL and recent BDS versions like 3.55: By default, the
`WHENEVER ANY ERROR` action is to `CONTINUE` the program flow. You can
change this behavior by setting the next FGLPROFILE entry to
true:

```
fglrun.mapAnyErrorToError = true
```

## Related links

**Related concepts**  

[Exceptions](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs.")

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
