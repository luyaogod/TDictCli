---
title: "Controlling the web component layout"
source: "fgl-topics/c_fgl_webcomponent_layout.html"
breadcrumb: "User interface > User interface programming > Web components > Controlling the web component layout"
type: "concept"
description: "Web components are usually complex widgets displaying detailed information, such as charts, graphs, or calendars. Such widgets are generally resizable. Therefore, the WEBCOMPONENT form item must be ..."
---

# Controlling the web component layout

Web components are usually complex widgets displaying detailed
information, such as charts, graphs, or calendars. Such widgets are generally resizable. Therefore,
the `WEBCOMPONENT` form item must be large and stretchable.

As web components are displayed in an individual web viewer, several layout concepts need to be
considered:

- The layout attributes of the `WEBCOMPONENT` form item, such as
  `SIZEPOLICY`, `SCROLLBARS`, `STRETCH` and
  `HEIGHT`.
- By default, a `WEBCOMPONENT` form item is stretchable
  (`STRETCH=BOTH`).
- By default, a `WEBCOMPONENT` form item gets a vertical scrollbar
  (`SCROLLBARS=VERTICAL`).
- The layout attributes of the HTML root and body element (viewport, height).
- The layout attributes of responsive elements such as SVG content.

## Child topics

- [Web component grid layout](2382-web-component-grid-layout.md)
- [Web component HTML layout](2383-web-component-html-layout.md)
