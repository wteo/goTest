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
			uuid uuid.UUID
			count int
		}

		tests := []test {
			{
				uuid: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17b"),
				count: 0,
			},
			{
				uuid: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				count: 999,
			},
		}

		for _, test := range tests {

			fetchedUUID := FetchFolderRequest{ OrgID: test.uuid}
			res, _ := GetAllFolders(&fetchedUUID)
			expectedCount := len(res.Folders)

			for _, folder := range res.Folders {
				if test.uuid != folder.OrgId {
					t.Errorf("Expected UUID: (%s) is not the same as actual UUID: (%s)", folder.OrgId, test.uuid)
				}
			}
			if expectedCount != test.count {
				t.Errorf("Expected count (%v), actual (%v)", expectedCount, test.count)
			}
		
		}
	})
}

