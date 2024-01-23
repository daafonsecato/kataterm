# Question
GET http://backend/question
POST https://backend/check_config
GET http://backend/skip_question
GET http://backend/config
POST https://backend/submit_answer
POST https://backend/stage_before_actions
GET https://backend/heart_beat
## Time
GET https://c22c7370220443a4.labs.kodekloud.com/remaining_time
Resp:
{"REMAININGTIME_SECONDS":3590,"REVIEWTIME_SECONDS":600}
## Multiple Choice 4 options
GET http://backend/question
    "text": "How many Endpoints are attached on the 'kubernetes' service?",
    "hint": "Run the command: 'kubectl describe service' and look at the 'Endpoints'.",
    "type": "multiple_choice",
    "options": [
        "0",
        "1",
        "2",
        "3"
    ],
    "answer": "1",
    "total_questions": 11,
    "current_question_number": 6,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null
    ]

{
    "text": "What is the name of the Network Policy?",
    "subtext": "",
    "hint": "Use the command `kubectl get netpol` and identify the name of available network policy. <br>\nNetwork policy is namespace-scoped so list in the `default` namespace.",
    "type": "multiple_choice",
    "image": "/images/kubernetes-ckad-network-policies-2.jpg",
    "options": [
        "network-policy",
        "policy-1",
        "payroll-policy",
        "deny-policy"
    ],
    "answer": "payroll-policy",
    "links": [
        {
            "type": "port",
            "name": "External Service",
            "port": "30080"
        },
        {
            "type": "port",
            "name": "Internal Portal",
            "port": "30082"
        }
    ],
    "total_questions": 10,
    "current_question_number": 2,
    "answer_statuses": [
        "completed",
        "current",
        null,
        null,
        null,
        null,
        null,
        null,
        null,
        null
    ]
}
## Multiple Choice 2 options
GET http://backend/question
{
    "text": "Are you able to accesss the Web App UI?",
    "subtext": "Try to access the Web Application UI using the tab 'simple-webapp-ui' above the terminal.",
    "type": "multiple_choice",
    "options": [
        "YES",
        "NO"
    ],
    "answer": "NO",
    "links": [
        {
            "type": "port",
            "name": "simple-webapp-ui",
            "port": "30080"
        }
    ],
    "total_questions": 11,
    "current_question_number": 9,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null
    ]
}
## Config_test Question
GET http://backend/question
{
    "text": "Create a new service to access the web application using the 'service-definition-1.yaml' file.",
    "subtext": "</br>'Name:' webapp-service</br>'Type:' NodePort</br>'targetPort:' 8080</br>'port:' 8080</br>'nodePort:' 30080</br> 'selector:'</br>&nbsp;&nbsp;'name:' simple-webapp",
    "hint": "Update the given values in the service definition file and create the service.",
    "type": "config_test",
    "tests": [
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "webapp-service",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].metadata.name",
            "state": "present",
            "spec": "Is service created?"
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "NodePort",
            "err_message": "Service type is not NodePort",
            "jmespath": "items[?metadata.name=='webapp-service'] | [?spec.type=='NodePort'] | [].spec.type",
            "state": "present",
            "spec": "Is it type nodeport?"
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 8080,
            "err_message": "Port not set to 8080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?port=='8080'] | [].port",
            "state": "present",
            "spec": "Is the port set?"
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 8080,
            "err_message": "TargetPort not set to 8080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?targetPort=='8080'] | [].targetPort",
            "state": "present",
            "spec": "Is the target port set?"
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 30080,
            "err_message": "NodePort not set to 30080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?nodePort=='30080'] | [].nodePort",
            "state": "present",
            "spec": "Is the node port set?"
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "simple-webapp",
            "err_message": "Selector not set to simple-webapp",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.selector.name",
            "state": "present",
            "spec": "Is the selector set?"
        }
    ],
    "solution": "",
    "links": [
        {
            "type": "port",
            "name": "simple-webapp-ui",
            "port": "30080"
        }
    ],
    "solution_file": "/course/certified-kubernetes-administrator/services-stable/solutions/solution-ques10.md",
    "total_questions": 11,
    "current_question_number": 10,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null
    ],
    "solution_md": "Update the '/root/service-definition-1.yaml' file as follows:\n\n'yaml\n---\napiVersion: v1\nkind: Service\nmetadata:\n  name: webapp-service\n  namespace: default\nspec:\n  ports:\n  - nodePort: 30080\n    port: 8080\n    targetPort: 8080\n  selector:\n    name: simple-webapp\n  type: NodePort\n'\n<br>\n\nRun the following command to create a 'webapp-service' service as follows: -\n'sh\nkubectl apply -f /root/service-definition-1.yaml\n'\n\n"
}

## Info Question
{
    "text": "Access the web application using the tab 'simple-webapp-ui' above the terminal window.",
    "type": "info",
    "options": [
        "Ok"
    ],
    "answer": "Ok",
    "links": [
        {
            "type": "port",
            "name": "simple-webapp-ui",
            "port": "30080"
        }
    ],
    "total_questions": 11,
    "current_question_number": 11,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current"
    ]
}
# Actions
## Check config_test question (not completed)
POST https://backend/check_config
Payload {configNumber: {isTrusted: true, _vts: 1703685500085}}
{
    "message": "Tasks not completed!",
    "additional_details": [
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "webapp-service",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].metadata.name",
            "state": "present",
            "spec": "Is service created?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "NodePort",
            "err_message": "Service type is not NodePort",
            "jmespath": "items[?metadata.name=='webapp-service'] | [?spec.type=='NodePort'] | [].spec.type",
            "state": "present",
            "spec": "Is it type nodeport?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 8080,
            "err_message": "Port not set to 8080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?port=='8080'] | [].port",
            "state": "present",
            "spec": "Is the port set?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 8080,
            "err_message": "TargetPort not set to 8080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?targetPort=='8080'] | [].targetPort",
            "state": "present",
            "spec": "Is the target port set?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 30080,
            "err_message": "NodePort not set to 30080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?nodePort=='30080'] | [].nodePort",
            "state": "present",
            "spec": "Is the node port set?",
            "test_result": false
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "simple-webapp",
            "err_message": "Selector not set to simple-webapp",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.selector.name",
            "state": "present",
            "spec": "Is the selector set?",
            "test_result": true
        }
    ]
}
## Check config_test question (completed)
POST https://backend/check_config
PAYLOAD {configNumber: {isTrusted: true, _vts: 1703685413763}}
{
    "message": "Tasks completed.",
    "additional_details": [
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "webapp-service",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].metadata.name",
            "state": "present",
            "spec": "Is service created?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "NodePort",
            "err_message": "Service type is not NodePort",
            "jmespath": "items[?metadata.name=='webapp-service'] | [?spec.type=='NodePort'] | [].spec.type",
            "state": "present",
            "spec": "Is it type nodeport?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 8080,
            "err_message": "Port not set to 8080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?port=='8080'] | [].port",
            "state": "present",
            "spec": "Is the port set?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 8080,
            "err_message": "TargetPort not set to 8080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?targetPort=='8080'] | [].targetPort",
            "state": "present",
            "spec": "Is the target port set?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": 30080,
            "err_message": "NodePort not set to 30080",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.ports[] | [?nodePort=='30080'] | [].nodePort",
            "state": "present",
            "spec": "Is the node port set?",
            "test_result": true
        },
        {
            "type": "kubernetes",
            "object": "Service",
            "name": "simple-webapp",
            "err_message": "Selector not set to simple-webapp",
            "jmespath": "items[?metadata.name=='webapp-service'] | [].spec.selector.name",
            "state": "present",
            "spec": "Is the selector set?",
            "test_result": true
        }
    ]
}

## Skip question
GET http://backend/skip_question
    {
        "total_questions": 11,
        "answer_statuses": [
            "completed",
            "completed",
            null,
            null,
            null,
            null,
            null,
            null,
            null,
            null,
            null
        ],
        "new_question_number": 3
}
## Config after some questions (Purpose TBD)
GET http://backend/config
Sometimes this answers an empty response
{
    "environment": "Prod",
    "quiz_mode": "standard",
    "review_question": false,
    "exam_evaluation_mode": "SHOW_RESULT_AND_SCORE",
    "under_review": false,
    "run_tests_one_at_a_time": false,
    "show_cmd_output_on_error": false,
    "layout": "horizontal",
    "start_from": 8,
    "panel_first": false
}
## Heart beat (Each 20 seconds)
GET https://backend/heart_beat
200 or 500

## Submit Answer
POST https://7d78542a085e4798.labs.kodekloud.com/submit_answer
Payload: {answer: "6443"}
200 Accepted
or
400 "Incorrect! Try Again."
## In case of stage before action step
GET http://backend/question
{
    "text": "How many Deployments exist on the system now?",
    "hint": "Run the command: 'kubectl get deployment' and count the number of pods.",
    "subtext": "In the current(default) namespace",
    "type": "multiple_choice",
    "options": [
        "0",
        "1",
        "2",
        "3",
        "4"
    ],
    "answer": "1",
    "before": [
        {
            "type": "kubernetes",
            "action": "create",
            "config": {
                "apiVersion": "apps/v1",
                "kind": "Deployment",
                "metadata": {
                    "name": "simple-webapp-deployment"
                },
                "spec": {
                    "replicas": 4,
                    "selector": {
                        "matchLabels": {
                            "name": "simple-webapp"
                        }
                    },
                    "template": {
                        "metadata": {
                            "labels": {
                                "name": "simple-webapp"
                            }
                        },
                        "spec": {
                            "containers": [
                                {
                                    "name": "simple-webapp",
                                    "image": "kodekloud/simple-webapp:red",
                                    "ports": [
                                        {
                                            "containerPort": 8080
                                        }
                                    ]
                                }
                            ]
                        }
                    }
                }
            }
        }
    ],
    "links": [
        {
            "type": "port",
            "name": "simple-webapp-ui",
            "port": "30080"
        }
    ],
    "total_questions": 11,
    "current_question_number": 7,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null
    ],
    "hasBeforeAction": true
}
### Post stage before actions
POST https://backend/stage_before_actions
PAYLOAD { question_number: 7 }
{
    "text": "How many Deployments exist on the system now?",
    "hint": "Run the command: 'kubectl get deployment' and count the number of pods.",
    "subtext": "In the current(default) namespace",
    "type": "multiple_choice",
    "options": [
        "0",
        "1",
        "2",
        "3",
        "4"
    ],
    "answer": "1",
    "before": [
        {
            "type": "kubernetes",
            "action": "create",
            "config": {
                "apiVersion": "apps/v1",
                "kind": "Deployment",
                "metadata": {
                    "name": "simple-webapp-deployment"
                },
                "spec": {
                    "replicas": 4,
                    "selector": {
                        "matchLabels": {
                            "name": "simple-webapp"
                        }
                    },
                    "template": {
                        "metadata": {
                            "labels": {
                                "name": "simple-webapp"
                            }
                        },
                        "spec": {
                            "containers": [
                                {
                                    "name": "simple-webapp",
                                    "image": "kodekloud/simple-webapp:red",
                                    "ports": [
                                        {
                                            "containerPort": 8080
                                        }
                                    ]
                                }
                            ]
                        }
                    }
                }
            }
        }
    ],
    "links": [
        {
            "type": "port",
            "name": "simple-webapp-ui",
            "port": "30080"
        }
    ],
    "hasBeforeAction": false,
    "total_questions": 11,
    "current_question_number": 7,
    "answer_statuses": [
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "completed",
        "current",
        null,
        null,
        null,
        null
    ]
}


# Manage session
## Successfully ended the session
GET https://manager.labs.kodekloud.com/session/end?lab_session_id=7d78542a085e4798&fpReqID=1703682898517.vUuVMW&type=end_of_lab
200
## Successfully created the session
GET https://manager.labs.kodekloud.com/?userid=8128david%40gmail.com&token=-Ws9upTnBhKcvIyC-qnx7g&environmentid=k3-single-node-ttyd-stable&labscenario=certified-kubernetes-administrator%2Fnamespaces-stable&coursename=CKA%20Certification%20Course%20%E2%80%93%20Certified%20Kubernetes%20Administrator&theme=light
200

## In case of many labs sessions in short period of time
GET https://manager.labs.kodekloud.com/session?userid=8128david%40gmail.com&token=PPyDknFOOGJf2Rh5uHSyHg&environmentid=docker-single-node-ttyd-19-03&coursename=Docker+Training+Course+for+the+Absolute+Beginner&theme=light&LAB_SCENARIO=docker-for-beginners%2Fdocker_images&fpReqID=1703682898517.vUuVMW
429
{
    "error": "You have made too many lab requests in a short period of time. Please wait for approximately 37 seconds before trying again."
}