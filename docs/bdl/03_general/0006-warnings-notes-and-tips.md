---
title: "Warnings, notes and tips"
source: "fgl-topics/c_fgl_DocConv_003.html"
breadcrumb: "General > Documentation conventions > Warnings, notes and tips"
type: "concept"
---

# Warnings, notes and tips

> Documentation notes focus on a technical fact you must be aware of.

## Important

An important note describes a fact that needs to be considered, before using the related
feature.

> **Important:**
>
> When a DATE, DATETIME or INTERVAL constant cannot be initialized correctly,
> it is set to NULL.

Some Genero features are not supported on all back-end or front-end platforms. The following
warnings about the limitation:

> **Important:**
>
> This feature is not supported on mobile platforms.

> **Important:**
>
> This feature is only for mobile platforms.

> **Important:**
>
> This feature is only for the GMA/Android™ platform.

> **Important:**
>
> This feature is only for the GMI/iOS platform.

Some Genero features are only to be used in a development context:

> **Important:**
>
> This feature is provided for development and
> debug purpose, and must not be used in a production environment.

Some Genero features are under construction:

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

Some Genero features are deprecated:

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

Some Genero features are desupported:

> **Important:**
>
> This feature is desupported. It may still be
> available, but can be removed in a next version.

Some Genero features are under construction:

> **Warning:**
>
> Genero vNext software and documentation are still
> under construction and subject of changes until GA. This documentation is provided for early access
> users, to better understand new features. Any suggestion regarding new features is welcome.

Some Genero features are still part of the syntax or library, but do nothing:

> **Important:**
>
> This feature is an empty shell. The syntax or
> API call is still available, but it does nothing.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This version of Genero Mobile (GMA/GMI) is desupported, use latest version of the product.

## Simple notes

Simple notes describe a fact you must be aware of, but does not represent any risk.

> **Note:**
>
> This feature has been introduced in version 2.50.

## Tip notes

Tip notes suggest an option or programming pattern to be used.

> **Tip:**
>
> Use dynamic arrays instead of static arrays to save memory usage.

## Important notes

Important notes bring your attention to facts that need to be considered with care.

> **Important:**
>
> A division by zero will result in a runtime error.

## Warning notes

Warning notes highlight an action that might cause something to break or lead to unexpected
program behavior.

> **Warning:**
>
> Using a high cost value for the salt is very CPU consuming, and can really
> slow down the application depending on the system it is running.
