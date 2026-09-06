---
title: "Returning TEXT/BYTE values from functions"
source: "fgl-topics/c_fgl_runtime_stack_return_text_byte.html"
breadcrumb: "Advanced features > Runtime stack > Returning TEXT/BYTE values from functions"
type: "concept"
description: "When returning a TEXT or BYTE value from a function, the locator is pushed in on the stack. Storage information of the TEXT / BYTE is defined in the locator structure and controlled with the LOCATE ..."
---

# Returning TEXT/BYTE values from functions

When returning a [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") or [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") value from a function, the locator is
pushed in on the stack.

Storage information of the `TEXT`/`BYTE` is defined in the locator
structure and controlled with the [`LOCATE`](../08_language-basics/0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.") instruction. The storage of the large object variable can be defined
in a function, initialize the object with a value, and return it:

```
MAIN
  DEFINE t TEXT
  LET t = init_text(t)
  DISPLAY "len = ", LENGTH(t)
END MAIN

FUNCTION init_text(t TEXT) RETURNS TEXT
  LOCATE t IN MEMORY
  LET t = "abc"
  RETURN t
END FUNCTION
```

The above sample will produce following ouput:

```
len =           3
```

## Related links

**Related concepts**  

[RETURN](../08_language-basics/0676-return.md "The RETURN instruction gives the control of execution back to the caller, optionally returning values on the stack.")
