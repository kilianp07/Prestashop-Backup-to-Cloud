# Utilisez une image de base Go pour la compilation
FROM golang:1.20 as builder
 
# Copiez le code source de votre projet dans le conteneur
WORKDIR /app
COPY . .

# Compilez votre projet Go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /bin/prestashop_backup

# Image de base minimale pour exécuter le binaire Go
FROM debian:bullseye-slim
LABEL org.opencontainers.image.source="https://github.com/kilianp07/Prestashop-Backup-to-Cloud"

# Copiez le binaire compilé depuis l'image précédente
COPY --from=builder /bin/prestashop_backup /bin/prestashop_backup

# Installez cron et configurez le cron job
RUN apt-get update && apt-get -y install cron && rm -rf /var/lib/apt/lists/*
COPY docker/cronjob /etc/cron.d/prestashop_backup-cron
RUN chmod 0644 /etc/cron.d/prestashop_backup-cron
RUN crontab /etc/cron.d/prestashop_backup-cron
RUN touch /var/log/cron.log



# Commande par défaut pour exécuter cron en arrière-plan
CMD cron && tail -f /var/log/cron.log
