# Question

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
