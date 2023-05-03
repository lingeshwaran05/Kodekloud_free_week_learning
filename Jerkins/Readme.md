### TASKS for creating and deploying multi-stage pipelines
1. create pipeline
# git pull
2. that contains workspace with sh command ```git pull```
# testing
3. test Dev stage contains sh command ```go test ./...```
# deploying
4. deploythe code on server with sh command ```go run main.go &```
# production build
5. Sh command ```git pull`` the production environment
# testing
6. test Dev stage contains sh command ```go test ./...```
# deploying
7. deploythe code on server with sh command ```go run main.go &`






## Install plugins
--> Manage Jerkins
--> Manage plugins
--> Serach SSH Build Agents
--> Serach Pipeline
--> Check those and click "Restart Jenkins when installation is complete and no jobs are running"

## add Credentials

--> Manage Jerkins
--> Manage Credentials
--> (global) --> Domains
--> Add Credentials
--> name and password
--> ok

## Add Nodes

1. Manage Jenkins
2. Manage Nodes and Clouds
3. New Node
4. Node name
5. Permanent Agent , click ok
6. enter Remote root directory
7. enter Labels
8. Launch Agents via SSH
9. enter Host and select the credentials you created.
10. Host Key Verification Strategy, select Non verifying verification strategy.
11. save

## Create and Configure the job:

1. click New Item
2. enter job name
3.  click Pipeline and ok
4. under pineline script enter the xml pipeline file
```
pipeline {
    agent none
    stages {
        //DEV
        stage('Build Dev') {
            agent {
              label {
                label 'dev'
                customWorkspace "/opt/go-app"
              }
            }
            steps {
                sh 'git pull'
            }
        }
        stage('Test Dev') {
            agent {
              label {
                label 'dev'
                customWorkspace "/opt/go-app"
              }
            }
            steps {
                sh 'go test ./...'
            }
        }
        stage('Deploy Dev') {
            agent {
              label {
                label 'dev'
                customWorkspace "/opt/go-app"
              }
            }
            steps {
              script {
                withEnv ( ['JENKINS_NODE_COOKIE=do_not_kill'] ) {
                  sh 'go run main.go &'
                  }
                }
            }
        }
        //PROD
        stage('Build Prod') {
            agent {
              label {
                label 'prod'
                customWorkspace "/opt/go-app"
              }
            }
            steps {
                sh 'git pull'
            }
        }
        stage('Test Prod') {
            agent {
              label {
                label 'prod'
                customWorkspace "/opt/go-app"
              }
            }
            steps {
                sh 'go test ./...'
            }
        }
        stage('Deploy Prod') {
            agent {
              label {
                label 'prod'
                customWorkspace "/opt/go-app"
              }
            }
            steps {
              script {
                withEnv ( ['JENKINS_NODE_COOKIE=do_not_kill'] ) {
                  sh 'go run main.go &'
                  }
                }
            }
        }
    }
}
```
5. save and build now

### Automatically creating pipelines once one of the ci/cd fails like docker pods


### COnsole output of the pipeline