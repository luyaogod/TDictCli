---
title: "Syntax of presentation styles file (.4st)"
source: "fgl-topics/c_fgl_presentation_styles_003.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Syntax of presentation styles file (.4st)"
type: "concept"
---

# Syntax of presentation styles file (.4st)

> A .4st presentation styles file is an XML file defining style attributes to be applied by front-ends.

## Syntax (.4st)

```
<StyleList>
  <Style name="style-identifier" >
    <StyleAttribute name="attribute-name" value="attribute-value" />
      [...]
  </Style>
  [...]
</StyleList>
```

where style-identifier
is:

```
{ element-type.style-name [ :pseudo-selector ]
| element-type [ :pseudo-selector ]
| .style-name [ :pseudo-selector ]
| *
}
```

where element-type is one
of:

```
{ Button | Edit | ButtonEdit | CheckBox | ComboBox
| RadioGroup | SpinEdit | DateEdit | DateTimeEdit
| Image | Label | Slider | TimeEdit | TextEdit
| ProgressBar | WebComponent
| Form | Folder | Page | Grid | Group | HBox | VBox
| ScrollGrid | Table | Tree
| Menu | Message | Action | MenuAction
| ToolBar | ToolBarSeparator | TopMenu
| UserInterface | Window
}
```

and where pseudo-selector is one
of:

```
{ focus | inactive | active
| query | display | input
| even | odd
| message | error | summaryLine
}
```

1. element-type is the type of the AUI tree element, such as
   `Edit`, `Window`. The element-type is case
   sensitive. When only using element-type, the style applies to all elements of
   this type. When combining element-type.style-name, it applies
   to elements of this element type and using the
   `STYLE="style-name"` form attribute. See also [Element types](1620-element-types.md "Which AUI tree elements can get a style?").
2. style-name is an explicit style name, that can be referenced in [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attributes of form items. The
   style-name is case sensitive. When only using .style-name, the
   style applies to all elements using the `STYLE="style-name"` form
   attribute. When combining element-type.style-name, it applies
   to elements of the specified type and using the
   `STYLE="style-name"` form attribute.
3. pseudo-selector indicates in what context the style should apply. See
   [Pseudo selectors](1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled.") for more details.
4. attribute-name defines the name of the style attribute.
5. attribute-value defines the value to be assigned to attribute-name.

## Syntax of attribute values

Presentation style attribute values are always specified as strings, for example:

```
<StyleAttribute name="fontFamily" value="Serif" />
```

Numeric values must be specified in quotes:

```
<StyleAttribute name="completionTimeout" value="60" />
```

Boolean values must be specified with the values `"yes"` or
`"no"`:

```
<StyleAttribute name="forceDefaultSettings" value="yes" />
```

For backward compatibility, the values `0/1` and `true/false` are
supported by some front-ends. Use exclusively the `yes/no` values.
