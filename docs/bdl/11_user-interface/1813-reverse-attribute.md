---
title: "REVERSE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_REVERSE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > REVERSE attribute"
type: "concept"
---

# REVERSE attribute

> The REVERSE attribute displays any value in the field in reverse video (dark characters in a bright field).

## Syntax

```
REVERSE
```

## Usage

Use the `REVERSE` attribute to highlight specific fields in your forms.

On graphical front-ends, the `REVERSE` attribute is rendered by using
the field `COLOR` attribute as background color. If the `COLOR`
attribute is not defined, the reverse color defaults to gray.

> **Note:**
>
> For the predefined (TTY)
> color names (like `blue`, `red`, `yellow`), the actual
> rendering color used by the front-ends is adapted to the reverse or normal mode: In normal mode,
> colors like `"yellow"` must be darker, to make the foreground text easily readable on
> an white background. In reverse mode (where the color specification applies to the background), the
> actual background color must be lighter, so that the black foreground text can be read more
> easily.

On character based terminals, the `REVERSE` video escape sequences must be
defined in the TERMINFO or TERMCAP databases.

## Example

```
EDIT f001 = customer.name, COLOR = BLUE, REVERSE;
```

## Related links

**Related concepts**  

[COLOR attribute](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element.")

[TERMINFO](../07_configuration/0502-terminfo.md "Defines the terminal capabilities database.")

[TERMCAP](../07_configuration/0501-termcap.md "Defines the termcap terminal capabilities database on UNIX platforms.")

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")
