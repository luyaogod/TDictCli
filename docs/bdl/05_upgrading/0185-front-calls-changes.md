---
title: "Front calls changes"
source: "fgl-topics/c_fgl_Migrate_to_310_frontcalls.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Front calls changes"
type: "concept"
---

# Front calls changes

> Modifications to consider when using front calls.

## `session.*` front call are desupported

The `session.setVar` and `session.getVar` front calls are no longer
supported in Genero 3.10: The `session` front calls rely on the GAS
"`wa`" protocol (html5Proxy). This protocol is only supported only by GWC-HTML5
client. html5Proxy is no longer delivered in Genero 3.10.

In order to store information on the front-end side, use the `localStorage` front
calls. For more details, see [New localStorage frontcalls](0200-new-localstorage-frontcalls.md "New localStorage frontcalls replace GAS specific session.setVar and session.getVar calls.").

## `standard.setWebComponentPath`

The [`standard.setWebComponentPath`](../15_library-reference/3416-standard-setwebcomponentpath.md "Defines the base path where web components are located.") front call is deprecated in Genero BDL 3.10.
To deploy your web components, consider using the solutions described in [Deploying the gICAPI web component files](../11_user-interface/2403-deploying-the-gicapi-web-component-files.md "Deploy web component files to the front-end platform before using gICAPI web components.").

## `standard.feInfo: outputMap`

Starting with Genero BDL 3.10, the `outputMap` information can no longer be used
with the [`standard.feInfo`](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.")
front call.

## `standard.cbSet` supported by GBC

Since GBC 1.00.35, the GBC supports the [`standard.cbSet`](../15_library-reference/3392-standard-cbset.md "Set the content of the clipboard.") front call.

## `mobile.isForeground` to check app foreground mode

Since Genero BDL 3.10.11 (GMA 1.30.10, GMI 1.30.11), the `mobile.isForeground`
front call can be used to check if the app is in foreground mode.

For more details, see [`mobile.isForeground`](../15_library-reference/3455-mobile-isforeground.md "Indicates if the mobile app is in foreground mode.") front call.

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")

[Web components](../11_user-interface/2378-web-components.md "This section describes how to use web components in your application.")
