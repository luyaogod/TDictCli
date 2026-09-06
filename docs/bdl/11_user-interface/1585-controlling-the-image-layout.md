---
title: "Controlling the image layout"
source: "fgl-topics/c_fgl_images_layout.html"
breadcrumb: "User interface > Form definitions > Using images > Controlling the image layout"
type: "concept"
---

# Controlling the image layout

> Explains how image form items can be sized in different front-end layout systems.

## Image sizing basics

It is important to differentiate the image and the image container (the widget): when designing a
form, you're defining the image container. The actual image that will be displayed in this container
can be smaller or larger. Genero provides several form file attributes, to control how the image and
its container are sized.

How an `IMAGE` item renders on the front-end screen depends on these
factors:

- The size of the form item tag in the [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") section, or the [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.") and [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.") attributes defined for the `IMAGE` item.
- The combination of image item attributes ([`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`AUTOSCALE`](1762-autoscale-attribute.md "The AUTOSCALE attribute causes the form element contents to automatically scale to the size given to the item."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.")). These attributes may have a limited effect depending on the
  front-end platform.
- The size of the [image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program."), when
  displayed (especially when `SIZEPOLICY=DYNAMIC/INITIAL`).
- The [`scaleIcon`](1651-toolbar-style-attributes.md "ToolBar presentation style attributes apply to the TOOLBAR element.") presentation style attribute, for elements using icons such as
  `BUTTON` or `TOOLBAR` items.

The typical layout settings of an `IMAGE` item is a combination of
`SIZEPOLICY=FIXED`, `AUTOSCALE` and optionally a
`WIDTH`/`HEIGHT`: The fixed size policy is to get an image size that
corresponds the rest of the form layout, while auto-scaling will make the image fit to its container
and avoid scrollbars with large images. The `WIDTH`/`HEIGHT`
attributes may be used to specify the size of the image container. However, the preferred way to
define the size of an image container is the item tag dimensions, in the layout section of the
form.

## Attributes to controlling the image size

The `AUTOSCALE` attribute indicates if the picture must be scaled to the
available space in the image item. The space is defined by the `SIZEPOLICY`
attribute, the `STRETCH` attribute, and the form item size (the form item tag in the
layout or the `WIDTH` and `HEIGHT` attributes).

`AUTOSCALE` is only useful if the size of the
image differs from the size of the container. `AUTOSCALE` is useless with
`SIZEPOLICY=DYNAMIC`, as the container always fits to the image
size.

The `STRETCH` attribute defines how the image item adapts to the parent container
when it is re-sized. The default is `NONE`.

`SIZEPOLICY` and the `WIDTH`/`HEIGHT` attributes
define the size of the container, not the size of the image.

The `SIZEPOLICY` attribute defines how the image widget gets its size,
depending on the context:

- When `SIZEPOLICY` is `INITIAL` (the default) and
  `AUTOSCALE` is not used, the size of the widget is defined by the first picture
  displayed in the form element. The size will not change if other pictures with different sizes
  display in the widget. If no initial image is displayed (the image field value is
  `NULL`), the form item does not take up space in the layout (and also does not adapt
  the size if an image is displayed later on).
- When `SIZEPOLICY` is `DYNAMIC`, the size of the widget is
  automatically adapted to the size of the pictures displayed in the image form item. The
  `AUTOSCALE` attribute has no effect.
- If `SIZEPOLICY` attribute is set to `FIXED`, the size of the
  widget is defined by the form specification file, either by the size of the
  item-tag in the layout, or by the `WIDTH` and
  `HEIGHT` attributes. With a fixed image widget size, if `AUTOSCALE` is
  not used, scrollbars may appear if the picture is greater than the widget.

By default, the size of the image widget defaults to the relative width and height defined by the
item-tag in the form layout section. The size of an image widget can also be
specified in the `WIDTH` and `HEIGHT` attributes, but these attributes
only have an effect when `SIZEPOLICY=FIXED`.

The `WIDTH` and `HEIGHT` attributes define the size of the
container, but they are dependent on the `SIZEPOLICY`. It means the image container
may grow or shrink even if `WIDTH` and `HEIGHT` are specified. If you
really want to have a container with a fixed size, you have to use `WIDTH` and
`HEIGHT` in combination with `SIZEPOLICY=FIXED`.

All image layout attributes (except `AUTOSCALE`) only have an impact on the
container size, not on the image size. If `AUTOSCALE` is defined, the image fits to
the container size, but without losing its initial proportionality. If `STRETCH`
allows the container to grow/reduce in the `X` or `Y` direction, the
image will grow with the container, but the original proportionality is always kept.

On some platforms, the image widgets automatically add a border to the source picture. For these
platforms, if the image form item is the same size as the image, you may need to increase the size
of the image form item to avoid automatic scrollbars. For example, if your image source has a size
of 500x500 pixels and the widget displays a border with a size of 1 pixel, you will have to set
`WIDTH` and `HEIGHT` to 502 pixels. If you do not, scrollbars will
appear or the image will shrink if `AUTOSCALE` is used. Alternatively, you can avoid
the image border with the `border` presentation style attribute.

## Related links

**Related concepts**  

[Form rendering basics](1539-form-rendering-basics.md "Get the essentials about form rendering.")

**Related reference**  

[Button style attributes](1631-button-style-attributes.md "Button presentation style attributes apply to BUTTON elements.")
