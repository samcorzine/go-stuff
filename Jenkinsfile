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
                def root = tool name: 'Go 1.8', type: 'go'
 
                // Export environment variables pointing to the directory where Go was installed
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'go version'
                }
                sh 'export GOPATH=$WORKSPACE/..'
                sh 'export PATH=$GOPATH:$PATH'

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
