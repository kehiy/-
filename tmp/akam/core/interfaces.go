package core

import "time"

type ICache interface {
	Set([]byte, []byte, time.Duration) error
	Have([]byte) bool
	Get([]byte) ([]byte, error)
	Delete([]byte) bool
}
