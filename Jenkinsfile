@Library('jenkins-go-pipeline') _

node{
    stage('checkout'){
        checkout scm
        print getScmType(scm)
        getEnvironment()
    }
}