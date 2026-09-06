---
title: "COLOR attribute"
source: "fgl-topics/c_fgl_FSFAttributes_COLOR.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > COLOR attribute"
type: "concept"
---

# COLOR attribute

> The COLOR attribute defines the foreground color of the text displayed by a form element.

## Syntax

```
COLOR = color-spec
```

1. color-spec can be: `BLACK`,
   `BLUE`, `CYAN`, `GREEN`, `MAGENTA`,
   `RED`, `WHITE`, and `YELLOW`.
2. color-spec can also be one of: `REVERSE`,
   `LEFT`, `BLINK`, and `UNDERLINE`, and it can be
   combined with a color name (`COLOR = RED REVERSE`).

## Usage

The `COLOR` attribute defines the logical color of a value displayed in a
field.

The fglform compiler ignores
`COLOR=WHITE` and `COLOR=BLACK`: These colors can be specified in the
.per file, but will not be written in the .42f file.

> **Important:**
>
> In GUI mode, form elements can also be decorated with
> presentation styles. Pay attention to the specific rules that apply when [combining TTY attributes and presentation
> styles](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.").

A color name like `RED`, `BLUE` can be combined
with an secondary keyword that must be one of: `REVERSE`, `LEFT`,
`BLINK`, and `UNDERLINE`. The secondary keyword can also be used
without a color name.

> **Note:**
>
> For backward compatibility, the color can be specified as a number:
> `0=WHITE, 1=YELLOW, 2=MAGENTA, 3=RED, 4=CYAN, 5=GREEN, 6=BLUE, 7=BLACK`. Note that
> fglform ignores `COLOR=0` or `COLOR=7`: like
> `COLOR=WHITE` and `COLOR=BLACK`.

## Example

```
EDIT f001 = customer.name, COLOR = RED;
```

## Related links

**Related concepts**  

[COLOR WHERE Attribute](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically.")
