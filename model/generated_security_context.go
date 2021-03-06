package model

const (
	SECURITY_CONTEXT_TYPE = "v1.SecurityContext"
)

type SecurityContext struct {
	Capabilities Capabilities `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`

	Privileged bool `json:"privileged,omitempty" yaml:"privileged,omitempty"`

	RunAsUser int64 `json:"runAsUser,omitempty" yaml:"run_as_user,omitempty"`

	SeLinuxOptions SELinuxOptions `json:"seLinuxOptions,omitempty" yaml:"se_linux_options,omitempty"`
}
