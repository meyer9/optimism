package eth

type ExecutionWitness struct {
	Keys  map[string]string `json:"keys"`
	Codes map[string]string `json:"codes"`
	State map[string]string `json:"state"`
}
