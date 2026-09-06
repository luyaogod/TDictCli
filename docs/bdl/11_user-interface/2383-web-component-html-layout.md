---
title: "Web component HTML layout"
source: "fgl-topics/c_fgl_webcomponent_layout_viewport.html"
breadcrumb: "User interface > User interface programming > Web components > Controlling the web component layout > Web component HTML layout"
type: "concept"
description: "Stretchable WEBCOMPOMENT form field In order to get a stretchable HTML content in your web component, start by defining the WEBCOMPONENT form field with the following attributes: WEBCOMPONENT ..."
---

# Web component HTML layout

## Stretchable WEBCOMPOMENT form field

In order to get a stretchable HTML content in your web component, start by defining the
`WEBCOMPONENT` form field with the following
attributes:

```
WEBCOMPONENT wc1=FORMONLY.wc1,
   COMPONENTTYPE="mywebcomp",
   SIZEPOLICY=FIXED,
   STRETCH=BOTH,
   SCROLLBARS=NONE;
```

## Viewport zooming on mobile devices

In order to avoid automatic viewport zooming with mobile applications, consider adding a meta tag
with `name='viewport'` in the HTML file of your gICAPI-based web components, with
initial and maximal scale attributes set to
1:

```
<meta name='viewport' content='initial-scale=1.0, maximum-scale=1.0' />
```

Don not use such responsive meta tag, if your web component isn't specifically designed to be
responsive.

## HTML body auto-resize

In order to force the resizing of the HTML content to the parent container, use the following CSS
style:

```
<style>

html,body {
  height:100%;
  padding:0;
  margin:0;
  border:0;
}

...

</style>
```

## Controlling SVG layout

When displaying SVG inside your web component, you need to define how the SVG adapts to the
parent container.

SVG root elements define their own viewport (`width` and `height`
attributes), and local layout rules through `viewBox` and
`preserveAspectRatio` attributes.

For general purpose, do not set the `width` and `height`
attributes in your root SVG element, and use `preserveAspectRatio="xMidYMid meet"`.

In the CSS style of the HTML page, define CSS styles for a parent div and for the svg root
element, depending on the needs (to get scrollbars for example):

```
<style>

...

.svg_container_nsb {
  overflow: hidden;
}

.svg_container_hsb {
  overflow-x: scroll;
  overflow-y: hidden;
}

.svg_container_vsb {
  overflow-x: hidden;
  overflow-y: scroll;
}

.root_svg_max_h {
  max-height: 100%;
}

.root_svg_h_max_h {
  height: 100%;
  max-height: 100%;
}

...

</style>
```

Control
the layout of the svg root element by using the appropriate CSS styles (typically set when building
your SVG by program):

SVG adapting to container, no scrollbars in div container (all SVG will
be
visible):

```
<body>
<div id="svg_container" class="svg_container_nsb">
   <svg . viewBox="-15 -10 158 143" preserveAspectRatio="xMidYMid meet" ... class="root_svg_max_h" ...>
      ...
   </svg>
</div>
</body>
```

SVG
adapting to container, with vertical scrollbars in the div container (SVG height is much bigger than
the width):

```
<body>
<div id="svg_container" class="svg_container_vsb">
   <svg ... viewBox="-20 -7 268 727" preserveAspectRatio="xMidYMin slice" ... >
      ...
   </svg>
</div>
</body>
```

SVG
adapting to container, with horizontal scrollbars in the div container (SVG width is much bigger than
the height):

```
<body>
<div id="svg_container" class="svg_container_hsb">
   <svg ... viewBox="-20 -7 268 727" preserveAspectRatio="xMidYMin slice" class="root_svg_h_max_h" ... >
      ...
   </svg>
</div>
</body>
```
