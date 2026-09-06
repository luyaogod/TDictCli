---
title: "Oracle DB scroll cursor emulation removal"
source: "fgl-topics/c_fgl_Migrate_to_300_oracle_scroll_emul.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Oracle DB scroll cursor emulation removal"
type: "concept"
description: "The scroll cursor emulation has been removed in the Oracle DB driver."
---

# Oracle DB scroll cursor emulation removal

> The scroll cursor emulation has been removed in the Oracle® DB driver.

Before Genero 3.00, it was possible to enable scrollable cursor emulation (with temporary
files) by defining the following FGLPROFILE entry:

```
dbi.database.mydbname.ora.cursor.scroll.emul = true
```

This feature was supported to workaround an Oracle DB bug in versions 8 and 9i. The bug no longer exist in recent Oracle DB versions and therefore the default native scrollable
cursor feature can be safely used.

If this FGLPROFILE entry is set, the runtime system will print a warning to
stderr.

## Related links

**Related concepts**  

[FGLPROFILE entries for core language](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.")

[Scrollable cursors](../10_sql-support/1021-scrollable-cursors.md "How scrollable cursors can be supported on different databases.")
