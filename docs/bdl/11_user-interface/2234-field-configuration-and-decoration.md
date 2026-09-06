---
title: "Field configuration and decoration"
source: "fgl-topics/c_fgl_prog_dialogs_field_configuration.html"
breadcrumb: "User interface > User interface programming > Input fields > Field configuration and decoration"
type: "concept"
---

# Field configuration and decoration

> Form fields can be customized with specific decoration and settings.

## Data formatting in fields

Data format for input and display of numeric (`DECIMAL`, `INTEGER`)
and `DATE` fields can be defined with the `FORMAT` attribute.

A default data format can be defined with environment variables ([DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values."), [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values."), etc)

## Forcing the input pattern

User input can be controlled with the [`PICTURE`](1807-picture-attribute.md "The PICTURE attribute specifies a character pattern for data entry in a text field, and prevents entry of values that conflict with the specified pattern.") attribute, to force alpha-numeric or numeric characters at a given
position. This is typically used for formatted data, such as credit card numbers or vehicle
identification numbers.

## Form field display size

If the default size of the form field widget is too large, provide a text example with the [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget.") attribute, to let the
front-end compute a specific size.

## Define field description

Form field description can be displayed to the user with the [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element.") attribute.

## Hide text for password input

In order to hide characters typed by the user, define the [`INVISIBLE`](1793-invisible-attribute.md "The INVISIBLE attribute prevents field data being readable on the screen.") attribute in the form
field.

## Force user to input value twice

When the [`VERIFY`](1845-verify-attribute.md "The VERIFY attribute requires users to enter data in the field twice to reduce the probability of erroneous data entry.") attribute is
define for a form field, the user must enter the value twice before leaving the field.

## Force uppercase or lowercase input

To force uppercase input, use the [`UPSHIFT`](1838-upshift-attribute.md "The UPSHIFT attribute forces character input to uppercase letters.") attribute, and to force lowercase use the [`DOWNSHIFT`](1776-downshift-attribute.md "The DOWNSHIFT attribute forces character input to lowercase letters.") attribute.

## Default form field values

When the input dialog is not using [`WITHOUT DEFAULTS`](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.") clause, the default value of a form field can be set with
the [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") attribute.

## Tabbing order of form fields

To bypass the default tabbing order, use the `TABINDEX` attribute, as described in
[Defining the tabbing order](2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.").

## Value input limits, boolean correspondence and steps

Depending on the form field item type, control the value limits with [`VALUEMIN`](1841-valuemin-attribute.md "The VALUEMIN attribute defines a lower limit of values displayed in widgets (such as progress bars).") / [`VALUEMAX`](1842-valuemax-attribute.md "The VALUEMAX attribute defines a upper limit of values displayed in widgets (such as progress bars).") attributes
(`SPINEDIT`, `SLIDER`), the incremental and decremental gap with [`STEP`](1820-step-attribute.md "The STEP attribute specifies how a value is increased or decreased in one step (by a mouse click or key up/down)."), and boolean correspondence with
[`VALUECHECKED`](1843-valuechecked-attribute.md "The VALUECHECKED attribute defines the value associated with a checkbox item when it is checked.") / [`VALUEUNCHECKED`](1844-valueunchecked-attribute.md "The VALUEUNCHECKED attribute defines the value associated with a checkbox item when it is not checked.")
(`CHECKBOX`).

## Form fields with button

To identify the action to be fired typically in a `BUTTONEDIT` field, define the
[`ACTION`](1758-action-attribute.md "The ACTION attribute defines the action associated with the form item.") attribute.

## Defining the TABLE/TREE column TITLE

With form fields used in a list container like `TABLE` / `TREE`, it
is possible to define the column title with the [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item.") attribute.

Consider using a `%"ident"` localized string.

## Control TABLE/TREE columns resize, visibility and sort selection

Form fields used in a `TABLE` / `TREE` list container are by
default resizable, hidable and can be used for sorting. This can be denied with
`UNSIZABLE`, `UNHIDABLE` and `UNSORTABLE`
attributes.

## CHECKBOX caption with TEXT

With form fields defined as `CHECKBOX` elements, labels can be defined by the
[`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") attribute.

Consider using a `%"ident"` localized string.

## The KEYBOARDHINT attribute

Especially for mobile devices, consider using the [`KEYBOARDHINT`](1798-keyboardhint-attribute.md "The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly.") attribute, to get
a specific keyboard for phone numbers, emails, URLs, etc.

## Using a presentation STYLE

Additional configuration and decoration can be define for a form field with the [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute, pointing to style
attributes defined in [.4st
files](1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements."). Form field text and background colors are a typical decoration that can be achieved
with styles.

## COMBOBOX and RADIOGROUP items

The selection items of a `COMBOBOX` or `RADIOGROUP` form field are
defined with the [`ITEMS`](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item.")
attribute.

Consider using `%"ident"` localized strings for item labels.

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")

[Field data type](2231-field-data-type.md "Depending on the type of dialog, the field data type is defined by program variables or form specification file.")
