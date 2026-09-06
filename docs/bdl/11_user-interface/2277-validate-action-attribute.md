---
title: "VALIDATE action attribute"
source: "fgl-topics/c_fgl_action_attribute_VALIDATE.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > VALIDATE action attribute"
type: "concept"
---

# VALIDATE action attribute

> The VALIDATE action attribute defines the data validation level for a given action.

## Syntax 1: (Dialog action handlers and form action defaults)

```
VALIDATE = NO
```

## Syntax 2 (Global .4ad action defaults file)

```
validate = "no"
```

## Usage

When the `VALIDATE` action attribute is set to `NO`, it indicates
that no data validation must occur for this action. However, current input buffer contains the text
modified by the user before triggering the action.

This attribute applies to the actions defined by the current dialog
in the current window. It can be specified as action default attribute in a global
.4ad file, in the `ACTION DEFAULTS` section of form files, or as
dialog action attribute (using `ON ACTION name
ATTRIBUTES(...)`).

## Example

```
-- As action handler attribute
ON ACTION undo ATTRIBUTES(VALIDATE=NO)

-- As action default
ACTION DEFAULTS
  ACTION undo (VALIDATE=NO)
END

-- In a global action defaults file
<ActionDefault name="undo" validate="no" ... />
```

## Related links

**Related concepts**  

[VALIDATE attribute](1839-validate-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.")

[Data validation at action invocation](2281-data-validation-at-action-invocation.md "The validate action attribute controls field validation when an action is fired.")
