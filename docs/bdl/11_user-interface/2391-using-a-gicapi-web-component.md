---
title: "Using a gICAPI web component"
source: "fgl-topics/t_fgl_webcomponent_gicapi.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component"
type: "task"
---

# Using a gICAPI web component

> This section describes how to add a gICAPI-based web component to your application.

To implement a gICAPI-based web component:

1. Identify the web component you want to use and get the source
   code (HTML, JavaScript,
   CSS).
2. Implement the gICAPI interface script for the web component.
3. Define the location where the front-end can find the gICAPI
   interface files. This depends on the front-end technology used by your application.
4. Define a `WEBCOMPONENT` field in the form file. Use the
   `COMPONENTTYPE` attribute to define the root HTML filename describing the gICAPI web
   component.
5. Use the web component in the dialog of the program.
6. If image resources are required by your web component, you must provide them as part of the
   gICAPI web component assets, or provide them from the program with a specific
   API.

Detailed information about these tasks are provided in the next topics.

## Child topics

- [HTML document and JavaScript for the gICAPI object](2392-html-document-and-javascript-for-the-gicapi-object.md): A gICAPI web component is identified by an HTML document containing the JavaScript interface (or a reference to the .js file).
- [The gICAPI web component interface script](2393-the-gicapi-web-component-interface-script.md): The gICAPI web components are controlled on the front-end through a gICAPI interface object, defined in a JavaScript script.
- [Deploying the gICAPI web component files](2403-deploying-the-gicapi-web-component-files.md): Deploy web component files to the front-end platform before using gICAPI web components.
- [Defining a gICAPI web component in forms](2407-defining-a-gicapi-web-component-in-forms.md): When defining a gICAPI web component in a form specification file, you can also provide a sizing policy and define additional properties.
- [Controlling the gICAPI web component in programs](2408-controlling-the-gicapi-web-component-in-programs.md)
- [Using image resources with the gICAPI web component](2409-using-image-resources-with-the-gicapi-web-component.md): This section explains how to use image resources in a gICAPI web component.
- [Examples](2410-examples.md): GICAPI Web Component usage examples.
