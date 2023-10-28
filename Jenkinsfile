pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build and Deploy') {
            steps {
                script {
                    // Define the location of your Docker Compose file
                    def composeFile = 'docker-compose.yml'

                    // Make sure Docker is available on the Jenkins agent
                    sh 'docker --version'

                    // Run Docker Compose
                    sh "docker-compose -f ${composeFile} up -d"
                }
            }
        }
    }

    post {
        success {
            // Clean up (stop and remove containers)
            script {
                docker.withServer('my-docker-host') {
                    // Define the location of your Docker Compose file
                    def composeFile = 'docker-compose.yml'

                    // Run Docker Compose
                    sh "docker-compose -f ${composeFile} down"
                }
            }
        }
    }
}
