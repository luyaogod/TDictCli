---
title: "reflect.Value.getCurrentValue"
source: "fgl-topics/c_fgl_ext_reflect_Value_getCurrentValue.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getCurrentValue"
type: "concept"
---

# reflect.Value.getCurrentValue

> Returns the actual (typed) value of this value object.

## Syntax

```
getCurrentValue( )
  RETURNS any-type
```

## Usage

The `getCurrentValue()` method returns the typed value of this
`reflect.Value` object.

The returned value is typed: The value pushed on the stack has the original type of the variable
or expression used to create this `reflect.Value` object.

The `getCurrentValue()` method can for example be used to set the value of an SQL
parameter with [`base.SqlHandle.setParameter()`](3037-base-sqlhandle-setparameter.md "Sets the value of an SQL parameter for this SQL handle."), in order to provide the type and value
expected by the SQL database engine.

## Example

```
IMPORT reflect
MAIN
    DEFINE dec DECIMAL(10,2)
    DEFINE mon MONEY(10,2)
    DEFINE dtm DATETIME YEAR TO SECOND
    LET dec = -123.45
    LET mon = 1200.45
    LET dtm = CURRENT
    CALL show_variable( reflect.Value.valueOf(dec) )
    CALL show_variable( reflect.Value.valueOf(mon) )
    CALL show_variable( reflect.Value.valueOf(dtm) )
END MAIN

FUNCTION show_variable( var reflect.Value )
    DEFINE t reflect.Type
    LET t = reflect.Type.typeOf( var.getCurrentValue() )
    DISPLAY "Type: ", t.toString(), COLUMN 35, " Value = ", var.getCurrentValue()
END FUNCTION
```

Shows:

```
Type: DECIMAL(10,2)                Value =      -123.45
Type: MONEY(10)                    Value =      $1200.45
Type: DATETIME YEAR TO SECOND      Value = 2021-07-18 11:12:34
```

## Related links

**Related concepts**  

[reflect.Value.toString](4378-reflect-value-tostring.md "Converts this reflect.Value to a character string.")

[Example 3: Generic SQL UPDATE](4406-example-3-generic-sql-update.md "Example 3: Generic SQL UPDATE")
