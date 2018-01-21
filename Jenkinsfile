import hudson.plugins.git.GitSCM

node{
    stage('checkout'){
        checkout scm
        print getScmType(scm)
    }
}

def getScmType(def scm) {

    if (scm.getClass() == GitSCM) {
        return 'git'
    }
    return 'nothing'
}