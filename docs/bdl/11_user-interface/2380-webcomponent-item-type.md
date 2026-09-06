---
title: "WEBCOMPONENT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_WEBCOMPONENT_2.html"
breadcrumb: "User interface > User interface programming > Web components > WEBCOMPONENT item type"
type: "concept"
---

# WEBCOMPONENT item type

> Defines a specialized form item that holds an external component.

## WEBCOMPONENT item basics

The `WEBCOMPONENT` form item defines a form field that will hold an
external component, implemented with a front-end plug-in mechanism.

![WEBCOMPONENT rendering](../_images/FormItemType_WEBCOMPONENT_1.jpg)

*WEBCOMPONENT form item type*

This topic describes the `WEBCOMPONENT` item type in form definition
files. For more details see [the chapter dedicated
to web component programming](2378-web-components.md "This section describes how to use web components in your application.").

## Defining a WEBCOMPONENT

The [`COMPONENTTYPE`](1771-componenttype-attribute.md "The COMPONENTTYPE attribute defines a name identifying the external widget for WEBCOMPONENT fields.")
attribute identifies gICAPI external objects to be used for the field. The [`PROPERTIES`](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics.") attribute is
typically used to define attributes that are specific to a given gICAPI-based web component.
For example, a chart component might have properties to define x-axis and y-axis labels.
For more details, see [Using a gICAPI web component](2391-using-a-gicapi-web-component.md "This section describes how to add a gICAPI-based web component to your application.").

If the `COMPONENTTYPE` attribute is not used, the web component will be a
URL-based web component. For more details, see [Using a URL-based web component](2385-using-a-url-based-web-component.md "This section describes how to add a URL-based web component to your application.").

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.").

## Where to use a WEBCOMPONENT

A `WEBCOMPONENT` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [WEBCOMPONENT item definition](1750-webcomponent-item-definition.md "Defines attributes for a generic form field that can receive an external widget.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Built-in Web Components

Genero BDL provides a set of ready-to-use web components, that are deployed by default.

For more details, see [Built-in web components](2415-built-in-web-components.md "Genero provides a set of ready-to-use web components.").

## Defining the widget size

The size of a `WEBCOMPONENT` widget can be controlled with several
attributes such as `SIZEPOLICY` and `STRETCH`.

For more details about webcomponent sizing, see [Controlling the web component layout](2381-controlling-the-web-component-layout.md).

## Related links

**Related concepts**  

[SIZEPOLICY attribute](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.")
