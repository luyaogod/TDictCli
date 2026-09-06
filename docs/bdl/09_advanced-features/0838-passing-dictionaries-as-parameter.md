---
title: "Passing dictionaries as parameter"
source: "fgl-topics/c_fgl_runtime_stack_params_dictionary.html"
breadcrumb: "Advanced features > Runtime stack > Passing dictionaries as parameter"
type: "concept"
description: "Passing a dictionary as a function parameter is legal and efficient. When passed as parameter, the runtime system pushes a reference of the dictionary on the stack, and the receiving local variables ..."
---

# Passing dictionaries as parameter

Passing a dictionary as a function parameter is legal and efficient. When passed as parameter,
the runtime system pushes a reference of the dictionary on the stack, and the receiving local
variables in the function can then manipulate the original data.

Returning a dictionary from a function is also possible: The runtime system pushes the reference
of the dictionary on the stack.

```
MAIN
  DEFINE dic DICTIONARY OF STRING
  DISPLAY dic.getLength()
  LET dic = init()
  DISPLAY dic.getLength()
  CALL modify(dic)
  DISPLAY dic["first"]
  DISPLAY dic["second"]
  DISPLAY dic.getLength()
END MAIN

FUNCTION init() RETURNS DICTIONARY OF STRING
  DEFINE x DICTIONARY OF STRING
  LET x["first"] = "abc"
  LET x["second"] = "def"
  RETURN x
END FUNCTION

FUNCTION modify(x DICTIONARY OF STRING) RETURNS ()
  LET x["first"] = "xyz"
  LET x["third"] = "ijk"
END FUNCTION
```

Output of the program:

```
          0
          2
xyz
def
          3
```

## Related links

**Related concepts**  

[Dictionaries](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.")
