---
title: "GROUP item type"
source: "fgl-topics/c_fgl_FormSpecFiles_GROUP.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > GROUP item type"
type: "concept"
---

# GROUP item type

> Defines a layout area to group other layout elements together.

## GROUP item basics

A `GROUP` form item type groups other form items together, typically in
a groupbox widget.

![GROUP rendering](../_images/FormItemType_GROUP_1.jpg)

*GROUP form item type*

## Defining an GROUP

The `GROUP` form item typically gets a [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") attribute, to
define the title of the group. Consider using [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") for this
attribute:

```
GROUP ...
  TEXT=%"customer.info";
```

Consider identifying group elements with a name, in order to manipulate the group
during program execution. For example use the [`ui.Form.setElementHidden()`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.") method to hide or show groups in a
form:

```
GROUP g1: g_cust_info, ... ;   -- grid-based layout
GROUP g_cust_info, ... ;       -- stack-based layout
```

Front-ends support different presentation and behavior options, which can be
controlled by a [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute.
For more details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and
[Group style attributes](1639-group-style-attributes.md "Group presentation style attributes apply to an GROUP element.").

## Where to use a GROUP

A `GROUP` form item can be defined in different ways:

1. As a [`GROUP`
   container](1719-group-container.md "Defines a layout area to group other layout elements together.") in a `LAYOUT` tree.
2. As a [`<GROUP >` layout
   tag](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.") with a [`GROUP` item
   definition](1738-group-item-definition.md "Defines attributes for a group-box layout tag.") in the `ATTRIBUTES` section.

## Group containers

In a `LAYOUT` tree with [`GROUP`
containers](1719-group-container.md "Defines a layout area to group other layout elements together."), if you want to include several children in a
`GROUP`, you can add a `VBOX` or `HBOX`
into the `GROUP`, to define how these form items are aligned.

> **Note:**
>
> When defining a `GROUP` container,
> you cannot set the [`GRIDCHILDRENINPARENT`](1781-gridchildreninparent-attribute.md "The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container.")
> attribute. This attribute makes sense only for a group item defined with a layout tag contained in a
> `GRID` area.

Consider using a group layout tag inside a `GRID` container, this
layout specification technique is often more appropriate to define forms:

```
GRID
{
<G g1             ><G g2        >
[l1  :f1          ][f4          ]
 ...
<G g3                            >
 ...
```

## Collapsible groups

By default, groups are not collapsible.

Some front-ends (see [Group
presentation style attributes reference](1639-group-style-attributes.md "Group presentation style attributes apply to an GROUP element.")) support the `collapsible`
presentation style attribute, to let end users expand/collapse `GROUP` elements in
your forms.

When a group is defined as collapsible, the `collapserPosition` style attribute
can be used to define the position of the collapser icon.

The `initiallyCollapsed` style attribute defines if the collapsible group is
collapsed or expanded when the form is displayed. Values can be `"yes"`,
`"no"`, `"never"` and `"always"`.

For example:

```
<Style name="Group.mystyle">
   <StyleAttribute name="collapsible" value="yes" />
   <StyleAttribute name="collapserPosition" value="right" />
   <StyleAttribute name="initiallyCollapsed" value="yes" />
</Style>
```

For more details, see [Group style attributes](1639-group-style-attributes.md "Group presentation style attributes apply to an GROUP element.").
