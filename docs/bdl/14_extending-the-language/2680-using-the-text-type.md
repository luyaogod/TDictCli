---
title: "Using the TEXT type"
source: "fgl-topics/c_fgl_JavaBridge_029.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Mapping native and Java data types > Using the TEXT type"
type: "concept"
description: "When calling a Java method with an expression evaluating to a TEXT , the runtime system converts the TEXT handle to an instance of the com.fourjs.fgl.lang.FglTextBlob class implemented in ..."
---

# Using the TEXT type

When calling a Java method with an expression evaluating to a [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data."), the
runtime system converts the `TEXT` handle to
an instance of the `com.fourjs.fgl.lang.FglTextBlob` class
implemented in $FGLDIR/lib/fgl.jar. You can then
manipulate the LOB from within the Java code.

You must add $FGLDIR/lib/fgl.jar to the class
path in order to compile Java code with `com.fourjs.fgl.lang.FglTextBlob` class.

The `com.fourjs.fgl.lang.FglTextBlob` class implements the following:

| Method | Description |
| --- | --- |
| void dispose() | Dereferences the underlying `TEXT` variable. |
| java.lang.String toString() | Converts the large text data to a simple `String`. |
| static FglTextBlob valueOf( java.lang.String val) | Creates a new `FglTextBlob` object from a `String`. |

In the Java code, you can
pass a `com.fourjs.fgl.lang.FglTextBlob` object
as in this example:

```
public static void useByte(FglTextBlob t) throws ParseException {
    String s = t.toString();
    ...
}
```

If you need to create a `com.fourjs.fgl.lang.FglTextBlob` object in your program,
you can use the `valueOf()` class method as in this
example:

```
IMPORT JAVA com.fourjs.fgl.lang.FglTextBlob
MAIN
  DEFINE jtext com.fourjs.fgl.lang.FglTextBlob
  LET jtext = FglTextBlob.valueOf("abcdef..........")
  DISPLAY jtext.toString()
END MAIN
```
