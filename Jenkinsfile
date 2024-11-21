pipeline {
    agent any
    tools {
        go '1.23.0'
    }
    environment  {
        DATABASE_URL = credentials('DATABASE_URL')
    }

    stages {
        stage('Checkout') {
            steps{
                checkout scm
            }
        }

        stage('Build') {
            steps {
                // Compilar el proyecto Go
                sh '''
                go clean
                go mod tidy
                go build -o app.exe .
                '''
            }
        }

        stage('Deploy') {
            steps {
                // Desplegar o ejecutar la aplicación (opcional, para pruebas locales)
                sh './app.exe'
            }
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