pipeline {
    agent any

    environment  {
        DATABASE_URL = credentials('DATABASE_URL')
    }

    stages {
        stage('Checkout') {
            checkout scm
        }

        stage('Build') {
            // Compilar el proyecto Go
                sh '''
                go clean
                go mod tidy
                go build -o app .
                '''
        }

        stage('Deploy') {
            // Desplegar o ejecutar la aplicación (opcional, para pruebas locales)
                sh './app'
        }
    }

    post {
        success {
            echo 'Pipeline success'
        }
        failure {
            echo 'somenthing failed. Check the logs'
        }
    }
}