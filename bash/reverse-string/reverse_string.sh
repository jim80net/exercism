#!/usr/bin/env bash

#perl -0777e "print scalar reverse \"$@\""
#rev <<< "$@"

reverse() {
  local original=$1
  local reversed=""
  local length=${#original}

  for(( index=length-1; index>=0; index-- ))
  do
    reversed+=${original:index:1}
  done

  echo $reversed
}

reverse "$@"
