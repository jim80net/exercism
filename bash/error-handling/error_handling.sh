#!/bin/bash

usage() {
  echo 'Usage: ./error_handling <greetee>'
  exit 1
}

[[ $# -eq 1 ]] && {
  echo "Hello, $1"
} || {
  usage
}

