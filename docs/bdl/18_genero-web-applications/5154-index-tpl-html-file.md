---
title: "index_tpl.html file"
source: "fgl-topics/c_fgl_gwa_index_tpl.html"
breadcrumb: "Genero Web applications > GWA Reference > index_tpl.html file"
type: "concept"
---

# index_tpl.html file

> Customizing the index_tpl.html file.

index\_tpl.html is located in the
gwa-install-dir/lib/gwa directory. The
`gwabuildtool` uses it to generate the index.html file as the
source for the GWA loading and ending pages.

Customizing this file is typically not necessary. You may only need to customize it if you need
to change the structure of the page—such as adding or removing elements, changing layout containers,
or adding a logo—to the existing `index_tpl.html`.
> **Warning:**
>
> Although this customization offers enhanced flexibility, it may be more vulnerable to disruptions
> from future updates. Therefore, it should be implemented with caution.

**Steps to customize index\_tpl.html**

1. Create a custom index\_tpl.html by copying the default template from the GWA
   installation directory
   (gwa-install-dir/lib/gwa/index\_tpl.html) to your
   programdir/gwa directory .
2. Update the `index_tpl.html` file with your changes.

   If you wish to add custom
   CSS, it's not mandatory to do so externally, as you can directly include the CSS within the file.
   However, you still have the option to use the `gwa_custom.css` file, referred to in
   [Customize the index page](5134-customize-the-index-page.md "The GWA allows customization of index.html using custom CSS and HTML files."), which will be injected using the `<!--
   GWA custom css injection -->` placeholder.

   > **Important:**
   >
   > When you create a custom HTML file at
   > programdir/gwa/index\_tpl.html, it will be used as the
   > template instead of the default one. Therefore, when modifying the
   > index\_tpl.html file, ensure that all templating comments remain intact.
   > Specifically, do not alter the injection points indicated by comments such as:
   >
   > - `<!-- GWA favicons injection -->`
   > - `<!-- GWA manifest injection -->`
   > - `<!-- GWA App Version injection -->`
   >
   > These comments mark the locations where the build tool adds essential content, including links to
   > icons, manifest files, and version information.

## Example index\_tpl.html file

> **Note:**
>
> The content shown is a complete sample; however, it is recommended to source the content directly
> from the gwa-install-dir/lib/gwa/index\_tpl.html, as both the
> content and format may change in future product releases.

```
<html>
<!--
 Property of Four Js*
 (c) Copyright Four Js 2024, 2025. All Rights Reserved.
 * Trademark of Four Js Development Tools Europe Ltd
   in the United States and elsewhere
 This file can be modified by licensees according to the
 product manual.
-->
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <!-- GWA favicons injection -->
  <!--
    the template comment below cause something like
    <link rel="manifest" href="gwa.651.123f64b89f8460ad24d05fbcade6f8e6.webmanifest">
  -->
  <!-- GWA manifest injection -->
  <!-- GWA App Version injection -->
  <style type="text/css">
  html {
    height:100%;
    font-family: Droid Sans, sans-serif;
  }
  body {
    height:100%;
    margin:0;
  }
  /* is added to the waiter frame after program did terminate */
  .hidden {
    display: none!important;
    visibility: hidden!important;
  }
  .gwa_waiter,.gwa_ending {
    height:100%;
    margin-left:1em;
    margin-right:1em;
    margin-top:0;
    margin-bottom:0;
    display:flex;
    align-items:center;
    justify-content:center;
    flex-direction:column
  }
  .gwa_ending {
    height:100%;
  }
  .gwa_title {
    width:100%;
    text-align:center;
    line-height:1em;
  }
  .gwa_progress {
    width:100%;
  }
  .gwa_status {
    width:100%;
    text-align:center;
    line-height:1em;
  }
  #endstatus {
     font-size:1.5em;
  }
  .gwa_button {
    font-size:1em;
    margin:1em;
    padding:0.3em;
  }
  </style>
  <!-- GWA custom css injection -->
  <!-- GWA serviceworker injection  -->
  <title><!-- GWA titleText injection --></title>
</head>
  <body>
  <!-- active in the beginning of the program -->
  <div class="gwa_waiter" id="waiter">
    <p class="gwa_title" id="title" >Please wait, <!-- GWA programOrTitleText injection --> is loading...</p>
    <progress class="gwa_progress" id="progress" value="0"></progress>
    <p class="gwa_status" id="empty"></p>
    <p class="gwa_status hidden" id="status"></p>
    <p class="gwa_status hidden" id="swstatus">SW Status</p>
  </div>
  <!-- is unhidden by removing the hidden class after the end of the Genero program -->
  <div class="gwa_ending hidden" id="ending">
    <p class="gwa_status" id="endstatus"><!-- GWA programOrTitleText injection --> ended.</p>
    <button class="gwa_button" id="reload">Reload Page</button>
    <p class="gwa_status hidden" id="debugEndedcloseBrowser">The debugger did disconnect. Please close this browser Tab and restart a new session on the command line for development/debugging.</p>
  </div>
  <!-- GWA WASM injection -->
  </body>
</html>
```

## Related links

**Related concepts**  

[Customize the index page](5134-customize-the-index-page.md "The GWA allows customization of index.html using custom CSS and HTML files.")

[gwa.webmanifest file](5153-gwa-webmanifest-file.md "Example of a customized gwa.webmanifest file, providing information about a Genero Web Application (GWA).")
