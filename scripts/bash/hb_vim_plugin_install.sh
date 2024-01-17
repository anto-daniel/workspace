#!/bin/bash

mkdir -p -v ~/.local/src
git clone https://github.com/mustache/vim-mustache-handlebars.git mustache.vim
mkdir -p -v ~/.vim/{syntax,ftdetect,ftplugin}
cp -R mustache.vim/syntax/* ~/.vim/syntax/
cp -R mustache.vim/ftdetect/* ~/.vim/ftdetect/
cp -R mustache.vim/ftplugin/* ~/.vim/ftplugin/
if [ $? -eq 0 ]; then
    echo "INFO: Copied mustache.vim to vim directory"
else
    echo "ERROR: Failed to copy mustache.vim to vim directory"
fi
