set nocompatible
filetype off

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'VundleVim/Vundle.vim'
Plugin 'fatih/vim-go'
Plugin 'scrooloose/nerdtree'
Plugin 'Shougo/neocomplete'
Plugin 'nsf/gocode', {'rtp': 'vim/'}
Plugin 'SirVer/ultisnips'

call vundle#end()

filetype plugin indent on

colorscheme codeschool
autocmd vimenter * NERDTree
let g:neocomplete#enable_at_startup = 1
set completeopt-=preview

set number
set relativenumber
set colorcolumn=80
set tabstop=4
set softtabstop=4 " ширина таба при использовании пробелов
set shiftwidth=4
set noexpandtab

set foldenable
set foldcolumn=2
set foldmethod=syntax

highlight Folded guibg=#141414 guifg=White
highlight FoldColumn guibg=DarkGrey guifg=White


"nnoremap <silent> <Leader>+ :exe "resize " . (winheight(0) * 3/2)<CR>
"nnoremap <silent> <Leader>- :exe "resize " . (winheight(0) * 2/3)<CR>

"set scrollopt+=hor
"set scrollopt-=ver
"set nowrap
"set mouse=a
"2split
"windo set scrollbind
"set statusline=%F%m%r%h%w[%{&ff}]%y%p%%\ %04l\/%04L\ Oc:Hx\ %b:%B%=%qF[%n]%t


" CTRL-s - сохранить файл
noremap <silent> <C-S> :update<CR>
vnoremap <silent> <C-S> <C-C>:update<CR>
inoremap <silent> <C-S> <C-O>:update<CR>

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
