---
title: "DateEdit style attributes"
source: "fgl-topics/r_fgl_presentation_styles_dateedit_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > DateEdit style attributes"
type: "reference"
---

# DateEdit style attributes

> DateEdit presentation style attributes apply to DATEEDIT elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `buttonIcon`

Defines the icon name to use for the button that opens the date picker.

## `calendarType`

Controls the type of calendar pop-up window.

This style attribute is only available when using the built-in date
picker. With the mobile theme of GBC, the native browser date picker is used, and built-in date
picker rendering attributes cannot be applied.

Possible values are `"dropdown"` and `"modal"`.

Default is `"modal"`.

When set to `"modal"`, the calendar is intrusive to the `DATEEDIT`
field: Once the calendar opens, the user must validate the date selection by clicking the OK button,
or keep the current field value by clicking the Cancel button.

When set to `"dropdown"`, the calendar is less intrusive: The user can type
directly into the `DATEEDIT` field. A single click on a date validates and closes the
calendar. There is no way to cancel the selected date.

## `daysOff`

Defines the days of the week that are grayed out.

This style attribute is only available when using the built-in date
picker. With the mobile theme of GBC, the native browser date picker is used, and built-in date
picker rendering attributes cannot be applied.

Possible values are `"monday"`, `"tuesday"`,
`"wednesday"`,  `"thursday"`, `"friday"`,
`"saturday"`, `"sunday"`.

Default is `"saturday sunday"`. The days of week can be combined, as
shown.

## `firstDayOfWeek`

Defines the first day of the week to be displayed in the calendar.

This style attribute is only available when using the built-in date
picker. With the mobile theme of GBC, the native browser date picker is used, and built-in date
picker rendering attributes cannot be applied.

Possible values are `"monday"`, `"tuesday"`,
`"wednesday"`,  `"thursday"`, `"friday"`,
`"saturday"`, `"sunday"`.

Default depends on the front-ends platform language settings: For example, the default first day
of week will be Sunday for an English/US locale, Monday for a French or German locale.

## `showCurrentMonthOnly`

Defines if dates of the previous and next months are shown.

This style attribute is only available when using the built-in date
picker. With the mobile theme of GBC, the native browser date picker is used, and built-in date
picker rendering attributes cannot be applied.

Values can be `"yes"`, `"no"` (default).

## `showWeekNumber`

Defines if the week numbers are displayed.

This style attribute is only available when using the built-in date
picker. With the mobile theme of GBC, the native browser date picker is used, and built-in date
picker rendering attributes cannot be applied.

Values can be `"yes"`, `"no"` (default).

## Related links

**Related concepts**  

[DATEEDIT item type](1688-dateedit-item-type.md "Defines a line-edit with a calendar widget to pick a date.")
