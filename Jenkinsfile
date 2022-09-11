pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo "build"
                sh 'printenv'
            }
        }
        stage('Test') {
            steps {
            sh 'pwd&&ls -alh'
                echo "test"
            }
        }
        stage('Test2') {
            steps {
            sh 'pwd&&ls -alh'
                echo "test"
            }
        }
        stage('Deploy') {
            steps {
                echo "deploy"
            }
        }
    }
}