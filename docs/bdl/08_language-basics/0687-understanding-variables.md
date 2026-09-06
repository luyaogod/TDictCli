---
title: "Understanding variables"
source: "fgl-topics/c_fgl_variables_002.html"
breadcrumb: "Language basics > Variables > Understanding variables"
type: "concept"
---

# Understanding variables

> This is an introduction to variables.

A variable is a program element that can hold volatile data. The
following list summarizes variables usage:

- Variables are declared with the `DEFINE` instruction or with the
  `VAR` instruction.
- With the `DEFINE` instruction,
  the scope of a variable can be [global, local to a module or
  local to a function](0693-declaration-context.md "A variable can be declared in different contexts, which defines its visibility.").
- With the [`VAR`](0689-var.md "The VAR instruction declares a program variable within a code block.") instruction, the
  scope of the variable is the code block where the instruction is used, inside a function.
- When defined at the module level, a variable can be declare as [`PRIVATE` or `PUBLIC`](0693-declaration-context.md "A variable can be declared in different contexts, which defines its visibility.").
- Variables can be defined as [records](0715-records.md "Records allow structured program variables definitions."), [arrays](0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements."), [dictionaries](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key."),
  etc.
- Variables containing database records are typicaly defined from a database schema with the [`LIKE` clause](0695-database-column-types.md "Simple variables and record structures can be defined from database columns types.").
- Variable can be defined from [user-types](0696-user-defined-types.md "User defined types help to centralize the definition of complex structured data types.").
- After a definition without initializer, variables get [default values](0697-variable-default-values.md "Variables get a default value when defined.") specific to their type.
- Default values (or `NULL`) can be assigned with the `INITIALIZE` instruction.
- In the definition of a variable, you can specify a value as [initializer](0691-variable-initializers.md "Variables can be initialized in their definition."), for simple and structured types.
- In the code, variable assignment is done with the `LET` instruction.
- Database validation rules can be applied with the `VALIDATE` instruction.
- Variables can be used as [SQL parameters](../10_sql-support/1117-using-program-variables-in-static-sql.md "Static SQL syntax supports the usage of program variables as SQL parameters.") or [fetch buffers](../10_sql-support/1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.") in SQL statements.
- Interactive instructions (dialogs) use program variables as [model](../11_user-interface/2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.") to hold the data.
- Variables can be defined with attributes by using the [`ATTRIBUTES()`](0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information.") clause.
