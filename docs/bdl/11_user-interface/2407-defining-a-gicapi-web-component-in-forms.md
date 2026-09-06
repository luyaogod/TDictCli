---
title: "Defining a gICAPI web component in forms"
source: "fgl-topics/c_fgl_webcomponent_gicapi_form_item.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Defining a gICAPI web component in forms"
type: "concept"
---

# Defining a gICAPI web component in forms

> When defining a gICAPI web component in a form specification file, you can also provide a sizing policy and define additional properties.

## Adding a WEBCOMPONENT to the form file

To define an gICAPI web component field, add a form field with the `WEBCOMPONENT`
item type and the `COMPONENTTYPE` attribute. The [`COMPONENTTYPE`](1771-componenttype-attribute.md "The COMPONENTTYPE attribute defines a name identifying the external widget for WEBCOMPONENT fields.") attribute is
mandatory when defining a gICAPI web component; it defines the root HTML filename describing the
gICAPI web component.

A web component field is typically defined with the
`FORMONLY` prefix, as the data for the field is rarely stored in a database
column.

## Sizing policy for web component fields

Web components are usually complex widgets displaying
detailed information, such as charts, graphs, or calendars, which are generally resizable. Use the
appropriate form item attributes to get the expected layout and behavior. For more details, see
[Controlling the web component layout](2381-controlling-the-web-component-layout.md).

## Defining gICAPI web component properties

Since web component field definitions are generic, you must use the [`PROPERTIES`](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics.") attribute to set
specific parameters for the component.

The `PROPERTIES` attribute can define a list of:

- simple properties ( `name = value` ),
- array properties ( `name = ( value1, value2, ... )` )
- map/dictionary properties ( `name=( name1=value1,name2=value2, ... )` )

where *name* is a simple identifier, and where *values* can
be numeric or string literals.

Component properties defined in the `PROPERTIES` attribute are transmitted to the web
component through the [`onProperty()`](2393-the-gicapi-web-component-interface-script.md) method of the `gICAPI` object.

The name of a property defined in the `PROPERTIES` attribute is converted to
lowercase by the form compiler. To avoid mistakes, a good programming pattern is to
define properties in lowercase, in both the [interface script](2393-the-gicapi-web-component-interface-script.md "The gICAPI web components are controlled on the front-end through a gICAPI interface object, defined in a JavaScript script.") and in the form
definition file. Property names are not checked at compile time, so nonexistent or
mistyped properties will be ignored at runtime.

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
         COMPONENTTYPE = "3DCharts",
         STRETCH = BOTH,
         PROPERTIES = ( type = "bars",
                        x_label = "Months",
                        y_label = "Sales" );
END
```

## Related links

**Related concepts**  

[WEBCOMPONENT item type](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.")
