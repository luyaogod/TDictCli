---
title: "Understanding web components"
source: "fgl-topics/c_fgl_webcomponent_intro.html"
breadcrumb: "User interface > User interface programming > Web components > Understanding web components"
type: "concept"
---

# Understanding web components

> External graphical components can be integrated into forms by using the WEBCOMPONENT form item type.

A `WEBCOMPONENT` form field is a form element that defines an area in the
form layout to hold an external component, typically not available as a native widget on the
front-end platform.

The purpose of a web component form field is not to replace a web browser, PDF viewer, or to
implement sub-forms in HTML. A `WEBCOMPONENT` field must behave as a single form
field and integrate with the other standard Genero form fields and form elements. When implementing
a [gICAPI webcomponent](2391-using-a-gicapi-web-component.md "This section describes how to add a gICAPI-based web component to your application."), the JavaScript code must
keep the focus in the area reserved for the field. For example, the JS code should not create
another window or tab with the `window.open()` API.

This section describes how to implement your own web component form fields. Genero provides a set
of ready-to-use web components that are available for all front-ends. For more details, see [Built-in web components](2415-built-in-web-components.md "Genero provides a set of ready-to-use web components.").

GBC runs the web component in an HTML iframe, with the limitations that come with the iframe
technology. For example, you should not mix http and https within an iframe. Refer to HTML standards
for more details.

Web component form fields are used to fulfill a specific display and/or input, with advanced and
powerful features which can bring added value to your applications. For example, you can find chart
and graph widgets, calendar widgets, drawing widgets, and more. Such specialized widgets are not
part of the standard GUI toolkits used by Genero front-ends. They need to be integrated as external
components with `WEBCOMPONENT` fields.

Web components available on the Internet may be licensed and their usage chargeable. It is
recommended to take the cost of external web components into account, before integrating these in
your application. Similarly, embedding third party JavaScript framework in self-designed web
components may have cost and licensing implications.

Web components can be implemented with two different techniques:

1. [Using a URL specification](2385-using-a-url-based-web-component.md "This section describes how to add a URL-based web component to your application."), by setting the URL
   as value of the `WEBCOMPONENT` field at runtime. This is the easiest way to implement
   a web component. The widget is controlled with URL values by the program, but requires some
   additional coding to handle URLs, instead of flat field values.
2. [Using an gICAPI object](2391-using-a-gicapi-web-component.md "This section describes how to add a gICAPI-based web component to your application.") (based on JavaScript),
   by defining the `COMPONENTTYPE` attribute in the form file. This kind of web
   component requires some JavaScript coding, to write a form field "plugin", which is usable in a
   normal dialog instruction, that behaves as all the other widgets in terms of value
   setting/getting.

The content and/or behavior of a web component can be controlled in the program code by using the
field value. To detect events inside the web component, the program dialogs must implement an
`ON CHANGE` control block, that will be fired immediately after a user action on the
web component.

## Related links

**Related concepts**  

[ON CHANGE block](1943-on-change-block.md "ON CHANGE block")
