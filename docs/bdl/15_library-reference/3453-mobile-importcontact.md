---
title: "mobile.importContact"
source: "fgl-topics/c_fgl_frontcall_mobile_importcontact.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.importContact"
type: "concept"
---

# mobile.importContact

> Creates a new contact, or merges to an existing entry, the contact details passed in a vCard string.

## Syntax

```
ui.Interface.frontCall("mobile", "importContact",
   [vcard], [result] )
```

1. vcard - Holds a vCard string to be imported into the device's contacts database.
2. result - Holds the completed vCard string.

## Usage

The "`importContact`" front call sends the vCard definition passed as
parameter to the mobile device.

If the contact import is canceled, the front-end returns `NULL`. Otherwise,
it returns the vCard data.

On iOS devices, the user has the choice to create a new contact, or complete an existing
contact entry. When creating a new entry, the contact input form is opened on the mobile
device, to let the user complete the default values passed as parameter. When merging
contact information to an existing entry, the user selects an entry from the contact list.
If the contact import is validated, the front call returns the completed vCard string.

On Android™ devices, this front call creates a new
contact entry directly in the mobile contact list, depending on the VCard definition passed as
parameter, no intermediate input form is presented to the end user. If the contact import is
validated, the front call returns the original vCard string passed as parameter.

> **Important:**
>
> For GMA / Android,
> using the `importContact` front call needs the
> `android.permission.WRITE_EXTERNAL_STORAGE` Dangerous Permission to be specified when
> building the APK. See [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more
> details.

## Example

```
DEFINE vcard, result STRING
LET vcard="BEGIN:VCARD\n"
      ||"VERSION:3.0\n"
      ||"N:Willi;;;;\n"
      ||"TEL;type=HOME;type=VOICE;type=pref:03812225610\n"
      ||"END:VCARD\n"
CALL ui.interface.frontcall("mobile","importContact",[vcard],[result])
```

## Related links

**Related concepts**  

[mobile.chooseContact](3442-mobile-choosecontact.md "Lets the user choose a contact from the mobile device contact list and returns the vCard.")

[mobile.newContact](3456-mobile-newcontact.md "Opens contact input form to create a new entry in the contact database.")
