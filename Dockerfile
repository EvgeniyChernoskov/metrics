FROM prom/prometheus
COPY ./config/prometheus.yml /etc/prometheus
EXPOSE 9090
