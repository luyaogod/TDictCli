---
title: "gICAPI.onProperty()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_onproperty.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.onProperty()"
type: "concept"
---

# gICAPI.onProperty()

> The gICAPI.onProperty() function is executed when web component properties change.

## Purpose of gICAPI.onProperty()

If defined, the `gICAPI.onProperty()` function is called to get the list of
properties of the `<PropertyDict >` AUI node. This node is created in the
.42f form file under the form field node, when using the [`PROPERTIES`](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics.") attribute in a
`WEBCOMPONENT` field.

`gICAPI.onProperty()` may also be called when `<PropertyDict >`
is changed at runtime, after the form was loaded.

This function gets a single string parameter, that holds all current properties defined for the
gICAPI web component.

The property list is provided as a JSON-formatted string.

Each time the `gICAPI.onProperty()` function is called, all properties are provided
in the JSON parameter.

For each `gICAPI.onProperty()` call, the complete list of existing properties is
passed to the function. In order to detect changes, compare the current set of properties with the
new set of properties passed to the `gICAPI.onProperty()` function.

## Converting properties JSON string to a JSON object

The `gICAPI.onProperty()` function gets a JSON formatted string, that can be
converted to a JSON object with `JSON.parse()`.

A typical `gICAPI.onProperty()` function starts as follows:

```
gICAPI.onProperty = function(properties) {
    var properties_object = JSON.parse( properties );
    ...
```

## Handling typed property values

In the `<PropertyDict>` nodes of the AUI tree, attribute type information is
not provided. Consequently, if the underlying component expects for example properties with JSON
booleans (true/false), your boolean properties need to be identified explicitly and converted to
JSON booleans.

In the AUI tree, boolean values are set to '0' or '1'. In JSON, a boolean must be
`true` or `false`.

For example, if you define the following `PROPERTIES` attribute for the
`WEBCOMPONENT` form
field:

```
PROPERTIES = (
    -- numeric
    height = 500,
    -- string
    theme = "modern",
    -- boolean
    menubar = TRUE,
    statusbar = FALSE,
    -- array
    toolbar = ( "undo redo", "print preview" )
  )
```

The resulting JSON string passed to the `onProperties()` function will look like
this:

```
{
  "height" : "500",
  "theme" : "modern",
  "menubar" : "1",
  "statusbar" : "0",
  "toolbar" : [ "undo redo", "print preview" ]
}
```

Therefore, the JSON string needs to be parsed and some fields must be converted, before it can be
used as a JSON object.

In order to replace AUI tree boolean fields to JSON booleans, you can implement a generic
function like the
following:

```
// Converts the JSON string from onProperty() into a JSON object.
//
// The function needs the list of properties that can be booleans.
//
// AUI booleans must be converted as follows:
//   .per source  :  PROPERTY ( menubar = FALSE )
//   .42f/AUI XML :  <Property name="menubar" value="0"/>
//   onProperty   :  'menubar' : '0'
//   JS/JSON      :  "menubar" : false
//
var aui_to_json = function(properties, pl_booleans) {
    var pso = JSON.parse(properties, function (key, value) {
        if (value && typeof value === 'string') {
            if ( pl_booleans.indexOf(key) > -1 ) {
                // Make a real boolean from the string value
                if (value === "1" || value === "true") {
                   return true;
                } else if (value === "0" || value === "false") {
                   return false;
                }
            }
        }
        return value;
    });
    return pso;
}
```

The above function can then be used in `onProperty()`, to get a well formatted
JSON object of properties.

## Example 1: Implementing gICAPI.onProperty()

This code example defines a default set of properties, that will be merged with a set of
properties passed to the `gICAPI.onProperty()` function.

Once the new set of properties is created, it is assigned to the underlying web component object
referenced here as `"component"`, by using the `resetProperties()`
function. Here we assume that `component.resetProperties()` expects a valid Java
Script object, with JSON arrays and JSON boolean values.

The `aui_to_json()` function is described in the previous section.

```
var properties = {
  options: [ "fastsearch", "textpattern", "colorpicker" ],
  resize: true,
  menubar: true,
  statusbar: false
};

var onICHostReady = function(version) {

    ...

    gICAPI.onProperty = function(ps) {
        var pso = aui_to_json( ps, ["menubar","toolbar","statusbar"] );
        jQuery.extend(properties, pso);
        component.resetProperties( pso );
    }

    ...

};
```

## Example 2: Setting WEBCOMPONENT properties at runtime

This example shows BDL code that can be used to change the `PropertyDict` AUI
node, in order to set the properties of a `WEBCOMPONENT` field at
runtime:

```
PRIVATE FUNCTION _get_property_node(fieldname, tc, property)
    DEFINE fieldname STRING,
           tc STRING, -- "P"=Property or "A"=PropertyArray
           property STRING
    DEFINE tagname STRING,
           w ui.Window,
           f ui.Form,
           n_ff, n_wc, n_pd, n_p om.DomNode,
           nl om.NodeList,
           is_new BOOLEAN
    CASE tc
      WHEN "P" LET tagname = "Property"
      WHEN "A" LET tagname = "PropertyArray"
      OTHERWISE
        DISPLAY "ERROR: Invalid node type:", tc
        EXIT PROGRAM 1
    END CASE
    LET w = ui.Window.getCurrent()
    LET f = w.getForm()
    LET n_ff = f.findNode("FormField", fieldname)
    LET n_wc = n_ff.getFirstChild()
    LET n_pd = n_wc.getFirstChild()
    IF n_pd IS NULL THEN
      LET n_pd = n_wc.createChild("PropertyDict")
    END IF
    LET nl = n_pd.selectByPath(SFMT("//%1[@name=\"%2\"]",tagname,property))
    IF nl.getLength() = 1 THEN
       LET n_p = nl.item(1)
       LET is_new = FALSE
    ELSE
       LET n_p=n_pd.createChild(tagname)
       CALL n_p.setAttribute("name", property)
       LET is_new = TRUE
    END IF
    RETURN n_p, is_new
END FUNCTION

PUBLIC FUNCTION setProperty(fieldname, property, value)
    DEFINE fieldname STRING,
           property STRING,
           value STRING
    DEFINE n om.DomNode,
           is_new BOOLEAN
    CALL _get_property_node(fieldname, "P", property)
         RETURNING n, is_new
    IF n IS NULL THEN
       DISPLAY "ERROR: Property node could not be found/created."
       EXIT PROGRAM 1
    END IF
    CALL n.setAttribute("value", value)
END FUNCTION

PUBLIC FUNCTION setPropertyBoolean(fieldname, property, value)
    DEFINE fieldname STRING,
           property STRING,
           value BOOLEAN
    DEFINE n om.DomNode,
           is_new BOOLEAN
    CALL _get_property_node(fieldname, "P", property)
         RETURNING n, is_new
    IF n IS NULL THEN
       DISPLAY "ERROR: Property node could not be found/created."
       EXIT PROGRAM 1
    END IF
    CALL n.setAttribute("value", IIF(value,"1","0"))
END FUNCTION

PUBLIC FUNCTION setPropertyArray(fieldname, property, value)
    DEFINE fieldname STRING,
           property STRING,
           value DYNAMIC ARRAY OF STRING
    DEFINE n, p, e om.DomNode,
           is_new BOOLEAN,
           i INTEGER
    CALL _get_property_node(fieldname, "A", property)
         RETURNING n, is_new
    IF n IS NULL THEN
       DISPLAY "ERROR: PropertyArray node could not be found/created."
       EXIT PROGRAM 1
    END IF
    IF NOT is_new THEN
       -- To rebuild child list, remove prop node, then re-create.
       LET p = n.getParent()
       CALL p.removeChild(n)
       LET n = p.createChild("PropertyArray")
       CALL n.setAttribute("name", property)
    END IF
    FOR i=1 TO value.getLength()
        LET e = n.createChild("Property")
        CALL e.setAttribute("value", value[i])
    END FOR
END FUNCTION
```

## Related links

**Related concepts**  

[PROPERTIES attribute](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics.")
