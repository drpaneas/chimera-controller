package admission

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/chimera-kube/chimera-controller/internal/pkg/constants"
)

func (r *AdmissionReconciler) reconcilePolicyServerService(ctx context.Context) error {
	err := r.Client.Create(ctx, r.service())
	if err == nil || apierrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}

func (r *AdmissionReconciler) service() *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.PolicyServerServiceName,
			Namespace: r.DeploymentsNamespace,
			Labels:    constants.PolicyServerLabels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port:     constants.PolicyServerPort,
					Protocol: corev1.ProtocolTCP,
				},
			},
			Selector: constants.PolicyServerLabels,
		},
	}
}
