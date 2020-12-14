package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransactionMarshal(t *testing.T) {
	var time = time.Now()
	var trans = &Transaction{
		ID:       1,
		DateTime: time,
		DebitID:  1,
		CreditID: 10,
		Amount:   1000.0,
		Note:     "test transaction",
	}

	bytes, err := json.Marshal(trans)
	if err != nil {
		t.Errorf("Error marshaling transaction: %v", err)
	}

	var str = fmt.Sprintf("%s", bytes)
	var assert = assert.New(t)

	assert.Contains(str, `{"id":1`)
	assert.Contains(str, time.Format("2006-01-02T15:04:05."))
	assert.Contains(str, `"debit_id":1`)
	assert.Contains(str, `"credit_id":10`)
	assert.Contains(str, `"amount":1000`)
	assert.Contains(str, `"note":"test transaction"}`)
}
