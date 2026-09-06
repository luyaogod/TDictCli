---
title: "The UNICODE Standard"
source: "fgl-topics/c_fgl_localization_011.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > The UNICODE Standard"
type: "concept"
description: "UNICODE is a standard specification to map all possible characters to a numeric value, in order to cover all possible languages in a unique character set. UNICODE defines the mapping of characters to ..."
---

# The UNICODE Standard

UNICODE is a standard specification to map all possible characters to a numeric
value, in order to cover all possible languages in a unique character set. UNICODE defines the
mapping of characters to a numeric whole number, but it does not define how these integers are
encoded in bytes.

The character encoding defines how the UNICODE number will be encoded in a byte or sequence of
bytes.

Several character encodings are based on the UNICODE standard, such as UTF-7, UTF-8, UTF-16,
UTF-32, UCS-2, and UCS-4. Each of these character sets use a different encoding method. For example,
with UTF-8, the letter **Æ** is encoded with two bytes as 0xC3 and 0xB6, while the same character
will be encoded 0x00C6 with UTF-16.

When Microsoft™
Windows® users talk about UNICODE, they typically mean
UCS-2 or UTF-16, while UNIX™ users typically mean UTF-8.

## Related links

**Related concepts**  

[Length semantics settings](0881-length-semantics-settings.md "Length semantics settings")

[When do I need a UNICODE character set?](0876-when-do-i-need-a-unicode-character-set.md "When do I need a UNICODE character set?")
