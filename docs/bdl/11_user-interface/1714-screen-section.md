---
title: "SCREEN section"
source: "fgl-topics/c_fgl_FormSpecFiles_SCREEN_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > SCREEN section"
type: "concept"
---

# SCREEN section

> The SCREEN section defines the form layout for TUI mode forms.

## Syntax

```
SCREEN
   [ SIZE lines [ BY chars ] ]
   [ TITLE "title" ]
   [ TAG "tag-string" ]
{
  { text | [ item-tag [ | item-tag ] [...] ] }
  [...]
}
[END]
```

1. lines is the number of characters the form can display vertically. The default
   is 24.
2. chars is the number of characters the form can display horizontally. The
   default is the maximum number of characters in any line of the screen definition.
3. title is the title for the top window.
4. tag-string is a user-defined string.
5. item-tag and text define form elements in the layout.

## Usage

The `SCREEN` section must be used to design TUI mode screens. For a GUI mode
application, use a `LAYOUT` instead.

The `SCREEN` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `SCREEN` section is mandatory, unless you use a `LAYOUT`
section.

The `END` keyword is optional.

The `SIZE lines [ BY chars ]` clause is
supported for backward compatibility. A good practice is to omit this clause, and let the form
compiler compute the size of the form, based on the content of the curly brackets.

The `TAG` attribute can be used to specify a string that will help to identify the
form at runtime. For more details about this attribute, see [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.").

Inside the `SCREEN` section, you can define the position of text labels and
form fields in the area delimited by the `{}` curly brackets.

Between the curly brackets, horizontal lines can be specified with a sequence of hyphen characters
(`-----`).

> **Tip:**
>
> Avoid Tab characters (ASCII 9) inside the area delimited by the curly brackets. If used, Tab
> characters will be replaced by 8 blank spaces by fglform.

You can include [graphics characters defined for your
terminal](1535-genero-specific-termcap-definitions.md) between the curly brackets to place rectangles in a screen form. Use
`\g` to start and end graphics mode. For example, `\g|\g` defines a
vertical segment. The following characters can be used to indicate the borders of rectangles:

- `p` upper-left corner.
- `q` upper-right corner.
- `b` left-lower corner.
- `d` left-right corner.
- `-` (hyphen) to mark horizontal segments.
- `|` (pipe) to mark vertical segments.

## Example

```
SCREEN TITLE "Customer info" TAG "regular"
{
  CustId : [f001   ] Name: [f002                ]
  Address: [f003                                ]
           [f003                                ]
  ------------------------------------------------
}
END
```

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")
