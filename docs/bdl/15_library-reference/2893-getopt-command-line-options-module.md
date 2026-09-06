---
title: "getopt: Command line options module"
source: "fgl-topics/r_fgl_utility_functions_getopt.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module"
type: "reference"
description: "Usage See Getopt module usage to understand how the getopt module can be used. Table 1. GetOpt types (getopt.4gl) Type Description TYPE GetoptOptions DYNAMIC ARRAY OF RECORD name STRING, description ..."
---

# getopt: Command line options module

## Usage

See [Getopt module usage](2894-getopt-module-usage.md "The getopt.4gl module provides command line argument processing.") to understand how the
`getopt` module can be used.

| Type | Description |
| --- | --- |
| TYPE GetoptOptions DYNAMIC ARRAY OF RECORD name STRING, description STRING, opt_char CHAR, arg_type INTEGER END RECORD | The GetoptOptions structured array type that holds the definition of command line options. |
| TYPE Getopt RECORD ... private members not documented here ... opt_ind INTEGER, opt_char CHAR, opt_arg STRING END RECORD | The Getopt structured type is used to process command line options. |

| Constant | Description |
| --- | --- |
| [GetOpt constants](2897-getopt-constants.md "List of predefined constants for the getopt API.") | List of predefined constants for the getopt API. |

| Function | Description |
| --- | --- |
| FUNCTION copyArguments( ind INTEGER ) RETURNS DYNAMIC ARRAY OF STRING | Returns a dynamic array of string with all command line arguments starting from the provided index. |

| Function | Description |
| --- | --- |
| FUNCTION (r Getopt) displayUsage( more_args STRING ) | Display the usage and command line option description to the standard output stream. |
| FUNCTION (r Getopt) getMoreArgumentCount( ) RETURNS INTEGER | Returns the number of command line arguments left to be processed after the known options. |
| FUNCTION (r Getopt) getMoreArgument( ind INTEGER ) RETURNS STRING | Returns the additional argument at the specified index. |
| FUNCTION (r Getopt) getopt( ) RETURNS INTEGER | Process the next command line option. |
| FUNCTION (r Getopt) initDefault( options GetoptOptions ) | Initializes a variable defined with the `Getopt` type. |
| FUNCTION (r Getopt) initialize( prog_name STRING, argv DYNAMIC ARRAY OF STRING, options GetoptOptions ) | Initializes a variable defined with the `Getopt` type for command line argument processing. |
| FUNCTION (r Getopt) invalidOptionSeen( ) RETURNS BOOLEAN | Checks if the command line options are misused. |
| FUNCTION (r Getopt) isEof( ) RETURNS BOOLEAN | Checks if there are more command line options to be read. |
| FUNCTION (r Getopt) isSuccess( ) RETURNS BOOLEAN | Checks if a command line option parsing succeeded. |
