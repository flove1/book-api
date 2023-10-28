pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build and Run Docker Compose') {
            steps {
                script {
                    // Set Docker environment (if needed)
                    docker.withServer('my-docker-host') {
                        // Pull the latest images and start services
                        sh 'docker-compose -f docker-compose.yml pull'
                        sh 'docker-compose -f docker-compose.yml up -d'
                    }
                }
            }
        }
    }

    post {
        always {
            // Clean up (stop and remove containers)
            script {
                docker.withServer('my-docker-host') {
                    sh 'docker-compose -f docker-compose.yml down'
                }
            }
        }
    }
}
