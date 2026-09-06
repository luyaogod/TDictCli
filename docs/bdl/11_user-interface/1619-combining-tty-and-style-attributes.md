---
title: "Combining TTY and style attributes"
source: "fgl-topics/c_fgl_presentation_styles_013.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Combining TTY and style attributes"
type: "concept"
---

# Combining TTY and style attributes

> TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.

We distinguish two kind of TTY attributes:

- specific TTY attributes are directly set in the element node in the AUI tree.
- inherited TTY attributes are taken from the parent AUI tree nodes (such as the
  `form` or `window` nodes).

The specific element TTY attributes can, for example, be defined with the [`COLOR`](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element.") attribute of form items. When
controlled by a dialog instruction such as `INPUT BY NAME`, form fields automatically
get specific TTY attributes, that may have been defined globally (`OPTIONS INPUT
ATTRIBUTES`), by the current window (`OPEN WINDOW ... ATTRIBUTES`), or by the
current form (`DISPLAY FORM ... ATTRIBUTES`), or by the current dialog (`INPUT
... ATTRIBUTES`). Array dialog cell attributes defined with the [`ui.Dialog.setArrayAttributes()`](../15_library-reference/3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs).") are also considered as specific
TTY attributes.

Specific TTY attributes defined for a form element have a higher priority
than style attributes, while inherited TTY attributes (set on one of the parent AUI
tree elements) have a lower priority than style attributes defined for the element.

To illustrate this rule, see the following form defining two static labels and two fields. All
items are using the "`mystyle`" presentation style, and some elements use a
specific TTY attribute with `COLOR=BLUE`:

```
LAYOUT
GRID
{
[lab01            :fld01                  ]
[lab02            :fld02                  ]
}
END
END
ATTRIBUTES
LABEL lab01: TEXT="Field 1:", COLOR = BLUE, STYLE = "mystyle";
EDIT fld01 = FORMONLY.field01, COLOR = BLUE, STYLE = "mystyle";
LABEL lab02: TEXT="Field 2:", STYLE = "mystyle";
EDIT fld02 = FORMONLY.field02, STYLE = "mystyle";
END
```

The program displays the form (`DISPLAY FORM`) with the
`ATTRIBUTES(RED)` clause, and the fields are used by an `INPUT`
dialog, with no `ATTRIBUTES` clause. The default TTY attributes to be used by the
`INPUT` dialog are taken from the `DISPLAY FORM`
instruction:

```
MAIN
    DEFINE rec RECORD
               field01 STRING,
               field02 STRING
           END RECORD

    CALL ui.Interface.loadStyles("ttyform")

    OPEN FORM f1 FROM "ttyform"
    DISPLAY FORM f1 ATTRIBUTES(RED)

    LET rec.field01 = "aaa"
    LET rec.field02 = "bbb"
    INPUT BY NAME rec.* WITHOUT DEFAULTS

END MAIN
```

The ttyform.4st styles file defines the "`mystyle`"
attributes as follows:

```
<StyleList>
  <Style name="Edit.mystyle">
    <StyleAttribute name="textColor" value="green" />
  </Style>
  <Style name="Label.mystyle">
    <StyleAttribute name="textColor" value="green" />
  </Style>
</StyleList>
```

With the above code sample, we get the following result:

![Simple form with labels and edit fields](../_images/TTYForm1.jpg)

  

- The text of the `lab01` label is displayed in blue, from the
  specific `COLOR=BLUE` attribute defined in the .per
  file. The style `Label.mystyle` is being ignored.
- The text in the form field `fld01` is displayed in blue, from the
  specific `COLOR=BLUE` attribute defined in the .per
  file. The style `Edit.mystyle` is being ignored.
- The text of the `lab02` label is displayed in green, from the style
  `Label.mystyle`. The TTY attribute `RED` from the `DISPLAY
  FORM` is ignored, because the label is not used by the dialog instruction and thus does not
  get that specific TTY attribute like fields do.
- The text in the form field `fld02` is displayed in red, the specific
  TTY attribute set by the dialog instruction because of the `ATTRIBUTES(RED)` clause
  used in `OPEN FORM`. The style `Edit.mystyle` is being ignored.
