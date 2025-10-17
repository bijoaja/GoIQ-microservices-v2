# FROM nginx:1.27-alpine

# # Hapus default config
# RUN rm /etc/nginx/conf.d/default.conf

# # Copy config custom
# COPY conf.d/default.conf /etc/nginx/conf.d/default.conf

# # Copy SSL certificates
# COPY ssl /etc/nginx/ssl

# EXPOSE 80
# EXPOSE 443

# CMD ["nginx", "-g", "daemon off;"]

FROM nginx:1.27-alpine

# Copy konfigurasi default sesuai mode (bisa di-overwrite lewat volume)
COPY conf.d /etc/nginx/conf.d

# Copy sertifikat SSL (optional untuk dev)
COPY ssl /etc/nginx/ssl

EXPOSE 80
EXPOSE 443
