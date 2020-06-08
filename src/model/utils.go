package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ModelUtils interface {
	EncodePassword() error
	IsValidPassword() error
	MascarateProvinceCity()
	MascarateHours()
}

func (self *UserAuth) EncodePassword() error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(self.Password), bcrypt.DefaultCost)
	if err != nil { return err }
	self.Password = string(hashedBytes[:])
	return nil
}
func (self UserAuth) IsValidPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(self.Password),[]byte(password))
}

func (self *BranchOffice) MascarateProvinceCity() {
	for _, value := range provinceCityColombia {
		if self.City == value.CodiCity {
			self.City = value.City
			self.Province = value.DepartmentName
		}
	}
}
func (self *BranchOffice) MascarateHours(formatEntry string, formatExit string) {
	loc, _ := time.LoadLocation("America/Bogota")
	t1,_ := time.ParseInLocation(time.RFC3339,self.ExitTime,loc)
	t2,_ := time.ParseInLocation(time.RFC3339,self.CheckInTime,loc)
	self.ExitTime = t1.Format(formatExit)
	self.CheckInTime = t2.Format(formatEntry)
}