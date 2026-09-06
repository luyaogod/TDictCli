---
title: "Deploying gICAPI assets with GAS"
source: "fgl-topics/c_fgl_webcomponent_gicapi_location_gas.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Deploying the gICAPI web component files > Deploying gICAPI assets with GAS"
type: "concept"
---

# Deploying gICAPI assets with GAS

> Using gICAPI web components through the GAS.

When using the Genero Application Server, the gICAPI web component files must be deployed as part
of the application program files.

To simplify deployment of gICAPI web components with the GAS, consider using the
fglgar utility. For more details, see [Packaging web applications](../13_programming-tools/2645-packaging-web-applications.md "Describes methods of packaging the runtime files and resources of your web applications and services using the fglgar tool.").

The .xcf configuration file of your application can define the base path to
search for HTML web component files. This base path is defined by the
`WEB_COMPONENT_DIRECTORY` entry of the `EXECUTION` element:

```
<APPLICATION ...
 <EXECUTION>
  ...
  <WEB_COMPONENT_DIRECTORY>$(res.fgldir)/webcomponents;
    $(application.path)/webcomponents</WEB_COMPONENT_DIRECTORY>
  ...
```

While specific web component deployment directories can be defined in the
.xcf file, consider using [the recommended gICAPI web component directory layout](2403-deploying-the-gicapi-web-component-files.md).

Genero BDL provides a set of standard web components in
$FGLDIR/webcomponents, that can be found with the default
as.xcf settings. If you want to use a standard Genero web component with the
GAS, and you have defined a specific `WEB_COMPONENT_DIRECTORY` entry in your
application .xcf configuration file, you must add
`$(res.fgldir)/webcomponents` to the `WEB_COMPONENT_DIRECTORY`
entry.

The HTML document must be located in a sub-directory below the base path, using the same name as
defined by the `COMPONENTTYPE` attribute. As result, the complete path to the HTML
document will be:

`base-path/component-type/component-type.html`

For example, if the form file defines the `COMPONENTTYPE` attribute as
follows:

```
WEBCOMPONENT wc = FORMOMLY.mychart,
   COMPONENTTYPE = "3DChart";
```

If `application.path` is "/opt/var/gas/appdata/app/myapp",
the HTML document will be found in:

- /opt/var/gas/appdata/app/myapp/webcomponents/3DChart/3DChart.html

The above .xcf example shows the default value of the
`WEB_COMPONENT_DIRECTORY` parameter that can be inherited by all application
configuration nodes. If your gICAPI web component files are located under
appdir/webcomponents, or if your programs use one of the
standard Genero web components provided in $FGLDIR/webcomponents, there is no
need to set the `WEB_COMPONENT_DIRECTORY` element in the .xcf
file.

## Related links

**Related concepts**  

[Introducing the GAS and JGAS](../13_programming-tools/2646-introducing-the-gas-and-jgas.md "The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.")
