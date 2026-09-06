---
title: "OPEN WINDOW attributes"
source: "fgl-topics/c_fgl_windows_and_forms_OPEN_WINDOW_attributes.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > OPEN WINDOW > OPEN WINDOW attributes"
type: "concept"
---

# OPEN WINDOW attributes

> List of attributes for the OPEN WINDOW instruction.

## Window attributes reference

| Attribute | Description |
| --- | --- |
| `TEXT = string` | Defines the default title of the window. When a form is displayed, the form title ([`LAYOUT(TEXT="mytitle")`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")) will be used as window title.We recommend that you define the window title in the form file. |
| `STYLE = string` | Defines the default style of the window. If the form defines a window style, ([`LAYOUT(WINDOWSTYLE="mystyle")`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")), it overwrites the default window style. The `WINDOWSTYLE` attribute is typically used when the form is displayed in the default `SCREEN` window. |
| `BLACK, BLUE, CYAN, GREEN, MAGENTA, RED, WHITE, YELLOW` | Default TTY color of the data displayed in the window. |
| `BOLD, DIM, INVISIBLE, NORMAL` | Default TTY font attribute of the data displayed in the window. |
| `REVERSE, BLINK, UNDERLINE` | Default TTY video attribute of the data displayed in the window. |
| `PROMPT LINE integer` | In character mode, indicates the position of the prompt line for this window. The position can be specified with `FIRST` and `LAST` predefined line positions. |
| `FORM LINE integer` | In character mode, indicates the position of the form line for this window. The position can be specified with `FIRST` and `LAST` predefined line positions. |
| `MENU LINE integer` | In character mode, indicates the position of the ring menu line for this window. The position can be specified with `FIRST` and `LAST` predefined line positions. |
| `MESSAGE LINE integer` | In character mode, indicates the position of the message line for this window. The position can be specified with `FIRST` and `LAST` predefined line positions. |
| `ERROR LINE integer` | In character mode, indicates the position of the error line for this window. The position can be specified with `FIRST` and `LAST` predefined line positions. |
| `COMMENT LINE {OFF\|integer}` | In character mode, indicates the position of the comment line or no comment line at all, for this window. The position can be specified with `FIRST` and `LAST` predefined line positions. See also [`OPTIONS COMMENT LINE`](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines."). |
| `BORDER` | Indicates if the window must be created with a border in character mode. A border frame is drawn outside the specified window. This means, that the window needs 2 additional lines and columns on the screen. |

## Defaut line positions for TUI mode

The default line positions in TUI mode are defined by the `OPTIONS` instruction.
See the [`OPTIONS` command reference topic](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.")
for more details.

## Combining TTY attributes and presentation styles

> **Important:**
>
> In GUI mode, form elements can also be decorated with
> presentation styles. Pay attention to the specific rules that apply when [combining TTY attributes and presentation
> styles](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.").
