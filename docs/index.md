# ${{values.component_id}} Documentation

{%- if values.description %}
## Description

description: ${{values.description | dump}}
{%- endif %}

## System

${{values.system | dump}}

## Team

${{values.owner | dump}}

This is a basic example of documentation.
