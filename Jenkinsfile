@Library('jenkins-go-pipeline') _

import com.privatesquare.pipeline.utils.Git

Git git

node{
    stage('checkout'){
        checkout scm
        print getScmType(scm)
        git.getUrl()
    }
}