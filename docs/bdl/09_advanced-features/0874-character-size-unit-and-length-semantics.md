---
title: "Character size unit and length semantics"
source: "fgl-topics/c_fgl_localization_037.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > Character size unit and length semantics"
type: "concept"
description: "When programming an application for a occidental language such as English, a single-byte character set can be used, and the logical size, storage size and print width of characters is the same. For ..."
---

# Character size unit and length semantics

When programming an application for a occidental language such as English, a single-byte
character set can be used, and the logical size, storage size and print width of characters is the
same. For example, in ISO-8859-1, the `ê` character takes one logical position, has a
storage size of one byte and a print width of one.

When programming an international application using multiple languages and a multibyte
character set encoding, you must distinguish three size units:

1. The size in character unit, to count or position logical characters used in a string. For
   example, the strings `abc` and `åôë` have both a length of 3, in
   character units.
2. The size in byte unit, used to encode the character in a given character set. For
   example, a Latin `ê` acute character will use a unique byte in the ISO-8859-1
   character set, but needs two bytes in UTF-8.
3. The size in width unit, used in formatting and alignments. The width is the length of the
   glyph/font of characters, especially in a fixed font. This concept is also known as
   "fullwidth" versus "halfwidth" characters. For example, the width of a
   Chinese logogram is twice the width of Latin characters such as A,B,C.

Working with byte units in a multibyte character set can be difficult: You need to calculate sizes,
lengths and substring offsets in a number of bytes, when the natural way is to count in characters.

Length semantics define the unit to be used for character data type definition, character string
lengths and positions.

With Byte Length Semantics, a length is expressed in bytes, while Character
Length Semantics counts in characters.

## Related links

**Related concepts**  

[Length semantics settings](0881-length-semantics-settings.md "Length semantics settings")
