FROM caddy:2.11.4

COPY ./public/ /usr/share/caddy
COPY ./Caddyfile /etc/caddy/Caddyfile
COPY ./redirect_routes /etc/caddy/redirect_routes
