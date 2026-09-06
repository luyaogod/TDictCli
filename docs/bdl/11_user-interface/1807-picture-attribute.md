---
title: "PICTURE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_PICTURE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > PICTURE attribute"
type: "concept"
---

# PICTURE attribute

> The PICTURE attribute specifies a character pattern for data entry in a text field, and prevents entry of values that conflict with the specified pattern.

## Syntax

```
PICTURE = "picture-string"
```

1. picture-string defines the data input pattern of the field.

## Usage

The `PICTURE` attribute defines an input mask for `CHAR`,
`VARCHAR` or `STRING` fields.

Do not use the `PICTURE` attribute for numeric or date/time fields.

This attribute is not used by the runtime system to validate the field.

> **Important:**
>
> For maximum security, when the data is stored in a database, define a `CHECK`
> constraint on the SQL column corresponding to the field, in order to deny values that do not match
> the `PICTURE` attribute, when an `INSERT` or `UPDATE`
> statement is executed.

picture-string can be any combination of characters, where the characters
"`A`", "`#`" and "`X`" have a special meaning.

- The placeholder "`X`" specifies any character at the given position, for a
  character width of 1, including space characters.
- The placeholder "`#`" specifies a digit character (1,2,3,…) at the given
  position. Spaces ar not allowed.
- The placeholder "`A`" specifies any alphabetic character (A, B, é, É), for a
  character width of 1, at the given position. Spaces ar not allowed.

The `"X"` and `"A"` placeholders represent characters in the
current [application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."), for a character width of
1. When using a locale charset allowing [fullwidth
characters](../09_advanced-features/0874-character-size-unit-and-length-semantics.md) like Chinese ideographs using a width of 2, two placeholders `"XX"`
or `"AA"` need to be specified to allow fullwidth character input.

Any character in the `PICTURE` specification different from "`A`",
"`X`" and "`#`" is treated as a literal. Such characters automatically
appear in the field and do not have to be entered by the user.

The `PICTURE` attribute does not require data entry into the entire field. It only
requires that whatever characters are entered conform to picture-string.

Combine the `PICTURE` attribute with the [`UPSHIFT`](1838-upshift-attribute.md "The UPSHIFT attribute forces character input to uppercase letters.") attribute, to allow only
uppercase letters for `X` and `A` place holders.

When `PICTURE` specifies input formats for `DATETIME` or
`INTERVAL` fields, the form compiler does not check the syntax of
picture-string. Any error in picture-string such as an
incorrect field separator, produces a runtime error.

The typical usage for the `PICTURE` attribute is for (fixed-length)
`CHAR` fields. It is not recommended to use `PICTURE` for other data
types, especially numeric or date/time fields: The current value of the field must always match the
`PICTURE` attribute.

Understand that the `PICTURE` attribute defines a mask for data entry. In order
to format fields when data is displayed to the field, use the `FORMAT` attribute
instead. `FORMAT` is typically used for numeric and date fields, while
`PICTURE` is typically used for formatted character string fields requiring
input control.

The `PICTURE` attribute is ignored in `CONSTRUCT` and
`DISPLAY` / `DISPLAY ARRAY` instructions. It only effects
`INPUT` and `INPUT ARRAY` dialogs.

## Example

```
EDIT f001 = carinfo.ident, PICTURE = "AA####-AA(X)";
```

## Related links

**Related concepts**  

[DATETIME qual1 TO qual2](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")

[INTERVAL qual1 TO qual2](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.")

[CHAR(size)](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.")

[FORMAT attribute](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.")
