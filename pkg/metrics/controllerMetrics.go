package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	UsersCreated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "users_created_total",
			Help: "Number of created users ",
		},
	)
	UsersCreatedErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "users_create_errors_total",
			Help: "Number of users created with errors",
		}, []string{"reason"},
	)
	UsersCreateTime = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "user_create_time_seconds",
		Help: "Histogram of user creation time in seconds",
	})
	UsersRemoved = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "users_removed_total",
			Help: "Number of removed users ",
		},
	)
	UsersRemovedErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "users_remove_errors_total",
			Help: "Number of users removed with errors",
		}, []string{"reason"},
	)
	UsersRemoveTime = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "user_remove_time_seconds",
		Help: "Histogram of user updating time in seconds",
	})
	DBProvisioningErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_provisioning_errors_total",
			Help: "Number of failed database provisioning",
		}, []string{"reason"},
	)
	DBCreated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "databases_created_total",
			Help: "Number of created databases ",
		},
	)
	PasswordRotated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "password_rotated_total",
			Help: "Number of rotated passwords",
		},
	)
	PasswordRotatedErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "password_rotate_errors_total",
			Help: "Number of passwords rotated with errors",
		}, []string{"reason"},
	)
	PasswordRotateTime = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "password_rotation_time_seconds",
		Help: "Histogram of password rotation time in seconds",
	})
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(UsersCreated, UsersCreatedErrors, UsersCreateTime)
	metrics.Registry.MustRegister(UsersRemoved, UsersRemovedErrors, UsersRemoveTime)
	metrics.Registry.MustRegister(DBCreated, DBProvisioningErrors)
	metrics.Registry.MustRegister(PasswordRotated, PasswordRotatedErrors, PasswordRotateTime)
}
