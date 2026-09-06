---
title: "VALIDATE"
source: "fgl-topics/c_fgl_variables_VALIDATE.html"
breadcrumb: "Language basics > Variables > VALIDATE"
type: "concept"
---

# VALIDATE

> The VALIDATE instructions checks a variable value based on database schema validation rules.

## Syntax

```
VALIDATE target [,...] LIKE
   {
      table.*
   |
      table.column
   }
```

1. target is the name of the variable to be validated.
2. If target is a record, you can use the star notation to validate all members
   in the record.
3. table.column can be any column reference defined in the database schema.

## Usage

The `VALIDATE` statement tests whether the value of the specified variable is
within the range of values for a corresponding column in [.val database schema file](../09_advanced-features/0796-column-validation-file-val.md "The .val database schema file holds functional and display attributes of database table columns.")
referenced by a [`SCHEMA`](../09_advanced-features/0793-schema.md "Defines the database schema files to be used for compilation.") clause.
If the value does not match any value defined in the `INCLUDE` attribute of the
corresponding column, the runtime system raises error [-1321](../15_library-reference/4483-genero-bdl-errors.md).

The argument of the `VALIDATE` instruction can be a simple variable, a record,
or an array element. If the target is a record, you can use the dot-star (`.*`)
notation to reference all [record](0715-records.md "Records allow structured program variables definitions.") members in the validation,
or specify a range of record members with the [`THRU`](0724-thru-through.md "The THRU keyword can be used to specify a set of members of a record.") clause.

## Example

```
SCHEMA stores
MAIN
  DEFINE addr LIKE customer.address1
  LET addr = "aaa"
  VALIDATE addr LIKE customer.address1
END MAIN
```

## Related links

**Related concepts**  

[Form-level validation rules](../11_user-interface/2240-form-level-validation-rules.md "Form-level validation rules can be defined for each field controlled by a dialog.")
