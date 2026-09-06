---
title: "When does type conversion occur?"
source: "fgl-topics/c_fgl_DataConversions_002.html"
breadcrumb: "Language basics > Type conversions > When does type conversion occur?"
type: "concept"
---

# When does type conversion occur?

> In Genero BDL, primitive data type conversion is implicit when possible.

The runtime system performs data conversion implicitly without objection, as long as the data
conversion is valid. A date value can be converted to a character string, but a character
string can only be converted to a date if the string represents a valid date in the current
date format settings (DBDATE).

Implicit data type conversion can for example occur in the following cases:

- In a [`LET`](0701-let.md "The LET statement assigns values to variables.") assignment,
- In an [expression](0592-expressions.md "Shows the possible expressions supported in the language."), when operands are not of
  the same data type,
- In `DISPLAY` instructions, or [`PRINT`](../12_reports/2491-print.md "Formats and prints a row of data in a report routine.") instructions in reports,
- In [dialogs](../11_user-interface/1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation."), when values must be converted
  to strings to be displayed in form fields,
- When passing and returning values to/from a [function](0761-functions.md "Describes user defined functions."),
- When serializing numeric values in [`UNLOAD`](../10_sql-support/1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database."), [JSON methods](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data."), etc.

In the next code example, implicit data type conversion occurs

1. When assigning the result of the `DECIMAL` expression to the
   `VARCHAR` variable `v`,
2. When assigning a `VARCHAR` value to the `DECIMAL` variable
   `d`,
3. When passing the `DECIMAL` value `d` to function
   `func()`, expecting a `VARCHAR`,
4. When returning the `VARCHAR` value from the `func()`
   function,
5. When displaying the `DECIMAL` value (formatting rules apply).

```
MAIN
  DEFINE v VARCHAR(50),
         d DECIMAL(10,2)
  LET v = 1234.50 * 2   -- 1.
  LET d = v             -- 2.
  LET d = func(d)       -- 3. and 4.
  DISPLAY d             -- 5.
END MAIN

FUNCTION func(v)
  DEFINE v VARCHAR(50)
  DISPLAY v
  RETURN v   -- 4.
END FUNCTION
```

## Related links

**Related concepts**  

[Runtime stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions.")
