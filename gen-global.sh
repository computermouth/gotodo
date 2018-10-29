#!/bin/bash

echo -e "package main"                      > global.go
echo -e ""                                 >> global.go
echo -e "var ("                            >> global.go

echo -e "\tJquery = \`"                    >> global.go
cat jquery.js                              >> global.go
echo -e "\`"                               >> global.go

echo -e "\tBootst = \`"                    >> global.go
cat bootst.js                              >> global.go
echo -e "\`"                               >> global.go

echo -e "\tCustom = \`"                    >> global.go
cat custom.js                              >> global.go
echo -e "\`"                               >> global.go

echo -e "\tCss = \`"                       >> global.go
cat custom.css                             >> global.go
echo -e "\`"                               >> global.go

echo -e ")"                                >> global.go
