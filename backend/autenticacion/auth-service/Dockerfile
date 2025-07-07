FROM node:20

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY certs /app/certs
COPY . .

EXPOSE 3001

CMD ["node", "start.js"]
