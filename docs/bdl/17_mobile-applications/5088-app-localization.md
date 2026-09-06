---
title: "App localization"
source: "fgl-topics/c_fgl_mobile_localization.html"
breadcrumb: "Mobile applications > App localization"
type: "concept"
---

# App localization

> Mobile apps can be designed to display localized texts based on the language selected on the device.

## Application locale definition

Make sure that your application locale is defined properly for development and deployment
environments.

With Genero mobile applications, the application locale must be UTF-8.

For more details, see [Quickstart guide for locale settings](../09_advanced-features/0866-quickstart-guide-for-locale-settings.md "This is a quick step-by-step guide to properly configure locale settings for your Genero application.").

## Using localized strings

Localized string files (.42s) must be deployed in directories matching the
language identifiers (`en` for English, `zh_TW` for simplified
Chinese, etc), beside the program module.

Default string files can be provided in APPDIR/defaults as fallback
directory, if the current mobile language does not match one of the language-specific directories
provided by the application.

The list of `.42s` files required by the application must be defined in the unique
fglprofile configuration file located beside the program module of your
application.

For more details, see [Localized string files on mobile devices](../09_advanced-features/0909-loading-localized-strings-at-runtime.md) and [Deploying mobile apps](5102-deploying-mobile-apps.md "This section describes how to build and deploy mobile apps with Genero.").

## Querying user's preferred language

The application can get the user's preferred language and regional settings as configured on the
mobile device, by using the `standard.feInfo` front call with the
`userPreferredLang` argument.

For more details, see [User's preferred language](../09_advanced-features/0892-user-s-preferred-language.md "An application can get the user's preferred language and territory as configured on the front-end platform.") and [standard.feInfo](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.").
