---
title: "Formatting data in Java code"
source: "fgl-topics/c_fgl_JavaBridge_032.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Formatting data in Java code"
type: "concept"
description: "To format numeric and date-time data in Java code, use the com.fourjs.fgl.lang.FglFormat class implemented in $FGLDIR/lib/fgl.jar . You must add $FGLDIR/lib/fgl.jar to the class path in order to ..."
---

# Formatting data in Java code

To format numeric and date-time data in Java code, use the
`com.fourjs.fgl.lang.FglFormat` class implemented in
$FGLDIR/lib/fgl.jar.

You must add $FGLDIR/lib/fgl.jar to the class
path in order to compile Java code with `com.fourjs.fgl.lang.FglFormat` class.

The `com.fourjs.fgl.lang.FglFormat` class provides
an interface to the data formatting functions of the runtime
system. This class is actually an equivalent of the [`USING`](../08_language-basics/0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") operator
in the language.

The `com.fourjs.fgl.lang.FglFormat` class implements
the following:

| Method | Description |
| --- | --- |
| static String format( int v, String fmt) | Formats the integer value provided as Java `int` based on fmt. Here fmt must specify a numeric format with `[$@*#&<()+-]` characters, same as in the `USING` operator. |
| static String format( double v, String fmt) | Formats the `FLOAT` value provided as Java double based on fmt. Here fmt must specify a numeric format with `[ $ @*#&<()+-. ,]` characters,same as in the `USING` operator. |
| static String format( FglDate v, String fmt) | Formats the `DATE` value provided as [`FglDate`](2677-using-the-date-type.md) based on fmt. Here fmt must specify a date format with `[mdy]`characters, same as in the `USING` operator. |
| static String format( FglDecimal v, String fmt) | Formats the `DECIMAL` value provided as [`FglDecimal`](2679-using-the-decimal-type.md), by using fmt. Here fmt must specify a numeric format with `[ $ @*# $<()+-. ,]` characters, same as in the `USING` operator. |

Example of Java code using
the `com.fourjs.fgl.lang.FglFormat` class:

```
public static void formatDecimal(FglDecimal dec){
    System.out.println( FglFormat.format(dec,"$#####&.&&" );
}
```
