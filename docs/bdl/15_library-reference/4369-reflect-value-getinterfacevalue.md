---
title: "reflect.Value.getInterfaceValue"
source: "fgl-topics/c_fgl_ext_reflect_Value_getInterfaceValue.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getInterfaceValue"
type: "concept"
---

# reflect.Value.getInterfaceValue

> Returns the value record referenced by this interface variable.

## Syntax

```
getInterfaceValue( )
  RETURNS reflect.Value
```

## Usage

The `getInterfaceValue()` method returns a `reflect.Value` object,
that is a reference to the variable referenced by this [interface](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") `reflect.Value` object.

The `reflect.Value` object used to call this method must have been created from an
[`INTERFACE`](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") variable, or is
a `reflect.Value` object returned from a method like [`valueOf()`](4357-reflect-value-valueof.md "Creates a new reflect.Value object as a reference to the original variable."), and references an
interface variable.

Program variables can be defined from types with
methods:

```
TYPE CustType RECORD
   ...
END RECORD
FUNCTION (self CustType) print() RETURNS ()
   ...
END FUNCTION
DEFINE c CustType  -- Variable created from type with methods
CALL c.print()
```

Interface variables can be defined from `INTERFACE` types and reference variables
with
methods:

```
TYPE DisplayInterface INTERFACE
    print()
END INTERFACE
DEFINE i DisplayInterface
LET i = c       -- Interface variable i references c
CALL i.print()  -- Can call c.print() indirectly
```

In
the above example, with a `reflect.Value` object referencing the interface variable
`i`, the `getInterfaceValue()` method returns the
`reflect.Value` object referencing the variable
`c`:

```
DEFINE vi reflect.Value
DEFINE ivi_c reflect.Value
LET vi = reflect.Value.valueOf(i)   -- This references i
LET i = c
LET ivi_c = vi.getInterfaceValue()  -- This references c
```

The new `reflect.Value` object references the original
variable: Any call to a reflection manipulation method modifies the underlying variable
directly.

The method throws error [-8112](4483-genero-bdl-errors.md),
if this value is not a `reflect.Value` object referencing an `INTERFACE`.

## Example

```
IMPORT reflect

TYPE CustType RECORD
    id INTEGER,
    name VARCHAR(50)
END RECORD

TYPE OrdListType RECORD
     cust_id INTEGER,
     orders DYNAMIC ARRAY OF RECORD
         id INTEGER,
         crea DATE,
         totam DECIMAL(10,2)
     END RECORD
END RECORD

TYPE DisplayInterface INTERFACE
    print()
END INTERFACE

FUNCTION (self CustType) print() RETURNS ()
    DISPLAY "Customer: ", self.name
END FUNCTION

FUNCTION (self OrdListType) print() RETURNS ()
    DEFINE x INTEGER
    DISPLAY "Orders for : ", self.cust_id
    FOR x=1 TO self.orders.getLength()
        DISPLAY self.orders[x].crea, ":", self.orders[x].totam
    END FOR
END FUNCTION

FUNCTION main()
    DEFINE i DisplayInterface
    DEFINE c CustType
    DEFINE ol OrdListType
    DEFINE vi reflect.Value
    DEFINE ivi_c reflect.Value
    DEFINE ivi_ol reflect.Value

    LET vi = reflect.Value.valueOf(i)
    DISPLAY "vi kind        : ", vi.getType().getKind()

    LET i = c
    LET ivi_c = vi.getInterfaceValue()
    DISPLAY "c field 1 type : ", ivi_c.getField(1).getType().toString()
    DISPLAY "c field 2 type : ", ivi_c.getField(2).getType().toString()

    LET i = ol
    LET ivi_ol = vi.getInterfaceValue()
    DISPLAY "ol field 1 type: ", ivi_ol.getField(1).getType().toString()
    DISPLAY "ol field 2 type: ", ivi_ol.getField(2).getType().toString()

    DISPLAY "ivi_c assignable from ivi_ol ? : ",
        IIF(ivi_c.getType().isAssignableFrom(ivi_ol.getType()),"YES","NO")

END FUNCTION
```

Shows:

```
vi kind        : INTERFACE
c field 1 type : INTEGER
c field 2 type : VARCHAR(50)
ol field 1 type: INTEGER
ol field 2 type: DYNAMIC ARRAY OF RECORD
ivi_c assignable from ivi_ol ? : NO
```
