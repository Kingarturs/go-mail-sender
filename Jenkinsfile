pipeline {
  agent {
    node {
      label 'node2'
    }

  }
  stages {
    stage('Helloworld') {
      steps {
        echo 'Hola mundo!'
      }
    }

    stage('Stage 2') {
      steps {
        echo 'Realizando etapa 2'
      }
    }

  }
}