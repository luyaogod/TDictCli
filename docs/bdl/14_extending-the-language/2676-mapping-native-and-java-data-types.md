---
title: "Mapping native and Java data types"
source: "fgl-topics/c_fgl_JavaBridge_023.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Mapping native and Java data types"
type: "concept"
description: "Java and Genero have different primitive data types. Unlike Genero, Java is a strongly typed language: You cannot call a method with a String if it was defined to get an int parameter. To call a Java ..."
---

# Mapping native and Java data types

Java and Genero have different primitive data types. Unlike Genero,
Java is a strongly typed language: You cannot call a method with a
`String` if it was defined to get an `int`
parameter.
To call a Java method, [Genero
native typed values](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") need to be converted to/from Java types
such as `byte`, `int`, `short`,
`char` or data objects such as `java.lang.String`.
If possible, the fglrun runtime system will do this conversion
implicitly.

The [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler
will raise the error [-6606](../15_library-reference/4483-genero-bdl-errors.md),
if the native data type does not match the Java (primitive) type, using
Widening Primitive Conversions. For example, passing a Genero
`DECIMAL` when a Java `double` is expected
will fail, but passing a `SMALLFLOAT` (equivalent to
`float`) when a Java `double` is expected
will compile and run.

Genero has advanced native data types such as [`DECIMAL`](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage."),
which do not have an equivalent primitive type or class in Java.
For such Genero types, you need to use a specific Java class provided
in the $FGLDIR/lib/fgl.jar package, like
`com.fourjs.fgl.lang.FglDecimal`.
You can then manipulate the Genero specific value in the Java code.

Genero also implements structured types with [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") definitions,
converted to `com.fourjs.fgl.lang.FglRecord` objects
for Java.

The [Genero arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") cannot be used to call Java methods.
You must use native Java arrays instead.

In some cases you need to explicitly cast with the new `CAST()` operator. See the
topic about [`CAST()` operator](2689-the-cast-operator.md) for more details.

The following tables show the implicit conversions done by the runtime system when a Java method
is called, or when a Java method returns a value or object reference:

| Genero data type | Java equivalent |
| --- | --- |
| [`BIGINT`](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.") | `long` (64-bit signed integer) |
| [`BOOLEAN`](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.") | `Boolean` |
| [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") | [`com.fourjs.fgl.lang.FglByteBlob`](2681-using-the-byte-type.md) |
| [`CHAR`](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.") | `java.lang.String` |
| [`DATE`](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") | [`com.fourjs.fgl.lang.FglDate`](2677-using-the-date-type.md) |
| [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") | [`com.fourjs.fgl.lang.FglDateTime`](2678-using-the-datetime-type.md) |
| [`DECIMAL`](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") | [`com.fourjs.fgl.lang.FglDecimal`](2679-using-the-decimal-type.md) |
| [`FLOAT`](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") | `double` (64-bit signed floating point number) |
| [`INTEGER`](../08_language-basics/0562-integer.md "The INTEGER data type is used for storing large whole numbers.") | `int` (32-bit signed integer) |
| [`INTERVAL`](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") | [`com.fourjs.fgl.lang.FglInterval`](2682-using-the-interval-type.md) |
| [`MONEY`](../08_language-basics/0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage.") | [`com.fourjs.fgl.lang.FglDecimal`](2679-using-the-decimal-type.md) |
| [`SMALLFLOAT`](../08_language-basics/0565-smallfloat.md "The SMALLFLOAT data type stores values as single-precision floating-point binary numbers with up to 8 significant digits.") | `float` (32-bit signed floating point number) |
| [`SMALLINT`](../08_language-basics/0566-smallint.md "The SMALLINT data type is used for storing small whole numbers.") | `short` (16-bit signed integer) |
| [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") | `java.lang.String` |
| [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") | [`com.fourjs.fgl.lang.FglTextBlob`](2680-using-the-text-type.md) |
| [`TINYINT`](../08_language-basics/0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers.") | `byte` (8-bit signed integer) |
| [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") | `java.lang.String` |

| Genero data type | Java equivalent |
| --- | --- |
| [`RECORD` structure](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") | [`com.fourjs.fgl.lang.FglRecord`](2684-using-genero-records.md) |
| [Java Array](2687-using-java-arrays.md) | This is a native Java Array |

| Genero data type |
| --- |
| [ARRAY structures](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.") |
| [Built-in classes](../09_advanced-features/0947-what-class-packages-exist.md "A set of utility packages including useful classes are part of the distribution.") |

## Child topics

- [Using the DATE type](2677-using-the-date-type.md)
- [Using the DATETIME type](2678-using-the-datetime-type.md)
- [Using the DECIMAL type](2679-using-the-decimal-type.md)
- [Using the TEXT type](2680-using-the-text-type.md)
- [Using the BYTE type](2681-using-the-byte-type.md)
- [Using the INTERVAL type](2682-using-the-interval-type.md)
