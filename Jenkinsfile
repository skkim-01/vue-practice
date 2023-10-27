pipeline {
  agent {
    kubernetes {
      yamlFile 'agentPod.yaml'
    }
  }
  stages {
    stage('test') {
      steps {
        container('docker-private') {
          sh 'ls -al'
          sh 'docker build -t skkim01/testflights:v1.0 .'
          sh 'docker images'
          sh 'docker push skkim01/testflights:v1.0'
        }
      }
    }
  }
}