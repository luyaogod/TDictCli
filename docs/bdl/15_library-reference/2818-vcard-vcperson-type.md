---
title: "VCard.VCPerson type"
source: "fgl-topics/c_fgl_utility_functions_vcard_vcperson.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.VCPerson type"
type: "concept"
---

# VCard.VCPerson type

> The VCPerson structured type holds vCard data.

## Syntax

```
TYPE VCPerson RECORD
    FirstName STRING, -- N[1]
    LastName STRING, -- N[2]
    MiddleName STRING, -- N[3]
    Prefix STRING, -- N[4]
    Suffix STRING, -- N[5]
    formattedName STRING, -- FN
    nickname STRING, -- NICKNAME
    jobTitle STRING, -- TITLE
    organization STRING, -- ORG.value[1]
    department STRING, -- ORG.value[2]
    birthday STRING, -- BDAY
    note STRING, -- NOTE
    address DYNAMIC ARRAY OF RECORD
        type STRING,
        PostOfficeBox, -- ADR[1]
        ExtendedAddress, -- ADR[2]
        Street, -- ADR[3]
        City, -- ADR[4]
        State, -- ADR[5]
        ZIP, -- ADR[6]
        Country STRING   -- ADR[7]
    END RECORD,
    phone DYNAMIC ARRAY OF RECORD
        type STRING,
        number STRING -- TEL
    END RECORD,
    email DYNAMIC ARRAY OF RECORD
        type STRING,
        value STRING -- EMAIL
  END RECORD
```

## Usage

This type defines a record structure to hold vCard information. It is used by VCard
functions such as [format\_person()](2819-vcard-format-person.md "Converts a VCPerson record to a vCard string representation vCard.").

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE p VCard.VCPerson
  LET p.FirstName = "Hans"
  LET p.LastName = "Mustermann"
END MAIN
```
