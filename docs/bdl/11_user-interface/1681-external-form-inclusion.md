---
title: "External form inclusion"
source: "fgl-topics/c_fgl_FormSpecFiles_form_inclusion.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > External form inclusion"
type: "concept"
---

# External form inclusion

> Form inclusion allows to reuse the same form part in different forms.

In some cases, application forms can become very complex, or can have a common layout part that
repeats across forms. In such case, some parts of the form can be defined in an external
.per file, that will be included in the final form by using the [`FORM` clause](1716-form-clause.md "Reuse the definition of a form in the current form.") inside the
`LAYOUT` section.

The external form parts can be controlled by a declarative dialog instruction, that can be
attached to any procedural dialog instruction, with the `SUBDIALOG` clause of
`DIALOG`.

```
LAYOUT
  VBOX
    GRID g1
    {
       Customer information
       Name:  [f001                 ]
       ...
    }
    END
    FORM "orders"
  END
END
```

## Related links

**Related concepts**  

[The SUBDIALOG clause](2087-the-subdialog-clause.md "The SUBDIALOG clause")
