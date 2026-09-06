---
title: "SCREEN RECORD fields specification"
source: "fgl-topics/c_fgl_MigI4GL_045.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > SCREEN RECORD fields specification"
type: "concept"
---

# SCREEN RECORD fields specification

> Genero BDL requires a comma separator in the field list of a screen record definition.

With IBM® Informix® 4GL the definition of a screen record in a .per form definition
file, does not require a comma to seperate field names. For example:

```
INSTRUCTIONS
SCREEN RECORD sr1 ( field1 field2 field3 );
END
```

The Genero BDL fglform compiler denies this syntax, and is considered as bad
practice. Always use comma to separate field names in the screen record definition.

## Related links

**Related concepts**  

[Screen records / arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
