package support

import (
	"context"
	"testing"

	"github.com/onsi/gomega"
	rayv1alpha1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestGetRayJob(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	_ = rayv1alpha1.AddToScheme(scheme)

	fakeRayJobs := []client.Object{
		&rayv1alpha1.RayJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-job-1",
				Namespace: "my-namespace",
			},
		},
	}

	fakeClient := NewFakeKubeClientWithScheme(scheme, fakeRayJobs...)

	rayJob := &rayv1alpha1.RayJob{}
	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "my-job-1", Namespace: "my-namespace"}, rayJob)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	g.Expect(rayJob.Name).To(gomega.Equal("my-job-1"))
	g.Expect(rayJob.Namespace).To(gomega.Equal("my-namespace"))
}

func TestGetRayCluster(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	_ = rayv1alpha1.AddToScheme(scheme)

	fakeRayCluster := []client.Object{
		&rayv1alpha1.RayCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-cluster-1",
				Namespace: "my-namespace",
			},
		},
	}

	fakeClient := NewFakeKubeClientWithScheme(scheme, fakeRayCluster...)

	raycluster := &rayv1alpha1.RayCluster{}
	err := fakeClient.Get(context.TODO(), client.ObjectKey{Name: "my-cluster-1", Namespace: "my-namespace"}, raycluster)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	g.Expect(raycluster.Name).To(gomega.Equal("my-cluster-1"))
	g.Expect(raycluster.Namespace).To(gomega.Equal("my-namespace"))
}
