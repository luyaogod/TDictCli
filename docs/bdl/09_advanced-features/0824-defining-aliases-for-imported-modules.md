---
title: "Defining aliases for imported modules"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_aliases.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Defining aliases for imported modules"
type: "concept"
---

# Defining aliases for imported modules

> Use aliases when package path or module names are too long.

When using packages a source directory tree can result in long package paths such as in the
following example. Using the package path and module as prefix in the rest of the code, can
overcrowd the source:

```
IMPORT FGL Financials.GeneralLedger.Models.Account
IMPORT FGL Financials.ProjectLedger.Models.Account
...
CALL Financials.GeneralLedger.Models.Account.validate()
CALL Financials.ProjectLedger.Models.Account.validate()
```

Module aliases can help in such case. Use the `AS` keyword at the end of the
`IMPORT FGL`
instruction:

```
IMPORT FGL Financials.GeneralLedger.Models.Account AS GLAccount
IMPORT FGL Financials.ProjectLedger.Models.Account AS PLAccount
...
CALL GLAccount.validate()
CALL PLAccount.validate()
```

## Related links

**Related concepts**  

[Using [package.]module prefix](0823-using-package-module-prefix.md "Resolve symbol name conflicts with module prefix.")
