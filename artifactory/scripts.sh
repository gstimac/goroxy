docker run --name artifactory -v ~/projects/goxy/artifactory/var/:/var/opt/jfrog/artifactory -d -p 8081:8081 -p 8082:8082 releases-docker.jfrog.io/jfrog/artifactory-oss:latest
helm repo add jfrog https://charts.jfrog.io
helm repo update
helm upgrade --install artifactory-oss --set artifactory.postgresql.postgresqlPassword=PopKopaProkop! --namespace artifactory-oss jfrog/artifactory-oss
