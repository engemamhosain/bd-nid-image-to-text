
pipeline {
    agent any
    tools {go "golang"}

  

    stages {
        stage('Build') {                
            steps {    
                sh 'cd tl_ops && go build'
            }            
        }


      
      stage('Deploy to dev') {
            when {
                expression {
                    script {
                        def ENV_MODE =System.getProperty('ENV_MODE')
                        return ENV_MODE == 'dev'
                    }
                }
            }
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'box_key_id', keyFileVariable: 'SSH_KEY')]) {
                    script {
                    
                        sshagent(['box_key_id']) {
                            sh "ssh -tt -i ${SSH_KEY} -p 22 -o StrictHostKeyChecking=no customer@domain-name 'sudo /bin/systemctl stop tl-ops.service' "
                            sh "cd mlkit && scp -i $SSH_KEY -P 22 -r conf mlkit customer@domain-name:/home/Go/mlkit"  
                          try {
                              
                            sh """
                                ssh -tt -i ${SSH_KEY} -p 22 -o StrictHostKeyChecking=no customer@domain-name \\
                                'sudo /bin/systemctl restart mlkit.service && 
                                 sudo /bin/systemctl daemon-reload &&
                                 sleep 5 && systemctl status mlkit.service && 
                                 timeout --foreground 10s  tail -f  /var/log/syslog'  
                               """  
                                
                        } catch (err) {
                            // Handle any errors here if necessary
                            echo "exit $err"
                        }
                       
                        
                        }
                    }
                }
            }
      }

      stage('Deploy to production') {
            when {
                expression {
                    script {
                        def ENV_MODE =System.getProperty('ENV_MODE')
                        return ENV_MODE == 'production'
                    }
                }
            }
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'prod-key', keyFileVariable: 'SSH_KEY')]) {
                    script {
                    
                        sshagent(['prod-key']) {
                            sh "ssh -tt -i ${SSH_KEY} -p 22 -o StrictHostKeyChecking=no user@domain-name: 'sudo systemctl stop mlkit.service ' "
                             sh "cd mlkit && scp -i $SSH_KEY -P 22 -r conf mlkit ubuntu@domain-name:/home/Go/mlkit"  
                          try {
                              
                            sh """
                                ssh -tt -i ${SSH_KEY} -p 36872 -o StrictHostKeyChecking=no user@domain-name \\
                                'sudo /bin/systemctl restart mlkit.service && 
                                 sudo /bin/systemctl daemon-reload &&
                                 sleep 10 && systemctl status mlkit.service && 
                                 timeout --foreground 10s  tail -f  /var/log/syslog'  
                               """  
                                
                        } catch (err) {
                            // Handle any errors here if necessary
                            echo "exit $err"
                        }
                       
                        
                        }
                    }
                }
            }
      }



    }
}
