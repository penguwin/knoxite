// +build backend

/*
 * knoxite
 *     Copyright (c) 2016-2020, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package ftp

import (
	"net/url"
	"os"
	"testing"

	"github.com/knoxite/knoxite/storage"
)

var (
	backendTest *storage.BackendTest
)

func TestMain(m *testing.M) {
	ftpurl := os.Getenv("KNOXITE_FTP_URL")
	if len(ftpurl) == 0 {
		panic("no backend configured")
	}

	backendTest = &storage.BackendTest{
		URL:         ftpurl,
		Protocols:   []string{"ftp"},
		Description: "FTP Storage",
		TearDown: func(tb *storage.BackendTest) {
			u, err := url.Parse(tb.URL)
			if err != nil {
				panic(err)
			}

			db := tb.Backend.(*FTPStorage)
			err = db.DeletePath(u.Path)
			if err != nil {
				panic(err)
			}
		},
	}

	storage.RunBackendTester(backendTest, m)
}

func TestStorageNewBackend(t *testing.T) {
	backendTest.NewBackendTest(t)
}

func TestStorageLocation(t *testing.T) {
	backendTest.LocationTest(t)
}

func TestStorageProtocols(t *testing.T) {
	backendTest.ProtocolsTest(t)
}

func TestStorageDescription(t *testing.T) {
	backendTest.DescriptionTest(t)
}

func TestStorageInitRepository(t *testing.T) {
	backendTest.InitRepositoryTest(t)
}

func TestStorageSaveRepository(t *testing.T) {
	backendTest.SaveRepositoryTest(t)
}

func TestStorageAvailableSpace(t *testing.T) {
	backendTest.AvailableSpaceTest(t)
}

func TestStorageSaveSnapshot(t *testing.T) {
	backendTest.SaveSnapshotTest(t)
}

func TestStorageStoreChunk(t *testing.T) {
	backendTest.StoreChunkTest(t)
}

func TestStorageDeleteChunk(t *testing.T) {
	backendTest.DeleteChunkTest(t)
}

func TestStorageLocking(t *testing.T) {
	backendTest.LockingTest(t)
}
