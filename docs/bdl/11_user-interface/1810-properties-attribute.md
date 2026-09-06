---
title: "PROPERTIES attribute"
source: "fgl-topics/c_fgl_FSFAttributes_PROPERTIES.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > PROPERTIES attribute"
type: "concept"
---

# PROPERTIES attribute

> The PROPERTIES attribute is used to define a list of widget-specific characteristics.

## Syntax

```
PROPERTIES = (
   { single-property
   | array-property
   | map-property
   } [,...]
 )
```

where single-property
is:

```
identifier = property-value
```

and array-property
is:

```
identifier = ( property-value [,...] )
```

and map-property
is:

```
identifier = (
    { single-property
    | map-property
    } [,...]
 )
```

and property-value
is:

```
{ boolean-value
| numeric-value
| [%]"string-value"
}
```

1. boolean-value is `true` or `false`.
2. numeric-value is an integer or decimal literal.
3. string-value is a string literal. `%` prefix can be used to
   define a localized string.

## Usage

The `PROPERTIES` attribute is typically used to define the widget-specific
attributes of a `WEBCOMPONENT` form item.

Property names and values are not checked, which allows you to freely set any characteristic of
an external widget. Verify that the web component implementation supports the specified
properties.

For more details about properties usage, see [Web components](2378-web-components.md "This section describes how to use web components in your application.").

## Example

```
WEBCOMPONENT f01 = FORMONLY.mycalendar,
   COMPONENTTYPE = "calendar",
   PROPERTIES = (
               type = "gregorian",
               caption = %"calendar.caption",
               week_start = 2,
               days_off = ( 1, 7 ),
               dates_off = ( "????-11-25", "????-06-20" ),
               day_titles = ( t1 = "Sunday",
                              t2 = "Monday",
                              t3 = "Tuesday",
                              t4 = "Wednesday",
                              t5 = "Thursday",
                              t6 = "Friday",
                              t7 = "Saturday" )
             );
```

## Related links

**Related concepts**  

[gICAPI.onProperty()](2398-gicapi-onproperty.md "The gICAPI.onProperty() function is executed when web component properties change.")

[WEBCOMPONENT item type](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.")
