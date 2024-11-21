pipeline {
    agent any
    tools {
        go '1.23.0' // Asegúrate de que Go esté correctamente instalado en Jenkins
    }
    environment {
        DATABASE_URL = credentials('DATABASE_URL')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                // Limpiar y preparar el proyecto
                sh '''
                go clean
                go mod tidy
                go build -o app .
                '''
            }
        }

        stage('Deploy') {
            steps {
                // Ejecutar el binario compilado para Linux
                sh './app'
            }
        }
    }

    post {
        success {
            echo 'Pipeline success'
        }
        failure {
            echo 'something failed. Check the logs'
        }
    }
}
