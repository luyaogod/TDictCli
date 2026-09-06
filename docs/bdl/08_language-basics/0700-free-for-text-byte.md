---
title: "FREE (for TEXT/BYTE)"
source: "fgl-topics/c_fgl_variables_FREE.html"
breadcrumb: "Language basics > Variables > FREE (for TEXT/BYTE)"
type: "concept"
---

# FREE (for TEXT/BYTE)

> The FREE statement releases resources allocated to the specified variable.

## Syntax

```
FREE target
```

1. target is the name of a `TEXT` or
   `BYTE` variable to be freed.

## Usage

When followed by a variable name, the `FREE` statement releases resources
allocated to store the data of [`TEXT`](0569-text.md "The TEXT data type stores large text data.")
and [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") variables.

If the `TEXT`/`BYTE` variable
was located in memory, the runtime system releases the memory.
If the variable was located in a file, the runtime system
deletes the file.

For variables declared in a local scope of reference, the resources are automatically freed by
the runtime system when returning from the function or [`MAIN`](../09_advanced-features/0789-the-main-block-function.md "The MAIN block is the starting point of the program.") block.

After freeing a `TEXT` or `BYTE` variable, it must be
re-configured with a new [`LOCATE`](0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.")
call.

Temporary files created by `LOCATE var IN FILE` (without
specifying an explicit filename) are automatically deleted when the program ends.

## Example

```
MAIN
  DEFINE ctext TEXT
  DATABASE stock
  LOCATE ctext IN FILE "/tmp/data1.txt"
  SELECT col1 INTO ctext FROM lobtab WHERE key=123
  FREE ctext
END MAIN
```

## Related links

**Related concepts**  

[FREE (result set cursor)](../10_sql-support/1154-free-result-set-cursor.md "Releases SQL cursor resources allocated by the DECLARE instruction.")
