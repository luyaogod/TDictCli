---
title: "COLOR WHERE Attribute"
source: "fgl-topics/c_fgl_FSFAttributes_COLOR_WHERE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > COLOR WHERE Attribute"
type: "concept"
---

# COLOR WHERE Attribute

> The COLOR WHERE attribute defines a condition to set the foreground color dynamically.

## Syntax

```
COLOR = color-spec [...] WHERE bool-expr
```

1. color-spec can be: `BLACK`,
   `BLUE`, `CYAN`, `GREEN`, `MAGENTA`,
   `RED`, `WHITE`, and `YELLOW`.
2. color-spec can also be one of: `REVERSE`,
   `LEFT`, `BLINK`, and `UNDERLINE`, and it can be
   combined with a color name (`COLOR = RED REVERSE`).
3. bool-expr defines a [boolean expression with a
   restricted syntax](1682-boolean-expressions-in-forms.md "Some form item definitions can include boolean expressions with a form file specific syntax."). This expression can only reference the current form field item tag.

## Usage

The attribute `COLOR WHERE` defines a conditional color. The color will be
applied if the condition is true.

The fglform compiler ignores
`COLOR=WHITE` and `COLOR=BLACK`: These colors can be specified in the
.per file, but will not be written in the .42f file.

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

The condition in `COLOR WHERE` can only reference the field for which the
attribute is set,using its item tag. See [Boolean expressions in forms](1682-boolean-expressions-in-forms.md "Some form item definitions can include boolean expressions with a form file specific syntax.") for more details.

With form fields such as `EDIT`, `BUTTONEDIT` the color
will be applied when leaving the field. The color will not change while editing the value.

Several `COLOR ... WHERE` attributes can be specified, to get different
colors depending on the field value.

When setting the color by program with `DISPLAY BY NAME varname
ATTRIBUTES(RED)`, the `COLOR … WHERE` rules will take precedence over
the `DISPLAY BY NAME` instruction.

> **Important:**
>
> In GUI mode, form elements can also be decorated with
> presentation styles. Pay attention to the specific rules that apply when [combining TTY attributes and presentation
> styles](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.").

## Example

```
EDIT f001 = item.price,
    COLOR = RED   WHERE f001 < 0,
    COLOR = BLUE  WHERE f001 = 0,
    COLOR = GREEN WHERE f001 > 0,
    ;
```

## Related links

**Related concepts**  

[COLOR attribute](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element.")
