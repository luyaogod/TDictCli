---
title: "mobile.newContact"
source: "fgl-topics/c_fgl_frontcall_mobile_newcontact.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.newContact"
type: "concept"
---

# mobile.newContact

> Opens contact input form to create a new entry in the contact database.

## Syntax

```
ui.Interface.frontCall("mobile", "newContact",
  [defaults],[vcard])
```

1. defaults - A vCard string with default values for the contact input form. A
   limited set of properties and types are supported.
2. vcard - Holds the vCard string of the new created contact.

## Usage

The "`newContact`" front call opens the contact input form on the mobile device,
with default values passed in the vCard structure of the first parameter. The user can then enter
contact information, and validate the input form to create the new entry in the mobile contact
database.

The vCard passed as first parameter defines default values for the contact input form. A limited
set of vCard properties and types are scanned to fill the contact input form:

- Properties: `"N"`, `"FN"`, `"TEL"`,
  `"EMAIL"`, `"ADR"`
- Types: `"CUSTOM"`, `"HOME"`, `"WORK"`,
  `"MOBILE"`, `"OTHER"`

If the contact creation is validated, the front call returns the completed vCard string. If
the contact import is canceled, the front-end returns `NULL`.

## Example

```
DEFINE defaults, vcard STRING
LET defaults="BEGIN:VCARD\n"
      ||"VERSION:3.0\n"
      ||"N:Willi;;;;\n"
      ||"TEL;type=MOBILE:03812225610\n"
      ||"END:VCARD\n"
CALL ui.interface.frontcall("mobile","newContact",[defaults],[vcard])
```

## Related links

**Related concepts**  

[mobile.chooseContact](3442-mobile-choosecontact.md "Lets the user choose a contact from the mobile device contact list and returns the vCard.")

[mobile.importContact](3453-mobile-importcontact.md "Creates a new contact, or merges to an existing entry, the contact details passed in a vCard string.")
