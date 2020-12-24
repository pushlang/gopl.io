set nocompatible
filetype off

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'VundleVim/Vundle.vim'
Plugin 'fatih/vim-go'
Plugin 'preservim/nerdtree'
"Plugin 'Shougo/neocomplete'
"Plugin 'nsf/gocode', {'rtp': 'vim/'}
Bundle 'Blackrush/vim-gocode'


call vundle#end()

filetype plugin indent on

colorscheme codeschool
"autocmd vimenter * NERDTree "Открывать при запуске
"let g:neocomplete#enable_at_startup = 1
set completeopt-=preview

set nu
set rnu
set tabstop=4
set softtabstop=4
set shiftwidth=4
set nowrap
