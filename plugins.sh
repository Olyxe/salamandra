#!/bin/bash

echo "Menu para instalar plugins"
echo " "

echo "Elija el nombre de tu plugin"
read name
echo "Buscando $name"



if [ "$name" == "hello" ]; then #Change "hello" 
    echo "$name encontrado"
    echo "$name instalando"
    
    # Reemplazar el archivo si ya existe
    curl -L -o ./plugins/$name.go https://raw.githubusercontent.com/Olyxe/salamandra/refs/heads/main/plugins/hello.go #Change the curl link and directory
    
    
    # Comprobar si el archivo fue descargado correctamente
    if [ $? -eq 0 ]; then
        echo "$name ha sido instalado correctamente."
        

    else
        echo "Hubo un error al instalar $name."
    fi

    echo "Empezando a construir $name"
    go build -buildmode=plugin -o ./plugins/hello.so plugins/hello.go #change the directory name files
else
    echo "$name no existe"
fi

#Copy the last conditional and paste, add the changes for you plugin
