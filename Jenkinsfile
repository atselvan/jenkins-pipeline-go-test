@Library('jenkins-go-pipeline') _

import com.privatesquare.pipeline.utils.Git

def git = new Git(this)

node{
    stage('checkout'){
        checkout scm
        print getScmType(scm)
        print git.getUrl()
        print git.createNextTag("1.0", "builder")
    }
}