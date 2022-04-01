package receipt

import (
	"time"

	"github.com/readreceipt/api/repository"
	"github.com/readreceipt/api/service/cache"
)

func IsRead(signature, id string) (bool, error) {
	exists, err := cache.Exists(signature)

	if err != nil {
		return false, err
	}

	if exists {
		return true, nil
	}

	if err := cache.Store(signature, true, time.Hour*72); err != nil {
		return false, err
	}

	isRead, err := repository.IsReceiptRead(id)

	if err != nil {
		return false, err
	}

	return isRead, nil
}
