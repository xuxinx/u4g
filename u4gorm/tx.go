package u4gorm

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/jinzhu/gorm"
)

func Transact(db *gorm.DB, f func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			if aerr, ok := r.(error); ok {
				err = aerr
			} else {
				err = fmt.Errorf("%+v", r)
			}
			log.Printf("transact recover error: %v", err)
			debug.PrintStack()
		}

		if err != nil {
			rberr := tx.Rollback().Error
			if rberr != nil {
				log.Printf("transact rollback failed: %v", rberr)
			}
		} else {
			err = tx.Commit().Error
		}
	}()

	return f(tx)
}
