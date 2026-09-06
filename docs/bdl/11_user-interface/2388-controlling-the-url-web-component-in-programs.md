---
title: "Controlling the URL web component in programs"
source: "fgl-topics/c_fgl_webcomponent_url_manip.html"
breadcrumb: "User interface > User interface programming > Web components > Using a URL-based web component > Controlling the URL web component in programs"
type: "concept"
---

# Controlling the URL web component in programs

> URL-based web components can be controlled with the field value and with front calls

## Detecting user interaction in a web component with ON CHANGE

The content of a URL-based web component is defined by the field value.

When the end user interacts with the content, and if the remote service points to a different
URL, the field value changes.

The URL change can be detected by implementing an [`ON CHANGE`](1943-on-change-block.md) control block for the web
component field.

> **Important:**
>
> Due to security rules in modern web browsers, URL-based web components may
> not trigger `ON CHANGE` block, if the web component URL contains a domain name
> different from the Genero application. Thus, any changes/navigation in cross domain location cannot
> be detected in the program. Only HTML pages in the same domain will trigger `ON
> CHANGE` blocks.

The `ON CHANGE` trigger will be fired immediately when the URL
changes:

```
DEFINE rec RECORD
           num INTEGER,
           name STRING,
           map STRING
       END RECORD
...
INPUT BY NAME rec.* WITHOUT DEFAULTS
    ATTRIBUTES(UNBUFFERED)
    ...
    ON CHANGE map
       CALL map_changed(rec.map)
    ...
```

## Controlling URL-based web components with front calls

The web component can be manipulated with specific front calls. The web component-specific front
calls are provided in the "`webcomponent`" front call module.

The [`webcomponent.call`](../15_library-reference/3423-webcomponent-call.md "Calls a JavaScript function through the web component.")
is a front call that can be used for general purposes. It takes as parameters the name of the form
field, a JavaScript function to call, and optional parameters as required. The JavaScript function
must be implemented in the HTML content pointed to by the URL of the web component field. The front
call returns the result of the JavaScript function.

```
DEFINE title STRING
CALL ui.Interface.frontCall("webcomponent", "call",
     ["formonly.url_field", "eval", "document.title"],
     [title] )
```

When passing a `RECORD` or `DYNAMIC ARRAY` as front call parameter,
it will be converted to a JSON string for the JavaScript function.

Some web component providers return key information in the title of the HTML document. The [`webcomponent.getTitle`](../15_library-reference/3425-webcomponent-gettitle.md "Returns the title of the HTML doc rendered by a web component.")
function is another useful `webcomponent` front call that can get the title of the
HTML document of the web component:

```
DEFINE info STRING
CALL ui.Interface.frontCall("webcomponent", "getTitle",
     ["formonly.url_field"], [info] )
```

## Related links

**Related concepts**  

[Web component front calls](../15_library-reference/3422-web-component-front-calls.md "This section describes web component specific front calls.")
