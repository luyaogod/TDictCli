---
title: "os.Path methods"
source: "fgl-topics/c_fgl_ext_os_path_004.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods"
type: "concept"
description: "Table 1. Class methods Name Description os.Path.atime ( path STRING) RETURNS STRING Returns the time of the last file access. os.Path.baseName ( path STRING) RETURNS STRING Returns the last element of ..."
---

# os.Path methods

| Name | Description |
| --- | --- |
| os.Path.atime( path STRING) RETURNS STRING | Returns the time of the last file access. |
| os.Path.baseName( path STRING) RETURNS STRING | Returns the last element of a path. |
| os.Path.chDir( path STRING) RETURNS INTEGER | Changes the current working directory. |
| os.Path.chOwn( path STRING, uid INTEGER, gui INTEGER ) RETURNS INTEGER | Changes the UNIX owner and group of a file. |
| os.Path.chRwx( path STRING, mode INTEGER ) RETURNS INTEGER | Changes the UNIX permissions of a file. |
| os.Path.chVolume( volume STRING) RETURNS INTEGER | Changes the current working volume. |
| os.Path.copy( fromPath STRING, toPath STRING ) RETURNS INTEGER | Creates a new file by copying an existing file. |
| os.Path.delete( path STRING) RETURNS INTEGER | Deletes a file or a directory. |
| os.Path.describeLastError( ) RETURNS STRING | Returns a description of the last error thrown by an `os.Path` method. |
| os.Path.dirClose( dirHandle INTEGER) | Closes the directory referenced by the directory opened by `os.Path.diropen()`. |
| os.Path.dirFMask( mask INTEGER) | Defines a filter mask for `os.Path.dirOpen()`. |
| os.Path.dirName( path STRING) RETURNS STRING | Returns all components of a path excluding the last one. |
| os.Path.dirNext( dirHandle INTEGER) RETURNS STRING | Reads the next entry in the directory opened with `os.Path.dirOpen()`. |
| os.Path.dirOpen( path STRING) RETURNS INTEGER | Opens a directory and returns an integer handle to this directory. |
| os.Path.dirSort( criteria STRING, order INTEGER ) | Defines the sort criteria and sort order for `os.Path.dirOpen()`. |
| os.Path.executable( path STRING) RETURNS INTEGER | Checks if a file is executable. |
| os.Path.exists( path STRING) RETURNS INTEGER | Checks if a file exists. |
| os.Path.extension( path STRING) RETURNS STRING | Returns the file extension. |
| os.Path.fullPath( path STRING) RETURNS STRING | Returns the canonical equivalent of a path. |
| os.Path.gid( path STRING) RETURNS INTEGER | Returns the UNIX group id of a file. |
| os.Path.getAccessTime( path STRING) RETURNS DATETIME YEAR TO FRACTION(5) | Returns the last time the file was accessed. |
| os.Path.getModificationTime( path STRING) RETURNS DATETIME YEAR TO FRACTION(5) | Returns the last time the file was modified. |
| os.Path.glob( pattern STRING ) RETURNS DYNAMIC ARRAY OF STRING | Returns a list of files matching the specified pattern. |
| os.Path.homeDir() RETURNS STRING | Returns the path to the HOME directory of the current user. |
| os.Path.isDirectory( path STRING) RETURNS BOOLEAN | Checks if a file is a directory. |
| os.Path.isFile( path STRING) RETURNS BOOLEAN | Checks if a file is a regular file. |
| os.Path.isHidden( path STRING) RETURNS BOOLEAN | Checks if a file is hidden. |
| os.Path.isLink( path STRING) RETURNS BOOLEAN | Checks if a file is an UNIX symbolic link. |
| os.Path.isRoot( path STRING) RETURNS BOOLEAN | Checks if a file path is a root path. |
| os.Path.isSameFile( path1 STRING, path2 STRING ) RETURNS BOOLEAN | Checks if two file paths point to the same file. |
| os.Path.join( begin STRING, end STRING ) RETURNS STRING | Joins two path segments adding the platform-dependent separator. |
| os.Path.makeTempName() RETURNS STRING | Generates a new file path to be used to create a temporary file or directory. |
| os.Path.mkDir( path STRING) RETURNS INTEGER | Creates a new directory. |
| os.Path.mtime( path STRING) RETURNS STRING | Returns the time of the last file modification. |
| os.Path.pathSeparator() RETURNS STRING | Returns the character used in environment variables to separate path elements. |
| os.Path.pathType( path STRING) RETURNS STRING | Checks if a path is a relative path or an absolute path. |
| os.Path.pwd() RETURNS STRING | Returns the current working directory. |
| os.Path.readable( path STRING) RETURNS INTEGER | Checks if a file is readable. |
| os.Path.rename( oldPath STRING, newPath STRING ) RETURNS INTEGER | Renames a file or a directory. |
| os.Path.rootDir() RETURNS STRING | Returns the root directory of the current working path. |
| os.Path.rootName( path STRING) RETURNS STRING | Returns the file path without the file extension of the last element of the file path. |
| os.Path.rwx( path STRING) RETURNS INTEGER | Returns the UNIX file permissions of a file. |
| os.Path.separator() RETURNS STRING | Returns the character used to separate path segments. |
| os.Path.setAccessTime( path STRING, time DATETIME YEAR TO FRACTION(5) ) | Sets the access time of a file. |
| os.Path.setModificationTime( path STRING, time DATETIME YEAR TO FRACTION(5) ) | Sets the modification time of a file. |
| os.Path.size( path STRING) RETURNS BIGINT | Returns the size of a file. |
| os.Path.type( path STRING) RETURNS STRING | Returns the file type as a string. |
| os.Path.uid( path STRING) RETURNS INTEGER | Returns the UNIX user id of a file. |
| os.Path.volumes() RETURNS STRING | Returns the available volumes. |
| os.Path.writable( path STRING) RETURNS INTEGER | Checks if a file is writable. |
