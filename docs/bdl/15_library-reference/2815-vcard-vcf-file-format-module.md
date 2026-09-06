---
title: "VCard: VCF file format module"
source: "fgl-topics/r_fgl_utility_functions_vcard.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module"
type: "reference"
description: "Table 1. vCard types (VCard.4gl) Type Description TYPE VCAddress RECORD PostOfficeBox, ExtendedAddress, -- apartment or suite number Street, City, State, ZIP, Country STRING -- , CountryCode STRING -- ..."
---

# VCard: VCF file format module

| Type | Description |
| --- | --- |
| TYPE VCAddress RECORD PostOfficeBox, ExtendedAddress, -- apartment or suite number Street, City, State, ZIP, Country STRING -- , CountryCode STRING -- X-ABADR:de END RECORD | The VCAddress structured type holds vCard address data. |
| TYPE VCName RECORD FirstName, LastName, MiddleName, Prefix, Suffix STRING --, FormattedName STRING END RECORD | The VCName structured type holds vCard data related to the person's name. |
| TYPE VCPerson RECORD FirstName STRING, -- N[1] LastName STRING, -- N[2] MiddleName STRING, -- N[3] Prefix STRING, -- N[4] Suffix STRING, -- N[5] formattedName STRING, -- FN nickname STRING, -- NICKNAME jobTitle STRING, -- TITLE organization STRING, -- ORG.value[1] department STRING, -- ORG.value[2] birthday STRING, -- BDAY note STRING, -- NOTE address DYNAMIC ARRAY OF RECORD type STRING, PostOfficeBox, -- ADR[1] ExtendedAddress, -- ADR[2] Street, -- ADR[3] City, -- ADR[4] State, -- ADR[5] ZIP, -- ADR[6] Country STRING -- ADR[7] END RECORD, phone DYNAMIC ARRAY OF RECORD type STRING, number STRING -- TEL END RECORD, email DYNAMIC ARRAY OF RECORD type STRING, value STRING -- EMAIL END RECORD | The VCPerson structured type holds vCard data. |

| Function | Description |
| --- | --- |
| FUNCTION format_person( person VCPerson ) RETURNS STRING | Converts a VCPerson record to a vCard string representation vCard. |
| FUNCTION scan_address( source STRING, type STRING ) RETURNS VCAddress | Extracts an address from a string representing a vCard. |
| FUNCTION scan_email( source STRING, type STRING ) RETURNS STRING | Extracts an email from a string representing a vCard. |
| FUNCTION scan_person( source STRING ) RETURNS VCPerson | Extracts person's data from a string representing a vCard. |
| FUNCTION scan_phone( source STRING, type STRING ) RETURNS STRING | Extracts a phone number from a string representing a vCard. |
