// Copyright 2020 NetApp, Inc. All Rights Reserved.

// +build linux

package utils

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDetermineNFSPackages(t *testing.T) {
	log.Debug("Running TestDetermineNFSPackages...")

	tests := map[string]struct {
		host             HostSystem
		expectedPackages []string
		errorExpected    bool
	}{
		"Ubuntu": {
			host: HostSystem{OS: SystemOS{
				Distro:  Ubuntu,
				Version: "18.04",
				Release: "18.04.1",
			}},
			expectedPackages: []string{"nfs-common"},
			errorExpected:    false,
		},
		"Centos": {
			host: HostSystem{OS: SystemOS{
				Distro:  Centos,
				Version: "8",
				Release: "8.1",
			}},
			expectedPackages: []string{"nfs-utils"},
			errorExpected:    false,
		},
		"RHEL": {
			host: HostSystem{OS: SystemOS{
				Distro:  RHEL,
				Version: "8",
				Release: "8.1",
			}},
			expectedPackages: []string{"nfs-utils"},
			errorExpected:    false,
		},
		"Foobar": {
			host: HostSystem{OS: SystemOS{
				Distro:  "Foobar",
				Version: "",
				Release: "",
			}},
			expectedPackages: nil,
			errorExpected:    true,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case '%s'", testName)

		packages, err := determineNFSPackages(context.Background(), test.host)
		assert.Equal(t, test.errorExpected, err != nil, "Unexpected error value")
		assert.ElementsMatch(t, test.expectedPackages, packages, "Incorrect packages returned")
	}
}

func TestGetPackageManagerForHost(t *testing.T) {
	log.Debug("Running TestGetPackageManagerForHost...")

	tests := map[string]struct {
		host          HostSystem
		expectedPM    string
		errorExpected bool
	}{
		"Ubuntu": {
			host: HostSystem{OS: SystemOS{
				Distro:  Ubuntu,
				Version: "18.04",
				Release: "18.04.1",
			}},
			expectedPM:    "apt",
			errorExpected: false,
		},
		"Centos7": {
			host: HostSystem{OS: SystemOS{
				Distro:  Centos,
				Version: "7",
				Release: "7.1",
			}},
			expectedPM:    "yum",
			errorExpected: false,
		},
		"RHEL7": {
			host: HostSystem{OS: SystemOS{
				Distro:  RHEL,
				Version: "7",
				Release: "7.1",
			}},
			expectedPM:    "yum",
			errorExpected: false,
		},
		"Centos8": {
			host: HostSystem{OS: SystemOS{
				Distro:  Centos,
				Version: "8",
				Release: "8.1",
			}},
			expectedPM:    "dnf",
			errorExpected: false,
		},
		"RHEL8": {
			host: HostSystem{OS: SystemOS{
				Distro:  RHEL,
				Version: "8",
				Release: "8.1",
			}},
			expectedPM:    "dnf",
			errorExpected: false,
		},
		"Foobar": {
			host: HostSystem{OS: SystemOS{
				Distro:  "Foobar",
				Version: "",
				Release: "",
			}},
			expectedPM:    "",
			errorExpected: true,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case '%s'", testName)

		pm, err := getPackageManagerForHost(context.Background(), test.host)
		assert.Equal(t, test.errorExpected, err != nil, "Unexpected error value")
		assert.Equal(t, test.expectedPM, pm, "Incorrect package manager returned")
	}
}
