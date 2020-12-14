package account

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountMarshal(t *testing.T) {
	var acc = &Account{
		ID:          1,
		Code:        "01",
		Name:        "test",
		Note:        "This is a testing account",
		Debit:       1000,
		Credit:      73.57,
		Operational: false,
		Inheritable: true,
	}

	var sub1 = &Account{
		ID:          2,
		ParentID:    1,
		Code:        "01.1",
		Name:        "child 1",
		Note:        "This is a testing nested account",
		Debit:       500,
		Credit:      73.57,
		Operational: false,
		Inheritable: true,
	}

	var sub2 = &Account{
		ID:          3,
		ParentID:    2,
		Code:        "01.1.1",
		Name:        "grandchild 1",
		Note:        "This is a testing nested account",
		Debit:       500,
		Credit:      73.57,
		Operational: true,
		Inheritable: false,
	}

	var sub3 = &Account{
		ID:          4,
		ParentID:    1,
		Code:        "01.2",
		Name:        "child 2",
		Note:        "This is a testing nested account",
		Debit:       500,
		Operational: true,
		Inheritable: false,
	}

	sub1.Children = []*Account{sub2}
	acc.Children = []*Account{sub1, sub3}

	bytes, err := json.Marshal(acc)
	if err != nil {
		t.Errorf("Error marshaling account: %v", err)
	}

	var str = fmt.Sprintf("%s", bytes)
	var assert = assert.New(t)

	// Children are nested, parent_id is omited
	assert.Contains(str, `{"id":1,"children":[{"id":2`)
	assert.Contains(str, `{"id":4,"parent_id":1,"code":`)
	assert.Contains(str, `{"id":2,"parent_id":1,"children":[{"id":3`)

	// Debits and credits are valid floats
	assert.Contains(str, `"debit":1000,"credit":73.57`)
	assert.Contains(str, `"debit":500,"credit":0`)
}

func TestAccountUnmarshal(t *testing.T) {
	var bytes = []byte(`{
		"id": 1,
		"code": "01",
		"parent_id": null,
		"name": "test",
		"inheritable": true,
		"debit": 73.57,
		"note": "parent account",
		"children": [
			{
				"id": 2,
				"parent_id": 1,
				"code": "01.1",
				"name": "child",
				"operational": true
			}
		]
	}`)

	var acc = &Account{}

	if err := json.Unmarshal(bytes, acc); err != nil {
		t.Errorf("Error unmarshaling account: %v", err)
	}

	var assert = assert.New(t)

	assert.Equal(acc.ID, uint64(1))
	assert.Equal(acc.Code, "01")
	assert.Equal(acc.Name, "test")
	assert.Equal(acc.Note, "parent account")
	assert.Equal(acc.Debit, 73.57)
	assert.Equal(acc.Credit, 0.0)
	assert.Equal(acc.Operational, false)
	assert.Equal(acc.Inheritable, true)
	assert.Equal(acc.ParentID, uint64(0))
	assert.Equal(acc.Children[0].ID, uint64(2))
	assert.Equal(acc.Children[0].ParentID, uint64(1))
	assert.Equal(acc.Children[0].Code, "01.1")
	assert.Equal(acc.Children[0].Operational, true)
	assert.Equal(acc.Children[0].Inheritable, false)
}
