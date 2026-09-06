---
title: "gwa.webmanifest file"
source: "fgl-topics/c_fgl_gwa_web_manifest.html"
breadcrumb: "Genero Web applications > GWA Reference > gwa.webmanifest file"
type: "concept"
---

# gwa.webmanifest file

> Example of a customized gwa.webmanifest file, providing information about a Genero Web Application (GWA).

A web application manifest is a JSON file that follows the web application manifest standard
specified by [W3C
manifest](https://w3c.github.io/manifest/) (external link). The following table outlines the key manifest members, their
descriptions, and important notes regarding their usage in the context of GWA app development.

| Manifest member | Description |
| --- | --- |
| `background_color` | This property can be set, but it does not have any visual effect on your GWA app. |
| `description` | This property is not set by `gwabuildtool`, but it could be used to provide a short description of your app in your custom web manifest. |
| `dir` | This property is currently not possible to set. |
| `display` | This property is not processed by `gwabuildtool`, but it can be set in your custom manifest according to the W3C specifications |
| `display_override` | This property is also not processed by `gwabuildtool`, but it can be set in your custom manifest according to the W3C specifications. |
| `icons` | This property specifies the icons for your application and can be set in your custom manifest. See customized gwa.webmanifest. |
| `id` | This property is not set by `gwabuildtool`. It should only be set if you are a web master expert, and it needs to contain the deployment target web server URL if set. |
| `name` | This property is automatically filled in by the `gwabuildtools --title` option, but it can be overridden in your custom manifest. |
| `name_localized` | This property is currently not possible to set. |
| `orientation` | This property is not processed by `gwabuildtool`, but it can be set in your custom manifest according to the W3C specifications. |
| `scope` | This property is automatically filled in by `gwabuildtool`, and it should only be changed if you are a web master expert. |
| `screenshots` | This property provides screenshots for your application and can be set in your custom manifest. See customized gwa.webmanifest. |
| `short_name` | This property is automatically filled in by the `gwabuildtools --title` option, but it can be overridden in your custom manifest. |
| `start_url` | This property is automatically set by `gwabuildtool`, and it should not be overridden in your custom manifest. |
| `theme_color` | This property sets the browser's decoration color when your app is installed as a desktop app (Chrome/Edge). |

Every GWA application must include a web application manifest file when deployed. A custom
manifest file, gwa.webmanifest, should be located in the
programdir/gwa directory of your project. The [gwabuildtool](5150-gwabuildtool.md "The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser.") reads the gwa.webmanifest file and
produces a
distribution\_dir/gwa.webmanifest.size.hash.webmanifest
file from it, enabling web servers to deliver the right image content.

If you don't provide a custom `gwa.webmanifest` file in
programdir/gwa, `gwabuildtool` creates a
functional default
gwa.webmanifest.size.hash.webmanifest
file in the distribution directory complete with Genero icons.

## customized gwa.webmanifest

This sample includes properties that define the visual elements of the application:

- `icons`: This property defines an array of icon objects, such as the
  specification for the favicon used in the address bar and a logo that serves as the application icon
  in the Finder or Explorer when the app is installed, each
  with properties for source, size, and type. For more information about setting icons, go to [Add a favicon](5133-add-a-favicon.md "This procedure shows you how to add a favicon image.").
- `screenshots`: This property defines an array of screenshot objects, each
  containing properties for the source, size, and type. These screenshots may be displayed when the
  installation button is clicked in the address bar (in Chrome/Edge) or when the browser prompts to
  install the app on the home screen (in Chrome on Android™ devices).

> **Note:**
>
> **Image paths**
>
> The `"src"` property specifies image paths that must
> be defined as either absolute or relative to the `gwa.webmanifest` file itself; this
> property does not refer to a web server path, as the build tool searches for these images based on
> the specified paths.

> **Note:**
>
> The comments in your custom manifest will not appear in the gwa\_dist file;
> they are removed during the build process.

```
{
    "_hint": "gwabuildtool places them in the root dir of the app with a unique MD5 hash",
    "_id": "/full/path/to/index.html?some_query=1",
    "start_url": "./index.html?some_query=1",
    "_name": "My GWA app",
    "_short_name": "MyApp",
    "display": "standalone",
    "orientation": "any",
    "background_color": "green",
    "theme_color": "yellow",
    "icons": [
        {
            "src": "icon1.png",
            "type": "image/png",
            "sizes": "192x192"
        },
        {
            "_hint": "the following image will be found"
        },
        {
            "src": "../gwa/icon2.png",
            "type": "image/png",
            "sizes": "512x512",
            "purpose": "any"
        },
        {
            "_hint": "the following image is a favicon, visible in the address bar"
        },
        {
            "GWA:linkRel": "icon",
            "src": "bug.png",
            "type": "image/png",
            "sizes": "690x690"
        }
    ],
    "description": "My app manifest.",
    "lang": "en",
    "categories": [
        "misc",
        "productivity",
        "utilities"
    ],
    "screenshots": [
        {
            "src": "Screenshot540x720.png",
            "type": "image/png",
            "sizes": "540x720",
            "form_factor": "narrow"
        },
        {
            "src": "Screenshot1024x576.png",
            "type": "image/png",
            "sizes": "1024x576",
            "form_factor": "wide"
        }
    ]
}
```

## Related links

**Related concepts**  

[index\_tpl.html file](5154-index-tpl-html-file.md "Customizing the index_tpl.html file.")

**Related tasks**  

[Add a favicon](5133-add-a-favicon.md "This procedure shows you how to add a favicon image.")
