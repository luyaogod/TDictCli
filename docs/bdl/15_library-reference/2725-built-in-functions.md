---
title: "Built-in functions"
source: "fgl-topics/r_fgl_BuiltInFunctions_list.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions"
type: "reference"
description: "Table 1. Built-in functions Function Description FUNCTION arg_val ( index INTEGER ) RETURNS STRING Returns a command line argument by position. FUNCTION arr_count () RETURNS INTEGER Returns the number ..."
---

# Built-in functions

| Function | Description |
| --- | --- |
| FUNCTION arg_val( index INTEGER ) RETURNS STRING | Returns a command line argument by position. |
| FUNCTION arr_count() RETURNS INTEGER | Returns the number of rows entered during an `INPUT ARRAY` statement. |
| FUNCTION arr_curr() RETURNS INTEGER | Returns the current row in a `DISPLAY ARRAY` or `INPUT ARRAY`. |
| FUNCTION downshift( s STRING ) RETURNS STRING | Converts a string to lowercase. |
| FUNCTION err_get( messageId INTEGER ) RETURNS STRING | Returns the text corresponding to an error number. |
| FUNCTION err_print( messageId INTEGER ) | Prints in the error line the text corresponding to an error number. |
| FUNCTION err_quit( messageId INTEGER ) | Prints in the error line the text corresponding to an error number and terminates the program. |
| FUNCTION errorlog( message STRING ) | Copies the string passed as parameter into the error log file. |
| FUNCTION fgl_buffertouched() RETURNS INTEGER | Returns `TRUE` if the input buffer was modified in the current field. |
| FUNCTION fgl_db_driver_type() RETURNS CHAR(3) | Returns the 3-letter identifier/code of the current database driver. |
| FUNCTION fgl_decimal_truncate( x DECIMAL, scale INTEGER ) RETURNS DECIMAL | Returns a decimal truncated to the precision passed as parameter. |
| FUNCTION fgl_decimal_sqrt( x DECIMAL ) RETURNS DECIMAL | Computes the square root of the decimal passed as parameter. |
| FUNCTION fgl_decimal_exp( x DECIMAL ) RETURNS DECIMAL | Returns the value of Euler's constant (e) raised to the power of the decimal passed as parameter. |
| FUNCTION fgl_decimal_logn( x DECIMAL ) RETURNS DECIMAL | Returns the natural logarithm of the decimal passed as parameter. |
| FUNCTION fgl_decimal_power( x DECIMAL, y DECIMAL ) RETURNS DECIMAL | Raises decimal to the power of the real exponent. |
| FUNCTION fgl_dialog_getbuffer() RETURNS STRING | Returns the text of the input buffer of the current field. |
| FUNCTION fgl_dialog_getbufferlength() RETURNS INTEGER | Returns the number of rows to feed a paged DISPLAY ARRAY. |
| FUNCTION fgl_dialog_getbufferstart() RETURNS INTEGER | Returns the row offset of the page to feed a paged display array. |
| FUNCTION fgl_dialog_getcursor() RETURNS INTEGER | Returns the position of the edit cursor in the current field. |
| FUNCTION fgl_dialog_getfieldname() RETURNS STRING | Returns the name of the current input field. |
| FUNCTION fgl_dialog_getkeylabel( keyName STRING ) RETURNS STRING | Returns the label associated to a key for the current interactive instruction. |
| FUNCTION fgl_dialog_getselectionend() RETURNS INTEGER | Returns the position of the last selected character in the current field. |
| FUNCTION fgl_dialog_infield( name STRING ) RETURNS INTEGER | This function checks for the current input field. |
| FUNCTION fgl_dialog_setbuffer( value STRING ) | Sets the input buffer of the current field. |
| FUNCTION fgl_dialog_setcurrline( screenLine INTEGER, row INTEGER ) | This function moves to a specific row in a record list. |
| FUNCTION fgl_dialog_setcursor( x INTEGER ) | This function sets the position of the edit cursor in the current field. |
| FUNCTION fgl_dialog_setfieldorder( constrainded INTEGER ) | This function enables or disables field order constraint. |
| FUNCTION fgl_dialog_setkeylabel( keyName STRING, text STRING ) | Sets the label associated to a key for the current interactive instruction. |
| FUNCTION fgl_dialog_setselection( start INTEGER, end INTEGER ) | Selects the text in the current field. |
| FUNCTION fgl_drawbox( height INTEGER, width INTEGER, posY INTEGER, posX INTEGER, color INTEGER ) | Draws a rectangle in the current window. |
| FUNCTION fgl_drawline( posY INTEGER, posX INTEGER, width INTEGER, type CHAR(1), color INTEGER) | Draws a line in the current window (TUI and traditional mode). |
| FUNCTION fgl_getenv( name STRING ) RETURNS STRING | Returns the value of the environment variable. |
| FUNCTION fgl_getfile( source STRING, target STRING ) | Retrieves a file from the front-end context to the virtual machine context. |
| FUNCTION fgl_gethelp( id INTEGER ) RETURNS STRING | Reads the current help file, returning help text based on the help identifier. |
| FUNCTION fgl_getkey() RETURNS INTEGER | Waits for a keystroke and returns the key number. |
| FUNCTION fgl_getkeylabel( keyName STRING ) RETURNS STRING | Returns the default label associated to a key. |
| FUNCTION fgl_getpid() RETURNS INTEGER | Returns the system process identifier. |
| FUNCTION fgl_getresource( name STRING ) RETURNS STRING | Returns the value of an FGLPROFILE entry. |
| FUNCTION fgl_getversion() RETURNS STRING | Returns the product version number of Genero. |
| FUNCTION fgl_getwin_height() RETURNS INTEGER | Returns the number of rows of the current window. |
| FUNCTION fgl_getwin_width() RETURNS INTEGER | Returns the width of the current window as a number of columns. |
| FUNCTION fgl_getwin_x() RETURNS INTEGER | Returns the horizontal position of the current window. |
| FUNCTION fgl_getwin_y() RETURNS INTEGER | Returns the vertical position of the current window. |
| FUNCTION fgl_keyval( keyName STRING ) RETURNS INTEGER | Returns the key code of a logical or physical key. |
| FUNCTION fgl_lastkey() RETURNS INTEGER | Returns the key code corresponding to the logical key that the user most recently typed in the form. |
| FUNCTION fgl_mblen( str STRING ) RETURNS INTEGER | Returns the number of bytes of the first character in a string. |
| FUNCTION fgl_putfile( source STRING, target STRING) | Transfers a file from the virtual machine context to the front-end context. |
| FUNCTION fgl_report_print_binary_file( path STRING ) | Prints a file containing binary data during a report. |
| FUNCTION fgl_report_set_document_handler( handler om.SaxDocumentHandler ) | Redirects the next report to an XML document handler. |
| FUNCTION fgl_scr_size( name STRING ) RETURNS INTEGER | Returns the size of the specified screen array in the current form. |
| FUNCTION fgl_set_arr_curr( row INTEGER ) | Moves to a specific row in a record list. |
| FUNCTION fgl_setenv( name STRING, value STRING ) | Sets the value of an environment variable. |
| FUNCTION fgl_setkeylabel( keyName STRING, text STRING ) | Sets the default label associated to a key. |
| FUNCTION fgl_setsize( height INTEGER, width INTEGER ) | Sets the size of the main application window. |
| FUNCTION fgl_settitle( newTitle STRING ) | Sets the title of the current application window. |
| FUNCTION fgl_sqldebug( level INTEGER ) | Sets the SQL debug level from program code. |
| FUNCTION fgl_system( program STRING ) | Runs a command on the application server. |
| FUNCTION fgl_width( str STRING ) RETURNS INTEGER | Returns the number of columns needed to represent the printed version of the expression. |
| FUNCTION fgl_window_getoption( name STRING ) RETURNS STRING | Returns attributes of the current window. |
| FUNCTION length( str STRING ) RETURNS INTEGER | Returns the number of characters in a string passed as parameter. |
| FUNCTION num_args() RETURNS INTEGER | Returns the number of program arguments. |
| FUNCTION scr_line() RETURNS INTEGER | Returns the index of the current row in the form screen array. |
| FUNCTION set_count( count INTEGER ) | Defines the number of rows containing explicit data in a static array used by the next dialog. |
| FUNCTION showhelp( number INTEGER ) | Displays a runtime help text. |
| FUNCTION startlog( path STRING ) | Initializes error logging and opens the error log file passed as the parameter. |
| FUNCTION upshift( s STRING ) RETURNS STRING | Converts a string to uppercase. |
