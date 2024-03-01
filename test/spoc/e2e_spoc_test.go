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

package main_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/release-utils/util"
)

//nolint:paralleltest // should not run in parallel
func TestSpoc(t *testing.T) {
	logrus.Info("Building demobinary...")
	cmd := exec.Command("go", "build", "demobinary.go")
	err := cmd.Run()
	require.Nil(t, err, "failed to build demobinary.go")
	err = util.CopyFileLocal("demobinary", "demobinary-child", true)
	require.Nil(t, err)
	err = os.Chmod("demobinary-child", 0o700)
	require.Nil(t, err)

	t.Run("record", recordTest)
}

func recordTest(t *testing.T) {
	t.Run("AppArmor", recordAppArmorTest)
	t.Run("Seccomp", recordSeccompTest)
}

func recordAppArmorTest(t *testing.T) {
	t.Run("files", func(t *testing.T) {
		cupaloy.SnapshotT(
			t,
			// this still shows a bug: the path to README should be resolved.
			record(t, "apparmor", "--file-read", "../../README.md", "--file-write", "/dev/null"),
			record(t, "apparmor", "--file-read", "/dev/null", "--file-write", "/dev/null"),
		)
	})
	t.Run("sockets", func(t *testing.T) {
		cupaloy.SnapshotT(
			t,
			record(t, "apparmor", "--net-tcp"),
			// Go is doing Go things and does some TCP syscalls when opening a UDP socket.
			// So the snapshot looks a bit weird here.
			record(t, "apparmor", "--net-udp"),
			record(t, "apparmor", "--net-icmp"),
		)
	})

	t.Run("subprocess", func(t *testing.T) {
		// Ensure that we can run subprocesses and that their action are recorded, too.
		cupaloy.SnapshotT(
			t,
			record(t, "apparmor", "./demobinary-child", "--file-read", "/dev/null"),
			record(t, "apparmor", "./demobinary", "--file-read", "/dev/null"),
		)
	})
}

func recordSeccompTest(t *testing.T) {
	// smoke test for seccomp
	record(t, "seccomp", "./demobinary-child", "--net-tcp")
	// TODO: add snapshot testing here - currently disabled because flaky
}

func runSpoc(t *testing.T, args ...string) []byte {
	t.Helper()
	args = append([]string{"../../build/spoc"}, args...)
	cmd := exec.Command("sudo", args...)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	require.Nil(t, err, "failed to run spoc")
	return out
}

func record(t *testing.T, typ string, args ...string) []byte {
	t.Helper()
	args = append([]string{
		"record", "-t", typ, "-o", "/dev/stdout", "--no-base-syscalls", "./demobinary",
	}, args...)
	return runSpoc(t, args...)
}
