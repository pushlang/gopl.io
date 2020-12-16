set nocompatible
filetype off

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'VundleVim/Vundle.Vim'

call vundle#end()

filetype plugin indent on

colorscheme codeschool

Plugin 'fatih/vim-go'
Plugin 'scrooloose/nerdtree'
Plugin 'Shougo/neocomplete'
" Plugin 'nsf/gocode', {'rtp': 'vim/'}

autocmd vimenter * NERDTree " Открывать дерево каталогов при запуске VIM
set number " номера строк
set colorcolumn=80 " вертикальная линия после 120 символов
set tabstop=4 " ширина таба
set softtabstop=4 " ширина таба при исп пробелов

let g:neocomplete#enable_at_startup = 1
set completeopt-=preview

" CTRL-s - сохранить файл
imap <C-S> <esc>:w<cr>i
nmap <C-S> :w<cr>
" noremap <silent> <C-s> :update<cr>
" vnoremap <silent> <C-s> <C-c>:update<cr>
" inoremap <silent> <C-s> <C-o>:update<cr>

" CTRL-F4 - закрыть окно
noremap <C-F4> <C-w>c
inoremap <C-F4> <C-o><C-w>c
cnoremap <C-F4> <C-c><C-w>
onoremap <C-F4> <C-c><C-w>c

" SHIFT-Del - "вырезание" в системный буфер
vnoremap <S-Del> "+x
" CTRL-Insert - копирование в системный буфер
vnoremap <C-Insert> "+y
" SHIFT-Insert - вставка из системного буфера
map <S-Insert>  "+gP
cmap <S-Insert> <C-R>+

" CTRL-z - отмена действия
noremap <C-z> u
inoremap <C-z> <C-O>u
" CTRL-y - вернуть отменённое назад
noremap <C-y> <C-R>
inoremap <C-y> <C-O><C-R>

" CTRL-d - дублирование текущей строки
imap <C-d> <esc>yypi