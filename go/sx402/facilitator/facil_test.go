package facilitator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/san-lab/sx402/all712"
)

func TestVoidJSON(t *testing.T) {
	empty := []byte("{}")
	envelope := all712.Envelope{}
	json.Unmarshal(empty, &envelope)
	fmt.Println(envelope.PaymentPayload)
}
