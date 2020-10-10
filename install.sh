#!/bin/bash

if [ -d "todo/" ]; then
    echo "Directory todo already exists."
    exit 1
fi

git clone https://github.com/dariubs/todo.git
cd todo
make && make install