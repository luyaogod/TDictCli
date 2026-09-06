---
title: "Ignored form definition attributes"
source: "fgl-topics/c_fgl_MigI4GL_034.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Ignored form definition attributes"
type: "concept"
description: "Field attributes inherited from the Informix SQL PERFORM syntax should be reviewed for necessity and handling."
---

# Ignored form definition attributes

> Field attributes inherited from the Informix® SQL PERFORM syntax should be reviewed for necessity and handling.

IBM® Informix 4GL
form specification file allows field attributes that are inherited from the Informix SQL PERFORM
syntax:

```
DATABASE FORMONLY 
SCREEN
{
[f01                           ]
}
END
ATTRIBUTES
f01 = FORMONLY.comment,
        QUERYCLEAR,
        NOUPDATE,
        ZEROFILL,
        RIGHT
      ;
END
```

With Genero BDL, use the `-W all` option of fglform to detect
such
attributes:

```
$ fglform -W all form.per
form.per:9:9:9:18:warning:(-8005) Deprecated feature: The QUERYCLEAR attribute is ignored
form.per:10:9:10:16:warning:(-8005) Deprecated feature: The NOUPDATE attribute is ignored
form.per:11:9:11:16:warning:(-8005) Deprecated feature: The ZEROFILL attribute is ignored
form.per:12:9:12:13:warning:(-8005) Deprecated feature: The RIGHT attribute is ignored
```

> **Tip:**
>
> Replace the `RIGHT` attribute by `JUSTIFY=RIGHT`
> Genero BDL attribute.

## Related links

**Related concepts**  

[Form item attributes](../11_user-interface/1753-form-item-attributes.md "The form item attributes reference.")

[fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.")
