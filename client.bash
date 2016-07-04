#!/bin/bash

SSH_OPTIONS=(
    -l arachnist itsumade.is-a.cat ~/bin/goup
)

set -o nounset
set -e

usage() {
    echo "Usage: ${0} </file/to/upload>"
}

if [[ "${#@}" -ne 1 ]]; then
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

"${INCMD[@]}" "${1}" | ssh "${SSH_OPTIONS[@]}" "${1}"
