---
title: "Order of precedence"
source: "fgl-topics/c_fgl_presentation_styles_precedence.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Order of precedence"
type: "concept"
---

# Order of precedence

> Style definitions are applied following and order of precedence.

If different styles can be applied to an element, the following priority is used to determine the
style definition to be applied:

1. `element-type.style-name:pseudo-selector`
2. `.style-name:pseudo-selector`
3. `element-type.style-name`
4. `element-type:pseudo-selector`
5. `:pseudo-selector`
6. `.style-name`
7. `element-type`
8. `*`

The precedence rules to apply styles can be specific to a front-end type. As a general rule,
Genero presentation styles precedence rules are similar to HTML/CSS precedence rules.

For example, consider an `Edit` element with the style attribute set to
'`mandatory`':

```
EDIT f1 = FORMONLY.cust_name, STYLE="mandatory"
```

With the following style definitions
(mystyles.4st):

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<StyleList>
  <Style name="Edit.mandatory:focus">
     <StyleAttribute name="backgroundColor" value="yellow" />
  </Style>
  <Style name=".mandatory:focus">
     <StyleAttribute name="backgroundColor" value="blue" />
  </Style>
  <Style name="Edit.mandatory">
     <StyleAttribute name="backgroundColor" value="green" />
  </Style>
  <Style name="Edit:focus">
     <StyleAttribute name="backgroundColor" value="red" />
  </Style>
  <Style name=":focus">
     <StyleAttribute name="backgroundColor" value="cyan" />
  </Style>
  <Style name=".mandatory">
     <StyleAttribute name="backgroundColor" value="magenta" />
  </Style>
  <Style name="*">
     <StyleAttribute name="backgroundColor" value="orange" />
  </Style>
</StyleList>
```

The style definitions are scanned in the following order:

1. `Edit.mandatory:focus`
2. `.mandatory:focus`
3. `Edit.mandatory`
4. `Edit:focus`
5. `:focus`
6. `.mandatory`
7. `Edit`
8. `*`

If the Edit field `f1` has the focus, with the mystyles.4st
definition file, the field background color will be yellow. If the Edit field `f1`
does not have the focus, the field background color will be green.
