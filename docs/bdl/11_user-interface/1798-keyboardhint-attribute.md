---
title: "KEYBOARDHINT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_KEYBOARDHINT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > KEYBOARDHINT attribute"
type: "concept"
---

# KEYBOARDHINT attribute

> The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly.

## Syntax

```
KEYBOARDHINT =
 { DEFAULT
 | NUMBER | NUMERIC | DECIMAL
 | PHONE | URL | EMAIL
 | SEARCH | TEXT
 | NONE }
```

## Usage

The `KEYBOARDHINT` attribute can be used to give a hint to the front-end,
regarding the kind of data the form field will contain. Based on this hint, the front-end will open
the virtual keyboard adapted to the data type; especially useful when designing application forms
for mobile platforms.

Valid values for `KEYBOARDHINT` are:

- `DEFAULT`: To enter values based on the type of the field variable. This
  is the default. See below for more details.
- `EMAIL`: To enter an email address.
- `NUMBER`: To enter a numeric value. Note that devices may or may not show a minus
  key. This is a synonym for DECIMAL.
- `PHONE`: To enter a phone number.
- `URL`: To enter a URL.
- `TEXT`: To enter any kind of text.
- `SEARCH`: To enter any text, with search option/button.
- `DECIMAL`: To enter real numbers with a decimal part. Note that devices may or
  may not show a minus key.
- `NUMERIC`: To enter integer numbers. Note that devices may or may not show a
  minus key.
- `NONE`: No virtual keybord is displayed. To be used when the form/app implements
  its own input control.

By default, or when `KEYBOARDHINT=DEFAULT` is used, the virtual
keyboard will depend on the data type of the program variable bound to the form
field:

- A numeric keyboard for integer number input is displayed when the type is
  `BOOLEAN`, `TINYINT`, `SMALLINT`,
  `INTEGER` or `BIGINT`;
- A numeric keyboard for real number input is displayed when the type is
  `SMALLFLOAT`, `FLOAT`, `DECIMAL` or
  `MONEY`.
- Any other type will result in a keyboard for regular text input.

## Example

```
EDIT f23 = customer.cust_phone, KEYBOARDHINT=PHONE;
```
