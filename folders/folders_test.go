// Renamed package to "folders" to access GetAllFolders function
package folders

import (
	"testing"

	"github.com/gofrs/uuid"
	// "github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/stretchr/testify/assert"
)


func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// your test/s here

		type test struct {
			actualUUID uuid.UUID
			count int
		}

		tests := []test {
			{
				actualUUID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17b"),
				count: 0,
			},
			{
				actualUUID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				count: 5,
			},
		}

		for _, tst := range tests {
			
			fetchedUUID := FetchFolderRequest{ OrgID: tst.actualUUID}
			res, _ := GetAllFolders(&fetchedUUID)

			for _, folder := range res.Folders {
				if tst.actualUUID != folder.OrgId {
					t.Errorf("Expected UUID: (%s) is not the same as actual UUID: (%s)", folder.OrgId, tst.actualUUID)
				}
			}
			if len(res.Folders) != tst.count {
				t.Errorf("Expected count (%v), actual (%v)", len(res.Folders), tst.count)
			}
		
		}
	})
}

