#!/bin/bash

printf "** Building linux/386\n"
go-linux-386 build -o bin/linux-386/der_shortcut github.com/zerklabs/gopin/der_shortcut
go-linux-386 build -o bin/linux-386/get_spki github.com/zerklabs/gopin/get_spki

printf "** Building linux/amd64\n"
go-linux-amd64 build -o bin/linux-amd64/der_shortcut github.com/zerklabs/gopin/der_shortcut
go-linux-amd64 build -o bin/linux-amd64/get_spki github.com/zerklabs/gopin/get_spki

printf "** Building windows/386\n"
go-windows-386 build -o bin/windows-386/der_shortcut.exe github.com/zerklabs/gopin/der_shortcut
go-windows-386 build -o bin/windows-386/get_spki.exe github.com/zerklabs/gopin/get_spki

printf "** Building windows/amd64\n"
go-windows-amd64 build -o bin/windows-amd64/der_shortcut.exe github.com/zerklabs/gopin/der_shortcut
go-windows-amd64 build -o bin/windows-amd64/get_spki.exe github.com/zerklabs/gopin/get_spki
