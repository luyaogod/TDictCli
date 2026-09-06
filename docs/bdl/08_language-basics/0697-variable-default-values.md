---
title: "Variable default values"
source: "fgl-topics/c_fgl_variables_008.html"
breadcrumb: "Language basics > Variables > Variable default values"
type: "concept"
---

# Variable default values

> Variables get a default value when defined.

By default, variables are initialized to a value depending on the variable data type.

For example, when defining an `INTEGER` variable, it gets zero as default value,
while a `DECIMAL` variable is set to
`NULL`:

```
MAIN
    DEFINE v1 INTEGER, v2 DECIMAL(10,2)
    DISPLAY v1  -- Shows "0"
    DISPLAY NVL(v2,"null") -- Shows "null"
END MAIN
```

The next table lists default initialization values for each [primitive data type](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."):

| Data type | Default Value |
| --- | --- |
| [`CHAR`](0557-char-size.md "The CHAR data type is a fixed-length character string data type.") | `NULL` |
| [`VARCHAR`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") | `NULL` |
| [`STRING`](0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") | `NULL` |
| [`BIGINT`](0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.") | Zero |
| [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers.") | Zero |
| [`SMALLINT`](0566-smallint.md "The SMALLINT data type is used for storing small whole numbers.") | Zero |
| [`TINYINT`](0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers.") | Zero |
| [`BOOLEAN`](0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE.") | `FALSE` |
| [`FLOAT`](0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") | Zero |
| [`SMALLFLOAT`](0565-smallfloat.md "The SMALLFLOAT data type stores values as single-precision floating-point binary numbers with up to 8 significant digits.") | Zero |
| [`DECIMAL`](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") | `NULL` |
| [`MONEY`](0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage.") | `NULL` |
| [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") | `1899-12-31` (zero days since Unix epoch -1 day) |
| [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") | `NULL` |
| [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") | `NULL` |
| [`TEXT`](0569-text.md "The TEXT data type stores large text data.") | `NULL`, must use [`LOCATE`](0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.") |
| [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") | `NULL`, must use [`LOCATE`](0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.") |

> **Note:**
>
> When adding new elements in a `DYNAMIC ARRAY`, the element fields are initialized
> to `NULL`, no matter the data type. For more details, see [`DYNAMIC ARRAY`
> usage](0734-dynamic-arrays.md).

## Related links

**Related concepts**  

[Variable initializers](0691-variable-initializers.md "Variables can be initialized in their definition.")
