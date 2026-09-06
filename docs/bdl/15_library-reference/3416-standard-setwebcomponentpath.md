---
title: "standard.setWebComponentPath"
source: "fgl-topics/c_fgl_frontcall_standard_setwebcomponentpath.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.setWebComponentPath"
type: "concept"
---

# standard.setWebComponentPath

> Defines the base path where web components are located.

> **Important:**
>
> This feature is provided for development and
> debug purpose, and must not be used in a production environment.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
ui.Interface.frontCall("standard", "setWebComponentPath",
  [path], [])
```

1. path - The base URL. For example,
   `"http://myserver/components"` or `"file:///c:/components"`.

## Usage

This front call defines the base path to find gICAPI web components files.

The `standard.setWebComponentPath` front call is deprecated. Consider using
deployment solutions described in [Deploying the gICAPI web component files](../11_user-interface/2403-deploying-the-gicapi-web-component-files.md "Deploy web component files to the front-end platform before using gICAPI web components.").

For the Genero Desktop Client, it defines the base path where web components are located, when
GDC is directly connected to the runtime system. This is ignored when GDC is connected to the
GAS.

For Genero Mobile, it sets the main web component lookup path. An URI is expected. For example,
"file:///data/data/com.fourjs.gma/cache/appdata/mywebcomponents"or
"http://mygas/mywebcomponents/".
