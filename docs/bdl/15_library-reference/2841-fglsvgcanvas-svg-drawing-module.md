---
title: "fglsvgcanvas: SVG drawing module"
source: "fgl-topics/r_fgl_utility_functions_fglsvgcanvas.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module"
type: "reference"
description: "Table 1. fglsvgcanvas types (fglsvgcanvas.4gl) Type Description TYPE t_svg_rect RECORD x DECIMAL, y DECIMAL, width DECIMAL, height DECIMAL END RECORD The t_svg_rect type defines the position and ..."
---

# fglsvgcanvas: SVG drawing module

| Type | Description |
| --- | --- |
| TYPE t_svg_rect RECORD x DECIMAL, y DECIMAL, width DECIMAL, height DECIMAL END RECORD | The `t_svg_rect` type defines the position and dimensions of a rectangle. |

| Constant | Description |
| --- | --- |
| [`SVGATT_*`](2843-fglsvgcanvas-svgatt-constants.md "List of predefined SVG attributes.") constants | List of predefined SVG attributes. |

| Function | Description |
| --- | --- |
| FUNCTION animateTransform( attributeName STRING, attributeType STRING, type STRING, from STRING, to STRING, by STRING, begin STRING, dur STRING, repeatCount STRING ) RETURNS om.DomNode | Produces an SVG `"animateTransform"` element. |
| FUNCTION circle( cx STRING, cy STRING, r STRING ) RETURNS om.DomNode | Produces an SVG `"circle"` element. |
| FUNCTION clean( cid SMALLINT ) | Deletes all SVG elements inside the SVG canvas. |
| FUNCTION clipPath_rect( id STRING, x STRING, y STRING, width STRING, height STRING ) RETURNS om.DomNode | Produces an SVG `"clipPath"` element with a "`rect`" element. |
| FUNCTION color_shade( source STRING, factor DECIMAL ) RETURNS STRING | Applies a shade factor to an RGB color. |
| FUNCTION color_tint( source STRING, factor DECIMAL ) RETURNS STRING | Applies a tint factor to an RGB color. |
| FUNCTION create( name STRING ) RETURNS SMALLINT | Creates a new SVG canvas handler. |
| FUNCTION createChars( value STRING ) RETURNS om.DomNode | Produces an SVG DOM text node. |
| FUNCTION createElement( tagName STRING, id STRING ) RETURNS om.DomNode | Produces an SVG DOM element with the tag name specified as parameter. |
| FUNCTION defs( id STRING ) RETURNS om.DomNode | Produces an SVG `"defs"` element. |
| FUNCTION destroy( cid SMALLINT ) | Releases resources allocated for the SVG canvas. |
| FUNCTION display( cid SMALLINT ) | Displays the SVG canvas. |
| FUNCTION ellipse( cx STRING, cy STRING, rx STRING, ry STRING ) RETURNS om.DomNode | Produces an SVG `"ellipse"` element. |
| FUNCTION filter( id STRING, x STRING, y STRING, width STRING, height STRING ) RETURNS om.DomNode | Produces the SVG `"filter"` element. |
| FUNCTION finalize( ) | Releases the fglsvgcanvas library. |
| FUNCTION g( id STRING ) RETURNS om.DomNode | Produces an SVG `"g"` element. |
| FUNCTION getBBox( cid SMALLINT, id STRING ) RETURNS t_svg_rect | Returns the bounding box of an SVG element. |
| FUNCTION getItemId( cid SMALLINT ) RETURNS STRING | Returns SVG element id after a user action. |
| FUNCTION image( href STRING, x STRING, y STRING, width STRING, height STRING, preserveAspectRatio STRING ) RETURNS om.DomNode | Produces an SVG `"image"` element. |
| FUNCTION initialize( ) | Prepares the fglsvgcanvas library for use. |
| FUNCTION line( x1 STRING, y1 STRING, x2 STRING, y2 STRING ) RETURNS om.DomNode | Produces an SVG `"line"` element. |
| FUNCTION linearGradient( id STRING, x1 STRING, y1 STRING, x2 STRING, y2 STRING, spreadMethod STRING, gradientTransform STRING, gradientUnits STRING ) RETURNS om.DomNode | Produces an SVG `"linearGradient"` element. |
| FUNCTION marker( id STRING, markerUnits STRING, refX STRING, refY STRING, markerWidth STRING, markerHeight STRING, orient STRING ) RETURNS om.DomNode | Produces an SVG `"marker"` element. |
| FUNCTION mask( id STRING, x STRING, y STRING, width STRING, height STRING, node om.DomNode, name STRING ) | Produces an SVG `"mask"` element. |
| FUNCTION nl_to_tspan( text om.DomNode, x STRING, y STRING, dx STRING, dy STRING, content STRING ) | Converts a string to an SVG `"text"` element with `"tspan"` sub-elements. |
| FUNCTION path( d STRING ) RETURNS om.DomNode | Produces an SVG `"path"` element. |
| FUNCTION pattern( id STRING, x STRING, y STRING, width STRING, height STRING, patternUnits STRING, patternContentUnits STRING, patternTransform STRING, preserveAspectRatio STRING ) RETURNS om.DomNode | Produces an SVG `"pattern"` element. |
| FUNCTION polygon( points STRING ) RETURNS om.DomNode | Produces an SVG `"polygon"` element. |
| FUNCTION polyline( points STRING ) RETURNS om.DomNode | Produces an SVG `"polyline"` element. |
| FUNCTION radialgradient( id STRING, cx STRING, cy STRING, fx STRING, fy STRING, r STRING, spreadMethod STRING, gradientTransform STRING, gradientUnits STRING ) RETURNS om.DomNode | Produces an SVG `"radialGradient"` element. |
| FUNCTION rect( x STRING, y STRING, width STRING, height STRING, rx STRING, ry STRING ) RETURNS om.DomNode | Produces an SVG `"rect"` element. |
| FUNCTION removeElement( node om.DomNode ) | Deletes an SVG element from the SVG canvas. |
| FUNCTION setAttributes( node om.DomNode, attrs om.SaxAttributes ) | Sets the SVG attributes from an attribute set. |
| FUNCTION setCurrent( cid SMALLINT ) | Selects the SVG canvas handler for subsequent SVG canvas API calls. |
| FUNCTION setRootSVGAttributes( id STRING, width STRING, height STRING, viewBox STRING, preserveAspectRatio STRING ) RETURNS om.DomNode | Produces the root SVG element. |
| FUNCTION stop( offset STRING, color STRING, opacity STRING ) RETURNS om.DomNode | Produces an SVG `"stop"` element for gradients. |
| FUNCTION styleAttributeList( attrs om.SaxAttributes ) RETURNS STRING | Builds a string with a list of attributes to be used in a `style` attribute. |
| FUNCTION styleDefinition( selector STRING, attrs om.SaxAttributes ) RETURNS STRING | Produces a CSS style definition with a selection and list of attributes. |
| FUNCTION styleList( content STRING ) RETURNS om.DomNode | Produces a CSS style list. |
| FUNCTION svg( id STRING, x STRING, y STRING, width STRING, height STRING, viewBox STRING, preserveAspectRatio STRING RETURNS om.DomNode | Produces an SVG `"svg"` element. |
| FUNCTION symbol( id STRING, width STRING, height STRING, viewBox STRING, preserveAspectRatio STRING, refX STRING, refY STRING ) RETURNS om.DomNode | Produces an SVG `"symbol"` element. |
| FUNCTION text( x STRING, y STRING, content STRING, class STRING ) RETURNS om.DomNode | Produces an SVG `"text"` element. |
| FUNCTION text_path( x STRING, y STRING, content STRING, path STRING class STRING ) RETURNS om.DomNode | Produces the SVG `"text"` element with a `"textPath"` sub-element. |
| FUNCTION text_tref( x STRING, y STRING, tref STRING, class STRING ) RETURNS om.DomNode | Produces the SVG `"text"` element with a `"tref"` sub-element. |
| FUNCTION title( text STRING ) RETURNS om.DomNode | Produces an SVG `"title"` element. |
| FUNCTION tspan( x STRING, y STRING, dx STRING, dy STRING, style STRING, content STRING ) RETURNS om.DomNode | Produces an SVG `"tspan"` element. |
| FUNCTION url( name STRING ) RETURNS STRING | Produces a `"url(#name)"` reference for SVG elements. |
| FUNCTION use( name STRING, x STRING, y STRING ) RETURNS om.DomNode | Produces an SVG `"use"` element. |
