# ${{values.component_id}}
## ${{values.description}}

[![main](https://github.com/${{values.destination.owner}}/${{values.component_id}}/actions/workflows/main.yml/badge.svg?branch=main&event=push)](https://github.com/${{values.destination.owner}}/${{values.component_id}}/actions/workflows/main.yml)
[![App Status](https://argocd.diegoluisi.eti.br/api/badge?name=${{values.env}}-${{values.component_id}}&revision=true)](https://argocd.diegoluisi.eti.br/applications/${{values.env}}-${{values.component_id}})
[![Go Report Card](https://goreportcard.com/badge/github.com/${{values.destination.owner}}/${{values.component_id}})](https://goreportcard.com/report/github.com/${{values.destination.owner}}/${{values.component_id}})
![GitHub repo size](https://img.shields.io/github/repo-size/${{values.destination.owner}}/${{values.component_id}})
![GitHub issues](https://img.shields.io/github/issues/${{values.destination.owner}}/${{values.component_id}})
![GitHub](https://img.shields.io/github/license/${{values.destination.owner}}/${{values.component_id}})
