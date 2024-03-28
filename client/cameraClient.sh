#! /bin/bash

dir=$(dirname $(readlink -f "$0"))

source "$dir/venv/bin/activate"
python "$dir/main.py"
