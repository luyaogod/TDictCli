---
title: "CAST() [function]"
source: "fgl-topics/c_fgl_operators_CAST.html"
breadcrumb: "Language basics > Operators > List of expression elements > Data type operators > CAST() [function]"
type: "concept"
---

# CAST() [function]

> The CAST operator converts a Java object to the user-defined type or Java class specified.

## Syntax

```
CAST( variable AS type )
```

1. variable is a variable referencing a Java object.
2. type is a user-defined type or a Java class.

## Usage

The `CAST()` operator is required when you want to assign an object reference
to variable defined with a user-defined type or Java class which requires narrowing
reference conversion.

In the next code example, when assigning a `java.lang.StringBuffer` reference to a
`java.lang.Object` variable, widening reference conversion occurs and no
`CAST()` operator is needed, but when assigning an `java.lang.Object`
reference to a `java.lang.StringBuffer` variable, you must cast the object reference
to a
`java.lang.StringBuffer`:

```
IMPORT JAVA java.lang.Object
IMPORT JAVA java.lang.StringBuffer
MAIN
  DEFINE sb1, sb2 java.lang.StringBuffer 
  DEFINE o java.lang.Object 
  LET sb1 = StringBuffer.create()
  LET o = sb1 -- Widening Reference Conversion does not need CAST()
  LET sb2 = CAST( o AS java.lang.StringBuffer ) -- Narrowing 
                  -- Reference Conversion needs CAST()
END MAIN
```

In order to cast an `fgl.FglRecord` object to a regular [`RECORD`](0715-records.md "Records allow structured program variables definitions."), you need to specify a user-defined
type ([`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")
definition):

```
IMPORT JAVA com.fourjs.fgl.lang.FglRecord
TYPE mytype RECORD f1, f2 INTEGER END RECORD
MAIN
  DEFINE r mytype
  DEFINE jr FglRecord
  LET jr = r
  LET r = CAST(jr AS mytype)
  -- This is denied:
  --   CAST(jr AS RECORD f1, f2 INTEGER END RECORD)
END MAIN
```

## Related links

**Related concepts**  

[The INSTANCEOF operator](../14_extending-the-language/2690-the-instanceof-operator.md "The INSTANCEOF operator")

[The Java interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
