import hudson.plugins.git.GitSCM

node{
    stage('checkout'){
        checkout scm
        print getScmType(scm)
        getEnvironment()
    }
}

def getEnvironment() {
    // Solution from Cloudbees Support: https://support.cloudbees.com/hc/en-us/articles/230610987-Pipeline-How-to-print-out-env-variables-available-in-a-build
    echo "buildVariables=${currentBuild.buildVariables}"
    sh 'env > env.txt'
    for (String i : readFile('env.txt').split("\r?\n")) {
        println i
    }
}

def getScmType(def scm) {

    if (scm.getClass() == GitSCM) {
        return 'git'
    } else if (scm.getClass() == SubversionSCM) {
        return 'svn'
    }
    return 'null'
}