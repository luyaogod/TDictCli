---
title: "Identifying Genero data types in Java code"
source: "fgl-topics/c_fgl_JavaBridge_030.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Identifying Genero data types in Java code"
type: "concept"
description: "Java data types and Genero data types are different. To identify Genero types in Java code, you can use the com.fourjs.fgl.lang.FglTypes class implemented in $FGLDIR/lib/fgl.jar . You can, for ..."
---

# Identifying Genero data types in Java code

Java data types and Genero data types are different. To identify Genero
types in Java code, you can use the `com.fourjs.fgl.lang.FglTypes`
class implemented in $FGLDIR/lib/fgl.jar.

You can, for example, identify the data type of a member of an [FglRecord object](2684-using-genero-records.md).

You must add $FGLDIR/lib/fgl.jar to the class
path in order to compile Java code with `com.fourjs.fgl.lang.FglType` class.

The `com.fourjs.fgl.lang.FglTypes` class implements the following:

| Field | Corresponding data type |
| --- | --- |
| final static int ARRAY | [ARRAY](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") object |
| final static int BIGINT | [BIGINT](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.") |
| final static int BOOLEAN | [BOOLEAN](../08_language-basics/0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE.") |
| final static int BYTE | [BYTE](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") |
| final static int CHAR | [CHAR](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.") |
| final static int DATE | [DATE](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") |
| final static int DATETIME | [DATETIME](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") |
| final static int DECIMAL | [DECIMAL](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") |
| final static int FGL_OBJECT | An FGL object like `base.Channel`. |
| final static int FLOAT | [FLOAT](../08_language-basics/0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") |
| final static int INT | [INTEGER](../08_language-basics/0562-integer.md "The INTEGER data type is used for storing large whole numbers.") |
| final static int INTERVAL | [INTERVAL](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") |
| final static int JAVA_OBJECT | A Java object like `java.lang.String`. |
| final static int MONEY | [MONEY](../08_language-basics/0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage.") |
| final static int RECORD | [RECORD](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") structure |
| final static int SMALLFLOAT | [SMALLFLOAT](../08_language-basics/0565-smallfloat.md "The SMALLFLOAT data type stores values as single-precision floating-point binary numbers with up to 8 significant digits.") |
| final static int SMALLINT | [SMALLINT](../08_language-basics/0566-smallint.md "The SMALLINT data type is used for storing small whole numbers.") |
| final static int STRING | [STRING](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") |
| final static int TEXT | [TEXT](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") |
| final static int TINYINT | [TINYINT](../08_language-basics/0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers.") |
| final static int VARCHAR | [VARCHAR](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") |
