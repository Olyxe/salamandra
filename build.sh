#!/bin/bash

# Colores ANSI
GREEN="\e[32m"
RED="\e[31m"
YELLOW="\e[33m"
BLUE="\e[34m"
RESET="\e[0m"

echo -e "${BLUE}# Version 1.0.0${RESET}"

# Verificar si git y go est�n instalados
if ! command -v git &>/dev/null; then
    echo -e "${RED}Error: git no esta instalado.${RESET}"
    exit 1
fi

if ! command -v go &>/dev/null; then
    echo -e "${RED}Error: Go no esta instalado.${RESET}"
    exit 1
fi

# Clonar el repositorio
echo -e "${YELLOW}Clonando el repositorio...${RESET}"
if git clone https://github.com/Olyxe/salamandra; then
    cd salamandra || exit
else
    echo -e "${RED}Error al clonar el repositorio.${RESET}"
    exit 1
fi

# Compilar el programa
echo -e "${YELLOW}Compilando el programa...${RESET}"
if go build -o salamandra; then
    echo -e "${GREEN}Compilacion exitosa.${RESET}"
else
    echo -e "${RED}Error en la compilacion.${RESET}"
    exit 1
fi

# Preguntar si desea hacer la instalaci�n global
echo -e "${YELLOW}Quieres que sea global?${RESET}"
echo -e "${GREEN}1) Si${RESET}"
echo -e "${RED}2) No${RESET}"
read -r opt

# Validar la entrada
case "$opt" in
    1)
        echo -e "${GREEN}Opcion elegida: S�${RESET}"
        sudo mv salamandra /usr/local/bin/
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}Instalacion global completada.${RESET}"
        else
            echo -e "${RED}Error al mover el binario.${RESET}"
            exit 1
        fi
        ;;
    2)
        echo -e "${RED}Opcion elegida: No${RESET}"
        ;;
    *)
        echo -e "${RED}Opcion no valida. Saliendo...${RESET}"
        exit 1
        ;;
esac

echo -e "${GREEN}Instalacion terminada.${RESET}"
