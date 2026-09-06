---
title: "Example 1: Record introspection and modification"
source: "fgl-topics/c_fgl_ext_reflect_example_1.html"
breadcrumb: "Library reference > Extension packages > The reflect package > Examples > Example 1: Record introspection and modification"
type: "concept"
description: "This example describes a RECORD variable with the reflection API. records.4gl source: IMPORT reflect PRIVATE FUNCTION _check_record(val reflect.Value) RETURNS () IF val.getType().getKind() != \"RECORD\" ..."
---

# Example 1: Record introspection and modification

This example describes a `RECORD` variable with the reflection API.

records.4gl source:

```
IMPORT reflect

PRIVATE FUNCTION _check_record(val reflect.Value) RETURNS ()
    IF val.getType().getKind() != "RECORD" THEN
        DISPLAY "ERROR: Expecting a RECORD"
        EXIT PROGRAM 1
    END IF
END FUNCTION

PRIVATE FUNCTION _show(rec reflect.Value, spaces STRING) RETURNS ()
    DEFINE tp reflect.Type
    DEFINE fv reflect.Value
    DEFINE fc, x SMALLINT
    DEFINE s STRING
    CALL _check_record(rec)
    LET tp = rec.getType()
    LET fc = tp.getFieldCount()
    FOR x = 1 TO fc
        LET fv = rec.getField(x)
        CASE fv.getType().getKind()
        WHEN "PRIMITIVE"
            LET s = fv.toString()
            DISPLAY spaces, tp.getFieldName(x), " : ", NVL(s,"NULL")
        WHEN "RECORD"
            DISPLAY spaces, tp.getFieldName(x), ":"
            CALL _show(fv, spaces||" ")
        WHEN "ARRAY"
            DISPLAY spaces, tp.getFieldName(x), ": array"
        WHEN "DICTIONARY"
            DISPLAY spaces, tp.getFieldName(x), ": dictionary"
        END CASE
    END FOR

END FUNCTION

PUBLIC FUNCTION show(rec reflect.Value) RETURNS ()
    CALL _show(rec, " ")
END FUNCTION

PUBLIC FUNCTION set_undef(rec reflect.Value) RETURNS ()
    DEFINE tp reflect.Type
    DEFINE fv reflect.Value
    DEFINE fc, x SMALLINT
    CALL _check_record(rec)
    LET tp = rec.getType()
    LET fc = tp.getFieldCount()
    FOR x = 1 TO fc
        LET fv = rec.getField(x)
        CASE fv.getType().getKind()
        WHEN "PRIMITIVE"
            CASE
            WHEN fv.getType().toString() MATCHES "*CHAR*"
                CALL fv.set( reflect.Value.copyOf("<undefined>") )
            END CASE
        WHEN "RECORD"
            CALL set_undef(fv)
        END CASE
    END FOR
END FUNCTION
```

main.4gl
source:

```
IMPORT reflect
IMPORT FGL records

DEFINE rec_cust RECORD
           pkey INTEGER,
           name VARCHAR(40),
           creation DATETIME YEAR TO SECOND,
           address RECORD
               street VARCHAR(200),
               pcode VARCHAR(5),
               city VARCHAR(50)
           END RECORD
       END RECORD
    
TYPE Circle RECORD
           x FLOAT,
           y FLOAT,
           r FLOAT
       END RECORD
DEFINE circle1 Circle

MAIN

    CALL records.set_undef( reflect.Value.valueOf(rec_cust) )
    LET rec_cust.pkey = 101 
    LET rec_cust.name = "Mike Fitzpatrick"
    LET rec_cust.address.pcode = "7236"
    LET rec_cust.address.city = "Redrock"
    LET rec_cust.creation = CURRENT
    DISPLAY "rec_cust:"
    CALL records.show( reflect.Value.valueOf(rec_cust) )
    
    LET circle1.x = 9834.23
    LET circle1.y = 1334.56
    LET circle1.r = 4.50
    DISPLAY "circle1:"
    CALL records.show( reflect.Value.valueOf(circle1) )

END MAIN
```

Output:

```
rec_cust:
 pkey : 101
 name : Mike Fitzpatrick
 creation : 2020-07-30 11:54:37
 address:
  street : <undefined>
  pcode : 7236
  city : Redrock
circle1:
 x : 9834.23
 y : 1334.56
 r : 4.5
```
