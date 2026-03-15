# 🎮 Game Vault API

Game Vault es una API REST desarrollada nativamente en Go (net/http) que permite a los usuarios gestionar su colección personal de videojuegos. El sistema se integra con la API pública de RAWG para buscar y obtener información real de los juegos, y utiliza SQL Server para almacenar la colección y progreso personal del usuario.

Este proyecto fue construido siguiendo principios de *Arquitectura por Capas / Hexagonal*, separando claramente los controladores (handlers), la lógica de negocio (services), el acceso a datos (repositories) y la configuración.

## 🛠️ Tecnologías Utilizadas

* *Lenguaje:* Go utilizando únicamente la librería estándar net/http para el ruteo.
* *Base de Datos:* SQL Server.
* *API Externa:* RAWG Video Games Database.
* *Testing:* Paquete estándar testing y net/http/httptest.

## ⚙️ Requisitos Previos

* [Go](https://go.dev/dl/) instalado (versión 1.22 o superior recomendada para soporte de ruteo nativo).
* SQL Server instalado localmente.
* Postman (para probar los endpoints con la colección adjunta).

## 🚀 Instalación y Configuración

*1. Clonar el repositorio*
\\\`bash
git clone <URL_DE_TU_REPOSITORIO>
cd GameVaultAPI
\\\`

*2. Configurar la Base de Datos*
* Abre tu gestor de base de datos (ej. SQL Server Management Studio).
* Crea una base de datos llamada EvaluacionGoland.
* Ejecuta el script SQL ubicado en la carpeta db/schema.sql para crear la tabla game_library.

*3. Variables de Entorno*
Crea un archivo .env en la raíz del proyecto con la siguiente configuración (ajusta DB_CONNECTION según tu entorno local):
\\\`env
PORT=8080
DB_CONNECTION="Server=DESKTOP-41S3RBP;Database=EvaluacionGoland;Integrated Security=True;TrustServerCertificate=True"
RAWG_API_KEY=945d345a57cc4c3fb7b4f67211edd4c8
RAWG_BASE_URL=https://api.rawg.io/api
\\\`

*4. Descargar dependencias y ejecutar*
\\\`bash
go mod tidy
go run main.go
\\\`
El servidor iniciará en http://localhost:8080.

## ⚠️ Solución de Problemas: Conexión a SQL Server

Tuve Errores de de conexión con la db por lo cual abrí el protocolo TCP/IP en SQL Server*, ya que viene desactivado por defecto:

## 🧪 Pruebas Unitarias

El proyecto incluye pruebas unitarias para la capa de servicios con una *cobertura superior al 80%*.

Para ejecutar las pruebas y ver la cobertura:
\\\`bash
go test -cover ./services
\\\`

## 📮 Colección de Postman

Para facilitar la revisión y evaluación de todos los endpoints, se ha incluido un archivo JSON con la colección completa de Postman. 

## 📡 Endpoints Principales

* GET /api/search?q={nombre} - Busca juegos en RAWG.
* GET /api/games/{id} - Obtiene detalles de un juego de RAWG.
* GET /api/library - Lista la colección local (soporta filtro ?status=).
* POST /api/library - Agrega un juego a la base de datos local.
* PUT /api/library/{id} - Actualiza estado, nota o puntaje.
* DELETE /api/library/{id} - Elimina un juego de la colección.
* GET /api/library/stats - Retorna estadísticas de la colección.