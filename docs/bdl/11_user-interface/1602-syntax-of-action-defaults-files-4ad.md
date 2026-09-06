---
title: "Syntax of action defaults files (.4ad)"
source: "fgl-topics/c_fgl_action_defaults_files_syntax.html"
breadcrumb: "User interface > Form definitions > Action defaults files > Syntax of action defaults files (.4ad)"
type: "concept"
---

# Syntax of action defaults files (.4ad)

> A .4ad action default file is an XML file defining default attributes for actions.

Action defaults are defined in the .4ad file with this
syntax:

```
<ActionDefaultList>
  <ActionDefault name="action-name" [ attribute="value" [...] ] />
  [...]
</ActionDefaultList>
```

1. action-name identifies the action.
2. attribute is the name of an attribute.
3. value defines the value to be assigned to attribute.
