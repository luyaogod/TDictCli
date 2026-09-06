---
title: "Static images"
source: "fgl-topics/c_fgl_images_static_images.html"
breadcrumb: "User interface > Form definitions > Using images > Static images"
type: "concept"
---

# Static images

> Describes how to decorate forms with icons.

## Static image usage context

Static images are application pictures that do not change during program executing,
like icons in toolbar buttons and window icons.

Static images can be defined in different contexts withing form definition, or
configuration files:

- Global application icon for platform window managers (taskbars), by using the
  [`ui.Interface.setImage()`](../15_library-reference/3119-ui-interface-setimage.md "Defines the icon image of the program.")
  method. The recommendation for mobile devices is that the application icon is provided in the
  installation package (.apk for Android™, .ipa for iOS).
- Window specific icons, with the `IMAGE` attribute in the [`LAYOUT` definition](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") of a form
  (recommended) or at runtime, with the [`ui.Window.setImage()`](../15_library-reference/3138-ui-window-setimage.md "Set the window icon.") method (if it must be changed during program
  execution).
- As default icon for action views, with the [`IMAGE`](2266-action-attributes-list.md) action configuration attribute (in
  action defaults for example).
- As specific action view icons, directly in the form item definition with the
  [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.")
  attribute (for toolbars, menu items, buttons, buttonedits, etc).
- Image form items (logos), defined by the [`IMAGE item-tag :
  item-name`](1739-image-item-definition.md "Defines attributes for an area that can display an image resource.") syntax, using the `IMAGE`
  attribute.
- Default treeview node icons, with the [`IMAGEEXPANDED`](1788-imageexpanded-attribute.md "The IMAGEEXPANDED attribute sets the global icon to be used when a tree node is expanded."), [`IMAGECOLLAPSED`](1787-imagecollapsed-attribute.md "The IMAGECOLLAPSED attribute sets the global icon to be used when a tree node is collapsed."), [`IMAGELEAF`](1789-imageleaf-attribute.md "The IMAGELEAF attribute defines the global icon for leaf nodes of a TREE container.")
  attributes of a [`TREE`](1749-tree-item-definition.md "Defines attributes for a tree layout tag.") container.

## Static image examples

The following code example, defines an `ITEM` toolbar element using a
icon, that is specified with the `IMAGE`
attribute:

```
TOOLBAR
   ITEM print ( TEXT="Print", IMAGE="printer" )
```

Next example defines a `BUTTONEDIT` form field with an icon named
"listchoice":

```
ATTRIBUTES
BUTTONEDIT f05 = customer.cust_city,
    ACTION=get_city,
    IMAGE="listchoice",
    ... ;
```

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
