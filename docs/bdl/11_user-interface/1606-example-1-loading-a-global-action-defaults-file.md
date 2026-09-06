---
title: "Example 1: Loading a global action defaults file"
source: "fgl-topics/c_fgl_action_defaults_files_example_1.html"
breadcrumb: "User interface > Form definitions > Action defaults files > Examples > Example 1: Loading a global action defaults file"
type: "concept"
description: "Some action defaults in XML format (exit action has Localized Strings ): <ActionDefaultList> <ActionDefault name=\"print\" text=\"Print\" image=\"printer\" comment=\"Print report\" /> <ActionDefault ..."
---

# Example 1: Loading a global action defaults file

Some action defaults in XML format (exit action has [Localized
Strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")):

```
<ActionDefaultList>
   <ActionDefault name="print" text="Print" image="printer"
       comment="Print report" />
   <ActionDefault name="modify" text="Update"
       comment="Update the record" />
   <ActionDefault name="exit" text="Quit" image="byebye" 
       comment="Exit the program" validate="no" >
       <LStr text="common.exit.text" />
   </ActionDefault>
</ActionDefaultList>
```

The program loading the action defaults file:

```
MAIN
  CALL ui.Interface.loadActionDefaults("mydefaults")
  OPEN FORM f FROM "myform"
  DISPLAY FORM f
  ...
END MAIN
```
