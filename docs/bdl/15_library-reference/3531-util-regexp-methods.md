---
title: "util.Regexp methods"
source: "fgl-topics/c_fgl_ext_util_Regexp_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods"
type: "concept"
---

# util.Regexp methods

> Methods for the util.Regexp class.

| Name | Description |
| --- | --- |
| util.Regexp.compile( expr STRING ) RETURNS util.Regexp | Compiles a regular expression and returns a `util.Regexp` object. |

| Name | Description |
| --- | --- |
| getMatch( subject STRING ) RETURNS STRING | Returns the text of the first (leftmost) match of the regular expression in the subject. |
| getMatchAll( subject STRING ) RETURNS DYNAMIC ARRAY OF STRING | Returns an array with all matches of the regular expression in the subject. |
| getMatchIndex( subject STRING ) RETURNS (INTEGER, INTEGER) | Returns the starting and ending position of the first (leftmost) match of the regular expression in the subject. |
| getMatchIndexAll( subject STRING ) RETURNS DYNAMIC ARRAY OF INTEGER | Returns an array with indexes to all matches of the regular expression in the subject. |
| getSubmatch( subject STRING ) RETURNS DYNAMIC ARRAY OF STRING | Returns an array with the sub-matches of the first match of regular expression in the subject. |
| getSubmatchAll( subject STRING ) RETURNS DYNAMIC ARRAY WITH DIMENSION 2 OF STRING | Returns 2-dimensional array with the sub-matches of all matches of regular expression in the subject. |
| getSubmatchIndex( subject STRING ) RETURNS DYNAMIC ARRAY OF INTEGER | Returns an array with the indexes of the sub-matches of the first match of regular expression in the subject. |
| getSubmatchIndexAll( subject STRING ) RETURNS DYNAMIC ARRAY WITH DIMENSION 2 OF INTEGER | Returns a 2-dimensional array with the indexes of the sub-matches of all matches of regular expression in the subject. |
| matches( subject STRING ) RETURNS BOOLEAN | Tests if a string contains any match of this regular expression. |
| replaceAll( subject STRING, replacement STRING ) RETURNS STRING | Substitutes all matches of the regular expression with the replacement string. |
| replaceFirst( subject STRING, replacement STRING ) RETURNS STRING | Substitutes the match of the regular expression with the replacement string. |
