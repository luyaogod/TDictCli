---
title: "HTML document and JavaScript for the gICAPI object"
source: "fgl-topics/c_fgl_webcomponent_gicapi_html.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > HTML document and JavaScript for the gICAPI object"
type: "concept"
---

# HTML document and JavaScript for the gICAPI object

> A gICAPI web component is identified by an HTML document containing the JavaScript interface (or a reference to the .js file).

The HTML document is defined by the [`COMPONENTTYPE`](1771-componenttype-attribute.md "The COMPONENTTYPE attribute defines a name identifying the external widget for WEBCOMPONENT fields.") attribute of the `WEBCOMPONENT` form field.
The name specified in this attribute will be used to identify the HTML
file:

```
WEBCOMPONENT wc = FORMOMLY.chart,
   COMPONENTTYPE = "mychart"; -- Identifies "mychart.html"
```

The HTML document must reference (or contain) the JavaScript implementing the gICAPI
interface:

```
<!DOCTYPE html>
<html>
<head>
  <title>The title</title>
  <script language="JavaScript" type="text/javascript" src="wc_echo.js"></script>
</head>
<body>
<div style="background-color:green;width:3000px;height:3000px;" >
here
</div>
</body>
</html>
```
