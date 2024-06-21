package controllers

import (
        "context"
        "testing"
        "time"

        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/client-go/kubernetes/scheme"
        ctrl "sigs.k8s.io/controller-runtime"
        "sigs.k8s.io/controller-runtime/pkg/client"
        "sigs.k8s.io/controller-runtime/pkg/client/fake"
        "sigs.k8s.io/controller-runtime/pkg/envtest"

        emailv1 "github.com/awesomeahi95/mailerlite/api/v1"
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
)

var _ = Describe("Email Controller", func() {
        var (
                testEnv   *envtest.Environment
                k8sClient client.Client
                ctx       context.Context
                cancel    context.CancelFunc
        )

        BeforeEach(func() {
                testEnv = &envtest.Environment{
                        CRDDirectoryPaths: []string{"../config/crd/bases"},
                }
                cfg, err := testEnv.Start()
                Expect(err).NotTo(HaveOccurred())
                Expect(cfg).NotTo(BeNil())

                err = emailv1.AddToScheme(scheme.Scheme)
                Expect(err).NotTo(HaveOccurred())

                k8sClient = fake.NewClientBuilder().WithScheme(scheme.Scheme).Build()

                mgr, err := ctrl.NewManager(cfg, ctrl.Options{
                        Scheme: scheme.Scheme,
                })
                Expect(err).ToNot(HaveOccurred())

                err = (&EmailReconciler{
                        Client: k8sClient,
                        Scheme: scheme.Scheme,
                }).SetupWithManager(mgr)
                Expect(err).ToNot(HaveOccurred())

                ctx, cancel = context.WithCancel(context.Background())
                go func() {
                        defer GinkgoRecover()
                        Expect(mgr.Start(ctx)).To(Succeed())
                }()
        })

        AfterEach(func() {
                cancel()
                err := testEnv.Stop()
                Expect(err).NotTo(HaveOccurred())
        })

        Context("When creating an Email resource", func() {
                It("Should update the status after sending the email via MailerSend", func() {
                        email := &emailv1.Email{
                                ObjectMeta: metav1.ObjectMeta{
                                        Name:      "sample-email",
                                        Namespace: "default",
                                },
                                Spec: emailv1.EmailSpec{
                                        SenderConfigRef: "sample-senderconfig",
                                        RecipientEmail:  "recipient@example.com",
                                        Subject:         "Sample Email",
                                        Body:            "This is a sample email.",
                                },
                        }
                        err := k8sClient.Create(context.Background(), email)
                        Expect(err).ToNot(HaveOccurred())

                        Eventually(func() string {
                                err := k8sClient.Get(context.Background(), client.ObjectKey{
                                        Namespace: email.Namespace,
                                        Name:      email.Name,
                                }, email)
                                Expect(err).ToNot(HaveOccurred())
                                return email.Status.DeliveryStatus
                        }, time.Second*10).Should(Equal("Sent"))
                })
        })
})

func TestEmailController(t *testing.T) {
        RegisterFailHandler(Fail)
        RunSpecs(t, "Email Controller Suite")
}
