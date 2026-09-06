---
title: "Form compiler checking invalid layout definition"
source: "fgl-topics/c_fgl_Migrate_to_220_invalid_layout_check.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Form compiler checking invalid layout definition"
type: "concept"
---

# Form compiler checking invalid layout definition

> It is better to identify form layout mistakes when the form is compiled, rather than at runtime.

Starting with version 2.20.05, the fglform compiler performs more layout checking than before.
Thus, existing (invalid) forms that compiled with prior versions of Genero may no longer
compile with 2.20.05. This strict checking is done to detect layout mistakes during form
design, instead of having the front-ends render invalid forms in a unknown manner at run
time.

For example, the following form definitions are invalid and will raise a compilation error with
fglform:

```
SCHEMA FORMONLY
LAYOUT
GRID
{
  [f01:  |f02    ]     -- HBox layout tags in lists are denied 
  [f01:  |f02    ]
  [f01:  |f02    ]
  [f01:  |f02    ]
}
END
END
```

```
SCHEMA FORMONLY
LAYOUT
GRID
{
  [f01    ]  [f02    ]
  [f01    ]       [f02    ]   -- Misaligned field tags (vertical)
  [f01    ]  [f02    ]
  [f01    ]       [f02    ]
}
END
END
```

```
SCHEMA FORMONLY
LAYOUT
GRID
{
  [f01][f01][f01]  [f01]   -- Misaligned field tags (horizontal)
}
END
END
```
