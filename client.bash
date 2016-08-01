#!/bin/bash

SSH_OPTIONS=(
    -l arachnist itsumade.is-a.cat ~/bin/goup
)

set -o nounset
set -e

usage() {
    echo "Usage: ${0} </file/to/upload>"
}

if [[ "${#@}" -eq 2 ]]; then
    filename="${2}"
elif [[ "${#@}" -eq 1 ]]; then
    filename="${1}"
else
    usage
    exit 1
fi

if [[ "${1}" =~ ^https?:// ]]; then
    INCMD=( curl )
elif type pv > /dev/null; then
    INCMD=( pv --width=80 )
else
    INCMD=( cat )
fi

"${INCMD[@]}" "${1}" | ssh "${SSH_OPTIONS[@]}" "${filename}"
