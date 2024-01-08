package support

import (
	"testing"

	"github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetPods(t *testing.T) {
	test := NewTest(t)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "test-namespace",
		},
	}

	test.client.Core().CoreV1().Pods("test-namespace").Create(test.ctx, pod, metav1.CreateOptions{})

	// Call the GetPods function with the fake client and namespace
	pods := GetPods(test, "test-namespace", metav1.ListOptions{})

	test.Expect(pods).Should(gomega.HaveLen(1), "Expected 1 pod, but got %d", len(pods))
	test.Expect(pods[0].Name).To(gomega.Equal("test-pod"), "Expected pod name 'test-pod', but got '%s'", pods[0].Name)
}

func TestGetNodes(t *testing.T) {
	test := NewTest(t)
	node := &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-node",
		},
	}

	test.client.Core().CoreV1().Nodes().Create(test.ctx, node, metav1.CreateOptions{})
	nodes := GetNodes(test)

	test.Expect(nodes).Should(gomega.HaveLen(1), "Expected 1 node, but got %d", len(nodes))
	test.Expect(nodes[0].Name).To(gomega.Equal("test-node"), "Expected node name 'test-node', but got '%s'", nodes[0].Name)

}
