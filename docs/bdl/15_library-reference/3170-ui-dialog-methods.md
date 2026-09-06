---
title: "ui.Dialog methods"
source: "fgl-topics/c_fgl_ClassDialog_methods.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods"
type: "concept"
---

# ui.Dialog methods

> Methods of the ui.Dialog class.

| Name | Description |
| --- | --- |
| ui.Dialog.createConstructByName( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD ) RETURNS ui.Dialog | Creates an `ui.Dialog` object to implement a dynamic `CONSTRUCT BY NAME`. |
| ui.Dialog.createMultipleDialog( ) RETURNS ui.Dialog | Creates an `ui.Dialog` object to implement a dynamic `DIALOG` multiple-dialog. |
| ui.Dialog.createDisplayArrayTo( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD, screenRecord STRING ) RETURNS ui.Dialog | Creates an `ui.Dialog` object to implement a dynamic `DISPLAY ARRAY TO`. |
| ui.Dialog.createInputArrayFrom( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD, screenRecord STRING ) RETURNS ui.Dialog | Creates an `ui.Dialog` object to implement a dynamic `INPUT ARRAY FROM`. |
| ui.Dialog.createInputByName( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD ) RETURNS ui.Dialog | Creates an `ui.Dialog` object to implement a dynamic `INPUT BY NAME`. |
| ui.Dialog.getCurrent() RETURNS ui.Dialog | Returns the current dialog object. |

| Name | Description |
| --- | --- |
| accept() | Validates and terminates the dialog. |
| ui.Dialog.addConstructByName( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD, name STRING ) | Adds a sub-dialog of type `CONSTRUCT BY NAME` to an existing `ui.Dialog` dynamic dialog. |
| ui.Dialog.addDisplayArrayTo( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD, screenRecord STRING ) | Adds a sub-dialog of type `DISPLAY ARRAY TO` to an existing `ui.Dialog` dynamic dialog. |
| addGroupBy( name STRING ) | Appends a grouping column of a list dialog. |
| ui.Dialog.addInputArrayFrom( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD, screenRecord STRING ) | Adds a sub-dialog of type `INPUT ARRAY FROM` to an existing `ui.Dialog` dynamic dialog. |
| ui.Dialog.addInputByName( fields DYNAMIC ARRAY OF RECORD name STRING, type STRING END RECORD, name STRING ) | Adds a sub-dialog of type `INPUT BY NAME` to an existing `ui.Dialog` dynamic dialog. |
| addTrigger( trigger STRING ) | Adds an event trigger to the dynamic dialog |
| appendNode( name STRING, parentIndex INTEGER ) RETURNS INTEGER | Appends a new node in the specified tree-view. |
| appendRow( name STRING ) | Appends a new row in the specified list. |
| arrayToVisualIndex( name STRING, arrayIndex INTEGER ) RETURNS INTEGER | Converts the program array index to the visual index for a given screen array. |
| cancel() | Cancels a parent dialog from a sub-dialog. |
| close() | Closes a dynamic dialog. |
| deleteAllRows( name STRING ) | Deletes all rows from the specified list. |
| deleteNode( name STRING, index INTEGER ) | Deletes a node from the specified tree-view. |
| deleteRow( name STRING, index INTEGER ) | Deletes a row from the specified list. |
| getArrayLength( name STRING ) RETURNS INTEGER | Returns the total number of rows in the specified list. |
| getCurrentItem() RETURNS STRING | Returns the current item having focus. |
| getCurrentRow( name STRING ) RETURNS INTEGER | Returns the current row of the specified list. |
| getEventDescription( ) RETURNS STRING | Returns a detailed description of the last event that has occurred in a dynamic dialog. |
| getFieldBuffer( name STRING ) RETURNS STRING | Returns the input buffer of the specified field. |
| getFieldTouched( formFieldList STRING ) RETURNS BOOLEAN | Returns the modification flag for a field. |
| getFieldValue( name STRING ) RETURNS fgl-type | Returns the value of a field controlled by a dynamic dialog. |
| getForm() RETURNS ui.Form | Returns the current form used by the dialog. |
| getQueryFromField( name STRING ) RETURNS STRING | Returns the SQL condition of a field used in a query by example dialog. |
| getSortKey( name STRING ) RETURNS STRING | Returns the name of the sort field selected by the user. |
| getSortKeyAt( name STRING, index INTEGER ) RETURNS STRING | Returns the name of field used as grouping or sorting column, for a given sort column position. |
| insertNode( name STRING, index INTEGER ) | Inserts a new node in the specified tree. |
| insertRow( name STRING, index INTEGER ) | Inserts a new row in the specified list. |
| isSortReverse( name STRING ) RETURNS BOOLEAN | Indicates the sort order direction (FALSE=ascending, TRUE=descending) |
| isSortReverseAt( name STRING, index: INTEGER ) RETURNS BOOLEAN | Indicates the sort order direction (FALSE=ascending, TRUE=descending), for a given sort column position. |
| isRowSelected( name STRING, row INTEGER ) RETURNS BOOLEAN | Queries row selection for a given list and row. |
| nextField( name STRING ) | Registers the next field to go to. |
| nextEvent() RETURNS STRING | Waits for a dialog event. |
| resetSort( name STRING ) | Resets the sort columns of a list dialog. |
| selectionToString( name STRING ) RETURNS STRING | Serializes data of the selected rows. |
| setActionActive( name STRING, active BOOLEAN ) | Enabling and disabling dialog actions. |
| setActionAttribute( action STRING, attrName STRING, attrValue STRING ) | Set an action attribute to configure a dynamic dialog. |
| setActionComment( name STRING, comment STRING ) | Set the comment/hint of a default action view. |
| setActionHidden( name STRING, hidden BOOLEAN ) | Showing or hiding a default action view. |
| setActionImage( name STRING, image STRING ) | Set the image of a default action view. |
| setActionText( name STRING, text STRING ) | Defining the text of a default action view. |
| setArrayAttributes( name STRING, attributes dynamic-array-type ) | Define cell decoration attributes array for the specified list (singular or multiple dialogs). |
| setArrayLength( name STRING, length INTEGER ) | Sets the number of rows in a `DISPLAY ARRAY` using paged mode. |
| setCellAttributes( attributes dynamic-array-type ) | Define cell decoration attributes array for the specified list (singular dialog only). |
| setColumnComparisonFunction( name STRING, compare FUNCTION ( s1 STRING, s2 STRING ) RETURNS INTEGER ) | Associates a comparison function to a form field of a list dialog. |
| setCompleterItems( items DYNAMIC ARRAY OF STRING ) | Define autocompletion items for a field defined with `COMPLETER` attribute. |
| setCurrentRow( name STRING, row INTEGER ) | Sets the current row in the specified list. |
| setFieldActive( formFieldList STRING, val BOOLEAN ) | Enable and disable form fields. |
| setFieldTouched( formFieldList STRING, val BOOLEAN ) | Sets the modification flag of the specified field. |
| setFieldValue( name STRING, value fgl-type ) | Sets the value of a field controlled by the dialog object. |
| setGroupBy( name STRING ) | Defines the grouping column of a list dialog. |
| setGroupByDesc( name STRING, val BOOLEAN ) | Defines sort order for a grouping column of a list dialog. |
| setSelectionMode( name STRING, mode INTEGER ) | Defines the row selection mode for the specified list. |
| setSelectionRange( name STRING, start INTEGER, end INTEGER, value BOOLEAN ) | Sets the row selection flags for a range of rows. |
| validate( formFieldList ) RETURNS INTEGER | Checks form level validation rules. |
| visualToArrayIndex( name STRING, visualIndex INTEGER ) RETURNS INTEGER | Converts the visual index to the program array index for a given screen array. |
