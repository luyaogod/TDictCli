---
title: "Interactive instructions"
source: "fgl-topics/c_fgl_intro_BDL_011.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > Interactive instructions"
type: "concept"
---

# Interactive instructions

> Control application forms with interactive instructions that perform field input and action handling.

These interactive instructions allow the program to respond to user actions
and data input. For example the `INPUT BY NAME` block controls a set of form
fields where the user can enter data:

```
DEFINE cust_rec RECORD LIKE customer.*
INPUT BY NAME cust_rec.*
    ...
    BEFORE FIELD cust_name
        ...
    ON ACTION print
        ...
END INPUT
```

## Related links

**Related concepts**  

[Dialog instructions](../11_user-interface/1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation.")
