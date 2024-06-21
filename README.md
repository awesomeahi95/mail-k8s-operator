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
- [Detailed Step-by-Step Instructions](#Detailed_Step-by-Step_Instructions)
</details>

## Task
<details open>
<a name="Task"></a>
<summary>Show/Hide</summary>
<br>

The purpose of this task is to create a Kubernetes operator to manage email sending using MailerSend and Mailgun. This operator allows you to define email configurations and emails as custom resources in Kubernetes and handles the sending of emails through the MailerSend API.

</details>

## Achieved Goals
<details open>
<a name="Achieved_Goals"></a>
<summary>Show/Hide</summary>
<br>

- Implemented a Kubernetes operator to manage email sending.
- Defined custom resources EmailSenderConfig and Email to configure and send emails.
- Developed a controller to handle email sending using the MailerSend API.
- Created Kubernetes manifests for deploying the operator.
- Tested the email sending functionality with MailerSend.
</details>

## What Can Be Improved
<details open>
<a name="What_Can_Be_Improved"></a>
<summary>Show/Hide</summary>
<br>

- Add support for multiple email providers, such as Mailgun, in addition to MailerSend.
- Implement more comprehensive error handling and retry mechanisms.
- Enhance logging and monitoring for better visibility and debugging.
- Add unit tests and integration tests to ensure the robustness of the operator.
- Optimize the Docker image for smaller size and faster deployment.
</details>

## Important Files
<details>
<a name="Important_Files"></a>
<summary>Show/Hide</summary>
<br>

- main.go: Entry point for the manager that starts the controller.
- api/v1/email_types.go: Contains the definitions for the custom resources EmailSenderConfig and Email.
- controllers/email_controller.go: Contains the logic for reconciling Email resources and sending emails through MailerSend.
- config/manager/manager.yaml: Kubernetes manifest for deploying the controller manager.
- config/crd/bases/email.mailerlitetask.com_emails.yaml: Custom Resource Definition (CRD) for the Email resource.
- config/crd/bases/email.mailerlitetask.com_emailsenderconfigs.yaml: Custom Resource Definition (CRD) for the EmailSenderConfig resource.
- config/rbac/role.yaml, config/rbac/role_binding.yaml, config/rbac/service_account.yaml: RBAC configuration for the operator.
- config/test/mailersend_emailsenderconfig.yaml: Sample EmailSenderConfig resource for MailerSend.
- config/test/mailersend_email.yaml: Sample Email resource for testing email sending.

</details>

<a name="Files_Directory_Tree"></a>
## Files Directory Tree
<details>
<summary>Show/Hide</summary>
<br>

mailer-operator/
├── api/
│   └── v1/
│       ├── email_types.go
│       └── zz_generated.deepcopy.go
├── config/
│   ├── crd/
│   │   └── bases/
│   │       ├── email.mailerlitetask.com_emails.yaml
│   │       └── email.mailerlitetask.com_emailsenderconfigs.yaml
│   ├── manager/
│   │   └── manager.yaml
│   ├── rbac/
│   │   ├── role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   └── test/
│       ├── mailersend_email.yaml
│       └── mailersend_emailsenderconfig.yaml
├── controllers/
│   └── email_controller.go
├── Dockerfile
├── go.mod
├── go.sum
└── main.go

</details>

## Detailed Step-by-Step Instructions
<details open>
<a name="Detailed_Step-by-Step_Instructions"></a>
<summary>Show/Hide</summary>


### Prerequisites
- Docker
- Minikube
- kubectl


#### - Windows
<details>
<summary>Show/Hide</summary>
- Install Docker Desktop from here.
- Install Minikube from here.
- Install kubectl from here.
</details>

#### - macOS
<details>
<summary>Show/Hide</summary>

- Install Docker Desktop from here.
- Install Minikube using Homebrew
```
brew install minikube
minikube start
```
- Install kubectl using Homebrew:
```
brew install kubectl
```

</details>

#### - Linux
<details>
<summary>Show/Hide</summary>

- Install Docker from here.
- Install Minikube from here.
- Install kubectl from here.

</details>

### Step 1: Clone the Repository

```
git clone https://github.com/awesomeahi95/email-operator.git
cd email-operator
```

### Step 2: Build and Push Docker Image
#### - Build the Docker image:

```
docker build -t ahilan95/email-operator:latest .
```

#### - Push the Docker image to your Docker repository:

```
docker push ahilan95/email-operator:latest
```

### Step 3: Setup Minikube and Kubernetes Resources
#### - Start Minikube:

```
minikube start
```

#### - Create the namespace:

```
kubectl create namespace mailer-operator-system
```

#### - Apply the CRDs:

```
kubectl apply -f config/crd/bases/email.mailerlitetask.com_emails.yaml
kubectl apply -f config/crd/bases/email.mailerlitetask.com_emailsenderconfigs.yaml
```

#### - Apply the RBAC configuration:

```
kubectl apply -f config/rbac/service_account.yaml
kubectl apply -f config/rbac/role.yaml
kubectl apply -f config/rbac/role_binding.yaml
```

#### - Apply the deployment:

```
kubectl apply -f config/manager/manager.yaml
```

### Step 4: Create Secrets and Resources
#### - Create Secret for MailerSend API Token

```
kubectl create secret generic mailersend-secret-token \
  --from-literal=api-token='your-api-token' \
  --from-literal=from-email='your-from-email' \
  -n mailer-operator-system
```

#### - Change Recipient Email to Preferred Email Address

config/test/mailersend_email.yaml
recipientEmail: <preferred email address>

#### - Apply the test EmailSenderConfig and Email resources:

```
kubectl apply -f config/test/mailersend_emailsenderconfig.yaml
kubectl apply -f config/test/mailersend_email.yaml
```

### Step 5: Verify the Deployment
#### - Check the logs of the controller manager to ensure emails are being sent:

```
kubectl logs -n mailer-operator-system -l control-plane=controller-manager -f
```

#### - Verify that the emails are send successfully to you given email address

</details>