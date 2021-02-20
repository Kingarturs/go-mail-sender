pipeline {
  agent {
    node {
      label 'master'
    }

  }
  stages {
    stage('Helloworld') {
      steps {
        echo 'Hola mundo!'
        bat 'ls'
      }
    }

    stage('Stage 2') {
      steps {
        echo 'Realizando etapa 2'
      }
    }

  }
}