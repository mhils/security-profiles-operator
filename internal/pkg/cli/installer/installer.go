/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package installer

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-logr/logr"
	"github.com/hairyhenderson/go-which"
	apparmorprofileapi "sigs.k8s.io/security-profiles-operator/api/apparmorprofile/v1alpha1"

	"sigs.k8s.io/security-profiles-operator/internal/pkg/artifact"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/cli"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/apparmorprofile"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/apparmorprofile/crd2armor"
)

// Installer is the main structure of this package.
type Installer struct {
	impl
	options *Options
}

// New returns a new Merger instance.
func New(options *Options) *Installer {
	return &Installer{
		impl:    &defaultImpl{},
		options: options,
	}
}

// Run the Merger.
func (p *Installer) Run() error {

	log.Printf("Reading %s", p.options.profilePath)
	content, err := p.ReadFile(p.options.profilePath)
	if err != nil {
		return fmt.Errorf("open profile: %w", err)
	}

	profile, err := artifact.ReadProfile(content)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", p.options.profilePath, err)
	}

	switch obj := profile.(type) {
	case *apparmorprofileapi.AppArmorProfile:
		manager := apparmorprofile.NewAppArmorProfileManager(logr.New(&cli.LogSink{}))
		if !manager.Enabled() {
			return fmt.Errorf("Insufficient permissions or AppArmor is unavailable.")
		}

		if obj.Spec.Policy == "" {
			var programName string
			if p.options.executablePath != "" {
				programName = p.options.executablePath
			} else {
				programName = obj.Name
			}
			if resolved := which.Which(programName); resolved != "" {
				programName = resolved
			}
			if programName == "" {
				return fmt.Errorf("cannot create apparmor profile with empty name")
			}
			log.Printf("Installing AppArmor profile for: %s", programName)
			obj.Spec.Policy, err = crd2armor.GenerateProfile(programName, &obj.Spec.Abstract)
			if err != nil {
				return fmt.Errorf("build raw apparmor profile: %w", err)
			}
		}

		_, err := manager.InstallProfile(obj)
		if err != nil && !strings.Contains(err.Error(), "AppArmorProfile name must match defined policy") {
			return fmt.Errorf("install apparmor profile: %w", err)
		}
	default:
		return fmt.Errorf("cannot install %T profile", obj)
	}

	return nil
}
