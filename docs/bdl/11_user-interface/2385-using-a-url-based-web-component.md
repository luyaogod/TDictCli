---
title: "Using a URL-based web component"
source: "fgl-topics/t_fgl_webcomponent_url.html"
breadcrumb: "User interface > User interface programming > Web components > Using a URL-based web component"
type: "task"
---

# Using a URL-based web component

> This section describes how to add a URL-based web component to your application.

To implement a URL-based web component:
> **Note:**
>
> For URL-based web components, many websites – for security reasons – are configured to not allow
> the site content to be displayed in an iframe; by itself, the Genero Browser Client cannot override
> this restriction set by the website. Some sites will provide an API that allow the embedding of the
> site content into an iframe; for example, Google Maps provides an API that allow you to share and
> embed in an iframe.

URL-based web components are hosted on a third party server and provide a specific service,
such as a geographical location on a map. Your application will be dependent on this external
service. Consider verifying its permanent availability.

1. Identify the URL of the hosted web component you want to use.
2. In the form file, define a `WEBCOMPONENT` field, without a
   `COMPONENTTYPE` attribute.
3. In the program, set the URL of the hosted web component in the form field value.
4. In the program, detect user interactions with an `ON CHANGE` control
   block, and control the URL-based web component with dedicated front calls.

Detailed information about these tasks are provided in the next topics.

## Child topics

- [Defining a URL-based web component in forms](2386-defining-a-url-based-web-component-in-forms.md): A URL-based web component needs to be declared in the form definition.
- [Specifying the URL source of a web component](2387-specifying-the-url-source-of-a-web-component.md): The content of URL-based web components is defined by the form field value. It can only be set by program.
- [Controlling the URL web component in programs](2388-controlling-the-url-web-component-in-programs.md): URL-based web components can be controlled with the field value and with front calls
- [Examples](2389-examples.md): URL-based Web Components usage examples.
