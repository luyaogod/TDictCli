---
title: "Using the DATE type"
source: "fgl-topics/c_fgl_JavaBridge_024.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Mapping native and Java data types > Using the DATE type"
type: "concept"
description: "When calling a Java method with an expression evaluating to a DATE , the runtime system converts the DATE value to an instance of the com.fourjs.fgl.lang.FglDate class implemented in ..."
---

# Using the DATE type

When calling a Java method with an expression evaluating to a [`DATE`](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation."), the
runtime system converts the `DATE` value to an
instance of the `com.fourjs.fgl.lang.FglDate` class
implemented in $FGLDIR/lib/fgl.jar. You
can then manipulate the date from within the Java code.

You must add $FGLDIR/lib/fgl.jar to the class
path in order to compile Java code with `com.fourjs.fgl.lang.FglDate`
class.

The `com.fourjs.fgl.lang.FglDate` class implements
following:

| Method | Description |
| --- | --- |
| java.lang.String toString() | Converts the `DATE` value to a String object representing the date in the format:YYYY-MM-DD |
| static FglDate valueOf( java.lang.String val) | Creates a new `FglDate` object from a String object representing a date in the format YYYY-MM-DD. |

In the Java code, you can
convert the `com.fourjs.fgl.lang.FglDate` to
a `java.util.Calendar` object as in this
example:

```
public static void useDate(FglDate d) throws ParseException {
    SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd");
    Calendar cal = Calendar.getInstance();
    cal.setTime( sdf.parse(d.toString()) );
    ...
}
```

If you need to create a `com.fourjs.fgl.lang.FglDate` object in your program, you
can use the `valueOf()` class method as in this
example:

```
IMPORT JAVA com.fourjs.fgl.lang.FglDate
MAIN
  DEFINE d com.fourjs.fgl.lang.FglDate
  LET d = FglDate.valueOf("2008-12-23")
  DISPLAY d.toString()
END MAIN
```
