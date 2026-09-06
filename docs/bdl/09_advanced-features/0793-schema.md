---
title: "SCHEMA"
source: "fgl-topics/c_fgl_DatabaseSchema_SCHEMA.html"
breadcrumb: "Advanced features > Database schema > SCHEMA"
type: "concept"
---

# SCHEMA

> Defines the database schema files to be used for compilation.

## Syntax 1: Database schema specification

```
SCHEMA dbname
```

1. dbname identifies the database schema file (.sch).

## Syntax 2: Database schema and default connection specification

```
[DESCRIBE] DATABASE dbname
```

1. dbname identifies the database schema file (.sch).

## Usage

The `SCHEMA dbname` instruction defines the database schema to
be used for compilation, where dbname identifies the name of the [database schema files (.sch)](0794-structure-of-database-schema-files.md "A database schema is composed by three files (.sch, .val, .att)").

The `[DESCRIBE] DATABASE` instruction defines the compilation
database schema, and the default connection for the `MAIN` block when the program
starts.
> **Tip:**
>
> Instead of `[DESCRIBE] DATABASE`, use the
> `SCHEMA` instruction: The `SCHEMA` instruction defines only the
> compilation database schema and allows to use a database schema name different from the connection
> database name.

The dbname database name must be expressed explicitly; it cannot be a variable
as in a `DATABASE` instruction inside a program block.

Use the `SCHEMA` instruction outside any program block, before a variable
declaration with [`DEFINE LIKE`](../08_language-basics/0695-database-column-types.md "Simple variables and record structures can be defined from database columns types.")
instructions. `SCHEMA` must precede any program block in each module that includes a
`DEFINE ... LIKE` declaration or [`INITIALIZE ... LIKE`](../08_language-basics/0698-initialize.md "The INITIALIZE instruction initializes program variables with NULL or default values.") and [`VALIDATE...LIKE`](../08_language-basics/0702-validate.md "The VALIDATE instructions checks a variable value based on database schema validation rules.") statements. It must
also precede any `DEFINE ... LIKE` declaration of module variables.

Database schema information such as data types for `DEFINE ... LIKE` are
taken from the schema files during compilation. Make sure that the database schema file of the
[development database](../03_general/0009-general-terms-used-in-this-documentation.md "This documentation uses general terms that must be clarified for a good understanding.") corresponds to the [production database](../03_general/0009-general-terms-used-in-this-documentation.md "This documentation uses general terms that must be clarified for a good understanding."); otherwise the program variables defined
in the p-code modules will not match the table structures of the production database.

For backward compatibility with IBM®
Informix® , dbname can be written
with different syntaxes. You can specify an Informix server, or a more complex Informix source
with a string like `"//server/database"`. Such database schema specification is
not recommended, as it prevents use of database type other than Informix:

```
  database
| database @ server
| "string"
```

When using a simple identifier for the database name, the compiler converts the name to
lowercase, before searching the schema file. However, if a double quoted string is used as
database name, the name will be used as is to find the schema file.

With the `SCHEMA` instruction, the name of the database schema during
development can be different from the name of the database source used at runtime.

To handle uppercase characters in the database name you must quote the name: `SCHEMA
"myDatabase"`

## Example

```
SCHEMA shop -- Compilation database schema
DEFINE rec RECORD LIKE customer.*
MAIN
   DATABASE shop -- Runtime database specification
   SELECT * INTO rec.* FROM customer WHERE custno=1
END MAIN
```

## Related links

**Related concepts**  

[DATABASE](../10_sql-support/1096-database.md "Opens a new database connection in unique-session mode.")
