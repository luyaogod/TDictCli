---
title: "mobile.chooseContact"
source: "fgl-topics/c_fgl_frontcall_mobile_choosecontact.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.chooseContact"
type: "concept"
---

# mobile.chooseContact

> Lets the user choose a contact from the mobile device contact list and returns the vCard.

## Syntax

```
ui.Interface.frontCall("mobile", "chooseContact",
   [], [result])
```

1. result - The vCard string from the device's contacts database.

## Usage

The "`chooseContact`" front call opens the mobile device contact
chooser, allows the user to select a contact and returns the contact as a vCard string.

> **Important:**
>
> To be usable with GBC, this front call needs a secure context / front-end
> connection, in addition to mobile device permissions, when required. A secure context
> is achieved by using the GAS via HTTPS, via localhost on the same machine, or running direct via
> GDC-UR. When using GDC-UR, additional security settings need to be enabled in the GDC configuration
> panel. See GDC documentation for more details.

> **Important:**
>
> For GMA / Android™,
> using the `chooseContact` front call needs the `android.permission.READ_CONTACTS`
> Dangerous Permission to be specified when building the APK. See
> [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more
> details. On Android 5.1 and lower
> (< API 23), use the `android.permission.GET_ACCOUNTS` permission.

If the user cancels the contact chooser, `NULL` is returned.

On Safari iOS, the `chooseContact` front call needs to be explicitly enabled in
`Safari Settings → Advanced → Experimental Features → Toggle Contact Picker
API`

The `chooseContact` front call returns a VCARD version 4, were the first name and
last name fields are optional.

## Example

```
DEFINE vcard STRING
CALL ui.Interface.frontCall("mobile", "chooseContact", [], [vcard] )
```

## Related links

**Related concepts**  

[mobile.importContact](3453-mobile-importcontact.md "Creates a new contact, or merges to an existing entry, the contact details passed in a vCard string.")

[mobile.newContact](3456-mobile-newcontact.md "Opens contact input form to create a new entry in the contact database.")
