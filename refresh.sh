#!/usr/bin/env bash

openapi-generator generate -i https://raw.githubusercontent.com/ccouzens/keycloak-openapi/refs/heads/main/keycloak/23.0.7.yml -g go -o . --skip-validate-spec
