#!/bin/bash

echo "Menu para instalar plugins"
echo " "

echo "Elija el nombre de tu plugin"
read name
echo "Buscando $name"

if [ "$name" == "hello" ]; then
    echo "$name encontrado"
    echo "$name instalando"
    
    # Reemplazar el archivo si ya existe
    curl -L -o ./plugins/hello.go https://raw.githubusercontent.com/Olyxe/salamandra/refs/heads/main/plugins/hello.go
    
    
    # Comprobar si el archivo fue descargado correctamente
    if [ $? -eq 0 ]; then
        echo "$name ha sido instalado correctamente."
        

    else
        echo "Hubo un error al instalar $name."
    fi

    echo "Empezando a construir $name"
    go build -buildmode=plugin -o ./plugins/hello.so plugins/hello.go
else
    echo "$name no existe"
fi
