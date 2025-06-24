FROM node:20

WORKDIR /app

COPY package.json ./
RUN npm install

COPY . .

EXPOSE 3001

CMD ["sh", "-c", "node init-db.js && npm start"]
