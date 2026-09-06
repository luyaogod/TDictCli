---
title: "Using the DECIMAL type"
source: "fgl-topics/c_fgl_JavaBridge_027.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Mapping native and Java data types > Using the DECIMAL type"
type: "concept"
description: "When calling a Java method with an expression evaluating to a DECIMAL , the runtime system converts the DECIMAL value to an instance of the com.fourjs.fgl.lang.FglDecimal class implemented in ..."
---

# Using the DECIMAL type

When calling a Java method with an expression evaluating to a [`DECIMAL`](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), the runtime
system converts the `DECIMAL` value to an instance of the
`com.fourjs.fgl.lang.FglDecimal` class implemented in
$FGLDIR/lib/fgl.jar. You can then manipulate the
`DECIMAL` from within the Java code.

You must add $FGLDIR/lib/fgl.jar to the class
path in order to compile Java code with `com.fourjs.fgl.lang.FglDecimal` class.

The `com.fourjs.fgl.lang.FglDecimal` class implements the following methods shown
in the table:

| Method | Description |
| --- | --- |
| static int encodeTypeQualifier( int precision, int scale) | Returns the encoded type qualifier for this decimal based on precision and scale.encoded qualifier = (precision \* 256) + scaleUse 255 as scale for floating point decimal. |
| java.lang.String toString() | Converts the `DECIMAL` value to a String object. |
| static FglDecimal valueOf(int val) | Creates a new `FglDecimal` object from an `int` value. |
| static FglDecimal valueOf( java.lang.String val) | Creates a new `FglDecimal` object from a `String` object representing a decimal value. |

In the Java code, you can convert the `com.fourjs.fgl.lang.FglDecimal` to a
`java.lang.BigDecimal` as in the following
example:

```
public static FglDecimal divide(FglDecimal d1, FglDecimal d2){
    BigDecimal bd1 = new BigDecimal(d1.toString());
    BigDecimal bd2 = new BigDecimal(d2.toString());
    BigDecimal res = bd1.divide(bd2, BigDecimal.ROUND_FLOOR);
    return FglDecimal.valueOf(res.toString());
}
```

If you need to create a `com.fourjs.fgl.lang.FglDecimal` object in your program,
you can use the `valueOf()` class method as in this
example:

```
IMPORT JAVA com.fourjs.fgl.lang.FglDecimal
MAIN
  DEFINE jdec com.fourjs.fgl.lang.FglDecimal
  LET jdec = FglDecimal.valueOf("123.45")
  DISPLAY jdec.toString()
END MAIN
```
