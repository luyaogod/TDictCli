---
title: "Internationalize your app"
source: "fgl-topics/c_fgl_gwa_internationalization.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Customizing GWA apps > Internationalize your app"
type: "concept"
---

# Internationalize your app

> GWA applications can be translated and modified for international markets.

The GWA provides support for language selection based on locale. For more details on
localization, go to [Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.").

You can use translation texts to provide localization in the following GWA components:

- Welcome Page
- Ending Page

## Locale files

You will find the language locale files in
gwa-install-dir/lib/gwa/locales, where the filenames can
follow either of the standard codes used for languages:

- xx.json: two character locale code used for languages
  like en.json, de.json
- xx-YY.json: the full locale code like
  en-US.json, de-DE.json,
  fr-FR.json.

Locale files are JSON files which have a typical JSON structure of keys and values in double
quotes separated by colons, ":". Translation key-value pairs are separated by commas. A sample from
the French locale file is shown. (Line breaks have
been added to improve readability.)

```
 // in fr.json
{
  "Reload Page": "Redémarrage de l'application",
  "Please wait": "Merci d'attendre",
  "is loading": "est entrain d'être chargée",
  "ended": "vient de se terminer",
  "SW Status": "Statut du Service Worker",
  "debugEndedCloseBrowser": "Le débogueur a été arreté. Si vous voulez redémarrer une nouvelle session, 
    merci de fermer cet onglet et de relancer le débogueur sur la ligne de commande."
}
```

The translation lookup first searches the full language maps, and if a key is not found
there, it then checks the 2-character language maps. You can also provide your custom translation in
a locale file in your application. For details, go to [Customize translation text](5136-customize-translation-text.md "This procedure shows you how custom translation text is provided in a locale file and referenced in the HTML code in the template file."). If a key exists both in an installation locale file
and in your custom locale file, the value from your custom locale file will overwrite it.

## i18n localized strings

To localize your application, you need compiled localized string files
(.42s) containing your translation text placed in the appropriate language
subdirectory of your GWA application. For example, for an English-US locale, this would be in
programdir/en\_US/mystrings.42s. For details on how to organize localized
strings, go to [Localized string files on mobile devices](../09_advanced-features/0909-loading-localized-strings-at-runtime.md).

In the demo/i18n, we have created language subdirectories with sample
translations for the string "TESTLABEL." The displayed value in the application form changes
according to the browser's language setting, such as "Hello en\_US" for English (US). GWA configures
the FGLRESOURCEPATH for the localized string files (.42s) based on the LANG
variable, as displayed in the demo program's initial message showing the values of LANG and
FGLRESOURCEPATH.

If you want to set the FGLRESOURCEPATH in your application while using localized strings, ensure
that the GWA-calculated value of FGLRESOURCEPATH is included in your fglprofile. For example:

```
mobile.environment.FGLRESOURCEPATH="$FGLRESOURCEPATH:mydir"
```

This will ensure that you get the correct lookup for the localized string files
(.42s), even when `mydir` is added to FGLRESOURCEPATH.

## Related links

**Related concepts**  

[Loading localized strings at runtime](../09_advanced-features/0909-loading-localized-strings-at-runtime.md "Understand the rules for using localized strings at runtime.")

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")

## Child topics

- [Customize translation text](5136-customize-translation-text.md): This procedure shows you how custom translation text is provided in a locale file and referenced in the HTML code in the template file.
