@Library('jenkins-go-pipeline') _

import com.privatesquare.pipeline.utils.Git

def git = new Git()

node{
    stage('checkout'){
        checkout scm
        print getScmType(scm)
        git.getUrl()
    }
}