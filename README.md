# Mailer Operator

## Contents
<details open>
<summary>Show/Hide</summary>
<br>

- [Task](#Task)
- [Achieved Goals](#Achieved_Goals)
- [What Can be Improved](#What_Can_Be_Improved)
- [Important Files](#Important_Files)
- [Files Directory Tree](#Files_Directory_Tree)
- [Instructions to Deploy Operator](#Instructions_to_Deploy_Operator)
</details>

<a name="Task"></a>
## Task
<details open>
<summary>Show/Hide</summary>
<br>

Create a Kubernetes operator that manages custom resources for configuring email sending and sending of emails via a transactional email provider like MailerSend. The operator should work cross namespace, and demonstrate sending from multiple providers, please use MailerSend and Mailgun for this task. This current operator allows you to define email configurations and emails as custom resources in Kubernetes and handles the sending of emails through the MailerSend API.

</details>

<a name="Achieved_Goals"></a>
## Achieved Goals
<details open>
<summary>Show/Hide</summary>
<br>

- Implemented a Kubernetes operator to manage email sending.
- Defined custom resources EmailSenderConfig and Email to configure and send emails.
- Developed a controller to handle email sending using the MailerSend API.
- Created Kubernetes manifests for deploying the operator.
- Tested the email sending functionality with MailerSend.
</details>

<a name="What_Can_Be_Improved"></a>
## What Can Be Improved
<details open>
<summary>Show/Hide</summary>
<br>

- Add support for multiple email providers, such as Mailgun, in addition to MailerSend.
- Implement more comprehensive error handling and retry mechanisms.
- Enhance logging and monitoring for better visibility and debugging.
- Add unit tests and integration tests to ensure the robustness of the operator.
- Optimize the Docker image for smaller size and faster deployment.
</details>

<a name="Important_Files"></a>
## Important Files
<details>
<summary>Show/Hide</summary>
<br>

- <strong>main.go</strong>: 
<br>Entry point for the manager that starts the controller.

- <strong>api/v1/email_types.go</strong>: 
<br>Contains the definitions for the custom resources EmailSenderConfig and Email.

- <strong>controllers/email_controller.go</strong>: 
<br>Contains the logic for reconciling Email resources and sending emails through MailerSend.

- <strong>config/manager/manager.yaml</strong>: 
<br>Kubernetes manifest for deploying the controller manager.

- <strong>config/crd/bases/email.mailerlitetask.com_emails.yaml</strong>: 
<br>Custom Resource Definition (CRD) for the Email resource.

- <strong>config/crd/bases/email.mailerlitetask.com_emailsenderconfigs.yaml</strong>: 
<br>Custom Resource Definition (CRD) for the EmailSenderConfig resource.

- <strong>config/rbac/role.yaml</strong>, <strong>config/rbac/role_binding.yaml</strong>, <strong>config/rbac/service_account.yaml</strong>: 
<br>RBAC configuration for the operator.

- <strong>config/test/mailersend_emailsenderconfig.yaml</strong>: 
<br>Sample EmailSenderConfig resource for MailerSend.

- <strong>config/test/mailersend_email.yaml</strong>: 
<br>Sample Email resource for testing email sending.

</details>

<a name="Files_Directory_Tree"></a>
## Files Directory Tree
<details>
<summary>Show/Hide</summary>
<br>

```plaintext
.
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── api
│   └── v1
│       ├── email_types.go
│       ├── emailsenderconfig_types.go
│       ├── groupversion_info.go
│       └── zz_generated.deepcopy.go
├── bin
│   └── manager
├── config
│   ├── controller_manager_config.yaml
│   ├── crd
│   │   ├── bases
│   │   │   ├── email.mailerlitetask.com_emails.yaml
│   │   │   ├── email.mailerlitetask.com_emailsenderconfigs.yaml
│   │   │   └── kustomization.yaml
│   │   ├── controller_manager_config_crd.yaml
│   │   └── kustomization.yaml
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_config_patch.yaml
│   │   └── manager_image_patch.yaml
│   ├── kustomization.yaml
│   ├── manager
│   │   ├── controller_manager_config.yaml
│   │   ├── kustomization.yaml
│   │   ├── manager.yaml
│   │   └── manager_config.yaml
│   ├── rbac
│   │   ├── auth_proxy_client_clusterrole.yaml
│   │   ├── auth_proxy_role.yaml
│   │   ├── auth_proxy_role_binding.yaml
│   │   ├── auth_proxy_service.yaml
│   │   ├── email_editor_role.yaml
│   │   ├── email_viewer_role.yaml
│   │   ├── emailsenderconfig_editor_role.yaml
│   │   ├── emailsenderconfig_viewer_role.yaml
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── mailer-operator-cluster-role-binding.yaml
│   │   ├── mailer-operator-cluster-role.yaml
│   │   ├── role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   ├── samples
│   │   ├── email_v1_email.yaml
│   │   ├── email_v1_emailsenderconfig.yaml
│   │   └── kustomization.yaml
│   └── test
│       ├── mailersend-test-pod.yaml
│       ├── mailersend_email.yaml
│       ├── mailersend_emailsenderconfig.yaml
│       ├── mailgun_email.yaml
│       ├── mailgun_emailsenderconfig.yaml
│       ├── new_mailersend_email.yaml
│       └── send_email.sh
├── controllers
│   ├── email_controller.go
│   ├── emailsenderconfig_controller.go
│   └── suite_test.go
├── cover.out
├── fetch-logs.sh
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── mailersend-secret-token.yaml
├── main.go
└── manager
```

</details>

<a name="Instructions_to_Deploy_Operator"></a>
## Instructions to Deploy Operator
<details open>
<summary>Show/Hide</summary>
<br>

### Prerequisites
- Docker
- Minikube
- kubectl


#### Windows
<details>
<summary>Show/Hide</summary>

- Install Docker Desktop from [here](https://docs.docker.com/desktop/install/windows-install/).
- Install Minikube from [here](https://minikube.sigs.k8s.io/docs/start/?arch=%2Fwindows%2Fx86-64%2Fstable%2F.exe+download).
- Install kubectl from [here](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/).

</details>

#### macOS
<details>
<summary>Show/Hide</summary>

- Install Docker Desktop from [here](https://docs.docker.com/desktop/install/mac-install/).
- Install Minikube from [here](https://minikube.sigs.k8s.io/docs/start/?arch=%2Fmacos%2Fx86-64%2Fstable%2Fbinary+download).
- Install kubectl from [here](https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/).

</details>

#### Linux
<details>
<summary>Show/Hide</summary>

- Install Docker from [here](https://docs.docker.com/desktop/install/linux-install/).
- Install Minikube from [here](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download).
- Install kubectl from [here](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/).

</details>

### Step 1: Clone the Repository

```
git clone https://github.com/awesomeahi95/email-operator.git
cd email-operator
```

### Step 2: Build Docker Image

```
docker build -t ahilan95/email-operator:latest .
```

### Step 3: Setup Minikube and Kubernetes Resources
#### Start Minikube:

```
minikube start
```

#### - Create the namespace:

```
kubectl create namespace mailer-operator-system
```

#### Apply the CRDs:

```
kubectl apply -f config/crd/bases/email.mailerlitetask.com_emails.yaml
kubectl apply -f config/crd/bases/email.mailerlitetask.com_emailsenderconfigs.yaml
```

#### Apply the RBAC configuration:

```
kubectl apply -f config/rbac/service_account.yaml
kubectl apply -f config/rbac/role.yaml
kubectl apply -f config/rbac/role_binding.yaml
```

#### Apply the deployment:

```
kubectl apply -f config/manager/manager.yaml
```

### Step 4: Create Secrets and Resources
#### Create Secret for MailerSend API Token

```
kubectl create secret generic mailersend-secret-token \
  --from-literal=api-token='your-api-token' \
  --from-literal=from-email='your-from-email' \
  -n mailer-operator-system
```

#### Change Recipient Email to Preferred Email Address

config/test/mailersend_email.yaml
```
spec:
  recipientEmail: your-preferred-email@example.com
```


#### Apply the test EmailSenderConfig and Email resources:

```
kubectl apply -f config/test/mailersend_emailsenderconfig.yaml
kubectl apply -f config/test/mailersend_email.yaml
```

### Step 5: Verify the Deployment
#### Check the logs of the controller manager to ensure emails are being sent:

```
kubectl logs -n mailer-operator-system -l control-plane=controller-manager -f
```

#### Verify that the emails are send successfully to you given email address

</details>