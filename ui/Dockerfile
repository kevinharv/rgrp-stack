# Node container to build application
FROM node:22-bookworm AS builder

WORKDIR /app
RUN npm install -g pnpm
COPY package*.json .
RUN pnpm install
COPY . .
RUN pnpm run build

# Python container to generate config file(s)
FROM python:3.12-bookworm AS scripter

RUN apt-get update && apt-get install openssl
RUN mkdir /certs && \
    openssl req -newkey rsa:4096 -x509 -sha512 -days 365 -nodes -out /certs/cert.pem -keyout /certs/key.pem \
    -subj "/CN=localhost"

WORKDIR /app
COPY --from=builder /app/dist/.vite/manifest.json manifest.json
COPY generate_nginx_directives.py .
RUN python generate_nginx_directives.py

# NGINX container to serve application content
FROM nginx:latest AS runner

# To-Do - create service account to run NGINX

RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=scripter /certs /etc/nginx/ssl
COPY --from=scripter /app/push_directives.conf /etc/nginx/conf.d/push_directives.conf 
COPY --from=builder /app/dist /usr/share/nginx/html
RUN rm -r /usr/share/nginx/html/.vite/

EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]