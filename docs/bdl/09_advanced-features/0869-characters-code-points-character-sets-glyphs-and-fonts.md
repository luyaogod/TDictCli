---
title: "Characters, code points, character sets, glyphs and fonts"
source: "fgl-topics/c_fgl_localization_006.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > Characters, code points, character sets, glyphs and fonts"
type: "concept"
description: "In computers, a character is the unit of information corresponding to a symbol of a natural language. This can be a letter, a digit, a punctuation mark, a mathematical or even musical symbol. To ..."
---

# Characters, code points, character sets, glyphs and fonts

In computers, a character is the unit of information corresponding to a symbol of a
natural language. This can be a letter, a digit, a punctuation mark, a mathematical or even musical
symbol.

To represent a character in memory or in a file, computers must encode the character in a
specific numeric value called code point. This code point uniquely identifies a
character in a given character set.

Mapping a character to a code point is called character encoding. The same code
point might represent a different character in several character sets.

The glyph is the graphical representation of the character. In other words, it's the
way the character is drawn on the screen or on a printer.

Computers implement the glyph of characters with fonts, by mapping a code point to a
bitmap image or drawing instructions based on math formulas or vector graphics.
