import hudson.plugins.git.GitSCM

node{
    stage('checkout'){
        def scm = git credentialsId: 'bitbucket', url: 'https://bitbucket.org/privatesquare/sonarqube-cli.git'
        print getScmType(scm)
    }
}

def getScmType(def scm) {

    if (scm.getClass() == GitSCM) {
        return 'git'
    }
    return 'nothing'
}