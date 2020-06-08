package employee

import (
	"github.com/kataras/iris/v12"
)

func IndexPost(ctx iris.Context) {
	/*var dataUserAuthValidator model.UserAuthValidator
	ctx.ReadJSON(&dataUserAuthValidator)
	// encode password
	encode, err := utils.EncodePassword(dataUserAuthValidator.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err})
	}
	dataUserAuthValidator.Password = encode
	fmt.Println(dataUserAuthValidator.Password)
	existError := dataUserAuthValidator.Validate()
	if len(existError) > 0 {
		ctx.StatusCode(iris.StatusNotAcceptable)
		ctx.JSON(iris.Map{"message": existError})
	}
	return*/
}