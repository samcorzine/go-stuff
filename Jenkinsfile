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
                sh 'export GOPATH=$WORKSPACE/..'
                sh 'export PATH=$GOPATH:$PATH'

                sh 'cd $GOPATH/src/github.com/coredns/coredns && make coredns'
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
