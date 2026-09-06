---
title: "Using the DATETIME type"
source: "fgl-topics/c_fgl_JavaBridge_025.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Mapping native and Java data types > Using the DATETIME type"
type: "concept"
description: "When calling a Java method with an expression evaluating to a DATETIME , the runtime system converts the DATETIME value to an instance of the com.fourjs.fgl.lang.FglDateTime class implemented in ..."
---

# Using the DATETIME type

When calling a Java method with an expression evaluating to a [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second."),
the runtime system converts the `DATETIME` value to an instance
of the `com.fourjs.fgl.lang.FglDateTime` class implemented in
$FGLDIR/lib/fgl.jar. You can then manipulate the
`DATETIME` from within the Java code.

You must add $FGLDIR/lib/fgl.jar to the class path in
order to compile Java code with `com.fourjs.fgl.lang.FglDateTime` class.

The `com.fourjs.fgl.lang.FglDateTime` class implements the
following:

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

| Method | Description |
| --- | --- |
| static int encodeTypeQualifier( int startUnit, int endUnit) | Returns the encoded type qualifier for a datetime with to datetime qualifiers passed:encoded qualifier = (length \* 256) + (startUnit \* 16) + endUnitWhere *length* defines the total number of significant digits in this time data.For example, with `DATETIME YEAR TO MINUTE`:startUnit = YEARlength = 12 (YYYYMMDDhhmm)endUnit = MINUTE |
| java.lang.String toString() | Converts the `DATETIME` value to a String object representing a datetime in the format `YYYY-MM-DD hh:mm:ss.fff`. |
| static FglDateTime valueOf( long milliseconds) | Creates a new `FglDateTime` object representing the specified number of milliseconds since the standard base time known as "the epoch", namely January 1, 1970, 00:00:00 GMT. |
| static FglDateTime valueOf( java.lang.String val) | Creates a new `FglDateTime` object from a String object representing a datetime value in the format: `YYYY-MM-DD hh:mm:ss.fff` |
| static FglDateTime valueOf( java.lang.String val, int startUnit, int endUnit) | Creates a new `FglDateTime` object from a String object representing a datetime value in the format `YYYY-MM-DD hh:mm:ss.fff`, using the qualifiers passed as parameter. |

In the Java code, you can
convert the `com.fourjs.fgl.lang.FglDateTime` to
a `java.util.Calendar` object as in this
example:

```
public static void useDatetime(FglDateTime dt) throws ParseException {
   SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss.SSS");
   Calendar cal = Calendar.getInstance();
   cal.setTime( sdf.parse(dt.toString()) );
   ...
}
```

If you need to create a `com.fourjs.fgl.lang.FglDateTime` object in your program,
you can use the `valueOf()` class method as in this
example:

```
IMPORT JAVA com.fourjs.fgl.lang.FglDateTime
MAIN
  DEFINE dt com.fourjs.fgl.lang.FglDateTime
  LET dt = FglDateTime.valueOf("2008-12-23 11:22:33.123")
  LET dt = FglDateTime.valueOf("11:22:33.123",
                FglDateTime.HOUR, FglDateTime.FRACTION3)
  DISPLAY dt.toString()
END MAIN
```

The `valueOf()` method expects a string representing
a complete date-time specification, from year to milliseconds,
equivalent to a `DATETIME YEAR TO FRACTION(3)` data
type.
