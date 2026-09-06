---
title: "Add a favicon"
source: "fgl-topics/t_fgl_gwa_change_favicon.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Customizing GWA apps > Add a favicon"
type: "task"
---

# Add a favicon

> This procedure shows you how to add a favicon image.

**About this task**

A favicon is an icon associated with a URL, displayed in various locations such as a browser's
address bar, tab, or alongside the site name in a bookmark list. Figure 1 shows a GWA app displayed using
the default favicon.

![Image shows default Genero icon](../_images/default-favicon.png)

*Default favicon*

The icon is referenced in the generated index.html file at the line
containing the `<link>` element, as shown in the following
example:

```
<link rel="icon"
href=logo512.png.63839.eb1d062b9dc0e90a6f05a237c8a2e3c8.png sizes="512x512"
type="image/png"/>
```

The favicon can be set in your
programdir/gwa/gwa.webmanifest file. For more details about
the manifest file, go to [gwa.webmanifest file](5153-gwa-webmanifest-file.md "Example of a customized gwa.webmanifest file, providing information about a Genero Web Application (GWA).").

**Steps to change the favicon**

1. Place your favicon image in your programdir/gwa
   directory.

   The image file for the favicon is typically in png format; however, other
   formats, such as svg, may also be used, depending on browser support. For
   example, Chrome and Edge support the svg format.
2. Open the programdir/gwa/gwa.webmanifest file in a text
   editor. 

   To set the favicon, you need to update or add the `icons` key in your
   manifest file. The following snippet illustrates how to define the `icons` property.
   For a complete sample, go to [gwa.webmanifest file](5153-gwa-webmanifest-file.md "Example of a customized gwa.webmanifest file, providing information about a Genero Web Application (GWA)."):

   ```
   {
     "_": "gwa.webmanifest file",
     "_hint": "To add comments, use underscore as the 1st letter of the option",
     "icons": [
       { "_hint":"next icon appears as link in the manifest"},
       { "src": "bug.png",
         "type": "image/png",
         "sizes": "690x690",
         "GWA:linkRel": "icon"
       }
     ]
   }
   ```

   Where:
   1. **`src` attribute**:

      - The `src` attribute specifies the path to the icon image file that will be used
        as the favicon. Both relative paths (to the programdir/gwa
        directory) and absolute paths are supported.
   2. **`type` and `sizes` attributes**:

      - The `type` attribute defines the MIME type of the icon image (for example,
        "image/png").
      - The `sizes` attribute specifies the actual pixel dimensions of your PNG file
        (width in pixels by height in pixels).
   3. **`GWA:linkRel` attribute**:

      - The `GWA:linkRel` attribute must be set to `icon` to identify the
        image used as the favicon in your application. In addition to "`icon`", other values
        such as "`apple-touch-icon`" can be used in the context of favicons.
      - This attribute is used internally by the `gwabuildtool` and does not appear in
        the generated web manifest file accessible to the browser. The value of the
        `GWA:linkRel` attribute corresponds to the `rel` attribute in the
        `<link>` element in the generated index.html.

**Build and test**

3. Follow the instructions to build and run the GWA in [Build and test the application](5138-build-and-test-the-application.md "Build a GWA application.")

   Check whether the favicon displays as expected.

   ![Images shows customized favicon](../_images/customized_favicon.png)

   *Customized favicon*

After completing the task above, you may want to consider that the icon can also be turned off by
using the `"GWA:removeFromManifest"` attribute. Setting this attribute to true will
omit the icon from the generated web manifest file, resulting in the creation of only a
`<link>` tag.

For example:

```
{ "_": "gwa.webmanifest file", 
  "icons": [
    { "_hint":"next icon appears only as a link and is removed from the resulting manifest"},
    { "src": "bug.png",
      "type": "image/png",
      "sizes": "690x690",
      "GWA:linkRel": "apple-touch-icon",
      "GWA:removeFromManifest":true
    }
  ]
}
```

## Related links

**Related concepts**  

[gwa.webmanifest file](5153-gwa-webmanifest-file.md "Example of a customized gwa.webmanifest file, providing information about a Genero Web Application (GWA).")

[Customize the index page](5134-customize-the-index-page.md "The GWA allows customization of index.html using custom CSS and HTML files.")

[Internationalize your app](5135-internationalize-your-app.md "GWA applications can be translated and modified for international markets.")

[Creating GWA apps with Genero](5126-creating-gwa-apps-with-genero.md "Prepare the environment to build GWA applications.")
