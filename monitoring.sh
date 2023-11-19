#!/bin/bash

sudo apt-get install -y apt-transport-https software-properties-common wget
sudo mkdir -p /etc/apt/keyrings/
wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null
echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list

sudo apt-get update
sudo apt-get install -y prometheus
sudo apt-get install -y prometheus-blackbox-exporter
sudo apt-get install -y grafana


cat <<EOF | sudo tee /etc/prometheus/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s
          
scrape_configs:
  - job_name: 'blackbox'
    metrics_path: /probe
    params:
      module: [http_2xx]
    static_configs:
      - targets: ['http://${app_ip}:8080/api/v1/healthcheck']

    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 127.0.0.1:9115

  - job_name: 'node_exporter'
    static_configs:
      - targets: ['${app_ip}:9100']

  - job_name: 'blog-api'
    static_configs:
      - targets: ['${app_ip}:8080']
EOF

cat <<EOF | sudo tee /etc/prometheus/blackbox.yml
modules:
  http_2xx:
    prober: http
    timeout: 5s
    http:
      valid_http_versions: [HTTP/1.1, HTTP/2]
      valid_status_codes: []
      method: GET
EOF

cat <<EOF | sudo tee /lib/systemd/system/blackbox.service
[Unit]
Description=Blackbox Exporter Service
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
User=blackbox
Group=blackbox
ExecStart=/usr/bin/prometheus-blackbox-exporter \
  --config.file=/etc/prometheus/blackbox.yml \
  --web.listen-address=":9115"

Restart=always

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload

sudo systemctl enable --now blackbox
sudo systemctl enable --now prometheus
sudo systemctl enable --now grafana-server