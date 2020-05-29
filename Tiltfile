# Setup custom docke build command for k8s deployments for backend and frontend
custom_build(
    'influx6/juber-app-backend',
    'docker buildx build --platform linux/arm,linux/arm64,linux/amd64 -t $EXPECTED_REF ./backend --push',
    ['./backend'],
    skips_local_docker=True
)

custom_build(
    'influx6/juber-app-frontend',
    'docker buildx build --platform linux/arm,linux/arm64,linux/amd64 -t $EXPECTED_REF ./frontend --push',
    ['./frontend'],
    skips_local_docker=True
)

# Setup k8s yaml files
k8s_yaml('infra/juber-app-namespace.yaml')
k8s_yaml('infra/backend-deployment.yaml')
k8s_yaml('infra/frontend-deployment.yaml')
k8s_yaml('infra/ingress-gateway.yaml')

