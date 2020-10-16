#! /bin/bash

[[ -z "${PREFIX}" ]] && PREFIX=/usr/local

SCRIPT_PATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

if [[ $(find "$SCRIPT_PATH"/bin -maxdepth 1 -perm -111 -type f | wc -l) -le 0 ]]; then
    echo "Error: No executable files found in directory ${SCRIPT_PATH}/bin !"
    exit 1;
fi

for f in $(find "$SCRIPT_PATH"/bin -maxdepth 1 -perm -111 -type f); do
    echo "Installing executable file $f to ${PREFIX}/bin";
    cp "${f}" ${PREFIX}/bin
done

echo "All executable files installed to ${PREFIX}/bin"