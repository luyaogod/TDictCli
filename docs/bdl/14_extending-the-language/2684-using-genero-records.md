---
title: "Using Genero records"
source: "fgl-topics/c_fgl_JavaBridge_031.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Using Genero records"
type: "concept"
description: "When passing a RECORD to a Java method, the runtime system converts the RECORD to an instance of the com.fourjs.fgl.lang.FglRecord class implemented in $FGLDIR/lib/fgl.jar . The FglRecord object is a ..."
---

# Using Genero records

When passing a [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") to a Java
method, the runtime system converts the `RECORD` to an instance of the
`com.fourjs.fgl.lang.FglRecord` class implemented in
$FGLDIR/lib/fgl.jar.

The `FglRecord` object is a copy of the `RECORD` variable;
structure and members of the `FglRecord` object can be read within the Java code, but
cannot be modified.

You must add $FGLDIR/lib/fgl.jar to the class path in order to compile Java
code with `com.fourjs.fgl.lang.FglRecord` class.

The `com.fourjs.fgl.lang.FglRecord` class implements the following methods:

| Method | Description |
| --- | --- |
| void dispose() | Dereferences the underlaying member variables. |
| double getDouble(int p) | Returns the double value of the record member at position p. |
| FglByteBlob getFglByteBlob(int p) | Returns the [`FglByteBlob`](2681-using-the-byte-type.md) value of the record member at position p. |
| FglDate getFglDate(int p) | Returns the [`FglDate`](2677-using-the-date-type.md) value of the record member at position p. |
| FglDateTime getFglDateTime(int p) | Returns the [`FglDateTime`](2678-using-the-datetime-type.md) value of the record member at position p. |
| FglDecimal getFglDecimal(int p) | Returns the [`FglDecimal`](2679-using-the-decimal-type.md) value of the record member at position p. |
| FglInterval getFglInterval(int p) | Returns the [`FglInterval`](2682-using-the-interval-type.md) value of the record member at position p. |
| FglTextBlob getFglTextBlob(int p) | Returns the [`FglTextBlob`](2680-using-the-text-type.md) value of the record member at position p. |
| int getFieldCount() | Returns the number of record members. |
| java.lang.String getFieldName(int p) | Returns the name of the record member at position p. |
| int getInt(int p) | Returns the int value of the record member at position p. |
| java.lang.String getString(int p) | Returns the String representation of the value of the record member at position p. |
| FglTypes getType(int p) | Returns the [`FglTypes`](2683-identifying-genero-data-types-in-java-code.md) constant of the record member at position p. |
| java.lang.String getTypeName(int p) | Returns the string representation of the [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") of the record member at position p. |
| int getTypeQualifier(int p) | Returns the encoded type qualifier of the record member at position p. |

In the Java code, use the query methods of the `com.fourjs.fgl.lang.FglRecord` to
identify the members of the
`RECORD`:

```
public static void showMemberTypes(FglRecord rec){
    int i;
    int n = rec.getFieldCount();
    for (i = 1; i <= n; i++)
        System.out.println( String.valueOf(i) + ":" +
        rec.getFieldName(i) + " / " + rec.getTypeName(i) );
}
```

When assigning a `RECORD` to a `com.fourjs.fgl.lang.FglRecord`,
widening conversion applies implicitly. But when assigning a
`com.fourjs.fgl.lang.FglRecord` to a `RECORD`, narrowing
conversion applies and you must explicitly [`CAST`](2689-the-cast-operator.md) the original object reference to the type of the
`RECORD`. The following example shows how to return an FglRecord object from a Java
method:

The PassRecord.4gl source:

```
IMPORT JAVA com.fourjs.fgl.lang.FglRecord
IMPORT JAVA UseRecord
MAIN
    TYPE type1 RECORD
               id INTEGER,
               name VARCHAR(50)
           END RECORD
    DEFINE rec1, rec2 type1
    LET rec1.id = 123
    LET rec1.name = "McFly"
    LET rec2 = CAST(UseRecord.getRecord(rec1) AS type1)
END MAIN
```

The UseRecord.java
source:

```
import com.fourjs.fgl.lang.FglRecord;
public class UseRecord{
    public static FglRecord getRecord(FglRecord rec){
        ...
        return rec;
    }
}
```
