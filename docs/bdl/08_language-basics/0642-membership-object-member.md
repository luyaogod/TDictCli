---
title: "Membership (object.member)"
source: "fgl-topics/c_fgl_operators_MEMBERSHIP.html"
breadcrumb: "Language basics > Operators > List of expression elements > Associative syntax operators > Membership (object.member)"
type: "concept"
---

# Membership (object.member)

> Separator for object members.

## Syntax

```
setname.element
```

## Usage

The period expression element specifies that its right-hand operand is a member of the set
whose name is its left-hand operand.

This notation is used to reference [`RECORD`](0715-records.md "Records allow structured program variables definitions.") members, [object and class
methods](../09_advanced-features/0942-oop-support.md "Describes Object Oriented Programming basics in the language."), as well as [module elements](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.").

## Example

```
IMPORT FGL customer_module
...
MAIN
  DEFINE rec RECORD
     n INTEGER,
     c CHAR(10)
  END RECORD
  DEFINE form ui.Form
  LET rec.n = 12345
  LET rec.c = "abcdef"
  ...
  CALL form.setElementHidden("page1")
  ...
  CALL customer_module.check(345)
  ...
END MAIN
```

## Related links

**Related concepts**  

[Built-in packages](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
