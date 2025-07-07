# Imagen base
FROM node:18

# Crear carpeta de trabajo
WORKDIR /app

# Copiar dependencias
COPY package*.json ./
RUN npm install

# Copiar el código fuente
COPY . .

# Puerto expuesto
EXPOSE 3014

# Comando de inicio
CMD ["node", "main.js"]
