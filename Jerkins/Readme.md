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
![Screenshot (385)](https://user-images.githubusercontent.com/76167753/235894112-c5e2d499-7260-4e05-b1c4-4eea226eee7a.png)

## add Credentials
![Screenshot (386)](https://user-images.githubusercontent.com/76167753/235894184-f0f05a0c-58ef-44f5-8f9d-50dee23480d0.png)

--> Manage Jerkins
--> Manage Credentials
--> (global) --> Domains
--> Add Credentials
--> name and password
--> ok

## Add Nodes

1. Manage Jenkins
2. Manage Nodes and Clouds
![Screenshot (387)](https://user-images.githubusercontent.com/76167753/235894292-e3033963-da1e-43c0-9786-ab21dfab543a.png)

4. New Node
![Screenshot (388)](https://user-images.githubusercontent.com/76167753/235894376-646029d6-9197-4527-aa65-8ae3813d407f.png)

6. Node name
7. Permanent Agent , click ok
8. enter Remote root directory
9. enter Labels
10. Launch Agents via SSH
11. enter Host and select the credentials you created.

![Screenshot (389)](https://user-images.githubusercontent.com/76167753/235894457-969843fe-a4bc-4d11-a804-7a5b9a8a7093.png)
![Screenshot (390)](https://user-images.githubusercontent.com/76167753/235894488-df727f6b-66bf-49bd-8dc2-2a778b5ddc69.png)
![Screenshot (391)](https://user-images.githubusercontent.com/76167753/235894526-1dcdc91d-d5d8-4935-aa6d-03914bf21df9.png)

13. Host Key Verification Strategy, select Non verifying verification strategy.
14. save
![Screenshot (392)](https://user-images.githubusercontent.com/76167753/235894568-d6c178f7-beb7-42f1-9b6e-c83bb8bda822.png)

## Create and Configure the job:

1. click New Item
2. enter job name
3.  click Pipeline and ok
4.  ![Screenshot (393)](https://user-images.githubusercontent.com/76167753/235894647-95b2265d-1760-4589-ab69-da572de85412.png)

5. under pineline script enter the xml pipeline file

![Screenshot (394)](https://user-images.githubusercontent.com/76167753/235894677-5cdcd97b-2f5c-444f-8e55-24282b5a23a6.png)
![Screenshot (395)](https://user-images.githubusercontent.com/76167753/235894696-2d93f0cf-2577-4587-9d54-73f6ff719393.png)

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
![Screenshot (396)](https://user-images.githubusercontent.com/76167753/235894722-1f9bbe80-fd36-431d-a174-93de181bedc7.png)
![Screenshot (397)](https://user-images.githubusercontent.com/76167753/235894749-95da9f91-49ac-4193-8d12-b2514a54728a.png)

### Automatically creating pipelines once one of the ci/cd fails like docker pods
![Screenshot (398)](https://user-images.githubusercontent.com/76167753/235894789-33f643c8-ba56-4415-b810-b47ca180f368.png)
![Screenshot (399)](https://user-images.githubusercontent.com/76167753/235894816-d17a1e90-a02d-4d0f-9959-62eb34b86b1a.png)


### COnsole output of the pipeline
![Screenshot (400)](https://user-images.githubusercontent.com/76167753/235894842-9703dfb5-8a88-415c-8d86-8c79133f4f63.png)
