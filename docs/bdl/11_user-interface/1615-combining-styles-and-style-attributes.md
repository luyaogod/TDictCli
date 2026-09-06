---
title: "Combining styles and style attributes"
source: "fgl-topics/c_fgl_presentation_styles_008.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Combining styles and style attributes"
type: "concept"
---

# Combining styles and style attributes

> Several styles can be combined, and style attributes can be mixed.

Several style names can be combined, by using the space character as a separator in the
`STYLE` attribute. In the following example, the `STYLE` attribute
references the `info` and `mandatory` style names:

```
EDIT f001 = customer.fname, STYLE = "info mandatory";
```

When several styles define attributes that do not enter in conflict, these style attributes are
combined. For example, the following .4st defines the `info` and
`mandatory` styles using different font attributes:

```
<StyleList>
  <Style name="Edit.info">
    <StyleAttribute name="fontFamily" value="monospace" />
  </Style>
  <Style name="Edit.mandatory">
    <StyleAttribute name="fontWeight" value="bold" />
  </Style>
</StyleList>
```

The `EDIT` widgets having a `STYLE` attribute set to `"info
mandatory"` will get all font attributes defined by the referenced styles.

## Related links

**Related concepts**  

[STYLE attribute](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.")
