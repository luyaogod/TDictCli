---
title: "Using the INTERVAL type"
source: "fgl-topics/c_fgl_JavaBridge_026.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Mapping native and Java data types > Using the INTERVAL type"
type: "concept"
description: "When calling a Java method with an expression evaluating to an INTERVAL , the runtime system converts the INTERVAL value to an instance of the com.fourjs.fgl.lang.FglInterval class implemented in ..."
---

# Using the INTERVAL type

When calling a Java method with an expression evaluating to an [`INTERVAL`](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units."),
the runtime system converts the `INTERVAL` value to an
instance of the `com.fourjs.fgl.lang.FglInterval` class
implemented in $FGLDIR/lib/fgl.jar. You can then
manipulate the `INTERVAL` from within the Java code.

You must add $FGLDIR/lib/fgl.jar to the class
path in order to compile Java code with `com.fourjs.fgl.lang.FglInterval` class.

The `com.fourjs.fgl.lang.FglInterval` class implements the following:

| Field | Description |
| --- | --- |
| final static int YEAR | Time qualifier for year |
| final static int MONTH | Time qualifier for month |
| final static int DAY | Time qualifier for day |
| final static int HOUR | Time qualifier for hour |
| final static int MINUTE | Time qualifier for minute |
| final static int SECOND | Time qualifier for second |
| final static int FRACTION | Time qualifier for fraction (start qualifier) |
| final static int FRACTION1 | Time qualifier for fraction(1) (end qualifier) |
| final static int FRACTION2 | Time qualifier for fraction(2) (end qualifier) |
| final static int FRACTION3 | Time qualifier for fraction(3) (end qualifier) |
| final static int FRACTION4 | Time qualifier for fraction(4) (end qualifier) |
| final static int FRACTION5 | Time qualifier for fraction(5) (end qualifier) |

| Methods | Description |
| --- | --- |
| static int encodeTypeQualifier( int startUnit, int length, int endUnit)) | Returns the encoded type qualifier for an interval with to interval qualifiers and length passed:encoded qualifier = (*length* \* 256) + (*startUnit* \* 16) + *endUnit*Where *length* defines the total number of significant digits in this time data.For example, with INTERVAL DAY(5) TO FRACTION3:*startUnit* = DAY*length* = 13 (DDDDhhmmssfff)*endUnit* = FRACTION3 |
| java.lang.String toString() | Converts the `INTERVAL` value to a `String` object representing an interval in default format. |
| static FglInterval valueOf( java.lang.String val) | Creates a new `FglInterval` object from a `String` object representing an interval value in the format:DD hh:mm:ss.fff |
| static FglInterval valueOf( java.lang.String val, int startUnit, int endUnit) | Creates a new `FglDateTime` object from a `String` object representing an interval value in standard format, using the qualifiers and precision passed as parameter. |

In the Java code, you can
pass a `com.fourjs.fgl.lang.FglInterval` object
as in this example:

```
public static void useInterval(FglInterval inv) throws ParseException {
    String s = inv.toString();
    ...
}
```

If you need to create a `com.fourjs.fgl.lang.FglInterval` object in your program,
you can use the `valueOf()` class method as in this
example:

```
IMPORT JAVA com.fourjs.fgl.lang.FglInterval
MAIN
  DEFINE inv com.fourjs.fgl.lang.FglInterval
  LET inv = FglInterval.valueOf("-510 12:33:45.123")
  DISPLAY inv.toString()
END MAIN
```
