# Node container to build application
FROM node:22-bookworm AS builder

# Generate SSL certificates for dev
RUN apt-get update && apt-get install openssl
RUN mkdir /certs && \
    openssl req -newkey rsa:4096 -x509 -sha512 -days 365 -nodes -out /certs/cert.pem -keyout /certs/key.pem \
    -subj "/CN=localhost"
RUN openssl rand -base64 32 > /certs/quic_host.key

# Build React app
WORKDIR /app
RUN npm install -g pnpm
COPY package*.json .
RUN pnpm install
COPY . .
RUN pnpm run build


# NGINX container to serve application content
FROM nginx:latest AS runner

RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /certs /etc/nginx/ssl
COPY --from=builder /app/dist /usr/share/nginx/html
RUN chown -R nginx:nginx /usr/share/nginx/html
RUN rm -r /usr/share/nginx/html/.vite/

EXPOSE 443/tcp
EXPOSE 443/udp
CMD ["nginx", "-g", "daemon off;"]