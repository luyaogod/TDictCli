---
title: "Dictionary in action"
source: "fgl-topics/c_fgl_Dictionary_usage.html"
breadcrumb: "Language basics > Dictionaries > Dictionary in action"
type: "concept"
description: "Defining dictionaries Dictionaries are defined with the DICTIONARY syntax. Dictionaries have no theoretical size limit. The elements of dictionaries are allocated automatically by the runtime system. ..."
---

# Dictionary in action

## Defining dictionaries

Dictionaries are defined with the
`DICTIONARY` syntax. Dictionaries have no theoretical size limit.

The elements
of dictionaries are allocated automatically by the runtime system. Dictionary elements are accessed
by key:

```
MAIN
  DEFINE dict DICTIONARY OF DECIMAL(10,2)
  LET dict["bananas"] = 10540.45
  LET dict["apples"]  =  3487.55
  LET dict["oranges"] =   234.10
  DISPLAY dict.getLength()  -- Shows 3
END MAIN
```

Dictionaries can be initialized in their definitions with [dictionary initializers](0745-dictionary-initializers.md "Dictionaries can be initialized in their definition.").

## Element keys

In a dictionary, elements are identified by a key. Similar to an array index, a key allows access
to the element it references.

A key must be a hashable character string with a given length. For example, `"Mike
Hurn"`, `"AXF98234"`.

Obviously, keys must be unique.

Keys are used in dictionary subscripts. The key value must be enclosed in square
brackets:

```
LET dict[ "the key" ].member = value
```

The `DICTIONARY` class provides following methods related to keys:

- To check if a key exists, use the [`contains()`](../15_library-reference/2958-dictionary-contains.md "Checks if an element with the given key exists in the dictionary.") method.
- To delete an element, use the [`remove()`](../15_library-reference/2962-dictionary-remove.md "Removes an element of the dictionary identified by the key.") method.
- To get all keys of a dictionary, use the [`getKeys()`](../15_library-reference/2961-dictionary-getkeys.md "Returns a dynamic array of all keys of the dictionary.") method.

## Element types

The elements of a dictionary variable are typically defined with simple data types, with a
structured record, or with a user-defined
type:

```
TYPE t_contact RECORD
       name VARCHAR(30),
       address VARCHAR(100),
       birth DATE
     END RECORD
MAIN
  DEFINE contact DICTIONARY OF t_contact
  LET contact["EFC456"].name = "Mike Campbell"
  LET contact["EFC456"].address = "5, Big tree St."
  LET contact["EFC456"].birth = MDY(10,23,1999)
  DISPLAY contact["EFC456"].*
END MAIN
```

Dictionaries can also be combined with another collection type, such as a dynamic array, or
another dictionary:

```
MAIN
    DEFINE d DICTIONARY OF DYNAMIC ARRAY OF DECIMAL(10,2)
    LET d["march"][12] = 456.25
END MAIN
```

## Automatic element allocation

When a dictionary element does not exist, it is automatically allocated before it is used.

For example, when you assign a dictionary element with the [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction by specifying a key that does not exist yet, the new
element is created automatically before assigning the value.

Similarly, when used in an expression, the element will be automatically allocated if it does not
exist, and the element value will evaluate to
`NULL`:

```
MAIN
  DEFINE dict DICTIONARY OF INTEGER
  DISPLAY (dict["unexisting"] IS NULL) -- shows 1 (TRUE)
END MAIN
```

Dictionary elements are also automatically created in a [`FOREACH`](../10_sql-support/1155-foreach-result-set-cursor.md "Processes a series of data rows returned from a database cursor.") loop or when dictionary
elements are used as r-values, for example in a `DISPLAY`.

Consider using the [`contains()`](../15_library-reference/2958-dictionary-contains.md "Checks if an element with the given key exists in the dictionary.") method, to check if a given element key is already existing.

```
MAIN
  DEFINE dict DICTIONARY OF INTEGER
  LET dict["id1"] = 33  -- Created automatically
  DISPLAY dict["id2"]   -- Created automatically
  IF dict.contains("id3") THEN
     DISPLAY dict["id3"]  -- Not displayed as it does not exist
  END IF
END MAIN
```

## Passing and returning dictionaries to functions

Dictionaries are passed (or returned) by reference to/from [functions](0761-functions.md "Describes user defined functions.").

The dictionary can be modified inside the called function, and the caller will see the
modifications.

```
MAIN
  DEFINE dict DICTIONARY OF INTEGER
  CALL fill(dict)
  DISPLAY dict.getLength() -- shows 2
END MAIN

FUNCTION fill(x)
  DEFINE x DICTIONARY OF INTEGER
  LET x["ABC"] = 123
  LET x["DEF"] = 456
END FUNCTION
```

## Comparing dictionaries

`DICTIONARY` variables are not comparable: Comparing two
`DICTIONARY` variable, or using a dictionary in any expression, will produce the
compiler error [-4340](../15_library-reference/4483-genero-bdl-errors.md):

```
MAIN
  DEFINE d1, d2 DICTIONARY OF INTEGER
  LET d1["ABC"] = 123
  LET d2["DEF"] = 456
  DISPLAY "1:", (d1 == d2)
| The variable 'd1' is too complex a type to be used in an expression.
| See error number -4340.
| The variable 'd2' is too complex a type to be used in an expression.
| See error number -4340.
END MAIN
```

When comparing records with a [`rec1.* == rec2.*` expression](0722-comparing-records.md "Records can be compared with the == comparison operator and the .* notation."), if the
record type contains `DICTIONARY` members, the comparison always evaluates to
`FALSE`:

```
MAIN
    DEFINE r1, r2 RECORD
               pkey INTEGER,
               dic DICTIONARY OF STRING
           END RECORD
    LET r1.pkey = 101
    LET r1.dic["aaa"] = "xxx"
    LET r2.pkey = 101
    LET r2.dic["aaa"] = "xxx"
    DISPLAY IIF( r1.* == r2.*, "TRUE", "FALSE" )  -- Shows "FALSE"
END MAIN
```

In order to compare the content of two dictionaries, write a loop to compare all elements
individually.

## Related links

**Related concepts**  

[DICTIONARY methods](../15_library-reference/2956-dictionary-methods.md "DICTIONARY methods")

[Attributes on dictionary definitions](0746-attributes-on-dictionary-definitions.md "Dictionaries can be defined with attributes, to complete the type description.")
