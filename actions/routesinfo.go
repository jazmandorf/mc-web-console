package actions

import (
	"log"
	"mc_web_console/models"
	"net/http"
	"reflect"

	"github.com/gobuffalo/buffalo"
)

// handler 전달을 위한 구조체
// router에 등록할 function 은 ( a action )을 붙여야 함.
// type actions struct{ buffalo.Context }
type actions struct{}

// 관리자 설정
func RoutesManager(app *buffalo.App) *buffalo.App {

	// TODO : DB에서 path와 handler를 가져오도록
	// 1. handler 구현한 뒤에 화면을 통해 db에 등록하는 process를 가져가도록
	// 2. 화면이 있으면 xxxform, json의 경우에는 /api/xxx 로 경로명 지정
	// ex ) "/api/<카테고리대분류>/<리소스>/<정의>" / "/<카테고리대분류>/<리소스>/mngform/"

	// ID           uuid.UUID `json:"id" db:"id"`
	// Method       string    `json:"method" db:"method"`
	// Path         string    `json:"path" db:"path"`
	// HandlerName  string    `json:"handler_name" db:"handler_name"`
	// ResourceName string    `json:"resource_name" db:"resource_name"`
	// PathName     string    `json:"path_name" db:"path_name"`
	// Aliases      string    `json:"aliases" db:"aliases"`
	// CreatedAt    time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`

	// routerList := []models.RouteInfo{
	// 	{Path: "/home1", HandlerName: "GetHome", Method: "GET"},
	// 	{Path: "/about1", HandlerName: "AboutHandler", Method: "GET"},
	// 	{Path: "/settings/resources/vpc/mngform", HandlerName: "VpcMngForm", Method: "GET"},

	// 	{Path: "/settings/resources/vpc/", HandlerName: "vpcList", Method: "GET"},
	// 	{Path: "/settings/resources/vpc/", HandlerName: "vpcReg", Method: "POST"},
	// 	{Path: "/settings/resources/vpc/id/{vNetId}", HandlerName: "vpcGet", Method: "GET"},
	// 	{Path: "/settings/resources/vpc/id/{vNetId}", HandlerName: "vpcDel", Method: "DELETE"},
	// 	{Path: "/settings/resources/vpc/region", HandlerName: "vpcListByRegion", Method: "DELETE"},
	// }

	routerList := models.RouteInfoes{}

	//err := models.DB.Find(&resultCredential, credential.ID)
	//if err != nil {
	//	return resultCredential, errors.WithStack(err)
	//}

	query := models.DB.Q()
	err := query.All(&routerList)
	if err != nil {
		log.Println("query err ", err)
		return app
	}

	for _, router := range routerList {
		//log.Println(router)
		// handlerFunction은 (a actions) function명 으로 정의 해야 함.
		handlerFunc := getHandlerFuncByName(router.HandlerName)
		if handlerFunc == nil {
			log.Println(router.HandlerName + " Handler not found")
			log.Println(router)
			continue
		}
		//log.Println(router)
		log.Println(router.Path + "   :   " + router.PathName + " : " + router.Method)
		//log.Println(handlerFunc)
		// 라우터 등록

		switch router.Method {
		case "GET":
			app.GET(router.Path, handlerFunc).Name(router.PathName)
		case "POST":
			app.POST(router.Path, handlerFunc).Name(router.PathName)
		case "PUT":
			app.PUT(router.Path, handlerFunc).Name(router.PathName)
		case "PATCH":
			app.PATCH(router.Path, handlerFunc).Name(router.PathName)
		case "HEAD":
			app.HEAD(router.Path, handlerFunc).Name(router.PathName)
		case "OPTIONS":
			app.OPTIONS(router.Path, handlerFunc).Name(router.PathName)
		case "DELETE":
			app.DELETE(router.Path, handlerFunc).Name(router.PathName)
		default:
			log.Println(" any begin~~~~~~~~~~~~~~~~~~~~~~~")
			log.Println(router)
			log.Println(" any end ~~~~~~~~~~~~~~~~~~~~~~~")
			app.ANY(router.Path, handlerFunc)
		}

	}

	return app
}

// Get the handler function by its name
func getHandlerFuncByName(handlerName string) buffalo.Handler {
	actions := &actions{}
	actionsType := reflect.TypeOf(actions)

	// 핸들러 함수 이름으로 메서드 가져오기
	method, ok := actionsType.MethodByName(handlerName)
	if !ok {
		// 핸들러 함수를 찾을 수 없는 경우에 대한 처리
		// 예: log.Fatal("Handler not found")
		return nil
	}

	// 핸들러 함수를 호출할 수 있는 함수 생성
	handlerFunc := func(c buffalo.Context) error {
		// 새로운 인스턴스 생성
		act := reflect.New(actionsType.Elem()).Interface()

		// 핸들러 함수 호출
		//return method.Func.Call([]reflect.Value{
		//	reflect.ValueOf(act),
		//	reflect.ValueOf(c),
		//})[0].Interface().(error)
		// 핸들러 함수 호출
		result := method.Func.Call([]reflect.Value{
			reflect.ValueOf(act),
			reflect.ValueOf(c),
		})

		log.Println("handlerFunc result ", result)
		if errVal := result[0]; !errVal.IsNil() {
			// Convert the error value to the error type
			err := errVal.Interface().(error)
			return err
		}
		return nil
	}

	//return handlerFunc
	return buffalo.Handler(handlerFunc)

	// // Get the method by name from the actions package
	// handlerMethod, found := reflect.TypeOf(actions{}).MethodByName(handlerName)
	// if !found {
	// 	log.Println("getHandlerFuncByName !found ", handlerName)
	// 	return nil
	// }

	// // Create a closure to execute the handler function
	// handlerFunc := func(c buffalo.Context) error {
	// 	// Instantiate the actions struct
	// 	act := actions{c}
	// 	// Call the handler method with the actions struct and context
	// 	return handlerMethod.Func.Call([]reflect.Value{reflect.ValueOf(act), reflect.ValueOf(c)})[0].Interface().(error)
	// }

	// return handlerFunc

	// handlerMethod := reflect.ValueOf(&actions{}).MethodByName(handlerName)
	// if !handlerMethod.IsValid() {
	// 	// Handler not found, handle the error
	// }

	// // Create and return the actual handler function
	// return func(c buffalo.Context) error {
	// 	// Instantiate the actions struct
	// 	act := &actions{}
	// 	// Invoke the handler function
	// 	return handlerMethod.Call([]reflect.Value{reflect.ValueOf(act), reflect.ValueOf(c)})[0].Interface().(error)
	// }

	//////////////
	// actionsType := reflect.TypeOf(actions{})
	// handlerMethod, ok := actionsType.MethodByName(handlerName)
	// if !ok {
	// 	log.Println("actionsType.MethodByName ", ok)
	// 	// Handler not found, handle the error
	// }

	// // Create and return the actual handler function
	// return func(c buffalo.Context) error {
	// 	// Instantiate the actions struct
	// 	act := reflect.New(actionsType).Interface()
	// 	log.Println("getHandlerFuncByName reflect.New ", act)
	// 	//log.Println("reflect.Value{reflect.ValueOf(act), reflect.ValueOf(c)}", reflect.Value{reflect.ValueOf(act), reflect.ValueOf(c)})
	// 	// Invoke the handler function
	// 	return handlerMethod.Func.Call([]reflect.Value{reflect.ValueOf(act), reflect.ValueOf(c)})[0].Interface().(error)
	// }
}

// 삭제 예정
func (a actions) GetHome(c buffalo.Context) error {
	log.Println("action GetHome")
	//return GetHome(c)
	return c.Render(200, r.String("Hello from GetHome"))
}

//func GetHome(c buffalo.Context) error {
//	log.Println("render GetHome")
//	return c.Render(200, r.String("Hello from GetHome"))
//}

func (a actions) AboutHandler(c buffalo.Context) error {
	//return c.Render(200, r.String("Hello from AboutHandler"))
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{
		"message":  "success",
		"status":   200,
		"VNetInfo": "aaa",
	}))
}
