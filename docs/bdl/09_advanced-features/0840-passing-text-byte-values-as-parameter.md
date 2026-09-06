---
title: "Passing TEXT/BYTE values as parameter"
source: "fgl-topics/c_fgl_runtime_stack_params_text_byte.html"
breadcrumb: "Advanced features > Runtime stack > Passing TEXT/BYTE values as parameter"
type: "concept"
description: "BYTE or TEXT data types define large data object (LOB) handlers internally implemented as \"locators\". When you pass a BYTE or TEXT to a function, the locator is pushed on the stack and popped to the ..."
---

# Passing TEXT/BYTE values as parameter

[`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") or [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") data types define large data object
(LOB) handlers internally implemented as "locators".

When you pass a `BYTE` or `TEXT` to a function, the locator is
pushed on the stack and popped to the receiving `BYTE` or `TEXT`
variable in the function. The actual LOB data is not copied, only the locator is passed by
value.

> **Important:**
>
> Since the information of the locator structure is copied (like the file
> name specified with a [`LOCATE IN FILE`](../08_language-basics/0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.")
> instruction). If you modify the locator storage information inside the function with a
> `LOCATE` instruction, the locator in the caller will become invalid. Therefore, only
> read and write the actual data of `BYTE` and `TEXT` parameters in
> functions, do not modify the storage.

Example:

```
MAIN
  DEFINE tx TEXT
  LOCATE tx IN MEMORY
  CALL load_text(tx,arg_val(1))
  DISPLAY "len = ", tx.getLength()
END MAIN

FUNCTION load_text(tx TEXT, fn STRING) RETURNS ()
  CALL tx.readFile(fn)
END FUNCTION
```
