---
title: "Returning dictionaries from functions"
source: "fgl-topics/c_fgl_runtime_stack_return_dictionary.html"
breadcrumb: "Advanced features > Runtime stack > Returning dictionaries from functions"
type: "concept"
description: "When returned by a function, dictionaries are pushed on the stack by reference. Therefore, you can create a dictionary in a function and return it to the caller for usage: MAIN DEFINE dic DICTIONARY ..."
---

# Returning dictionaries from functions

When returned by a function, [dictionaries](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") are pushed on
the stack by reference.

Therefore, you can create a dictionary in a function and return it to the caller for usage:

```
MAIN
  DEFINE dic DICTIONARY OF STRING
  LET dic = create_dictionary(10)
  DISPLAY dic["item3"]
  DISPLAY dic["item10"]
END MAIN

FUNCTION create_dictionary(n INTEGER) RETURNS DICTIONARY OF STRING
  DEFINE i INTEGER
  DEFINE dic DICTIONARY OF STRING
  FOR i=1 TO n
     LET dic[SFMT("item%1",i)] = SFMT("This is item %1",i)
  END FOR
  RETURN dic
END FUNCTION
```

## Related links

**Related concepts**  

[RETURN](../08_language-basics/0676-return.md "The RETURN instruction gives the control of execution back to the caller, optionally returning values on the stack.")
