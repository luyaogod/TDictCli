---
title: "Screen readers"
source: "fgl-topics/c_fgl_Accessibility_005.html"
breadcrumb: "User interface > Form definitions > Accessibility guidelines > Screen readers"
type: "concept"
---

# Screen readers

> How to integrate with platform screen readers?

## Understanding screen readers

Screen readers are special system applications that transform the application's graphical
user interface into speech. The behavior may change between screen reader implementations, but,
basically, each widget is named and described by speech. On some workstation operating systems,
special keyboard shortcuts are available to trigger the complete enumeration of all the components
of a window, or to describe only the component having the current focus.

## Providing form item descriptions to screen readers

Screen readers use special bindings to get the information that they need (name, full
description, hierarchy, triggered actions, and so on) about each graphical component of the entire
graphical user interface. It is up to the programmer to provide these bindings to the screen reader,
but most of the work is already done by the front-end.

To allow assistive technologies such as screen readers, The GBC front end implements "Accessible
Rich Internet Applications" attributes and roles. For more details about this HTML standard feature,
see [WAI-ARIA](https://www.w3.org/TR/wai-aria-1.1/#intro_ria_accessibility).

Form elements can be defined with [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") and [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item.") attributes, to provide speech information to screen readers, keeping
in mind that the text will be spoken.

Programmers can provide description for each form element, to provide speech information to
screen readers with the [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") and [`TITLE`](1829-title-attribute.md "The TITLE attribute defines the title of a form item.") form attributes, keeping in mind that the text will be spoken.

Most of the form items are supported: All kind of form field, static labels, static images,
and action-based items (such as buttons); some containers (`GROUP` and
`FOLDER`) work out of the box as soon as their `TEXT`
attributes are set.

For [accessibility
compliance](1589-accessibility-guidelines.md "This section describes the best practices to make your application accessible to disabled people."), when a form field is defined with the [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") attribute, the GBC
front-end automatically sets the HTML attribute [WAI-ARIA `aria-required`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-required) to
`true`. Conversely, the `aria-required` attribute is not set when
using the [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") form field
attribute, as its functional purpose differs from standard ARIA requirement states.

## Examples

In an action defaults file (mydefaults.4ad)

```
<ActionDefaultList>
  <ActionDefault name="new" text="New" image="new.svg"
     comment="Create a new database"
   acceleratorName="Control-N" />
  <ActionDefault name="open" text="Open" image="open.svg"
     comment="Open an existing database"
   acceleratorName="Control-O" />
  <ActionDefault name="save" text="Save" image="save.svg"
     comment="Save the current database"
   acceleratorName="Control-S" />
  ...
```

In field definitions on a form specification file (myform.per)

```
ATTRIBUTES
  EDIT login_name = formonly.login_name, NOT NULL,
    COMMENT="Login name of the current user";
  EDIT password = formonly.password, NOT NULL, INVISIBLE, VERIFY,
    COMMENT="Password of the current user";
  EDIT first_name = formonly.first_name, NOT NULL,
    COMMENT="First name of the current user";
  EDIT last_name = formonly.last_name, NOT NULL,
    COMMENT="Last name of the current user";
  DATEEDIT birthdate = formonly.birthdate, FORMAT="mm/dd/yyyy",
    COMMENT="Date of birth of the current user";
  EDIT email = formonly.email,
    COMMENT="E-mail of the current user";
END -- ATTRIBUTES
```

In this form specification file, the `COMMENT` attribute is used for both the accessible
name and the accessible description.
