#!/bin/bash

# Versi�n 1.0.0
echo "Clonando el repositorio..."
git clone https://github.com/Olyxe/salamandra
cd salamandra || exit

echo "Compilando el programa..."
go build -o salamandra

echo "�Quieres que sea global?"
echo "1 [y] 2 [n]"
read -r opt

if [ "$opt" -eq 1 ]; then
    echo "Opci�n elegida: $opt"
    sudo mv salamandra /usr/local/bin/
else
    echo "Opci�n elegida: $opt"
fi

echo "Instalaci�n terminada"
