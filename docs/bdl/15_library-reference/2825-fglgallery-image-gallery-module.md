---
title: "fglgallery: Image gallery module"
source: "fgl-topics/r_fgl_utility_functions_fglgallery.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module"
type: "reference"
description: "Table 1. fglgallery types (fglgallery.4gl) Type Description TYPE t_struct_value RECORD current INTEGER, selected DYNAMIC ARRAY OF INTEGER END RECORD The t_struct_value type holds image selection data. ..."
---

# fglgallery: Image gallery module

| Type | Description |
| --- | --- |
| TYPE t_struct_value RECORD current INTEGER, selected DYNAMIC ARRAY OF INTEGER END RECORD | The `t_struct_value` type holds image selection data. |

| Constant | Description |
| --- | --- |
| FUNCTION addImage( id SMALLINT, path STRING, title STRING ) | Adds a picture resource to an fglgallery. |
| FUNCTION clean( id SMALLINT ) | Removes all pictures from an fglgallery. |
| FUNCTION create( name STRING ) RETURNS SMALLINT | Creates a new fglgallery handle. |
| FUNCTION deleteImages( id SMALLINT, indexes DYNAMIC ARRAY OF INTEGER ) RETURNS STRING | Deletes pictures used in an fglgallery. |
| FUNCTION destroy( id SMALLINT ) | Frees resources allocated for an fglgallery. |
| FUNCTION display( id SMALLINT, type INTEGER, size INTEGER ) | Displays an fglgallery to the end user. |
| FUNCTION finalize( ) | Releases the fglgallery library. |
| FUNCTION flush( id SMALLINT ) | Displays new added images to the end user. |
| FUNCTION getImageCount( id SMALLINT ) RETURNS INTEGER | Returns the number of pictures in an fglgallery. |
| FUNCTION getPath( id SMALLINT, index STRING ) RETURNS STRING | Returns the URL of a picture in an fglgallery. |
| FUNCTION getTitle( id SMALLINT, index STRING ) RETURNS STRING | Returns the description of a picture in an fglgallery. |
| FUNCTION initialize( ) | Prepares the fglgallery library for use. |
| FUNCTION setImageAspectRatio( id SMALLINT, ratio DECIMAL(5,2) ) | Prepares the fglgallery library for use. |
| FUNCTION setMultipleSelection( id SMALLINT, on BOOLEAN ) | Prepares the fglgallery library for use. |
