---
title: "PIXELWIDTH / PIXELHEIGHT are deprecated"
source: "fgl-topics/c_fgl_Migrate_to_200_pixelwidth_pixelheight.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > PIXELWIDTH / PIXELHEIGHT are deprecated"
type: "concept"
---

# PIXELWIDTH / PIXELHEIGHT are deprecated

> Use the WIDTH and HEIGHT attributes to specify the size of an image.

Before version 2.00, the `PIXELWIDTH` and `PIXELHEIGHT` attributes
were used to specify the real size of an `IMAGE` form item.

Starting with version 2.00, you must use the `WIDTH attribute` and `HEIGHT attribute` to specify the size of an
image:

In the .per form file:

```
IMAGE img1 = FORMONLY.image1,
      HEIGHT = 100 PIXELS,
      WIDTH = 100 PIXELS;
```

The `PIXELWIDTH` and `PIXELHEIGHT` attributes are still
supported by the form compiler, but are deprecated and will be removed in a future
version.

## Related links

**Related concepts**  

[IMAGE item type](../11_user-interface/1695-image-item-type.md "Defines an area that can display an image resource.")
