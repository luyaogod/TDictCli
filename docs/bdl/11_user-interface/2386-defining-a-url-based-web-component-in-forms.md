---
title: "Defining a URL-based web component in forms"
source: "fgl-topics/c_fgl_webcomponent_url_form_item.html"
breadcrumb: "User interface > User interface programming > Web components > Using a URL-based web component > Defining a URL-based web component in forms"
type: "concept"
---

# Defining a URL-based web component in forms

> A URL-based web component needs to be declared in the form definition.

## Adding a WEBCOMPONENT to the form file

To define a URL-based web component field, add a form field with the
[`WEBCOMPONENT` item type](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component."),
without the `COMPONENTTYPE` attribute:

```
WEBCOMPONENT f001 = FORMONLY.mymap;
```

A web component field is typically defined with the
`FORMONLY` prefix, as the data for the field is rarely stored in a database
column.

The field type (and its corresponding program variable) must be a character string type. Consider
using the `STRING` type to avoid any size limitation for the URL specification.

## Sizing policy for web component fields

Web components are usually complex widgets displaying
detailed information, such as charts, graphs, or calendars, which are generally resizable. Use the
appropriate form item attributes to get the expected layout and behavior. For more details, see
[Controlling the web component layout](2381-controlling-the-web-component-layout.md).

## Example

```
LAYOUT
GRID
{
[wc                       ]
[                         ]
[                         ]
[                         ]
[                         ]
}
END
END
ATTRIBUTES
WEBCOMPONENT wc = FORMONLY.mychart,
         STRETCH = BOTH;
END
```

## Related links

**Related concepts**  

[Example 1: URL-based web component using Wikipedia](2390-example-1-url-based-web-component-using-wikipedia.md "Example 1: URL-based web component using Wikipedia")
