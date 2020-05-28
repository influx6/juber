k8s_yaml('infra/juber-app-namespace.yaml')
k8s_yaml('infra/backend-deployment.yaml')
k8s_yaml('infra/frontend-deployment.yaml')

# Setup docker-files to be built and used in generating yaml
docker_build('influx6/juber-app-backend', './backend')
docker_build('influx6/juber-app-frontend', './frontend')
