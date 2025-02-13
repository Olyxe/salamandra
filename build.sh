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
    echo -e "${RED}Error: git no est� instalado.${RESET}"
    exit 1
fi

if ! command -v go &>/dev/null; then
    echo -e "${RED}Error: Go no est� instalado.${RESET}"
    exit 1
fi

# Clonar el repositorio
echo -e "${YELLOW}Clonando el repositorio...${RESET}"

REPO_URL="https://github.com/Olyxe/salamandra"
INCLUDED_FOLDERS="cmd plugins"
EXCLUDED_FOLDER="web"

echo "Clonando el repositorio..."
if git clone --no-checkout "$REPO_URL"; then
    cd salamandra || exit 1

    # Inicializar sparse-checkout
    git sparse-checkout init --cone

    # Incluir las carpetas deseadas y excluir la carpeta web
    git sparse-checkout set $INCLUDED_FOLDERS "!$EXCLUDED_FOLDER" --skip-checks

    # Realizar el checkout
    git checkout

    echo -e "${GREEN}Clonaci�n completada con solo las carpetas '$INCLUDED_FOLDERS'.${RESET}"
else
    echo -e "${RED}Error al clonar el repositorio.${RESET}"
    exit 1
fi

# Compilar el programa
echo -e "${YELLOW}Compilando el programa...${RESET}"
if go build -o salamandra; then
    echo -e "${GREEN}Compilaci�n exitosa.${RESET}"
else
    echo -e "${RED}Error en la compilaci�n.${RESET}"
    exit 1
fi

# Preguntar si desea hacer la instalaci�n global
echo -e "${YELLOW}�Quieres que sea global?${RESET}"
echo -e "${GREEN}1) S�${RESET}"
echo -e "${RED}2) No${RESET}"
read -r opt

# Validar la entrada
case "$opt" in
    1)
        echo -e "${GREEN}Opci�n elegida: S�${RESET}"
        sudo mv salamandra /usr/local/bin/
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}Instalaci�n global completada.${RESET}"
            if rm -rf salamandra/; then
                echo "Desistalando la carpeta de instalacion"
            else
                echo "No se pudo desinstalar la carpeta de instalacion"
            fi
            
        else
            echo -e "${RED}Error al mover el binario.${RESET}"
            exit 1
        fi
        
        ;;
    2)
        echo -e "${RED}Opci�n elegida: No${RESET}"
        ;;
    *)
        echo -e "${RED}Opci�n no v�lida. Saliendo...${RESET}"
        exit 1
        ;;
esac

echo -e "${GREEN}Instalaci�n terminada.${RESET}"
