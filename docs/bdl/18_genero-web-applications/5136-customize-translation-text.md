---
title: "Customize translation text"
source: "fgl-topics/t_fgl_gwa_customize_translation_text.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Customizing GWA apps > Internationalize your app > Customize translation text"
type: "task"
---

# Customize translation text

> This procedure shows you how custom translation text is provided in a locale file and referenced in the HTML code in the template file.

**About this task**

You can use translation texts to provide localization in the following GWA components:

- Welcome Page
- Ending Page

If you want to personalize the translatable text for your GWA app, start by creating your custom
locale file. Then, place it in the programdir/gwa/locales
directory. Ensure your JSON file is structured correctly, using straight quoted key-value pairs in a
comma-separated list. You can use any of the formats shown for the key.

```
// en.json
{
    "key1": "value1",
    "key with spaces": "value2",
    "if.someone.likes.dots.as.key": "value3"
}
```

**Steps**

1. Add a new locale file.

   For example, you need an Italian locale file for your translation text.

   1. Create a file named it.json and save it to your
      programdir/gwa/locales directory.

      Note you can use the same names as locale files in
      gwa-install-dir/lib/gwa/locales, because if a key exists in
      an installation locale file, the value from your custom locale file will overwrite it. For more
      information on locale filenames, go to [Loading localized strings at runtime](../09_advanced-features/0909-loading-localized-strings-at-runtime.md "Understand the rules for using localized strings at runtime.").

      On the other hand, if a locale is not provided by Four Js in an installation locale file, create
      a new locale file to add it in your
      programdir/gwa/locales.
   2. Add an Italian entry for this string, with the following content: 

      ```
      // it.json 
      { 
      "Please wait": "Attendere prego" 
      }
      ```
2. Add translatable keys in the HTML template file, `index_tpl.html`, for your
   translation text.

   Translatable keys use the pattern: `@{{myTanslateKey}}`,
   where myTanslateKey is the name of a key in the JSON file.

   ```
   <!-- index_tpl.html -->
   <!-- ... -->

   <p class="gwa_title" id="title" >@{{Please wait}}, 
             <!-- GWA programOrTitleText injection --> @{{is loading}}...</p>

   <!-- ... -->
   ```

   The template string, `@{{Please wait}}` will be
   translated utilizing the locale lookup mechanism of GWA when the application runs in the
   browser.
3. Test your customization.
   1. Set your browser to Italian.
   2. Start the application.

      To run your application, you will need to compile the sources and build it. If you have
      make installed, run `make gwa.run`; otherwise, follow the
      instructions to build and run the GWA in [Build and test the application](5138-build-and-test-the-application.md "Build a GWA application.").

   You should see "Attendere prego" appear in the loading message.
