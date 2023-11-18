pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build and Test Go Project') {
            steps {
                script {
                    sh 'go version'
                    sh 'go get -v ./...'
                    sh 'go test -v ./...'
                }
            }
        }

        stage('Test Docker Compose Configuration') {
            steps {
                script {
                    sh 'docker-compose --version'
                    sh 'docker-compose config -q'
                }
            }
        }

        stage('Deploy with Terraform') {
            steps {
                // Deploy infrastructure with Terraform
                script {
                    sh 'terraform init'
                    sh 'terraform apply -auto-approve'
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed. Check the logs for details.'
        }
    }
}
