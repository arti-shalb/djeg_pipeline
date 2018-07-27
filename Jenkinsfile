pipeline {
    agent any
    stages {
        stage('checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'cp -R docker/* .'
                sh "go build main.go"
            }
        }
        stage('Test') {
            steps {
                sh './main tests'
            }
        }
    }

    post {
        always {
            junit 'result.txt'
        }
    }
}
