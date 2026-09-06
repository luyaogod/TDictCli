---
title: "Example 2: Finding rows in an array"
source: "fgl-topics/c_fgl_ext_reflect_example_2.html"
breadcrumb: "Library reference > Extension packages > The reflect package > Examples > Example 2: Finding rows in an array"
type: "concept"
description: "This example shows how to implement a function to find rows in an array. arrays.4gl source: IMPORT reflect PUBLIC CONSTANT FC_EQUALS = \"=\" PUBLIC CONSTANT FC_BEGINS_WITH = \"b\" PUBLIC CONSTANT ..."
---

# Example 2: Finding rows in an array

This example shows how to implement a function to find rows in an array.

arrays.4gl source:

```
IMPORT reflect

PUBLIC CONSTANT FC_EQUALS       = "="
PUBLIC CONSTANT FC_BEGINS_WITH  = "b"
PUBLIC CONSTANT FC_ENDS_WITH    = "e"
PUBLIC CONSTANT FC_LESS_THAN    = "<"
PUBLIC CONSTANT FC_GREATER_THAN = ">"
PUBLIC CONSTANT FC_CONTAINS     = "c"
PUBLIC CONSTANT FC_NOT_CONTAINS = "!c"

PUBLIC TYPE FindCondition RECORD
    field STRING,
    type STRING,
    value STRING,
    ignoreCase BOOLEAN
END RECORD
PUBLIC TYPE FindConditionList DYNAMIC ARRAY OF FindCondition

PRIVATE FUNCTION _conditionMatches(
    value reflect.Value, condition FindCondition INOUT)
    RETURNS BOOLEAN
    DEFINE s, v STRING
    IF value.getType().getKind() != "PRIMITIVE" THEN
        DISPLAY "ERROR: The record members must be of primitive types"
        EXIT PROGRAM 1
    END IF
    LET s = value.toString()
    LET v = condition.value
    IF condition.ignoreCase THEN
        LET s = s.toLowerCase()
        LET v = condition.value.toLowerCase()
    END IF
    CASE condition.type
        WHEN FC_EQUALS
            RETURN s == v
        WHEN FC_BEGINS_WITH
            RETURN s MATCHES (v || "*")
        WHEN FC_ENDS_WITH
            RETURN s MATCHES ("*" || v)
        WHEN FC_CONTAINS
            RETURN s MATCHES ("*" || v || "*")
        WHEN FC_NOT_CONTAINS
            RETURN s NOT MATCHES ("*" || v || "*")
        WHEN FC_LESS_THAN
            RETURN s < v
        WHEN FC_GREATER_THAN
            RETURN s > v
        OTHERWISE
            DISPLAY "ERROR: Invalid condition type"
            EXIT PROGRAM 1
    END CASE
    RETURN FALSE
END FUNCTION

PRIVATE FUNCTION _checkArray( a reflect.Value ) RETURNS ()
    IF a.getType().getKind() != "ARRAY" THEN
        DISPLAY "ERROR: Expecting an array to search in"
        EXIT PROGRAM 1
    END IF
    IF a.getType().getElementType().getKind() != "RECORD" THEN
        DISPLAY "ERROR: The element type of the array must be a record"
        EXIT PROGRAM 1
    END IF
END FUNCTION

PRIVATE FUNCTION _getFieldIndex(t reflect.Type, n STRING) RETURNS INT
    DEFINE x, c INT
    LET c = t.getFieldCount()
    FOR x = 1 TO c
        IF t.getFieldName(x) == n THEN
            EXIT FOR
        END IF
    END FOR
    IF x > c THEN
        DISPLAY "ERROR: Illegal field name: ", n
        EXIT PROGRAM 1
    END IF
    RETURN x
END FUNCTION

FUNCTION search(a reflect.Value, condition FindCondition INOUT, fromIndex INT)
    RETURNS INT
    DEFINE fx, x, n INT
    CALL _checkArray(a)
    LET fx = _getFieldIndex(a.getType().getElementType(), condition.field)
    LET n = a.getLength()
    FOR x = fromIndex TO n
        IF _conditionMatches(a.getArrayElement(x).getField(fx), condition) THEN
            RETURN x
        END IF
    END FOR
    RETURN 0
END FUNCTION

FUNCTION searchMultiCond( a reflect.Value, conditions FindConditionList,
                          all BOOLEAN, fromIndex INT)
    RETURNS INT
    DEFINE fx DYNAMIC ARRAY OF INT
    DEFINE m, x, c, n INT
    DEFINE element reflect.Value
    DEFINE r BOOLEAN
    CALL _checkArray(a)
    LET m = conditions.getLength()
    FOR x = 1 TO m
        LET fx[x] = _getFieldIndex(a.getType().getElementType(), conditions[x].field)
    END FOR
    LET n = a.getLength()
    FOR x = fromIndex TO n
        LET element = a.getArrayElement(x)
        LET r = TRUE
        FOR c = 1 TO m
            IF _conditionMatches(element.getField(fx[c]), conditions[c]) THEN
                IF NOT all THEN RETURN x END IF
            ELSE
                LET r = FALSE
                IF all THEN EXIT FOR END IF
            END IF
        END FOR
        IF r THEN RETURN x END IF
    END FOR
    RETURN 0
END FUNCTION
```

main.4gl source:

```
IMPORT util
IMPORT reflect
IMPORT FGL arrays

TYPE Customer RECORD
    pkey INTEGER,
    name VARCHAR(50),
    address VARCHAR(200)
END RECORD
TYPE CustomerList DYNAMIC ARRAY OF Customer

MAIN
    DEFINE custlist CustomerList = [
            (pkey:101, name:"Mike"),
            (pkey:102, name:"Scott"),
            (pkey:103, name:"Peter"),
            (pkey:155, name:"Bruce"),
            (pkey:299, name:"Phil"),
            (pkey:302, name:"Scott")
        ]
    DEFINE c arrays.FindCondition
    DEFINE cl arrays.FindConditionList
    DEFINE x INT

    LET c.field="name"
    LET c.type=FC_EQUALS
    LET c.value="Scott"
    LET c.ignoreCase=TRUE
    DISPLAY "search 1:", util.JSON.stringify(c)
    LET x = 1
    WHILE x>0
        LET x = arrays.search(reflect.Value.valueOf(custlist), c, x)
        IF x>0 THEN
           DISPLAY "  ",util.JSON.stringify(custlist[x])
           LET x = x + 1
        END IF
    END WHILE

    LET cl[1].field="pkey"
    LET cl[1].type=FC_GREATER_THAN
    LET cl[1].value=100
    LET cl[2].field="name"
    LET cl[2].type=FC_CONTAINS
    LET cl[2].value="e"
    DISPLAY "search 2:", util.JSON.stringify(cl)
    LET x = 1
    WHILE x>0
        LET x = arrays.searchMultiCond(reflect.Value.valueOf(custlist), cl, TRUE, x)
        IF x>0 THEN
           DISPLAY "  ",util.JSON.stringify(custlist[x])
           LET x = x + 1
        END IF
    END WHILE

END MAIN
```

Output:

```
search 1:{"field":"name","type":"=","value":"Scott","ignoreCase":true}
  {"pkey":102,"name":"Scott"}
  {"pkey":302,"name":"Scott"}
search 2:[{"field":"pkey","type":">","value":"100"},{"field":"name","type":"c","value":"e"}]
  {"pkey":101,"name":"Mike"}
  {"pkey":103,"name":"Peter"}
  {"pkey":155,"name":"Bruce"}
```
