pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.10-alpine3.7'
                }
            }
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            agent {
                docker {
                    image 'golang:1.10-alpine3.7'
                }
            }
            steps {
                sh 'go test'
            }
        }
    }
}