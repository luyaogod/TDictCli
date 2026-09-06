---
title: "gICAPI.UseGbcThemeVariables()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_usegbcthemevariables.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.UseGbcThemeVariables()"
type: "concept"
---

# gICAPI.UseGbcThemeVariables()

> The gICAPI.UseGbcThemeVariables() function can be used to declare a set of GBC theme variables to be referenced in the CSS of a web component.

## Purpose of gICAPI.UseGbcThemeVariables()

The values defined for GBC theme variables can be propagated into the HTML/CSS code of your
gICAPI webcomponent.

This allows to adapt the rendering of the webcomponent to the settings used in GBC.

## Providing the list of GBC theme variables

The `gICAPI.UseGbcThemeVariables()` function must be called with an array of
strings, declaring the GBC theme variables to be used in the CSS of your web component.

The variable names referenced in the array passed in the `UseGbcThemeVariables()`
call must be defined in your theme.scss.json GBC theme file.

In the CSS of your webcomponent HTML code, the value of a GBC theme variable can then be
referenced by prefixing the variable name with two minus signs, in the `var()`
function:

```
  "some-var-name"    =>    var(--some-var-name)
```

## Example

The following code example uses the `gICAPI.UseGbcThemeVariables()` function to
declare a set of background color GBC theme variables, to be used in CSS definitions:

```
var onICHostReady = function(version) {

    ...

    gICAPI.UseGbcThemeVariables([
        "theme-primary-background-color",
        "theme-secondary-background-color",
        "theme-field-background-color"
    ]);

    ...

};

...

.container {
  background-color: var(--theme-field-background-color);
}
```
