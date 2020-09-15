#!/usr/bin/groovy
@Library('utilities@2020-06-29_117')
import com.bosch.jenkins.Utilities
@Library('utilities@2020-06-29_117')
import com.bosch.jenkins.Utilities
import groovy.transform.Field

@Field
def utils = new Utilities(this)

String GIT_REPO = 'products.bosch-si.com/stash/scm/bisa/ory-fosite.git'

try {
    node() {
        timestamps {
            ansiColor('xterm') {
                utils.addToolToPath(['go'])

                git poll: true, url: "https://${GIT_REPO}".toString(), credentialsId: "GIT_USER", branch: "token_exchange_build"

                stage('Test') {
                    echo 'Running tests'
                    sh 'go test -v ./...'
                }
            }
        }
    }
} catch (e) {
    print "Pipeline failed: " + e
    utils.sendFailureMail("suite-portal@bosch.io")
    /* Must re-throw exception to propagate error */
    throw e
}
