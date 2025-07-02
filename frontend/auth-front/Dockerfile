# Construcción
FROM node:20 AS builder
WORKDIR /app
COPY . .
RUN npm install && npm run build

# Servir con nginx
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
